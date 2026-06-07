package entity

import "time"

type QrData struct {
	SignatureID string    `json:"signature_id"`
	Name        string    `json:"name"`
	Position    string    `json:"position"`
	Company     string    `json:"company"`
	ValidUntil  time.Time `json:"valid_until"`
	VerifyURL   string    `json:"verify_url"`
}

type SignatureEntity struct {
	SignatureID string
	OwnerID     string
	OwnerType   string
	Signature   string
	QrData      QrData
	QrImage     string
	IsActive    bool
}

type UpdateSignatureEntity struct {
	SignatureID string
	OwnerID     *string
	OwnerType   *string
	Signature   *string
	QrData      *QrData
	QrImage     *string
}