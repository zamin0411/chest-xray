package model

type Patient struct {
	ID      string `gorm:"column:patient_id"`
	Name    string `gorm:"column:patient_name"`
	Sex     string `gorm:"column:patient_sex"`
	Allergy string `gorm:"column:patient_allergy"`
}
