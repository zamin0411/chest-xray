package handler

import (
	"strconv"

	"chest-xray/database"
	"chest-xray/model"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	if uid != n {
		return false
	}

	return true
}

func validUser(username string, password string) bool {
	db := database.DB
	var doctor model.Doctor
	db.First(&doctor, username)
	if doctor.Username == "" {
		return false
	}
	if !CheckPasswordHash(password, doctor.Password) {
		return false
	}
	return true
}

func GetDoctor(c *fiber.Ctx) error {
	username := c.Params("username")
	db := database.DB
	var doctor model.Doctor
	db.Find(&doctor, username)
	if doctor.Username == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with username", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Doctor found", "data": doctor})
}

func GetDoctors(c *fiber.Ctx) error {
	db := database.DB
	var doctors []model.Doctor
	db.Find(&doctors)
	return c.JSON(fiber.Map{"status": "success", "message": "Doctors found", "data": doctors})
}
