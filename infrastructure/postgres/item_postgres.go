package postgres

import (
	"context"

	"gorm.io/gorm"

	"Backend-Warehouse/domain/entity"
	"Backend-Warehouse/domain/repository"
	"Backend-Warehouse/infrastructure/model"
	"Backend-Warehouse/interface/mapper"
)

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{
		db: db,
	}
}

func (r *ItemRepository) CreateItem(ctx context.Context, item *entity.ItemEntity) error {
	itemStorage := mapper.ToItemStorage(item)

	if err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(itemStorage).Error; err != nil {
			return repository.HandleDBError(err)
		}

		return nil
	}); err != nil { return err }

	return nil
}

func (r *ItemRepository) CreateCategory(ctx context.Context, category *entity.CategoryEntity) error {
	categoryStorage := mapper.ToCategoryStorage(category)

	if err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(categoryStorage).Error; err != nil {
			return repository.HandleDBError(err)
		}

		return nil
	}); err != nil { return err }

	return nil
}

func (r *ItemRepository) GetItems(ctx context.Context) ([]entity.ItemEntity, error) {
	var itemStorages []model.ItemStorage

	err := r.db.WithContext(ctx).
		Preload("GoodsReceiptDetails").
		Order("created_at ASC").
		Find(&itemStorages).Error

	if err != nil {
        return nil, repository.HandleDBError(err)
    }

	var items []entity.ItemEntity
    for _, itemStorage := range itemStorages {
        item := mapper.ToItemEntity(&itemStorage)
        items = append(items, *item)
    }

    return items, nil
}

func (r *ItemRepository) GetItemByItemID(ctx context.Context, itemID string) (*entity.ItemEntity, error) {
	var itemStorage model.ItemStorage

	err := r.db.WithContext(ctx).
		Preload("GoodsReceiptDetails").
		Where("item_id = ? ", itemID).
		Order("created_at ASC").
		Find(&itemStorage).Error

	if err != nil {
		return nil, repository.HandleDBError(err)
	}

	item := mapper.ToItemEntity(&itemStorage)

	return item, nil
}

func (r *ItemRepository) UpdateItem(ctx context.Context, update *entity.UpdateItemEntity) error {
	fields := map[string]any{}

	if update.ItemCode     != nil && *update.ItemCode    != "" { fields["item_code"]     = *update.ItemCode    }
	if update.ItemName     != nil && *update.ItemName    != "" { fields["item_name"]     = *update.ItemName    }
	if update.CategoryID   != nil && *update.CategoryID  != "" { fields["category_id"]   = *update.CategoryID  }
	if update.UnitOfStock  != nil && *update.UnitOfStock != "" { fields["unit_of_stock"] = *update.UnitOfStock }
	if update.Stock        != nil { 
		fields["stock"] = *update.Stock 
	}
	if update.MinimumStock != nil { 
		fields["minimum_stock"] = *update.MinimumStock 
	}

	if len(fields) == 0 {
        return nil
    }

	result := r.db.WithContext(ctx).
        Model(&model.ItemStorage{}).
        Where("item_id = ?", update.ItemID).
        Updates(fields)

    if result.Error != nil {
        return repository.HandleDBError(result.Error)
    }

    return nil
}

func (r *ItemRepository) GetCategories(ctx context.Context) ([]entity.CategoryEntity, error) {
	var categoryStorages []model.CategoryStorage

	err := r.db.WithContext(ctx).
		Preload("Items").Preload("Items.GoodsReceiptDetails").
		Order("created_at ASC").
		Find(&categoryStorages).Error

	if err != nil {
        return nil, repository.HandleDBError(err)
    }

	var categories []entity.CategoryEntity
    for _, categoryStorage := range categoryStorages {
        category := mapper.ToCategoryEntity(&categoryStorage)
        categories = append(categories, *category)
    }

    return categories, nil
}

func (r *ItemRepository) GetCategoryByCategoryID(ctx context.Context, categoryID string) (*entity.CategoryEntity, error) {
	var categoryStorage model.CategoryStorage

	err := r.db.WithContext(ctx).
		Preload("Items").Preload("Items.GoodsReceiptDetails").
		Where("category_id = ? ", categoryID).
		Order("created_at ASC").
		Find(&categoryStorage).Error

	if err != nil {
		return nil, repository.HandleDBError(err)
	}

	category := mapper.ToCategoryEntity(&categoryStorage)

	return category, nil
}

func (r *ItemRepository) UpdateCategory(ctx context.Context, update *entity.UpdateCategoryEntity) error {
	fields := map[string]any{}

	if update.CategoryName != nil && *update.CategoryName != "" { fields["category_name"] = *update.CategoryName }

	if len(fields) == 0 {
        return nil
    }

	result := r.db.WithContext(ctx).
        Model(&model.CategoryStorage{}).
        Where("category_id = ?", update.CategoryID).
        Updates(fields)

    if result.Error != nil {
        return repository.HandleDBError(result.Error)
    }

    return nil
}