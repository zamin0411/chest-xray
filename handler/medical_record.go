package handler

import (
	"chest-xray/database"
	"chest-xray/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetMedicalRecordsByDoctorName(c *fiber.Ctx) error {
	name := c.Params("name")
	db := database.DB
	var doctors []model.Doctor
	var records []model.MedicalRecord
	// db.Model(&model.Doctor{}).Preload("MedicalRecords").Joins("INNER JOIN medical_record_has_doctor md ON md.doctor_id = doctor.doctor_id").
	// 	Joins("INNER JOIN medical_record mr ON mr.medical_record_id = md.medical_record_id").Where("doctor.doctor_name = ?", name).Find(&doctors)
	db.Model(&model.Doctor{}).Preload("MedicalRecords").Find(&doctors)
	db.Model(&model.MedicalRecord{}).Preload("Doctors").Find(&records)
	// db.Preload("MedicalRecords").Find(&doctors)
	// db.Table("doctor").Joins("JOIN medical_record_has_doctor ON medical_record_has_doctor.doctor_id = doctor.doctor_id").Joins("JOIN medical_record ON medical_record_has_doctor.medical_record_id = medical_record.medical_record_id").Where("doctor_name = ?", name).Scan(&doctors)
	// db.Joins("JOIN medical_record_has_doctor md ON md.doctor_id = doctor.doctor_id JOIN medical_record mr.medical_record_id = md.medical_record_id AND doctor.doctor_name = ? ", name).Group("doctor.doctor_id")
	fmt.Print(records)

	fmt.Print(doctors)
	return c.JSON(fiber.Map{"status": "success", "message": "records found", "data": doctors})

}
