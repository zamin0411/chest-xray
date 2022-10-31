package handler

import (
	"chest-xray/database"
	"chest-xray/model"

	"github.com/gofiber/fiber/v2"
)

func GetOBJModel(c *fiber.Ctx) error {
	recordId := c.Params("recordId")
	db := database.DB
	var model model.OBJModel
	if err := db.First(&model, "medical_record_id = ?", recordId).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"code": c.Response().StatusCode(), "data": err})
	}

	return c.JSON(fiber.Map{"code": c.Response().StatusCode(), "data": model})
}

func SaveOBJModel(c *fiber.Ctx) error {
	db := database.DB
	var model model.OBJModel
	if err := c.BodyParser(&model); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"code": c.Response().StatusCode(), "data": err})
	}

	if err := db.Create(&model).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"code": c.Response().StatusCode(), "data": err})
	}

	return c.JSON(fiber.Map{"code": c.Response().StatusCode(), "data": model})
}
