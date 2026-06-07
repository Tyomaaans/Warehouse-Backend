package usecase

import (
	"context"

	"github.com/go-playground/validator/v10"

	"Backend-Warehouse/domain/repository"
	"Backend-Warehouse/interface/dto"
	"Backend-Warehouse/interface/mapper"
	jsonValidator "Backend-Warehouse/validator"
)

type SupplierUsecase struct {
	spplrRepo repository.SupplierRepository
	validate  *validator.Validate
}

func NewSupplierUsecase(
	spplrRepo repository.SupplierRepository,
	validate *validator.Validate,
) *SupplierUsecase {
	return &SupplierUsecase{
		spplrRepo: spplrRepo,
		validate:  validate,
	}
}

func (uc *SupplierUsecase) RegisterSupplier(ctx context.Context, req dto.RegisterSupplierRequest) error {
	if err := jsonValidator.ValidateStruct(uc.validate, req); err != nil {
		return err
	}

	spplr := mapper.ToEntityFromRegisterSupplier(req)
	if err := uc.spplrRepo.CreateSupplier(ctx, spplr); err != nil {
		return err
	}

	return nil
}

func (uc *SupplierUsecase) GetSuppliers(ctx context.Context) ([]dto.SuppliersResponse, error) {
	spplrs, err := uc.spplrRepo.GetSuppliers(ctx)
	if err != nil {
		return nil, err
	}

	res := mapper.ToSuppliersResponse(spplrs)

    return res, nil
}

func (uc *SupplierUsecase) GetSupplier(ctx context.Context, req dto.GetSupplierRequest) (*dto.SuppliersResponse, error) {
	spplrs, err := uc.spplrRepo.GetSupplierBySupplierID(ctx, req.SupplierID)
	if err != nil {
		return nil, err
	}

	res := mapper.ToSupplierResponse(spplrs)

	return &res, nil
}

func (uc *SupplierUsecase) UpdateSupplier(ctx context.Context, req dto.UpdateSupplierRequest) error {
	if err := jsonValidator.ValidateStruct(uc.validate, req); err != nil {
		return err
	}

	update := mapper.ToUpdateSupplierEntity(req)
	if err := uc.spplrRepo.UpdateSupplier(ctx, update); err != nil {
		return err
	}

	return nil
}