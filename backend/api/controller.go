package api

import "github.com/gofiber/fiber/v2"

func SetupAPI(app *fiber.App) {
	api := app.Group("/api")

	// users
	api.Get("/users", GetUsers)
	api.Post("/users", CreateUser)
	api.Get("/users/:id", GetUser)
	api.Put("/users/:id", UpdateUser)
	api.Delete("/users/:id", DeleteUser)

	// tweets
	api.Get("/tweets", GetTweets)
	api.Post("/tweets", CreateTweet)
	api.Get("/tweets/:id", GetTweet)
	api.Put("/tweets/:id", UpdateTweet)
	api.Delete("/tweets/:id", DeleteTweet)
}
