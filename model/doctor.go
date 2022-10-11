package model

type Doctor struct {
	ID             string `gorm:"primaryKey;column:doctor_id"`
	IDT            string `json:"id" gorm:"column:doctor_id_text"`
	Username       string `json:"username" gorm:"column:doctor_username"`
	Password       string `gorm:"column:doctor_password"`
	FullName       string `json:"fullName" gorm:"column:doctor_name"`
	Sex            string `json:"sex" gorm:"column:doctor_sex"`
	Specialization string `json:"specialization" gorm:"column:doctor_specialization"`
	// MedicalRecords []*MedicalRecord `json:"medicalRecords" gorm:"many2many:medical_record_has_doctor"`
}
