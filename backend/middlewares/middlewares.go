package middlewares

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

func enableCache(app *fiber.App) {
	log.Println("[Middlewares] Enabling cache")

	app.Use(cache.New(cache.Config{
		Expiration:   3600 * time.Second,
		CacheControl: true,
	}))
}

func enableCompress(app *fiber.App) {
	log.Println("[Middlewares] Enabling compress")

	app.Use(compress.New())
}

func EnableMiddlewares(app *fiber.App) {

	// enable cache
	enableCache(app)

	// enable compress
	enableCompress(app)
}
