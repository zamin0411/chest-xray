package handler

import (
	"chest-xray/database"
	"chest-xray/model"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func GetOBJModel(c *fiber.Ctx) error {
	db := database.DB

	if c.Body() == nil {
		var models []model.OBJModel
		if err := db.Find(&models).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"code": c.Response().StatusCode(), "message": "Failed Selecting Data", "data": err})
		}
		return c.JSON(fiber.Map{"code": c.Response().StatusCode(), "message": "Successful!", "data": models})
	} else {
		var obj model.OBJModel
		if err := c.BodyParser(&obj); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"code": c.Response().StatusCode(), "message": "Unprocessable Entity", "data": err})
		}
		recordId := obj.MedicalRecordID
		if err := db.First(&obj, "medical_record_id = ?", recordId).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"code": c.Response().StatusCode(), "message": "Failed Selecting Data", "data": err})
		}
		return c.JSON(fiber.Map{"code": c.Response().StatusCode(), "message": "Successful!", "data": obj})
	}
}

func SaveOBJModel(c *fiber.Ctx) error {
	db := database.DB
	var model model.OBJModel
	if err := c.BodyParser(&model); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"code": c.Response().StatusCode(), "message": "Unprocessable Entity", "data": err})
	}

	if err := db.Create(&model).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"code": c.Response().StatusCode(), "message": "Failed Inserting Data", "data": err})
	}

	var v [][]float64
	var f [][]int64
	json.Unmarshal(model.V, &v)
	json.Unmarshal(model.F, &f)
	fileName := fmt.Sprintf("files/%s.obj", model.ID)
	file, err := os.Create(fileName)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"code": c.Response().StatusCode, "message": "Failed Creating File", "data": err})
	}
	defer file.Close()
	for _, h := range v {
		for j, c := range h {
			if j == 0 {
				fmt.Fprint(file, "v ")
			}
			if j == 2 {
				fmt.Fprintln(file, c)
			} else {
				fmt.Fprintf(file, "%f ", c)
			}
		}
	}

	for _, h := range f {
		for j, c := range h {
			if j == 0 {
				fmt.Fprint(file, "f ")
			}
			if j == 2 {
				fmt.Fprintln(file, c)
			} else {
				fmt.Fprintf(file, "%d ", c)
			}
		}
	}

	// postBody, _ := json.Marshal(model)
	// responseBody := bytes.NewBuffer(postBody)

	// res, err := http.Post("", "application/json", responseBody)

	return c.JSON(fiber.Map{"code": c.Response().StatusCode(), "message": "Successful", "data": model})
}
