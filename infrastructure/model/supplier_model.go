package model

import (
	"time"

	"gorm.io/gorm"
)

type SupplierStorage struct {
	SupplierID   string         `gorm:"primaryKey"`
	Name         string         `gorm:"not null"`
	Company      string         `gorm:"not null"`
	Email        string         `gorm:"uniqueIndex;not null"`
	Phone        string         `gorm:"uniqueIndex;not null"`
	Address      string         `gorm:"not null"`
	CreatedAt    time.Time      `gorm:"not null"`
	UpdatedAt    time.Time      `gorm:"not null"` 
	DeletedAt    gorm.DeletedAt `gorm:"index"`

	GoodsReceipts []GoodsReceiptStorage     `gorm:"foreignKey:SupplierID"`
	Signatures    []SignatureStorage        `gorm:"polymorphic:Owner"`
}