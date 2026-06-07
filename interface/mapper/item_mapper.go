package mapper

import (
	"Backend-Warehouse/domain/entity"
	"Backend-Warehouse/infrastructure/model"
	"Backend-Warehouse/interface/dto"

	"github.com/google/uuid"
)

// Mapper To Storage

func ToItemStorage(e *entity.ItemEntity) *model.ItemStorage {
	item := &model.ItemStorage{
		ItemID:       e.ItemID,
		ItemCode:     e.ItemCode,
		ItemName:     e.ItemName,
		CategoryID:   e.CategoryID,
		UnitOfStock:  e.UnitOfStock,
		Stock:        e.Stock,
		MinimumStock: e.MinimumStock,
	}

	if e.GoodsReceiptDetails != nil {
		item.GoodsReceiptDetails = toGoodsReceiptDetailStorages(e.GoodsReceiptDetails)
	}

	return item
}

func ToCategoryStorage(e *entity.CategoryEntity) *model.CategoryStorage {
	category := &model.CategoryStorage{
		CategoryID:   e.CategoryID,
		CategoryName: e.CategoryName,
	}

	if e.Items != nil {
		category.Items = toItemStorages(e.Items)
	}

	return category
}

// Mapper To Entity

func ToItemEntity(s *model.ItemStorage) *entity.ItemEntity {
	item := &entity.ItemEntity{
		ItemID:       s.ItemID,
		ItemCode:     s.ItemCode,
		ItemName:     s.ItemName,
		CategoryID:   s.CategoryID,
		UnitOfStock:  s.UnitOfStock,
		Stock:        s.Stock,
		MinimumStock: s.MinimumStock,
	}

	if s.GoodsReceiptDetails != nil {
		item.GoodsReceiptDetails = toGoodsReceiptDetailEntities(s.GoodsReceiptDetails)
	}

	return item
}

func ToCategoryEntity(s *model.CategoryStorage) *entity.CategoryEntity {
	category := &entity.CategoryEntity{
		CategoryID:   s.CategoryID,
		CategoryName: s.CategoryName,
	}

	if s.Items != nil {
		category.Items = toItemEntities(s.Items)
	}

	return category
}

func ToUpdateItemEntity(req dto.UpdateItemRequest) *entity.UpdateItemEntity {
	return &entity.UpdateItemEntity{
		ItemID:       req.ItemID,
		ItemCode:     req.ItemCode,
		ItemName:     req.ItemName,
		CategoryID:   req.CategoryID,
		UnitOfStock:  req.UnitOfStock,
		Stock:        req.Stock,
		MinimumStock: req.MinimumStock,
	}
}

func ToUpdateCategoryEntity(req dto.UpdateCategoryRequest) *entity.UpdateCategoryEntity {
	return &entity.UpdateCategoryEntity{
		CategoryID:   req.CategoryID,
		CategoryName: req.CategoryName,
	}
}

// Helpers Slice Converter

func toItemEntities(s []model.ItemStorage) []entity.ItemEntity {
    result := make([]entity.ItemEntity, 0, len(s))
    for _, item := range s {
        result = append(result, *ToItemEntity(&item))
    }
    return result
}

func toItemStorages(e []entity.ItemEntity) []model.ItemStorage {
    result := make([]model.ItemStorage, 0, len(e))
    for _, item := range e {
        result = append(result, *ToItemStorage(&item))
    }
    return result
}

func toItemEntitiesRequest(categoryID string, items []dto.AddItemsRequest) []entity.ItemEntity {
	result := make([]entity.ItemEntity, len(items))

    for i, item := range items {
        result[i] = entity.ItemEntity{
            ItemID:       uuid.NewString(),
            ItemCode:     item.ItemCode,
            ItemName:     item.ItemName,
            CategoryID:   categoryID,
            UnitOfStock:  item.UnitOfStock,
            Stock:        item.Stock,
            MinimumStock: item.MinimumStock,
        }
    }

    return result
}

func toItemsResponse(items []entity.ItemEntity) []dto.ItemsResponse {
	response := make([]dto.ItemsResponse, len(items))

	for i, item := range items {
        response[i] = dto.ItemsResponse{
            ItemID:              item.ItemID,
            ItemCode:            item.ItemCode,
            ItemName:            item.ItemName,
            CategoryID:          item.CategoryID,
            UnitOfStock:         item.UnitOfStock,
            Stock:               item.Stock,
            MinimumStock:        item.MinimumStock,
			GoodsReceiptDetails: toGoodsReceiptDetailResponses(item.GoodsReceiptDetails),
        }
    }

	return response
}

// Mapper for DTO

func ToEntityFromAddItemRequest(req dto.AddItemRequest) *entity.ItemEntity {
    return &entity.ItemEntity{
        ItemID:       uuid.NewString(),
        ItemCode:     req.ItemCode,
        ItemName:     req.ItemName,
        CategoryID:   req.CategoryID,
        UnitOfStock:  req.UnitOfStock,
        Stock:        req.Stock,
        MinimumStock: req.MinimumStock,
    }
}

func ToEntityFromAddCategoryRequest(req dto.AddCategoryRequest) *entity.CategoryEntity {
	categoryID := uuid.NewString()
    
    return &entity.CategoryEntity{
        CategoryID:   categoryID,
        CategoryName: req.CategoryName,
        Items:        toItemEntitiesRequest(categoryID, req.Items),
    }
}

func ToItemResponse(e *entity.ItemEntity) dto.ItemsResponse {
	return dto.ItemsResponse{
		ItemID:              e.ItemID,
		ItemCode:            e.ItemCode,
		ItemName:            e.ItemName,
		CategoryID:          e.CategoryID,
		UnitOfStock:         e.UnitOfStock,
		Stock:               e.Stock,
		MinimumStock:        e.MinimumStock,
		GoodsReceiptDetails: toGoodsReceiptDetailResponses(e.GoodsReceiptDetails),
	}
}

func ToItemsResponse(e []entity.ItemEntity) []dto.ItemsResponse {
	response := make([]dto.ItemsResponse, 0, len(e))

	for _, item := range e {
		response = append(response, ToItemResponse(&item))
	}

	return response
}

func ToCategoryResponse(e *entity.CategoryEntity) dto.CategoriesResponse {
	return dto.CategoriesResponse{
		CategoryID:   e.CategoryID,
		CategoryName: e.CategoryName,
		Items:        toItemsResponse(e.Items),
	}
} 


func ToCategoriesResponse(e []entity.CategoryEntity) []dto.CategoriesResponse {
	response := make([]dto.CategoriesResponse, 0, len(e))

	for _, category := range e {
		response = append(response, ToCategoryResponse(&category))
	}

	return response
}
