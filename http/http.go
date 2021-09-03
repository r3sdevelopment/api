package http

import (
	"fmt"
	"log"

	"api/config"
	"api/routes"
	"api/http/middleware"

	"github.com/gofiber/fiber/v2"
)

func setUpRoutes(server *fiber.App) {
	server.Get("/hello", routes.Hello)
	server.Get("/allbooks", routes.AllBooks)
	server.Post("/addbook", routes.AddBook)
	server.Post("/book", routes.Book)
	server.Put("/update", routes.Update)
	server.Delete("/delete", routes.Delete)

	server.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
}

func Start(cfg *config.Config) {
	server := fiber.New()

	middleware.SetUpMiddleware(cfg, server)
	setUpRoutes(server)

	log.Fatal(server.Listen(fmt.Sprintf(":%s", cfg.HTTP.Port)))
}
