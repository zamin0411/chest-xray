package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClinicalExamination struct {
	ID               uuid.UUID    `gorm:"column:clinical_examination_id" json:"id"`
	DateTime         *time.Time   `gorm:"column:clinical_examination_datetime" json:"datetime"`
	GeneralCondition string       `gorm:"column:clinical_examination_general_condition" json:"generalCondition"`
	AreaDetail       []AreaDetail `json:"areaDetail"`
}

func (ex *ClinicalExamination) BeforeCreate(tx *gorm.DB) (err error) {
	ex.ID = uuid.New()
	return
}
