package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DiseaseHistory struct {
	ID              uuid.UUID `gorm:"column:disease_history_id" json:"id"`
	Stage           uint      `gorm:"column:disease_history_stage" json:"stage"`
	Detail          string    `gorm:"column:disease_history_detail" json:"detail"`
	MedicalRecordID string    `gorm:"column:medical_record_id" json:"medicalRecordId"`
}

func (history *DiseaseHistory) BeforeCreate(tx *gorm.DB) (err error) {
	history.ID = uuid.New()
	return
}
