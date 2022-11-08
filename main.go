package main

import (
	"log"

	"chest-xray/database"
	"chest-xray/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	database.Connect()
	router.SetupRoutes(app)
	app.Static("/", "./files")
	log.Fatal(app.Listen(":3000"))
}
