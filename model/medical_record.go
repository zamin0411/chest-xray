package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MedicalRecord struct {
	ID                       uuid.UUID                `gorm:"primaryKey;column:medical_record_id" json:"id"`
	HospitalizationDateTime  time.Time                `gorm:"column:medical_record_hospitalization_datetime" json:"hospitalizationDateTime"`
	PresentingComplaint      string                   `gorm:"column:medical_record_presenting_complaint" json:"presentingComplaint"`
	HospitalizationStatus    string                   `gorm:"column:medical_record_hospitalization_status" json:"hospitalizationStatus"`
	InitialDiagnosis         string                   `gorm:"column:medical_record_initial_diagnosis" json:"initialDiagnosis"`
	DifferentialDiagnosis    string                   `gorm:"column:medical_record_differential_diagnosis" json:"differentialDiagnosis"`
	PrimaryDiagnosis         string                   `gorm:"column:medical_record_primary_diagnosis" json:"primaryDiagnosis"`
	Treatment                string                   `gorm:"column:medical_record_treatment" json:"treatment"`
	Doctors                  []*Doctor                `gorm:"many2many:medical_record_has_doctor" json:"doctors"`
	MedicalRecordSummaryID   uuid.UUID                `gorm:"column:medical_record_summary_id" json:"-"`
	MedicalRecordSummary     MedicalRecordSummary     `json:"medicalRecordSummary"`
	ClinicalExaminationID    uuid.UUID                `gorm:"column:clinical_examination_id" json:"-"`
	ClinicalExamination      ClinicalExamination      `json:"clinicalExamination"`
	CompleteBloodCountTestID sql.NullString           `gorm:"column:complete_blood_count_test_id" json:"-"`
	DiseaseHistory           []DiseaseHistory         `json:"diseaseHistory"`
	PatientID                uuid.UUID                `gorm:"column:patient_id" json:"-"`
	Patient                  Patient                  `json:"patient"`
	SubclinicalExaminations  []SubclinicalExamination `json:"subclinicalExaminations"`
	DiagnosticAnalytics      []DiagnosticAnalytics    `json:"diagnosticAnalytics"`
}

func (record *MedicalRecord) BeforeCreate(tx *gorm.DB) (err error) {
	if record.ID == uuid.Nil {
		record.ID = uuid.New()
	}
	return
}
