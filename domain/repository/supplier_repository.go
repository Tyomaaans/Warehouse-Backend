package repository

import (
	"Backend-Warehouse/domain/entity"
	"context"
)

type SupplierRepository interface {
	CreateSupplier(ctx context.Context, supplier *entity.SupplierEntity) error
	GetSuppliers(ctx context.Context) ([]entity.SupplierEntity, error)
	GetSupplierBySupplierID(ctx context.Context, supplierID string) (*entity.SupplierEntity, error)
	UpdateSupplier(ctx context.Context, update *entity.UpdateSupplierEntity) error
}