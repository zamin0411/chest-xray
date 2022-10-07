package model

type Doctor struct {
	ID             string           `gorm:"column:doctor_id_text"`
	Username       string           `gorm:"column:doctor_username"`
	Password       string           `gorm:"column:doctor_password"`
	FullName       string           `gorm:"column:doctor_name"`
	Sex            string           `gorm:"column:doctor_sex"`
	Specialization string           `gorm:"column:doctor_specialization"`
	MedicalRecords *[]MedicalRecord `gorm:"many2many:medical_record_has_doctor"`
}
