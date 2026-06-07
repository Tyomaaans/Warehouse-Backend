package entity

type EmployeeEntity struct {
	EmployeeID string
	Name       string
	Email      string
	Username   string
	Phone      string
	Address    string
	Role       string
	Avatar     string
	Password   string

	GoodsReceipts []GoodsReceiptEntity
	GoodsIssues   []GoodsIssueEntity
	Signatures    []SignatureEntity
}

type UpdateEmployeeEntityForUser struct {
	EmployeeID  string
	Username   *string
	Phone      *string
	Address    *string
	Avatar     *string
}

type UpdateEmployeeEntityForAdmin struct {
	EmployeeID string
	Name       *string
	Email      *string
	Username   *string
	Phone      *string
	Address    *string
	Role       *string
	Avatar     *string
	Password   *string
}