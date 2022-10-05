package model

type Doctor struct {
	ID             string `json:"id" gorm:"column:doctor_id_text"`
	Username       string `json:"username" gorm:"column:doctor_username"`
	Password       string `json:"password" gorm:"column:doctor_password"`
	FullName       string `json:"fullName" gorm:"column:doctor_name"`
	Sex            string `json:"sex" gorm:"column:doctor_sex"`
	Specialization string `json:"specialization" gorm:"column:doctor_specialization"`
}
