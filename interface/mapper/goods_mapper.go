package mapper

import (
	"Backend-Warehouse/domain/entity"
	"Backend-Warehouse/infrastructure/model"
	"Backend-Warehouse/interface/dto"
	"fmt"
	"math/rand/v2"
	"strings"
	"time"

	"github.com/google/uuid"
)

// ========================
// Entity <-> Storage
// ========================

func ToGoodsReceiptStorage(e *entity.GoodsReceiptEntity) *model.GoodsReceiptStorage {
	return &model.GoodsReceiptStorage{
		GoodsReceiptID:      e.GoodsReceiptID,
		Invoice:             e.Invoice,
		SupplierID:          e.SupplierID,
		DateOfEntry:         e.DateOfEntry,
		EmployeeID:          e.EmployeeID,
		Status:              e.Status,
		GoodsReceiptDetails: toGoodsReceiptDetailStorages(e.GoodsReceiptDetails),
	}
}

func ToGoodsReceiptDetailStorage(e entity.GoodsReceiptDetailEntity) *model.GoodsReceiptDetailStorage {
	return &model.GoodsReceiptDetailStorage{
		GoodsReceiptDetailID: e.GoodsReceiptDetailID,
		GoodsReceiptID:       e.GoodsReceiptID,
		ItemID:               e.ItemID,
		Qty:                  e.Qty,
	}
}

func ToGoodsReceiptEntity(s *model.GoodsReceiptStorage) *entity.GoodsReceiptEntity {
	return &entity.GoodsReceiptEntity{
		GoodsReceiptID:      s.GoodsReceiptID,
		Invoice:             s.Invoice,
		SupplierID:          s.SupplierID,
		DateOfEntry:         s.DateOfEntry,
		EmployeeID:          s.EmployeeID,
		Status:              s.Status,
		GoodsReceiptDetails: toGoodsReceiptDetailEntities(s.GoodsReceiptDetails),
	}
}

func ToGoodsReceiptDetailEntity(s model.GoodsReceiptDetailStorage) *entity.GoodsReceiptDetailEntity {
	return &entity.GoodsReceiptDetailEntity{
		GoodsReceiptDetailID: s.GoodsReceiptDetailID,
		GoodsReceiptID:       s.GoodsReceiptID,
		ItemID:               s.ItemID,
		Qty:                  s.Qty,
	}
}

func ToGoodsIssueStorage(e *entity.GoodsIssueEntity) *model.GoodsIssueStorage {
	return &model.GoodsIssueStorage{
		GoodsIssueID:      e.GoodsIssueID,
		Invoice:           e.Invoice,
		ExitDestination:   e.ExitDestination,
		DateOfExit:        e.DateOfExit,
		EmployeeID:        e.EmployeeID,
		Status:            e.Status,
		GoodsIssueDetails: toGoodsIssueDetailStorages(e.GoodsIssueDetails),
	}
}

func ToGoodsIssueDetailStorage(e entity.GoodsIssueDetailEntity) *model.GoodsIssueDetailStorage {
	return &model.GoodsIssueDetailStorage{
		GoodsIssueDetailID: e.GoodsIssueDetailID,
		GoodsIssueID:       e.GoodsIssueID,
		ItemID:             e.ItemID,
		Qty:                e.Qty,
	}
}

func ToGoodsIssueEntity(s *model.GoodsIssueStorage) *entity.GoodsIssueEntity {
	return &entity.GoodsIssueEntity{
		GoodsIssueID:      s.GoodsIssueID,
		Invoice:           s.Invoice,
		ExitDestination:   s.ExitDestination,
		DateOfExit:        s.DateOfExit,
		EmployeeID:        s.EmployeeID,
		Status:            s.Status,
		GoodsIssueDetails: toGoodsIssueDetailEntities(s.GoodsIssueDetails),
	}
}

func ToGoodsIssueDetailEntity(s model.GoodsIssueDetailStorage) *entity.GoodsIssueDetailEntity {
	return &entity.GoodsIssueDetailEntity{
		GoodsIssueDetailID: s.GoodsIssueDetailID,
		GoodsIssueID:       s.GoodsIssueID,
		ItemID:             s.ItemID,
		Qty:                s.Qty,
	}
}

// ========================
// DTO -> Entity
// ========================

func ToEntityFromAddGoodsReceiptRequest(req dto.AddGoodsReceiptRequest) *entity.GoodsReceiptEntity {
	goodsReceiptID := uuid.NewString()

	return &entity.GoodsReceiptEntity{
		GoodsReceiptID:      goodsReceiptID,
		Invoice:             generateInvoiceReceiptID("GDRC", req.SupplierID, req.EmployeeID),
		SupplierID:          req.SupplierID,
		DateOfEntry:         req.DateOfEntry,
		EmployeeID:          req.EmployeeID,
		Status:              "pending",
		GoodsReceiptDetails: toDetailReceiptEntitiesFromRequest(goodsReceiptID, req.GoodsReceiptDetails),
	}
}

func ToEntityFromAddGoodsIssueRequest(req dto.AddGoodsIssueRequest) *entity.GoodsIssueEntity {
	goodsIssueID := uuid.NewString()

	return &entity.GoodsIssueEntity{
		GoodsIssueID:      goodsIssueID,
		Invoice:           generateInvoiceIssueID("GDIS", req.EmployeeID),
		ExitDestination:   req.ExitDestination,
		DateOfExit:        req.DateOfExit,
		EmployeeID:        req.EmployeeID,
		Status:            "pending",
		GoodsIssueDetails: toDetailIssueEntitiesFromRequest(goodsIssueID, req.GoodsIssueDetails),
	}
}

// ========================
// Entity -> DTO Response
// ========================

func ToGoodsReceiptResponseList(e []entity.GoodsReceiptEntity) []dto.GoodsReceiptResponse {
	responses := make([]dto.GoodsReceiptResponse, 0, len(e))

	for _, response := range e {
		responses = append(responses, ToGoodsReceiptResponse(&response))
	}

	return responses
}

func ToGoodsReceiptResponse(e *entity.GoodsReceiptEntity) dto.GoodsReceiptResponse {
	return dto.GoodsReceiptResponse{
		GoodsReceiptID:      e.GoodsReceiptID,
		Invoice:             e.Invoice,
		SupplierID:          e.SupplierID,
		DateOfEntry:         e.DateOfEntry,
		EmployeeID:          e.EmployeeID,
		Status:              e.Status,
		GoodsReceiptDetails: toGoodsReceiptDetailResponses(e.GoodsReceiptDetails),
	}
}

func ToGoodsIssueResponseList(e []entity.GoodsIssueEntity) []dto.GoodsIssueResponse {
	responses := make([]dto.GoodsIssueResponse, 0, len(e))
	
	for _, response := range e {
		responses = append(responses, ToGoodsIssueResponse(&response))
	}
	
	return responses
}

func ToGoodsIssueResponse(e *entity.GoodsIssueEntity) dto.GoodsIssueResponse {
	return dto.GoodsIssueResponse{
		GoodsIssueID:        e.GoodsIssueID,
		Invoice:             e.Invoice,
		ExitDestination:     e.ExitDestination,
		DateOfExit:          e.DateOfExit,
		EmployeeID:          e.EmployeeID,
		Status:              e.Status,
		GoodsIssueDetails:   toGoodsIssueDetailResponses(e.GoodsIssueDetails),
	}
}

// ========================
// Unexported Helpers
// ========================

func toDetailReceiptEntitiesFromRequest(goodsReceiptID string, reqs []dto.AddGoodsReceiptDetailsRequest) []entity.GoodsReceiptDetailEntity {
	details := make([]entity.GoodsReceiptDetailEntity, len(reqs))

	for i, detail := range reqs {
		details[i] = entity.GoodsReceiptDetailEntity{
			GoodsReceiptDetailID: uuid.NewString(),
			GoodsReceiptID:       goodsReceiptID,
			ItemID:               detail.ItemID,
			Qty:                  detail.Qty,
		}
	}

	return details
}

func toDetailIssueEntitiesFromRequest(goodsIssueID string, reqs []dto.AddGoodsIssueDetailRequest) []entity.GoodsIssueDetailEntity {
	details := make([]entity.GoodsIssueDetailEntity, len(reqs))
	
	for i, detail := range reqs {
		details[i] = entity.GoodsIssueDetailEntity{
			GoodsIssueDetailID: uuid.NewString(),
			GoodsIssueID:       goodsIssueID,
			ItemID:             detail.ItemID,
			Qty:                detail.Qty,
		}
	}

	return details
}

func toGoodsReceiptDetailResponses(e []entity.GoodsReceiptDetailEntity) []dto.GoodsReceiptDetailResponse {
	responses := make([]dto.GoodsReceiptDetailResponse, len(e))

	for i, response := range e {
		responses[i] = dto.GoodsReceiptDetailResponse{
			GoodsReceiptDetailID: response.GoodsReceiptDetailID,
			GoodsReceiptID:       response.GoodsReceiptID,
			ItemID:               response.ItemID,
			Qty:                  response.Qty,
		}
	}

	return responses
}

func toGoodsIssueDetailResponses(e []entity.GoodsIssueDetailEntity) []dto.GoodsIssueDetailResponse {
	responses := make([]dto.GoodsIssueDetailResponse, len(e))

	for i, response := range e {
		responses[i] = dto.GoodsIssueDetailResponse{
			GoodsIssueDetailID:   response.GoodsIssueDetailID,
			GoodsIssueID:         response.GoodsIssueID,
			ItemID:               response.ItemID,
			Qty:                  response.Qty,
		}
	}

	return responses
}

func toGoodsReceiptEntities(storages []model.GoodsReceiptStorage) []entity.GoodsReceiptEntity {
	result := make([]entity.GoodsReceiptEntity, len(storages))
	for i, s := range storages {
		result[i] = *ToGoodsReceiptEntity(&s)
	}
	return result
}

func toGoodsReceiptStorages(entities []entity.GoodsReceiptEntity) []model.GoodsReceiptStorage {
	result := make([]model.GoodsReceiptStorage, len(entities))
	for i, e := range entities {
		result[i] = *ToGoodsReceiptStorage(&e)
	}
	return result
}

func toGoodsReceiptDetailEntities(storages []model.GoodsReceiptDetailStorage) []entity.GoodsReceiptDetailEntity {
	result := make([]entity.GoodsReceiptDetailEntity, len(storages))
	for i, s := range storages {
		result[i] = *ToGoodsReceiptDetailEntity(s)
	}
	return result
}

func toGoodsReceiptDetailStorages(entities []entity.GoodsReceiptDetailEntity) []model.GoodsReceiptDetailStorage {
	result := make([]model.GoodsReceiptDetailStorage, len(entities))
	for i, e := range entities {
		result[i] = *ToGoodsReceiptDetailStorage(e)
	}
	return result
}

func toGoodsIssueEntities(storages []model.GoodsIssueStorage) []entity.GoodsIssueEntity {
	result := make([]entity.GoodsIssueEntity, len(storages))
	for i, s := range storages {
		result[i] = *ToGoodsIssueEntity(&s)
	}
	return result
}

func toGoodsIssueStorages(entities []entity.GoodsIssueEntity) []model.GoodsIssueStorage {
	result := make([]model.GoodsIssueStorage, len(entities))
	for i, e := range entities {
		result[i] = *ToGoodsIssueStorage(&e)
	}
	return result
}

func toGoodsIssueDetailEntities(storages []model.GoodsIssueDetailStorage) []entity.GoodsIssueDetailEntity {
	result := make([]entity.GoodsIssueDetailEntity, len(storages))
	for i, s := range storages {
		result[i] = *ToGoodsIssueDetailEntity(s)
	}
	return result
}

func toGoodsIssueDetailStorages(entities []entity.GoodsIssueDetailEntity) []model.GoodsIssueDetailStorage {
	result := make([]model.GoodsIssueDetailStorage, len(entities))
	for i, e := range entities {
		result[i] = *ToGoodsIssueDetailStorage(e)
	}
	return result
}

func generateInvoiceReceiptID(typeReq, supplierID, employeeID string) string {
	const (
		charSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
		randLen   = 6
		idPadLen  = 4
	)

	date := time.Now().Format("020106")

	var sb strings.Builder
	sb.Grow(randLen)
	for range randLen {
		sb.WriteByte(charSet[rand.IntN(len(charSet))])
	}

	sup := truncatePad(supplierID, idPadLen)
	emp := truncatePad(employeeID, idPadLen)

	return fmt.Sprintf("INV/%s/%s/%s/%s/%s", date, typeReq, sup, emp, sb.String())
}

func generateInvoiceIssueID(typeReq, employeeID string) string {
	const (
		charSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
		randLen   = 6
		idPadLen  = 4
	)

	date := time.Now().Format("020106")

	var sb strings.Builder
	sb.Grow(randLen)
	for range randLen {
		sb.WriteByte(charSet[rand.IntN(len(charSet))])
	}

	emp := truncatePad(employeeID, idPadLen)

	return fmt.Sprintf("INV/%s/%s/%s/%s/%s", date, typeReq, truncatePad(sb.String(), idPadLen), emp, sb.String())
}

func truncatePad(s string, n int) string {
	s = strings.ToUpper(fmt.Sprintf("%0*s", n, s))
	if len(s) > n {
		return s[len(s)-n:]
	}
	return s
}