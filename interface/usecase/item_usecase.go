package usecase

import (
	"context"

	"github.com/go-playground/validator/v10"

	"Backend-Warehouse/domain/repository"
	"Backend-Warehouse/interface/dto"
	"Backend-Warehouse/interface/mapper"
	jsonValidator "Backend-Warehouse/validator"
)

type ItemUseCase struct {
	itemRepo repository.ItemRepository
	validate *validator.Validate
}

func NewItemUseCase(
	itemRepo repository.ItemRepository,
	validate *validator.Validate,
) *ItemUseCase {
	return &ItemUseCase{
		itemRepo: itemRepo,
		validate: validate,
	}
}

func (uc *ItemUseCase) AddItem(ctx context.Context, req dto.AddItemRequest) error {
	if err := jsonValidator.ValidateStruct(uc.validate, req); err != nil {
		return err
	}

	item := mapper.ToEntityFromAddItemRequest(req)
	if err := uc.itemRepo.CreateItem(ctx, item); err != nil {
		return err
	}

	return nil
}

func (uc *ItemUseCase) AddCategory(ctx context.Context, req dto.AddCategoryRequest) error {
	if err := jsonValidator.ValidateStruct(uc.validate, req); err != nil {
		return err
	}

	category := mapper.ToEntityFromAddCategoryRequest(req)
	if err := uc.itemRepo.CreateCategory(ctx, category); err != nil {
		return err
	}

	return nil
}

func (uc *ItemUseCase) GetItems(ctx context.Context) ([]dto.ItemsResponse, error) {
	items, err := uc.itemRepo.GetItems(ctx)
    if err != nil {
        return nil, err
    }

    res := mapper.ToItemsResponse(items)

    return res, nil
}

func (uc *ItemUseCase) GetItem(ctx context.Context, req dto.GetItemOrCategoryByID) (*dto.ItemsResponse, error) {
	items, err := uc.itemRepo.GetItemByItemID(ctx, req.ItemID)
	if err != nil {
		return nil, err
	}

	res := mapper.ToItemResponse(items)

	return &res, nil
}

func (uc *ItemUseCase) GetCategories(ctx context.Context) ([]dto.CategoriesResponse, error) {
	categories, err := uc.itemRepo.GetCategories(ctx)
	if err != nil {
		return nil, err
	}

	res := mapper.ToCategoriesResponse(categories)

	return res, nil
}

func (uc *ItemUseCase) GetCategory(ctx context.Context, req dto.GetItemOrCategoryByID) (*dto.CategoriesResponse, error) {
	categories, err := uc.itemRepo.GetCategoryByCategoryID(ctx, req.CategoryID)
	if err != nil {
		return nil, err
	}

	res := mapper.ToCategoryResponse(categories)

	return &res, nil
}

func (uc *ItemUseCase) UpdateItem(ctx context.Context, req dto.UpdateItemRequest) error {
	if err := jsonValidator.ValidateStruct(uc.validate, req); err != nil {
		return err
	}

	update := mapper.ToUpdateItemEntity(req)
	if err := uc.itemRepo.UpdateItem(ctx, update); err != nil {
		return err
	}

	return nil
}

func (uc *ItemUseCase) UpdateCategory(ctx context.Context, req dto.UpdateCategoryRequest) error {
	if err := jsonValidator.ValidateStruct(uc.validate, req); err != nil {
		return err
	}

	update := mapper.ToUpdateCategoryEntity(req)
	if err := uc.itemRepo.UpdateCategory(ctx, update); err != nil {
		return err
	}

	return nil
}