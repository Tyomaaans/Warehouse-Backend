package repository

import (
	"Backend-Warehouse/domain/entity"
	"context"
)

type GoodsRepository interface {
	CreateGoodsReceipt(ctx context.Context, receipt *entity.GoodsReceiptEntity) error
	GetGoodsReceipts(ctx context.Context) ([]entity.GoodsReceiptEntity, error)
	GetGoodsReceiptByGoodsReceiptID(ctx context.Context, goodsReceiptID string) (*entity.GoodsReceiptEntity, error)
	ApproveGoodsReceipt(ctx context.Context, receipt *entity.GoodsReceiptEntity) error

	CreateGoodsIssue(ctx context.Context, issue *entity.GoodsIssueEntity) error
	GetGoodsIssues(ctx context.Context) ([]entity.GoodsIssueEntity, error)
	GetGoodsIssueByGoodsIssueID(ctx context.Context, goodsIssueID string) (*entity.GoodsIssueEntity, error)
	ApproveGoodsIssue(ctx context.Context, issue *entity.GoodsIssueEntity) error
}