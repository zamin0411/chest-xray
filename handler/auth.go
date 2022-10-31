package handler

import (
	"errors"
	"fmt"
	"time"

	"chest-xray/config"
	"chest-xray/database"
	"chest-xray/model"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getDoctorByUsername(u string) (*model.Doctor, error) {
	db := database.DB
	var doctor model.Doctor
	if err := db.Table("doctor").Where("doctor_username = ?", u).First(&doctor).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &doctor, nil
}

// Login get user and password
func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	type DoctorData struct {
		ID       string `json:"ID"`
		Username string `json:"username"`
	}
	input := new(LoginInput)
	var ud DoctorData
	fmt.Print(c.Body())
	if err := c.BodyParser(&input); err != nil {
		// return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"code": c.Response().StatusCode(), "message": "Error on login request", "data": err})
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	username := input.Username
	pass := input.Password

	user, err := getDoctorByUsername(username)

	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"code": c.Response().StatusCode(), "message": "Error on username", "data": err})
	}

	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"code": c.Response().StatusCode(), "message": "Invalid username or password", "data": nil})
	}

	if user != nil {
		ud = DoctorData{
			ID:       user.ID,
			Username: user.Username,
		}

	}

	if pass != user.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid username or password", "data": nil})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = ud.Username
	claims["id"] = ud.ID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = logSession(t)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t, "message": "Login successfully!", "status": "success", "code": c.Response().StatusCode()})
}

func LoginWithToken(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)

	if err := logSession(token.Raw); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": token.Raw, "message": "Login successfully!", "status": "success", "code": c.Response().StatusCode()})
}

func logSession(token string) error {
	session := model.LoginSession{
		Token: token,
	}

	db := database.DB

	if err := db.Create(&session).Error; err != nil {
		return err
	}

	return nil

}
