package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Patient struct {
	ID               uuid.UUID  `gorm:"column:patient_id" json:"id"`
	Name             string     `gorm:"column:patient_name" json:"name"`
	Sex              uint       `gorm:"column:patient_sex" json:"sex"`
	DoB              *time.Time `gorm:"column:patient_date_of_birth" json:"dob"`
	Allergy          string     `gorm:"column:patient_allergy" json:"allergy"`
	Vaccination      string     `gorm:"column:patient_vaccination" json:"vaccination"`
	HealthBackground string     `gorm:"column:patient_health_background" json:"healthBackground"`
	Address          string     `gorm:"column:patient_address" json:"address"`
}

func (patient *Patient) BeforeCreate(tx *gorm.DB) (err error) {
	if patient.ID == uuid.Nil {
		patient.ID = uuid.New()
	}
	return
}
