package main

import (
	"backend/api"
	"backend/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()
	api.SetupAPI(app)
	log.Fatal(app.Listen(":8080"))
}
