package model

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/datatypes"
)

type SignatureStorage struct {
	SignatureID string         `gorm:"primaryKey"`
	OwnerID     string         `gorm:"not null"`
	OwnerType   string         `gorm:"not null"`
	Signature   string         `gorm:"not null"`
	QrData      datatypes.JSON `gorm:"type:jsonb;not null"`
	QrImage     string         `gorm:"not null"`
	IsActive    bool           `gorm:"default:true;not null"`
	CreatedAt   time.Time      `gorm:"not null"`
	UpdatedAt   time.Time      `gorm:"not null"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}