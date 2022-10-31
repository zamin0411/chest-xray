package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubclinicalExamination struct {
	ID              uuid.UUID `gorm:"column:subclinical_examination_id" json:"id"`
	Test            string    `gorm:"column:subclinical_examination_test" json:"test"`
	Type            string    `gorm:"column:subclinical_examination_type" json:"type"`
	Result          string    `gorm:"column:subclinical_examination_result" json:"result"`
	Conclusion      string    `gorm:"column:subclinical_examination_conclusion" json:"conclusion"`
	Image           []byte    `gorm:"column:subclinical_examination_image" json:"image"`
	MedicalRecordID uuid.UUID `gorm:"column:medical_record_id" json:"-"`
}

func (examination *SubclinicalExamination) BeforeCreate(tx *gorm.DB) (err error) {
	if examination.ID == uuid.Nil {
		examination.ID = uuid.New()
	}
	return
}
