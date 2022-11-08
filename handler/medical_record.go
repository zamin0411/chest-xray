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

	var records []model.MedicalRecord

	// db.Model(&model.MedicalRecord{}).Preload("Doctors").Find(&RecordsData{}).Find(&records)
	db.Model(&model.MedicalRecord{}).Preload("Doctors").
		Joins("JOIN medical_record_has_doctor md ON md.medical_record_id = medical_record.medical_record_id").
		Joins("JOIN doctor d ON d.doctor_id = md.doctor_id").
		Where("d.doctor_name = ?", name).Find(&records)
	// db.Model(&model.Doctor{}).Find(&DoctorsData{})
	// db.Model(&model.Doctor{}).Preload("MedicalRecords").Find(&doctors)
	// db.Model(&model.MedicalRecord{}).Preload("Doctors").Scan(&records)
	// db.Table("doctor").Joins("JOIN medical_record_has_doctor ON medical_record_has_doctor.doctor_id = doctor.doctor_id").Joins("JOIN medical_record ON medical_record_has_doctor.medical_record_id = medical_record.medical_record_id").Where("doctor_name = ?", name).Scan(&doctors)
	// db.Joins("JOIN medical_record_has_doctor md ON md.doctor_id = doctor.doctor_id JOIN medical_record mr.medical_record_id = md.medical_record_id AND doctor.doctor_name = ? ", name).Group("doctor.doctor_id")
	fmt.Print(records)
	if records == nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No medical record found with username", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "records found", "data": records})

}

func CreateMedicalRecord(c *fiber.Ctx) error {
	db := database.DB
	var record model.MedicalRecord
	if err := c.BodyParser(&record); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"code": c.Response().StatusCode(), "message": "review your input", "data": err})
	}

	if err := db.Create(&record).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"code": c.Response().StatusCode(), "message": "server error", "data": err})
	}

	return c.JSON(fiber.Map{"code": c.Response().StatusCode(), "message": "created successfully", "data": record})
}

func GetAllMedicalRecords(c *fiber.Ctx) error {
	db := database.DB
	var records []model.MedicalRecord
	db.Preload("MedicalRecordSummary").Preload("Doctors").Preload("ClinicalExamination").Preload("DiseaseHistory").Preload("Patient").Preload("SubclinicalExaminations").Preload("DiagnosticAnalytics").Find(&records)
	return c.JSON(fiber.Map{"code": c.Response().StatusCode(), "message": "Successful", "data": records})

}
