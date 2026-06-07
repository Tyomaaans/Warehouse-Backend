package usecase

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"

	"Backend-Warehouse/domain/repository"
	"Backend-Warehouse/interface/dto"
	"Backend-Warehouse/interface/mapper"
	jsonValidator "Backend-Warehouse/validator"
)

type GoodsUseCase struct {
	goodsRepo repository.GoodsRepository
	validate *validator.Validate
}

func NewGoodsUseCase(
	goodsRepo repository.GoodsRepository,
	validate *validator.Validate,
) *GoodsUseCase {
	return &GoodsUseCase{
		goodsRepo: goodsRepo,
		validate:  validate,
	}
}

func (uc *GoodsUseCase) AddGoodsReceipt(ctx context.Context, req dto.AddGoodsReceiptRequest) error {
	if err := jsonValidator.ValidateStruct(uc.validate, req); err != nil {
		return err
	}

	goodsReceipt := mapper.ToEntityFromAddGoodsReceiptRequest(req)
	if err := uc.goodsRepo.CreateGoodsReceipt(ctx, goodsReceipt); err != nil {
		return err
	}

	return nil
}

func (uc *GoodsUseCase) GetGoodsReceipts(ctx context.Context) ([]dto.GoodsReceiptResponse, error) {
	goodsReceipt, err := uc.goodsRepo.GetGoodsReceipts(ctx)
	if err != nil {
		return nil, err
	}

	res := mapper.ToGoodsReceiptResponseList(goodsReceipt)

	return res, nil
}

func (uc *GoodsUseCase) GetGoodsReceipt(ctx context.Context, req dto.GetOrApproveGoodsReceiptOrIssueRequest) (*dto.GoodsReceiptResponse, error) {
	goodsReceipt, err := uc.goodsRepo.GetGoodsReceiptByGoodsReceiptID(ctx, req.GoodsReceiptID)
	if err != nil {
		return nil, err
	}

	res := mapper.ToGoodsReceiptResponse(goodsReceipt)

	return &res, nil
}

func (uc *GoodsUseCase) ApproveGoodsReceipt(ctx context.Context, req dto.GetOrApproveGoodsReceiptOrIssueRequest) error {
    goodsReceipt, err := uc.goodsRepo.GetGoodsReceiptByGoodsReceiptID(ctx, req.GoodsReceiptID)
    if err != nil {
        return err
    }

    if goodsReceipt.Status != "pending" {
        return errors.New("goods receipt status is not pending!")
    }

    return uc.goodsRepo.ApproveGoodsReceipt(ctx, goodsReceipt)
}

func (uc *GoodsUseCase) AddGoodsIssue(ctx context.Context, req dto.AddGoodsIssueRequest) error {
	if err := jsonValidator.ValidateStruct(uc.validate, req); err != nil {
		return err
	}

	goodsIssue := mapper.ToEntityFromAddGoodsIssueRequest(req)
	if err := uc.goodsRepo.CreateGoodsIssue(ctx, goodsIssue); err != nil {
		return err
	}

	return nil
}

func (uc *GoodsUseCase) GetGoodsIssues(ctx context.Context) ([]dto.GoodsIssueResponse, error) {
	goodsIssue, err := uc.goodsRepo.GetGoodsIssues(ctx)
	if err != nil {
		return nil, err
	}

	res := mapper.ToGoodsIssueResponseList(goodsIssue)

	return res, nil
}

func (uc *GoodsUseCase) GetGoodsIssue(ctx context.Context, req dto.GetOrApproveGoodsReceiptOrIssueRequest) (*dto.GoodsIssueResponse, error) {
	goodsIssue, err := uc.goodsRepo.GetGoodsIssueByGoodsIssueID(ctx, req.GoodsIssueID)
	if err != nil {
		return nil, err
	}

	res := mapper.ToGoodsIssueResponse(goodsIssue)

	return &res, nil
}

func (uc *GoodsUseCase) ApproveGoodsIssue(ctx context.Context, req dto.GetOrApproveGoodsReceiptOrIssueRequest) error {
    goodsIssue, err := uc.goodsRepo.GetGoodsIssueByGoodsIssueID(ctx, req.GoodsIssueID)
    if err != nil {
        return err
    }

    if goodsIssue.Status != "pending" {
        return errors.New("goods issue status is not pending!")
    }

    return uc.goodsRepo.ApproveGoodsIssue(ctx, goodsIssue)
}