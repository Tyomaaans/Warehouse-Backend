package model

import (
	"time"

	"gorm.io/gorm"
)

type ItemStorage struct {
	ItemID       string         `gorm:"primaryKey"`
	ItemCode     string         `gorm:"uniqueIndex;not null"`
	ItemName     string         `gorm:"uniqueIndex;not null"`
	CategoryID   string         `gorm:"not null"`
	UnitOfStock  string         `gorm:"not null"`
	Stock        int64          `gorm:"not null"`
	MinimumStock int64          `gorm:"not null"`
	CreatedAt    time.Time      `gorm:"not null"`
	UpdatedAt    time.Time      `gorm:"not null"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`

	GoodsReceiptDetails []GoodsReceiptDetailStorage `gorm:"foreignKey:ItemID"`
	GoodsIssueDetails   []GoodsIssueDetailStorage   `gorm:"foreignKey:ItemID"`
}

type CategoryStorage struct {
	CategoryID   string         `gorm:"primaryKey"`
	CategoryName string         `gorm:"uniqueIndex;not null"`
	CreatedAt    time.Time      `gorm:"not null"`
	UpdatedAt    time.Time      `gorm:"not null"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`

	Items []ItemStorage `gorm:"foreignKey:CategoryID"`
}