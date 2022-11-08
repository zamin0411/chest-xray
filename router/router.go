package router

import (
	"chest-xray/handler"
	"chest-xray/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", middleware.LoginWithToken(), handler.LoginWithToken)

	// Doctor
	doctor := api.Group("/doctors")
	doctor.Get("/", handler.GetDoctors)

	// Medical Record
	record := api.Group("/records")
	record.Post("/", handler.CreateMedicalRecord)
	record.Get("/", handler.GetAllMedicalRecords)

	// OBJ Model
	model := record.Group("/models")
	model.Get("/", handler.GetOBJModel)
	model.Post("/", handler.SaveOBJModel)

}
