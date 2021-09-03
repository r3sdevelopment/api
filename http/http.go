package http

import (
	"fmt"
	"log"

	"api/config"
	"api/http/middleware"
	"api/http/routes"

	"github.com/gofiber/fiber/v2"
)

func Start(cfg *config.Config) {
	server := fiber.New()

	middleware.SetUpMiddleware(cfg, server)
	routes.SetUpRoutes(server)

	log.Fatal(server.Listen(fmt.Sprintf(":%s", cfg.HTTP.Port)))
}
