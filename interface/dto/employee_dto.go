package dto

// DTO Employee Request

type RegisterEmployeeRequest struct {
	Name            string `json:"name"             validate:"required,alphaspaceunicode"`
	Email           string `json:"email"            validate:"required,email"`
	Username        string `json:"username"         validate:"required,username,min=4"`
	Phone           string `json:"phone"            validate:"required,phone"`
	Address         string `json:"address"          validate:"required"`
	Role            string `json:"role"             validate:"required,oneof='Warehouse Admin' 'Warehouse Staff' 'Warehouse Supervisor' Purchasing Management"`
	Password        string `json:"password"         validate:"required,password,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

type LoginEmployeeRequest struct {
	Username string `json:"username" validate:"omitempty"`
	Email    string `json:"email"    validate:"omitempty,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateEmployeeForUserRequest struct {
	EmployeeID string  `json:"-" uri:"employeeID"`
	Username   *string `json:"username" validate:"omitempty,username,min=4"`
	Phone      *string `json:"phone"    validate:"omitempty,phone"`
	Address    *string `json:"address"  validate:"omitempty"`
	Avatar     *string `json:"avatar"   validate:"omitempty,url"`
}

type UpdateEmployeeForAdminRequest struct {
	EmployeeID      string  `json:"-" uri:"employeeID"`
	Name            *string `json:"name"             validate:"omitempty,alphaspaceunicode"`
	Email           *string `json:"email"            validate:"omitempty,email"`
	Username        *string `json:"username"         validate:"omitempty,username,min=4"`
	Phone           *string `json:"phone"            validate:"omitempty,phone"`
	Address         *string `json:"address"          validate:"omitempty"`
	Role            *string `json:"role"             validate:"omitempty,oneof='Warehouse Admin' 'Warehouse Staff' 'Warehouse Supervisor' Purchasing Management"`
	Avatar          *string `json:"avatar"           validate:"omitempty,url"`
	Password        *string `json:"password"         validate:"omitempty,password,min=8"`
	ConfirmPassword *string `json:"confirm_password" validate:"omitempty"`
}

// DTO Employee Response

type LoginEmployeeResponse struct {
	AccessToken    string           `json:"access_token"`
	EmployeProfile EmployeeResponse `json:"employee_profile"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type EmployeeResponse struct {
	EmployeeID string `json:"employee_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	Role       string `json:"role"`
	Avatar     string `json:"avatar"`

	GoodsReceipts []GoodsReceiptResponse `json:"goods_receipt"`
	GoodsIssues   []GoodsIssueResponse   `json:"goods_issue"`
}

type UploadAvatarResponse struct {
	URL string `json:"url"`
}