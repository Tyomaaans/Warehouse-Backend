package dto

// DTO Request

type AddItemRequest struct {
	ItemCode     string `json:"item_code"     validate:"required"`
	ItemName     string `json:"item_name"     validate:"required"`
	CategoryID   string `json:"category_id"   validate:"required"`
	UnitOfStock  string `json:"unit_of_stock" validate:"required,oneof=Batang Lembar Pcs Roll"`
	Stock        int64  `json:"stock"         validate:"required,gte=0"`
	MinimumStock int64  `json:"minimum_stock" validate:"required,gte=0"`
}

type AddItemsRequest struct {
	ItemCode     string `json:"item_code"     validate:"required"`
	ItemName     string `json:"item_name"     validate:"required"`
	UnitOfStock  string `json:"unit_of_stock" validate:"required,oneof=Batang Lembar Pcs Roll"`
	Stock        int64  `json:"stock"         validate:"required,gte=0"`
	MinimumStock int64  `json:"minimum_stock" validate:"required,gte=0"`
}

type AddCategoryRequest struct {
	CategoryName string `json:"category_name" validate:"required,alphaspace"`

	Items []AddItemsRequest `json:"items" validate:"required,dive"`
}

type UpdateItemRequest struct {
	ItemID       string  `json:"-"             uri:"itemID"`
	ItemCode     *string `json:"item_code"     validate:"omitempty"`
	ItemName     *string `json:"item_name"     validate:"omitempty,alphaspace"`
	CategoryID   *string `json:"category_id"   validate:"omitempty"`
	UnitOfStock  *string `json:"unit_of_stock" validate:"omitempty,oneof=Batang Lembar Pcs Roll"`
	Stock        *int64  `json:"stock"         validate:"omitempty,gte=0"`
	MinimumStock *int64  `json:"minimum_stock" validate:"omitempty,gte=0"`
}

type UpdateCategoryRequest struct {
	CategoryID   string  `json:"-" uri:"categoryID"`
	CategoryName *string `json:"category_name" validate:"omitempty,alphaspace"`
}

type GetItemOrCategoryByID struct {
	ItemID     string `json:"-" uri:"itemID"`
	CategoryID string `json:"-" uri:"categoryID"`
}

// DTO Response

type ItemsResponse struct {
	ItemID       string `json:"item_id"`
	ItemCode     string `json:"item_code"`
	ItemName     string `json:"item_name"`
	CategoryID   string `json:"category_id"`
	UnitOfStock  string `json:"unit_of_stock"`
	Stock        int64  `json:"stock"`
	MinimumStock int64  `json:"minimum_stock"`

	GoodsReceiptDetails []GoodsReceiptDetailResponse `json:"goods_receipt_details"`
}

type CategoriesResponse struct {
	CategoryID   string `json:"categor_id"`
	CategoryName string `json:"category_name"`

	Items []ItemsResponse `json:"items"`
}