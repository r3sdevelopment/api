package middleware

import (
	"api/config"
	"api/keycloak"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func HardenSecurity(c *fiber.Ctx) error {
	// Set some security headers:
	c.Set("X-XSS-Protection", "1; mode=block")
	c.Set("X-Content-Type-Options", "nosniff")
	c.Set("X-Download-Options", "noopen")
	c.Set("Strict-Transport-Security", "max-age=5184000")
	c.Set("X-Frame-Options", "SAMEORIGIN")
	c.Set("X-DNS-Prefetch-Control", "off")

	return c.Next()
}

func SetUpMiddleware(_ *config.Config, s *fiber.App, k *keycloak.Keycloak) {
	s.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000, https://noc.r3s.dev",
		AllowHeaders:     "Authorization,Content-Type",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowCredentials: true,
		// ExposeHeaders: "Authorization,Content-Length",
		MaxAge: 5600,
	}))
	s.Use(HardenSecurity)
	s.Use(k.ApplyMiddleware())

	s.Use(logger.New())
}
