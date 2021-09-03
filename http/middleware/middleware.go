package middleware

import (
	"api/config"
	"api/keycloak"
	"fmt"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
)

type RealmRole map[string][]string
type ResourceRole map[string]map[string][]string
type jwtCustomClaims struct {
	*jwt.StandardClaims
	RealmAccess    RealmRole    `json:"realm_access"`
	ResourceAccess ResourceRole `json:"resource_access"`
}
type Header struct {
	Authorization string
}

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

func KeycloakMW(keycloak *keycloak.Keycloak) fiber.Handler {
	// Return new handler
	return func(c *fiber.Ctx) error {
		auth := c.Get(fiber.HeaderAuthorization)

		fmt.Printf("Token: %s", auth)

		return c.Next()
	}
}

func SetUpMiddleware(cfg *config.Config, server *fiber.App, keycloak *keycloak.Keycloak) {
	config := cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Authorization",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowCredentials: true,
		// ExposeHeaders: "Authorization,Content-Length",
		MaxAge: 5600,
	}

	server.Use(cors.New(config))
	server.Use(HardenSecurity)
	server.Use(KeycloakMW(keycloak))

	if cfg.HTTP.Env == "development" {
		server.Use(fiberLogger.New())
	}
}
