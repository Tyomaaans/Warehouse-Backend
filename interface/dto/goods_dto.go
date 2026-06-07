package dto

import "time"

type AddGoodsReceiptRequest struct {
    SupplierID  string    `json:"supplier_id"   validate:"required,uuid4"`
    EmployeeID  string    `json:"-"`
    DateOfEntry time.Time `json:"date_of_entry" validate:"required"`
    
    GoodsReceiptDetails []AddGoodsReceiptDetailsRequest `json:"goods_receipt_details" validate:"required,dive"`
}

type AddGoodsReceiptDetailsRequest struct {
    ItemID string `json:"item_id" validate:"required,uuid4"`
    Qty    int64  `json:"qty"     validate:"required,gte=0"`
}

type AddGoodsIssueRequest struct {
	EmployeeID      string    `json:"-"`
	ExitDestination string    `json:"exit_destination" validate:"required"`
	DateOfExit      time.Time `json:"date_of_exit"  validate:"required"`

	GoodsIssueDetails []AddGoodsIssueDetailRequest `json:"goods_issue_details" validate:"required,dive"`
}

type AddGoodsIssueDetailRequest struct {
	ItemID string `json:"item_id" validate:"required,uuid4"`
    Qty    int64  `json:"qty"     validate:"required,gte=0"`
}

type GetOrApproveGoodsReceiptOrIssueRequest struct {
	GoodsReceiptID string `json:"-" uri:"goodsReceiptID"`
	GoodsIssueID   string `json:"-" uri:"goodsIssueID"`
}

// DTO Goods Response

type GoodsReceiptResponse struct {
	GoodsReceiptID string    `json:"goods_receipt_id"`
	Invoice        string    `json:"invoice"`
	SupplierID     string    `json:"supplier_id"`
	DateOfEntry    time.Time `json:"date_of_entry"`
	EmployeeID     string    `json:"employee_id"`
	Status         string    `json:"status"`

	GoodsReceiptDetails []GoodsReceiptDetailResponse `json:"goods_receipt_details"`
}

type GoodsReceiptDetailResponse struct {
	GoodsReceiptDetailID string `json:"goods_receipt_detail_id"`
	GoodsReceiptID       string `json:"goods_receipt_id"`
	ItemID               string `json:"item_id"`
	Qty                  int64  `json:"qty"`
}

type GoodsIssueResponse struct {
	GoodsIssueID    string    `json:"goods_issue_id"`
	Invoice         string    `json:"invoice"`
	ExitDestination string    `json:"exit_destination"`
	DateOfExit      time.Time `json:"date_of_exit"`
	EmployeeID      string    `json:"employee_id"`
	Status          string    `json:"status"`

	GoodsIssueDetails []GoodsIssueDetailResponse `json:"goods_issue_details"`
}

type GoodsIssueDetailResponse struct {
	GoodsIssueDetailID string `json:"goods_issue_detail_id"`
	GoodsIssueID       string `json:"goods_issue_id"`
	ItemID             string `json:"item_id"`
	Qty                int64  `json:"qty"`
}