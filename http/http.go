package http

import (
	"fmt"
	"log"

	"api/config"
	mw "api/http/middleware"
	r "api/http/router"
	"api/keycloak"

	"github.com/gofiber/fiber/v2"
)

func Start(c *config.Config, k *keycloak.Keycloak) {
	s := fiber.New()

	mw.SetUpMiddleware(c, s, k)
	r.SetUpRoutes(s)

	log.Fatal(s.Listen(fmt.Sprintf(":%s", c.HTTP.Port)))
}
