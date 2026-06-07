package postgres

import (
	"Backend-Warehouse/domain/entity"
	"Backend-Warehouse/domain/repository"
	"Backend-Warehouse/infrastructure/model"
	"Backend-Warehouse/interface/mapper"
	"context"

	"gorm.io/gorm"
)

type SupplierRepository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) *SupplierRepository {
	return &SupplierRepository{
		db: db,
	}
}

func (r *SupplierRepository) CreateSupplier(ctx context.Context, supplier *entity.SupplierEntity) error {
	supplierStorage := mapper.ToSupplierStorage(supplier)

	if err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(supplierStorage).Error; err != nil {
			return repository.HandleDBError(err)
		}

		return nil
	}); err != nil { return err }

	return nil
}

func (r *SupplierRepository) GetSuppliers(ctx context.Context) ([]entity.SupplierEntity, error) {
	var supplierStorages []model.SupplierStorage

	err := r.db.WithContext(ctx).
		Preload("GoodsReceipts").
		Order("created_at ASC").
		Find(&supplierStorages).Error

	if err != nil {
        return nil, repository.HandleDBError(err)
    }

	var suppliers []entity.SupplierEntity
    for _, supplierStorage := range supplierStorages {
        supplier := mapper.ToSupplierEntity(&supplierStorage)
        suppliers = append(suppliers, *supplier)
    }

    return suppliers, nil
}

func (r *SupplierRepository) GetSupplierBySupplierID(ctx context.Context, supplierID string) (*entity.SupplierEntity, error) {
	var supplierStorage model.SupplierStorage

	err := r.db.WithContext(ctx).
		Preload("GoodsReceipts").
		Order("created_at ASC").
		Find(&supplierStorage).Error

	if err != nil {
		return nil, repository.HandleDBError(err)
	}

	supplier := mapper.ToSupplierEntity(&supplierStorage)

	return supplier, nil
}

func (r *SupplierRepository) UpdateSupplier(ctx context.Context, update *entity.UpdateSupplierEntity) error {
	fields := map[string]any{}

	if update.Name    != nil && *update.Name    != "" { fields["name"]    = *update.Name    }
	if update.Company != nil && *update.Company != "" { fields["company"] = *update.Company }
	if update.Email   != nil && *update.Email   != "" { fields["email"]   = *update.Email   }
	if update.Phone   != nil && *update.Phone   != "" { fields["phone"]   = *update.Phone   }
	if update.Address != nil && *update.Address != "" { fields["address"] = *update.Address }

	if len(fields) == 0 {
        return nil
    }

	result := r.db.WithContext(ctx).
        Model(&model.SupplierStorage{}).
        Where("supplier_id = ?", update.SupplierID).
        Updates(fields)

    if result.Error != nil {
        return repository.HandleDBError(result.Error)
    }

    return nil
}