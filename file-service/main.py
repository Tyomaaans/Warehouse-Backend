import io
import json
import os
import sys
import uuid
from contextlib import asynccontextmanager
from typing import List

import qrcode
from fastapi import FastAPI, File, Form, HTTPException, UploadFile
from fastapi.responses import JSONResponse
from minio import Minio
from minio.error import S3Error
from pydantic import BaseModel
from pydantic_settings import BaseSettings


# ── Settings ──────────────────────────────────────────────────
class Settings(BaseSettings):
    minio_public_url: str
    minio_endpoint: str
    minio_access_key: str
    minio_secret_key: str
    minio_avatar_bucket: str
    minio_signature_bucket: str
    minio_qr_bucket: str

    class Config:
        env_file = ".env"


def load_settings() -> Settings:
    try:
        return Settings()
    except Exception as e:
        print(f"ERROR: Missing environment variables: {e}")
        sys.exit(1)


settings = load_settings()

MAX_FILE_SIZE = 2 * 1024 * 1024  # 2 MB
ALLOWED_TYPES = {"image/jpeg", "image/png", "image/webp"}
ALLOWED_EXTS  = {"jpg", "jpeg", "png", "webp"}


# ── MinIO Client ──────────────────────────────────────────────
minio_client = Minio(
    settings.minio_endpoint,
    access_key=settings.minio_access_key,
    secret_key=settings.minio_secret_key,
    secure=False,
)


def ensure_bucket(bucket: str):
    try:
        if not minio_client.bucket_exists(bucket):
            minio_client.make_bucket(bucket)
            print(f"Bucket '{bucket}' created")
        else:
            print(f"Bucket '{bucket}' already exists")
    except S3Error as e:
        print(f"ERROR: Failed to ensure bucket '{bucket}': {e}")
        sys.exit(1)


# ── Lifespan ──────────────────────────────────────────────────
@asynccontextmanager
async def lifespan(app: FastAPI):
    ensure_bucket(settings.minio_avatar_bucket)
    ensure_bucket(settings.minio_signature_bucket)
    ensure_bucket(settings.minio_qr_bucket)
    yield


app = FastAPI(lifespan=lifespan)


# ── Helpers ───────────────────────────────────────────────────
def upload_to_minio(bucket: str, object_name: str, data: bytes, content_type: str) -> str:
    try:
        minio_client.put_object(
            bucket,
            object_name,
            data=io.BytesIO(data),
            length=len(data),
            content_type=content_type,
        )
    except S3Error as e:
        raise HTTPException(status_code=500, detail=f"MinIO error: {e}")

    return f"{settings.minio_public_url}/{bucket}/{object_name}"


def delete_from_minio(bucket: str, object_name: str):
    try:
        minio_client.remove_object(bucket, object_name)
    except S3Error as e:
        print(f"WARNING: Failed to delete '{object_name}' from '{bucket}': {e}")


def extract_object_name(url: str, bucket: str) -> str:
    prefix = f"{settings.minio_public_url}/{bucket}/"
    return url.replace(prefix, "") if url.startswith(prefix) else ""


def validate_image_file(file: UploadFile) -> bytes:
    if file.content_type not in ALLOWED_TYPES:
        raise HTTPException(400, f"Only {', '.join(ALLOWED_TYPES)} allowed")

    ext = file.filename.rsplit(".", 1)[-1].lower() if "." in file.filename else ""
    if ext not in ALLOWED_EXTS:
        raise HTTPException(400, f"Invalid file extension: .{ext}")

    contents = file.file.read()
    if len(contents) > MAX_FILE_SIZE:
        raise HTTPException(400, f"File too large, max {MAX_FILE_SIZE // (1024 * 1024)}MB")

    return contents


# ── Models ────────────────────────────────────────────────────
class QrDataPayload(BaseModel):
    signature_id: str
    name: str
    position: str
    company: str
    valid_until: str
    verify_url: str


class DeleteFilesRequest(BaseModel):
    urls: List[str]


# ── Routes ────────────────────────────────────────────────────
@app.get("/health")
def health():
    return {"status": "ok"}


@app.post("/upload-avatar")
def upload_avatar(file: UploadFile = File(...)):
    contents = validate_image_file(file)

    ext = file.filename.rsplit(".", 1)[-1].lower()
    object_name = f"{uuid.uuid4()}.{ext}"

    url = upload_to_minio(
        bucket=settings.minio_avatar_bucket,
        object_name=object_name,
        data=contents,
        content_type=file.content_type,
    )

    return {"url": url}


@app.post("/generate-signature")
def generate_signature(
    signature: UploadFile = File(...),
    qr_data: str = Form(...),
):
    sig_contents = validate_image_file(signature)

    try:
        qr_payload = QrDataPayload(**json.loads(qr_data))
    except Exception as e:
        raise HTTPException(400, f"Invalid qr_data: {e}")

    sig_ext = signature.filename.rsplit(".", 1)[-1].lower()
    sig_object_name = f"{uuid.uuid4()}.{sig_ext}"

    signature_url = upload_to_minio(
        bucket=settings.minio_signature_bucket,
        object_name=sig_object_name,
        data=sig_contents,
        content_type=signature.content_type,
    )

    try:
        qr = qrcode.QRCode(
            version=1,
            error_correction=qrcode.constants.ERROR_CORRECT_H,
            box_size=10,
            border=4,
        )
        qr.add_data(qr_payload.model_dump_json())
        qr.make(fit=True)

        qr_img = qr.make_image(fill_color="black", back_color="white").convert("RGB")
        qr_buffer = io.BytesIO()
        qr_img.save(qr_buffer, format="PNG")
        qr_bytes = qr_buffer.getvalue()
    except Exception as e:
        delete_from_minio(settings.minio_signature_bucket, sig_object_name)
        raise HTTPException(500, f"Failed to generate QR: {e}")

    qr_object_name = f"{uuid.uuid4()}.png"

    try:
        qr_url = upload_to_minio(
            bucket=settings.minio_qr_bucket,
            object_name=qr_object_name,
            data=qr_bytes,
            content_type="image/png",
        )
    except Exception as e:
        delete_from_minio(settings.minio_signature_bucket, sig_object_name)
        raise HTTPException(500, f"Failed to upload QR: {e}")

    return {
        "signature_url": signature_url,
        "qr_url": qr_url,
    }


@app.delete("/delete-files")
def delete_files(payload: DeleteFilesRequest):
    for url in payload.urls:
        for bucket in [
            settings.minio_avatar_bucket,
            settings.minio_signature_bucket,
            settings.minio_qr_bucket,
        ]:
            object_name = extract_object_name(url, bucket)
            if object_name:
                delete_from_minio(bucket, object_name)
                break

    return {"message": "files deleted"}