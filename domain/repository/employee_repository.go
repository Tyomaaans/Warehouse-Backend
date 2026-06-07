package repository

import (
	"Backend-Warehouse/domain/entity"
	"context"
)

type EmployeeRepository interface {
	CreateEmployee(ctx context.Context, employee *entity.EmployeeEntity) error
	GetEmployeeByEmployeeID(ctx context.Context, empID string) (*entity.EmployeeEntity, error)
	GetEmployeeByEmployeeUsername(ctx context.Context, empUsername string) (*entity.EmployeeEntity, error)
	GetEmployeeByEmployeeEmail(ctx context.Context, empEmail string) (*entity.EmployeeEntity, error)
	GetEmployees(ctx context.Context) ([]entity.EmployeeEntity, error)
	UpdateEmployeeForUser(ctx context.Context, update *entity.UpdateEmployeeEntityForUser) error
	UpdateEmployeeForAdmin(ctx context.Context, update *entity.UpdateEmployeeEntityForAdmin) error
}