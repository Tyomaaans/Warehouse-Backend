package model

import (
	"time"

	"gorm.io/gorm"
)

type EmployeeStorage struct {
	EmployeeID string         `gorm:"primaryKey"`
	Name       string         `gorm:"not null"`
	Email      string         `gorm:"uniqueIndex;not null"`
	Username   string         `gorm:"uniqueIndex;not null"`
	Phone      string         `gorm:"uniqueIndex;not null"`
	Address    string         `gorm:"not null"`
	Role       string         `gorm:"not null"`
	Avatar     string         `gorm:"default:''"`
	Password   string         `gorm:"not null"`
	CreatedAt  time.Time      `gorm:"not null"`
	UpdatedAt  time.Time      `gorm:"not null"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`

	GoodsReceipts []GoodsReceiptStorage    `gorm:"foreignKey:EmployeeID"`
	GoodsIssues   []GoodsIssueStorage      `gorm:"foreignKey:EmployeeID"`
	Signatures    []SignatureStorage       `gorm:"polymorphic:Owner"`
}