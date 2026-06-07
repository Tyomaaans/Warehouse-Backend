package repository

import (
	"context"

	"Backend-Warehouse/domain/entity"
)

type ItemRepository interface {
	CreateItem(ctx context.Context, item *entity.ItemEntity) error
	GetItems(ctx context.Context) ([]entity.ItemEntity, error)
	GetItemByItemID(ctx context.Context, itemID string) (*entity.ItemEntity, error)
	UpdateItem(ctx context.Context, update *entity.UpdateItemEntity) error

	CreateCategory(ctx context.Context, category *entity.CategoryEntity) error
	GetCategories(ctx context.Context) ([]entity.CategoryEntity, error)
	GetCategoryByCategoryID(ctx context.Context, categoryID string) (*entity.CategoryEntity, error)
	UpdateCategory(ctx context.Context, update *entity.UpdateCategoryEntity) error
}