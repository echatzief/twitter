package main

import (
	"backend/api"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	api.SetupAPI(app)
	log.Fatal(app.Listen(":8080"))
}
