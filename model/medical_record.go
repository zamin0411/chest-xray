package model

import "time"

type MedicalRecord struct {
	ID                      string     `gorm:"primaryKey;column:medical_record_id"`
	IDT                     string     `json:"id" gorm:"column:medical_record_id_text"`
	HospitalizationDatetime *time.Time `json:"datetime" gorm:"column:medical_record_hospitalization_datetime"`
	PresentingComplaint     string     `json:"presentingComplaint" gorm:"column:medical_record_presenting_complaint"`
	HospitalizationStatus   string     `json:"status" gorm:"column:medical_record_hospitalization_status"`
	InitialDiagnosis        string     `json:"initialDiagnosis" gorm:"column:medical_record_initial_diagnosis"`
	DifferentialDiagnosis   string     `json:"differentialDiagnosis" gorm:"column:medical_record_differential_diagnosis"`
	PrimaryDiagnosis        string     `json:"primaryDiagnosis" gorm:"column:medical_record_primary_diagnosis"`
	Treatment               string     `json:"treatment" gorm:"column:medical_record_treatment"`
	Doctors                 []*Doctor  `json:"doctors" gorm:"many2many:medical_record_has_doctor"`
}
