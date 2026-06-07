package entity

type ItemEntity struct {
	ItemID       string         
	ItemCode     string         
	ItemName     string         
	CategoryID   string         
	UnitOfStock  string         
	Stock        int64          
	MinimumStock int64          

	GoodsReceiptDetails []GoodsReceiptDetailEntity
}

type CategoryEntity struct {
	CategoryID   string         
	CategoryName string         

	Items []ItemEntity
}

type UpdateItemEntity struct {
	ItemID       string
	ItemCode     *string
	ItemName     *string         
	CategoryID   *string         
	UnitOfStock  *string         
	Stock        *int64          
	MinimumStock *int64 
}

type UpdateCategoryEntity struct {
	CategoryID   string
	CategoryName *string
}