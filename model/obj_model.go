package model

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type OBJModel struct {
	ID              uuid.UUID      `gorm:"column:obj_model_id" json:"id"`
	MedicalRecordID uuid.UUID      `gorm:"column:medical_record_id" json:"-"`
	V               datatypes.JSON `gorm:"column:obj_model_v" json:"v"`
	F               datatypes.JSON `gorm:"column:obj_model_f" json:"f"`
}

func (model *OBJModel) BeforeCreate(tx *gorm.DB) (err error) {
	if model.ID == uuid.Nil {
		model.ID = uuid.New()
	}
	return
}
