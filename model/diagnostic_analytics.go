package model

type DiagnosticAnalytics struct {
	ID              string `gorm:"column:diagnostic_analytics_id" json:"id"`
	MedicalRecordID string `gorm:"column:medical_record_id" json:"-"`
	Thesis          string `gorm:"column:diagnostic_analytics_thesis" json:"thesis"`
	Argument        string `gorm:"column:diagnostic_analytics_argument" json:"argument"`
}
