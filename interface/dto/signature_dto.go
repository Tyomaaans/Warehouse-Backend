package dto

import "time"

type RegisterSignatureRequest struct {
	OwnerID   string        `json:"owner_id" validate:"required"`
	OwnerType string        `json:"owner_type" validate:"required,oneof=employees suppliers"`
	Signature string        `json:"signature" validate:"required"`
	QrData    QrDataRequest `json:"qr_data" validate:"required"`
}

type QrDataRequest struct {
	SignatureID string    `json:"signature_id" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	Position    string    `json:"position" validate:"required"`
	Company     string    `json:"company" validate:"required"`
	ValidUntil  time.Time `json:"valid_until" validate:"required,datetime=2006-01-02"`
	VerifyURL   string    `json:"verify_url" validate:"required"`
}

type SignatureResponse struct {
    SignatureURL string `json:"signature_url"`
    QrURL        string `json:"qr_url"`
}