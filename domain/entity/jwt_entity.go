package entity

import (
	"time"
)

type Claims struct {
	EmployeeID string
	Username   string
	Role       string
	ExpiresAt  time.Time
	IssuedAt   time.Time
}

type TokenPair struct {
	EmployeeID   string
	Username     string
	AccessToken  string
	RefreshToken string
}

const (
	RoleWHAdmin    = "Warehouse Admin"
	RoleWHSpv      = "Warehouse Supervisor"
	RoleWHStaff    = "Warehouse Staff"
	RolePurchasing = "Purchasing"
	RoleMangement  = "Management"
)