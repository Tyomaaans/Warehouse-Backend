package mapper

import (
	"encoding/json"

	"gorm.io/datatypes"

	"Backend-Warehouse/domain/entity"
	"Backend-Warehouse/infrastructure/model"
)

func ToSignatureStorage(e *entity.SignatureEntity) *model.SignatureStorage {
	qrData, _ := json.Marshal(e.QrData)

    return &model.SignatureStorage{
        SignatureID: e.SignatureID,
        OwnerID:     e.OwnerID,
        OwnerType:   e.OwnerType,
        Signature:   e.Signature,
        QrData:      datatypes.JSON(qrData),
        QrImage:     e.QrImage,
        IsActive:    e.IsActive,
    }
}

func ToSignatureEntity(s *model.SignatureStorage) *entity.SignatureEntity {
	var qrData entity.QrData
    json.Unmarshal(s.QrData, &qrData)

    return &entity.SignatureEntity{
        SignatureID: s.SignatureID,
        OwnerID:     s.OwnerID,
        OwnerType:   s.OwnerType,
        Signature:   s.Signature,
        QrData:      qrData,
        QrImage:     s.QrImage,
        IsActive:    s.IsActive,
    }
}

func toSignatureStorages(entities []entity.SignatureEntity) []model.SignatureStorage {
	result := make([]model.SignatureStorage, len(entities))

	for i, e := range entities {
		result[i] = *ToSignatureStorage(&e)
	}
	
	return result
}

func toSignatureEntities(storages []model.SignatureStorage) []entity.SignatureEntity {
	result := make([]entity.SignatureEntity, len(storages))

	for i, e := range storages {
		result[i] = *ToSignatureEntity(&e)
	}
	
	return result
}