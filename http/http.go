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
	s := fiber.New(fiber.Config{
    Prefork:       true,
    CaseSensitive: true,
    StrictRouting: false,
    ServerHeader:  "r3s.dev",
    AppName: "api.r3s.dev",
})


	mw.SetUpMiddleware(c, s, k)
	r.SetUpRoutes(s, k)

	log.Fatal(s.Listen(fmt.Sprintf("%s:%s", c.HTTP.IP, c.HTTP.Port)))
}
