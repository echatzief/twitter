package main

import (
	"backend/api"
	"backend/database"
	"backend/middlewares"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()

	// enable middlewares
	middlewares.EnableMiddlewares(app)

	// setup api
	api.SetupAPI(app)

	log.Fatal(app.Listen(":8080"))
}
