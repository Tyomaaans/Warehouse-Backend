package postgres

import (
	"context"

	"gorm.io/gorm"

	"Backend-Warehouse/domain/entity"
	"Backend-Warehouse/domain/repository"
	"Backend-Warehouse/infrastructure/model"
	"Backend-Warehouse/interface/mapper"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{
		db: db,
	}
}

func (r *EmployeeRepository) CreateEmployee(ctx context.Context, employee *entity.EmployeeEntity) error {
	employeeStorage := mapper.ToEmployeeStorage(employee)

	if err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
        if err := tx.Create(employeeStorage).Error; err != nil {
            return repository.HandleDBError(err)
        }

        return nil

    }); err != nil { return err }

    return nil
}

func (r *EmployeeRepository) GetEmployeeByEmployeeID(ctx context.Context, empID string) (*entity.EmployeeEntity, error) {
	var empStorage model.EmployeeStorage

	err := r.db.WithContext(ctx).
		Preload("GoodsReceipts").Preload("GoodsIssues").Preload("GoodsReceipts.GoodsReceiptDetails").Preload("GoodsIssues.GoodsIssueDetails").
		Where("employee_id = ?", empID).
		First(&empStorage).Error

	if err != nil {
		return nil, repository.HandleDBError(err)
	}

	emp := mapper.ToEmployeeEntity(&empStorage)

	return emp, nil
}

func (r *EmployeeRepository) GetEmployeeByEmployeeUsername(ctx context.Context, empUsername string) (*entity.EmployeeEntity, error) {
	var empStorage model.EmployeeStorage

	err := r.db.WithContext(ctx).
		Preload("GoodsReceipts").Preload("GoodsIssues").Preload("GoodsReceipts.GoodsReceiptDetails").Preload("GoodsIssues.GoodsIssueDetails").
		Where("username = ?", empUsername).
		First(&empStorage).Error

	if err != nil {
		return nil, repository.HandleDBError(err)
	}

	emp := mapper.ToEmployeeEntity(&empStorage)

	return emp, nil
}

func (r *EmployeeRepository) GetEmployeeByEmployeeEmail(ctx context.Context, empEmail string) (*entity.EmployeeEntity, error) {
	var empStorage model.EmployeeStorage

	err := r.db.WithContext(ctx).
		Preload("GoodsReceipts").Preload("GoodsIssues").Preload("GoodsReceipts.GoodsReceiptDetails").Preload("GoodsIssues.GoodsIssueDetails").
		Where("email = ?", empEmail).
		First(&empStorage).Error

	if err != nil {
		return nil, repository.HandleDBError(err)
	}

	emp := mapper.ToEmployeeEntity(&empStorage)

	return emp, nil
}

func (r *EmployeeRepository) GetEmployees(ctx context.Context) ([]entity.EmployeeEntity, error) {
	var empStorages []model.EmployeeStorage

	err := r.db.WithContext(ctx).
		Preload("GoodsReceipts").Preload("GoodsIssues").Preload("GoodsReceipts.GoodsReceiptDetails").Preload("GoodsIssues.GoodsIssueDetails").
		Order("created_at ASC").
		Find(&empStorages).Error

	if err != nil {
		return nil, repository.HandleDBError(err)
	}

	var emps []entity.EmployeeEntity
	for _, empStorage := range empStorages {
		emp := mapper.ToEmployeeEntity(&empStorage)
		emps = append(emps, *emp)
	}

	return emps, nil
}

func (r *EmployeeRepository) UpdateEmployeeForUser(ctx context.Context, update *entity.UpdateEmployeeEntityForUser) error {
	fields := map[string]any{}

	if update.Username != nil && *update.Username != "" { fields["username"] = *update.Username }
	if update.Phone    != nil && *update.Phone    != "" { fields["phone"]    = *update.Phone    }
	if update.Address  != nil && *update.Address  != "" { fields["address"]  = *update.Address  }
	if update.Avatar   != nil && *update.Avatar   != "" { fields["avatar"]   = *update.Avatar   }

	if len(fields) == 0 {
        return nil
    }

	result := r.db.WithContext(ctx).
        Model(&model.EmployeeStorage{}).
        Where("employee_id = ?", update.EmployeeID).
        Updates(fields)

    if result.Error != nil {
        return repository.HandleDBError(result.Error)
    }

    return nil
}

func (r *EmployeeRepository) UpdateEmployeeForAdmin(ctx context.Context, update *entity.UpdateEmployeeEntityForAdmin) error {
	fields := map[string]any{}

	if update.Name     != nil && *update.Name     != "" { fields["name"]     = *update.Name     }
	if update.Email    != nil && *update.Email    != "" { fields["email"]    = *update.Email    }
	if update.Username != nil && *update.Username != "" { fields["username"] = *update.Username }
	if update.Phone    != nil && *update.Phone    != "" { fields["phone"]    = *update.Phone    }
	if update.Address  != nil && *update.Address  != "" { fields["address"]  = *update.Address  }
	if update.Role     != nil && *update.Role     != "" { fields["role"]     = *update.Role     }
	if update.Avatar   != nil && *update.Avatar   != "" { fields["avatar"]   = *update.Avatar   }
	if update.Password != nil && *update.Password != "" { fields["password"] = *update.Password }

	if len(fields) == 0 {
        return nil
    }

	result := r.db.WithContext(ctx).
        Model(&model.EmployeeStorage{}).
        Where("employee_id = ?", update.EmployeeID).
        Updates(fields)

    if result.Error != nil {
        return repository.HandleDBError(result.Error)
    }

    return nil
}