package model

type Doctor struct {
	ID             string           `gorm:"primaryKey;column:doctor_id" json:"id"`
	Username       string           `gorm:"column:doctor_username" json:"username"`
	Password       string           `gorm:"column:doctor_password"`
	FullName       string           `gorm:"column:doctor_name" json:"fullName"`
	Sex            string           `gorm:"column:doctor_sex" json:"sex"`
	Specialization string           `gorm:"column:doctor_specialization" json:"specialization"`
	MedicalRecords []*MedicalRecord `gorm:"many2many:medical_record_has_doctor" json:"medicalRecords"`
}
