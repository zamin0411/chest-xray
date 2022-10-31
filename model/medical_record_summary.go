package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MedicalRecordSummary struct {
	ID                 uuid.UUID `gorm:"column:medical_record_summary_id" json:"id"`
	Sign               string    `gorm:"column:medical_record_summary_sign" json:"sign"`
	Symptom            string    `gorm:"column:medical_record_summary_symptom" json:"symptom"`
	GeneralInformation string    `gorm:"column:medical_record_summary_general_information" json:"generalInformation"`
}

func (summary *MedicalRecordSummary) BeforeCreate(tx *gorm.DB) (err error) {
	if summary.ID == uuid.Nil {
		summary.ID = uuid.New()
	}
	return
}
