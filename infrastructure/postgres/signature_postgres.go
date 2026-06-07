package postgres

import (
	"context"

	"gorm.io/gorm"

	"Backend-Warehouse/domain/entity"
	"Backend-Warehouse/domain/repository"
	"Backend-Warehouse/infrastructure/model"
	"Backend-Warehouse/interface/mapper"
)

type SignatureRepository struct {
    db *gorm.DB
}

func NewSignatureRepository(db *gorm.DB) *SignatureRepository {
    return &SignatureRepository{
		db: db,
	}
}

func (r *SignatureRepository) BeginTx(ctx context.Context) *gorm.DB {
    return r.db.WithContext(ctx).Begin()
}

func (r *SignatureRepository) RegisterWithTx(ctx context.Context, tx *gorm.DB, signature *entity.SignatureEntity) error {
    storage := mapper.ToSignatureStorage(signature)
    return tx.WithContext(ctx).Create(storage).Error
}

func (r *SignatureRepository) DeactivateAllWithTx(ctx context.Context, tx *gorm.DB, ownerID, ownerType string) error {
    return tx.WithContext(ctx).
        Model(&model.SignatureStorage{}).
        Where("owner_id = ? AND owner_type = ? AND is_active = true", ownerID, ownerType).
        Update("is_active", false).Error
}

func (r *SignatureRepository) GetActive(ctx context.Context, ownerID, ownerType string) (*entity.SignatureEntity, error) {
    var storage model.SignatureStorage
    err := r.db.WithContext(ctx).
        Where("owner_id = ? AND owner_type = ? AND is_active = true", ownerID, ownerType).
        First(&storage).Error
    if err != nil {
        return nil, repository.HandleDBError(err)
    }
    return mapper.ToSignatureEntity(&storage), nil
}

func (r *SignatureRepository) GetAll(ctx context.Context, ownerID, ownerType string) ([]entity.SignatureEntity, error) {
    var storages []model.SignatureStorage
    err := r.db.WithContext(ctx).
        Where("owner_id = ? AND owner_type = ?", ownerID, ownerType).
        Order("created_at DESC").
        Find(&storages).Error
    if err != nil {
        return nil, repository.HandleDBError(err)
    }

    entities := make([]entity.SignatureEntity, len(storages))
    for i, s := range storages {
        entities[i] = *mapper.ToSignatureEntity(&s)
    }
    return entities, nil
}