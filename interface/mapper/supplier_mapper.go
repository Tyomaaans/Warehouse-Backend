package mapper

import (
	"Backend-Warehouse/domain/entity"
	"Backend-Warehouse/infrastructure/model"
	"Backend-Warehouse/interface/dto"

	"github.com/google/uuid"
)

// Mapper To Storage

func ToSupplierStorage(e *entity.SupplierEntity) *model.SupplierStorage {
	supplier := &model.SupplierStorage{
		SupplierID:    e.SupplierID,
		Name:          e.Name,
		Company:       e.Company,
		Email:         e.Email,
		Phone:         e.Phone,
		Address:       e.Address,
		GoodsReceipts: toGoodsReceiptStorages(e.GoodsReceipts),
		Signatures:    toSignatureStorages(e.Signatures),
	}

	return supplier
}

// Mapper To Entity

func ToSupplierEntity(s *model.SupplierStorage) *entity.SupplierEntity {
	supplier := &entity.SupplierEntity{
		SupplierID:    s.SupplierID,
		Name:          s.Name,
		Company:       s.Company,
		Email:         s.Email,
		Phone:         s.Phone,
		Address:       s.Address,
		GoodsReceipts: toGoodsReceiptEntities(s.GoodsReceipts),
		Signatures:    toSignatureEntities(s.Signatures),
	}

	return supplier
}

func ToUpdateSupplierEntity(req dto.UpdateSupplierRequest) *entity.UpdateSupplierEntity {
	return &entity.UpdateSupplierEntity{
		SupplierID: req.SupplierID,
		Name:       req.Name,
		Company:    req.Company,
		Email:      req.Email,
		Phone:      req.Phone,
		Address:    req.Address,
	}
}

// Mapper for DTO

func ToEntityFromRegisterSupplier(req dto.RegisterSupplierRequest) *entity.SupplierEntity {
	return &entity.SupplierEntity{
		SupplierID: uuid.NewString(),
		Name:       req.Name,
		Company:    req.Company,
		Email:      req.Email,
		Phone:      req.Phone,
		Address:    req.Address,
	}
}

func ToSupplierResponse(e *entity.SupplierEntity) dto.SuppliersResponse{
	return dto.SuppliersResponse{
		SupplierID:    e.SupplierID,
		Name:          e.Name,
		Company:       e.Company,
		Email:         e.Email,
		Phone:         e.Phone,
		Address:       e.Address,
		GoodsReceipts: ToGoodsReceiptResponseList(e.GoodsReceipts),
	}
}

func ToSuppliersResponse(e []entity.SupplierEntity) []dto.SuppliersResponse {
	response := make([]dto.SuppliersResponse, 0, len(e))
	
	for _, supplier := range e {
		response = append(response, ToSupplierResponse(&supplier))
	}

	return response
}