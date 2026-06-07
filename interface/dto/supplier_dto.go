package dto

type RegisterSupplierRequest struct {
	Name    string `json:"name"    validate:"required,alphaspaceunicode"`
	Company string `json:"company" validate:"required"`
	Email   string `json:"email"   validate:"required,email"`
	Phone   string `json:"phone"   validate:"required,phone"`
	Address string `json:"address" validate:"required"`
}

type UpdateSupplierRequest struct {
	SupplierID string  `json:"-" uri:"supplierID"`
	Name       *string `json:"name"    validate:"omitempty,alphaspaceunicode"`
	Company    *string `json:"company" validate:"omitempty"`
	Email      *string `json:"email"   validate:"omitempty,email"`
	Phone      *string `json:"phone"   validate:"omitempty,phone"`
	Address    *string `json:"address" validate:"omitempty"`
}

type GetSupplierRequest struct {
	SupplierID string `json:"-" uri:"supplierID"`
}

type SuppliersResponse struct {
	SupplierID string `json:"supplier_id"`
	Name       string `json:"name"`
	Company    string `json:"company"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`

	GoodsReceipts []GoodsReceiptResponse `json:"goods_receipt"`
}