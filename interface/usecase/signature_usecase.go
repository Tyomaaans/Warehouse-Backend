package usecase

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"time"

	"github.com/google/uuid"

	"Backend-Warehouse/domain/entity"
	"Backend-Warehouse/domain/repository"
	"Backend-Warehouse/interface/dto"
	"Backend-Warehouse/interface/httpclient"
)

type SignatureUseCase struct {
	signatureRepo repository.SignatureRepository
	employeeRepo  repository.EmployeeRepository
	supplierRepo  repository.SupplierRepository
	fileClient    *httpclient.PythonClient
}

func NewSignatureUseCase(
	signatureRepo repository.SignatureRepository,
	employeeRepo  repository.EmployeeRepository,
	supplierRepo  repository.SupplierRepository,
	fileClient    *httpclient.PythonClient,
) *SignatureUseCase {
	return &SignatureUseCase{
		signatureRepo: signatureRepo,
		employeeRepo:  employeeRepo,
		supplierRepo:  supplierRepo,
		fileClient:    fileClient,	
	}
}

func (uc *SignatureUseCase) Register(ctx context.Context, employeeID string, file multipart.File, header *multipart.FileHeader) error {
    const ownerType = "employees"

    // 1. Get employee data
    employee, err := uc.employeeRepo.GetEmployeeByEmployeeID(ctx, employeeID)
    if err != nil {
        return fmt.Errorf("employee not found: %w", err)
    }

	log.Println("employee get success")

    // 2. Build QrData
    signatureID := uuid.New().String()
    qrData := dto.QrDataRequest{
        SignatureID: signatureID,
        Name:        employee.Name,
        Position:    employee.Role,
        Company:     "Warehouse Company",
        ValidUntil:  time.Now().AddDate(1, 0, 0),
        VerifyURL:   fmt.Sprintf("https://yourcompany.com/verify/%s", signatureID),
    }

	log.Println("employee build qr data success")

    // 3. Call Python service (upload signature + generate QR → MinIO)
    pythonResp, err := uc.fileClient.GenerateSignature(ctx, file, header, qrData)
    if err != nil {
        // File gagal upload → tidak ada yang masuk DB, aman
        return fmt.Errorf("failed to generate signature: %w", err)
    }

	log.Println("generate signature success")

    // 4. Begin transaction — deactivate lama + insert baru harus atomic
    tx := uc.signatureRepo.BeginTx(ctx)
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

    // 5. Deactivate semua signature lama
    if err := uc.signatureRepo.DeactivateAllWithTx(ctx, tx, employeeID, ownerType); err != nil {
        tx.Rollback()
        // TODO: hapus file yang sudah terupload di MinIO (kompensasi)
        return fmt.Errorf("failed to deactivate old signatures: %w", err)
    }

    // 6. Insert signature baru
    signatureEntity := &entity.SignatureEntity{
        SignatureID: signatureID,
        OwnerID:     employeeID,
        OwnerType:   ownerType,
        Signature:   pythonResp.SignatureURL,
        QrData: entity.QrData{
            SignatureID: qrData.SignatureID,
            Name:        qrData.Name,
            Position:    qrData.Position,
            Company:     qrData.Company,
            ValidUntil:  qrData.ValidUntil,
            VerifyURL:   qrData.VerifyURL,
        },
        QrImage:  pythonResp.QrURL,
        IsActive: true,
    }

    if err := uc.signatureRepo.RegisterWithTx(ctx, tx, signatureEntity); err != nil {
        tx.Rollback()
        // TODO: hapus file yang sudah terupload di MinIO (kompensasi)
        return fmt.Errorf("failed to save signature: %w", err)
    }

    // 7. Commit
    if err := tx.Commit().Error; err != nil {
        tx.Rollback()
        // TODO: hapus file yang sudah terupload di MinIO (kompensasi)
        return fmt.Errorf("failed to commit transaction: %w", err)
    }

    return nil
}

func (uc *SignatureUseCase) GetActive(ctx context.Context, employeeID string) (*entity.SignatureEntity, error) {
    return uc.signatureRepo.GetActive(ctx, employeeID, "employees")
}

func (uc *SignatureUseCase) GetAll(ctx context.Context, employeeID string) ([]entity.SignatureEntity, error) {
    return uc.signatureRepo.GetAll(ctx, employeeID, "employees")
}