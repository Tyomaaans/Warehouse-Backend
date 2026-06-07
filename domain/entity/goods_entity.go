package entity

import "time"

// Goods Receipts

type GoodsReceiptEntity struct {
	GoodsReceiptID string        
	Invoice        string         
	SupplierID     string         
	DateOfEntry    time.Time      
	EmployeeID     string         
	Status         string         

	GoodsReceiptDetails []GoodsReceiptDetailEntity
}

type GoodsReceiptDetailEntity struct {
	GoodsReceiptDetailID string       
	GoodsReceiptID       string         
	ItemID               string         
	Qty                  int64          
}

// Goods Issues

type GoodsIssueEntity struct {
	GoodsIssueID    string        
	Invoice         string        
	ExitDestination string         
	DateOfExit      time.Time      
	EmployeeID      string         
	Status          string         

	GoodsIssueDetails []GoodsIssueDetailEntity 
}

type GoodsIssueDetailEntity struct {
	GoodsIssueDetailID string         
	GoodsIssueID       string         
	ItemID             string         
	Qty                int64          
}
