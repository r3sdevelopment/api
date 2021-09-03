package http

import (
	"fmt"
	"log"

	"api/config"
	"api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setUpMiddleware(server *fiber.App) {
	config := cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowCredentials: true,
	}

	server.Use(cors.New(config))
}

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

	setUpMiddleware(server)
	setUpRoutes(server)

	log.Fatal(server.Listen(fmt.Sprintf(":%s", cfg.HTTP.Port)))
}
