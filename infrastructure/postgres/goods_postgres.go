package postgres

import (
	"Backend-Warehouse/domain/entity"
	"Backend-Warehouse/domain/repository"
	"Backend-Warehouse/infrastructure/model"
	"Backend-Warehouse/interface/mapper"
	"context"
	"errors"

	"gorm.io/gorm"
)

type GoodsRepository struct {
	db *gorm.DB
}

func NewGoodsRepository(db *gorm.DB) *GoodsRepository {
	return &GoodsRepository{
		db: db,
	}
}

func (r *GoodsRepository) CreateGoodsReceipt(ctx context.Context, receipt *entity.GoodsReceiptEntity) error {
	goodsReceiptStorage := mapper.ToGoodsReceiptStorage(receipt)

	if err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
        if err := tx.Create(goodsReceiptStorage).Error; err != nil {
            return repository.HandleDBError(err)
        }

        return nil

    }); err != nil { return err }

    return nil
}

func (r *GoodsRepository) GetGoodsReceipts(ctx context.Context) ([]entity.GoodsReceiptEntity, error) {
	var goodsReceiptStorages []model.GoodsReceiptStorage

	err := r.db.WithContext(ctx).
		Preload("GoodsReceiptDetails").
		Order("created_at ASC").
		Find(&goodsReceiptStorages).Error

	if err != nil {
        return nil, repository.HandleDBError(err)
    }

	var goodsReceipts []entity.GoodsReceiptEntity
    for _, goodsReceiptStorage := range goodsReceiptStorages {
        item := mapper.ToGoodsReceiptEntity(&goodsReceiptStorage)
        goodsReceipts = append(goodsReceipts, *item)
    }

    return goodsReceipts, nil
}

func (r *GoodsRepository) GetGoodsReceiptByGoodsReceiptID(ctx context.Context, goodsReceiptID string) (*entity.GoodsReceiptEntity, error) {
	var goodsReceiptStorage model.GoodsReceiptStorage

	err := r.db.WithContext(ctx).
		Preload("GoodsReceiptDetails").
		Where("goods_receipt_id = ?", goodsReceiptID).
		First(&goodsReceiptStorage).Error

	if err != nil {
		return nil, repository.HandleDBError(err)
	}

	goodsReceipt := mapper.ToGoodsReceiptEntity(&goodsReceiptStorage)

	return goodsReceipt, nil
}

func (r *GoodsRepository) ApproveGoodsReceipt(ctx context.Context, receipt *entity.GoodsReceiptEntity) error {
	goodsReceiptStorage := mapper.ToGoodsReceiptStorage(receipt)
	
	if len(goodsReceiptStorage.GoodsReceiptDetails) == 0 {
		return errors.New("details cannot be empty!")
	}

    return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
        for _, detail := range goodsReceiptStorage.GoodsReceiptDetails {
            result := tx.Model(&model.ItemStorage{}).
                Where("item_id = ?", detail.ItemID).
                Update("stock", gorm.Expr("stock + ?", detail.Qty))
				
			if result.Error != nil {
				return repository.HandleDBError(result.Error)
			}

			if result.RowsAffected == 0 {
				return errors.New("goods receipt already approved or not found!")
			}
        }

        result := tx.Model(&model.GoodsReceiptStorage{}).
            Where("goods_receipt_id = ? AND status = ?", goodsReceiptStorage.GoodsReceiptDetails[0].GoodsReceiptID, "pending").
            Update("status", "approved")
		
		if result.Error != nil {
			return repository.HandleDBError(result.Error)
		}

		if result.RowsAffected == 0 {
			return errors.New("goods receipt already approved or not found!")
		}

        return nil
    })
}

func (r *GoodsRepository) CreateGoodsIssue(ctx context.Context, issue *entity.GoodsIssueEntity) error {
	goodsIssueStorage := mapper.ToGoodsIssueStorage(issue)

	if err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
        if err := tx.Create(goodsIssueStorage).Error; err != nil {
            return repository.HandleDBError(err)
        }

        return nil

    }); err != nil { return err }

    return nil
}

func (r *GoodsRepository) GetGoodsIssues(ctx context.Context) ([]entity.GoodsIssueEntity, error) {
	var goodsIssueStorages []model.GoodsIssueStorage

	err := r.db.WithContext(ctx).
		Preload("GoodsIssueDetails").
		Order("created_at ASC").
		Find(&goodsIssueStorages).Error

	if err != nil {
        return nil, repository.HandleDBError(err)
    }

	var goodsIssues []entity.GoodsIssueEntity
    for _, goodsIssueStorage := range goodsIssueStorages {
        item := mapper.ToGoodsIssueEntity(&goodsIssueStorage)
        goodsIssues = append(goodsIssues, *item)
    }

    return goodsIssues, nil
}

func (r *GoodsRepository) GetGoodsIssueByGoodsIssueID(ctx context.Context, goodsIssueID string) (*entity.GoodsIssueEntity, error) {
	var goodsIssueStorage model.GoodsIssueStorage

	err := r.db.WithContext(ctx).
		Preload("GoodsIssueDetails").
		Where("goods_issue_id = ?", goodsIssueID).
		First(&goodsIssueStorage).Error

	if err != nil {
		return nil, repository.HandleDBError(err)
	}

	goodsIssue := mapper.ToGoodsIssueEntity(&goodsIssueStorage)

	return goodsIssue, nil
}

func (r *GoodsRepository) ApproveGoodsIssue(ctx context.Context, issue *entity.GoodsIssueEntity) error {
	goodsIssueStorage := mapper.ToGoodsIssueStorage(issue)
	
	if len(goodsIssueStorage.GoodsIssueDetails) == 0 {
		return errors.New("details cannot be empty!")
	}

    return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
        for _, detail := range goodsIssueStorage.GoodsIssueDetails {
            result := tx.Model(&model.ItemStorage{}).
                Where("item_id = ?", detail.ItemID).
                Update("stock", gorm.Expr("stock - ?", detail.Qty))
				
			if result.Error != nil {
				return repository.HandleDBError(result.Error)
			}

			if result.RowsAffected == 0 {
				return errors.New("goods issue already approved or not found!")
			}
        }

        result := tx.Model(&model.GoodsIssueStorage{}).
            Where("goods_issue_id = ? AND status = ?", goodsIssueStorage.GoodsIssueDetails[0].GoodsIssueID, "pending").
            Update("status", "approved")
		
		if result.Error != nil {
			return repository.HandleDBError(result.Error)
		}

		if result.RowsAffected == 0 {
			return errors.New("goods issue already approved or not found!")
		}

        return nil
    })
}