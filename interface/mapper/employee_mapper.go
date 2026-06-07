package mapper

import (
	"Backend-Warehouse/domain/entity"
	"Backend-Warehouse/infrastructure/model"
	"Backend-Warehouse/interface/dto"

	"github.com/google/uuid"
)

// Mapper To Storage

func ToEmployeeStorage(e *entity.EmployeeEntity) *model.EmployeeStorage {
	employee := &model.EmployeeStorage{
		EmployeeID:    e.EmployeeID,
		Name:          e.Name,
		Email:         e.Email,
		Username:      e.Username,
		Phone:         e.Phone,
		Address:       e.Address,
		Role:          e.Role,
		Avatar:        e.Avatar,
		Password:      e.Password,
		GoodsReceipts: toGoodsReceiptStorages(e.GoodsReceipts),
		GoodsIssues:   toGoodsIssueStorages(e.GoodsIssues),
		Signatures:    toSignatureStorages(e.Signatures),
	}

	return employee
}

// Mapper To Entity

func ToEmployeeEntity(s *model.EmployeeStorage) *entity.EmployeeEntity {
	employee := &entity.EmployeeEntity{
		EmployeeID:    s.EmployeeID,
		Name:          s.Name,
		Email:         s.Email,
		Username:      s.Username,
		Phone:         s.Phone,
		Address:       s.Address,
		Role:          s.Role,
		Avatar:        s.Avatar,
		Password:      s.Password,
		GoodsReceipts: toGoodsReceiptEntities(s.GoodsReceipts),
		GoodsIssues:   toGoodsIssueEntities(s.GoodsIssues),
		Signatures:    toSignatureEntities(s.Signatures),
	}

	return employee
}

func ToUpdateEmployeeEntityForUser(req dto.UpdateEmployeeForUserRequest) *entity.UpdateEmployeeEntityForUser {
	return &entity.UpdateEmployeeEntityForUser{
		EmployeeID: req.EmployeeID,
		Username:   req.Username,
		Phone:      req.Phone,
		Address:    req.Address,
		Avatar:     req.Avatar,
	}
}

func ToUpdateEmployeeEntityForAdmin(req dto.UpdateEmployeeForAdminRequest) *entity.UpdateEmployeeEntityForAdmin {
	return &entity.UpdateEmployeeEntityForAdmin{
		EmployeeID: req.EmployeeID,
		Name:       req.Name,
		Email:      req.Email,
		Username:   req.Username,
		Phone:      req.Phone,
		Address:    req.Address,
		Role:       req.Role,
		Avatar:     req.Avatar,
		Password:   req.Password,
	}
}

// Mapper for DTO

func ToEntityFromRegisterEmployeeRequest(req dto.RegisterEmployeeRequest) *entity.EmployeeEntity {
	return &entity.EmployeeEntity{
		EmployeeID: uuid.NewString(),
		Name:       req.Name,
		Email:      req.Email,
		Username:   req.Username,
		Phone:      req.Phone,
		Address:    req.Address,
		Role:       req.Role,
		Password:   req.Password,
	}
}

func ToEmployeeResponseList(e []entity.EmployeeEntity) []dto.EmployeeResponse {
	responses := make([]dto.EmployeeResponse, 0, len(e))

	for _, response := range e {
		responses = append(responses, ToEmployeeResponse(&response))
	}

	return responses
}

func ToEmployeeResponse(e *entity.EmployeeEntity) dto.EmployeeResponse {
	return dto.EmployeeResponse{
		EmployeeID:    e.EmployeeID,
		Name:          e.Name,
		Email:         e.Email,
		Username:      e.Username,
		Phone:         e.Phone,
		Address:       e.Address,
		Role:          e.Role,
		Avatar:        e.Avatar,
		GoodsReceipts: ToGoodsReceiptResponseList(e.GoodsReceipts),
		GoodsIssues:   ToGoodsIssueResponseList(e.GoodsIssues),
	}
}