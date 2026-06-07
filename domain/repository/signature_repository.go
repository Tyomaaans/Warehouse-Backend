package repository

import (
	"context"

	"gorm.io/gorm"

	"Backend-Warehouse/domain/entity"
)

type SignatureRepository interface {
    RegisterWithTx(ctx context.Context, tx *gorm.DB, signature *entity.SignatureEntity) error
    DeactivateAllWithTx(ctx context.Context, tx *gorm.DB, ownerID, ownerType string) error
    GetActive(ctx context.Context, ownerID, ownerType string) (*entity.SignatureEntity, error)
    GetAll(ctx context.Context, ownerID, ownerType string) ([]entity.SignatureEntity, error)
    BeginTx(ctx context.Context) *gorm.DB
}