package entity

type SupplierEntity struct {
	SupplierID string
	Name       string
	Company    string
	Email      string
	Phone      string
	Address    string

	GoodsReceipts []GoodsReceiptEntity
	Signatures    []SignatureEntity
}

type UpdateSupplierEntity struct {
	SupplierID string
	Name       *string
	Company    *string
	Email      *string
	Phone      *string
	Address    *string
}