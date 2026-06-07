package model

import (
	"time"

	"gorm.io/gorm"
)

type GoodsReceiptStorage struct {
	GoodsReceiptID string         `gorm:"primaryKey"`
	Invoice        string         `gorm:"uniqueIndex;not null"`
	SupplierID     string         `gorm:"not null"`
	DateOfEntry    time.Time      `gorm:"not null"`
	EmployeeID     string         `gorm:"not null"`
	Status         string         `gorm:"not null"`
	CreatedAt      time.Time      `gorm:"not null"`
	UpdatedAt      time.Time      `gorm:"not null"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`

	GoodsReceiptDetails []GoodsReceiptDetailStorage `gorm:"foreignKey:GoodsReceiptID"`
}

type GoodsReceiptDetailStorage struct {
	GoodsReceiptDetailID string         `gorm:"primaryKey"`
	GoodsReceiptID       string         `gorm:"not null"`
	ItemID               string         `gorm:"not null"`
	Qty                  int64          `gorm:"not null"`
	CreatedAt            time.Time      `gorm:"not null"`
	UpdatedAt            time.Time      `gorm:"not null"`
	DeletedAt            gorm.DeletedAt `gorm:"index"`
}

type GoodsIssueStorage struct {
	GoodsIssueID    string         `gorm:"primaryKey"`
	Invoice         string         `gorm:"uniqueIndex;not null"`
	ExitDestination string         `gorm:"not null"`
	DateOfExit      time.Time      `gorm:"not null"`
	EmployeeID      string         `gorm:"not null"`
	Status          string         `gorm:"not null"`
	CreatedAt       time.Time      `gorm:"not null"`
	UpdatedAt       time.Time      `gorm:"not null"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`

	GoodsIssueDetails []GoodsIssueDetailStorage `gorm:"foreignKey:GoodsIssueID"`
}

type GoodsIssueDetailStorage struct {
	GoodsIssueDetailID string         `gorm:"primaryKey"`
	GoodsIssueID       string         `gorm:"not null"`
	ItemID             string         `gorm:"not null"`
	Qty                int64          `gorm:"not null"`
	CreatedAt          time.Time      `gorm:"not null"`
	UpdatedAt          time.Time      `gorm:"not null"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}
