package middleware

import (
	"api/config"
	"api/keycloak"
	"fmt"
	"log"
	"strings"

	"github.com/MicahParks/keyfunc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt/v4"
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

func KeycloakMW(k *keycloak.Keycloak) fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqToken := c.Get(fiber.HeaderAuthorization)
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]

		jwks, err := keyfunc.Get(k.JwksUrl)
		if err != nil {
			log.Fatalf("Failed to get the JWKs from the given URL.\nError:%s\n", err.Error())
		}

		// Parse the JWT.
		token, err := jwt.Parse(reqToken, jwks.KeyFunc)
		if err != nil {
			fmt.Printf("failed to parse token: %s", err)
		}

		if !token.Valid {
			fmt.Printf("Invalid token")
		}

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
