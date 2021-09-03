package middleware

import (
	"api/config"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2"
)

func SetUpMiddleware(cfg *config.Config, server *fiber.App) {
	config := cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowCredentials: true,
	}

	server.Use(cors.New(config))

	if cfg.HTTP.Env == "development" {
		// server.Use(middlewares.RequestLogger)
	}
}
