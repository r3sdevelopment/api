package router

import (
	d "api/db"
	"api/keycloak"
	"api/pkg/entities"
	"api/pkg/post"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(s *fiber.App, k *keycloak.Keycloak) {
	postCollection := d.DB
	postRepo := post.NewRepo(postCollection)
	postService := post.NewService(postRepo)

	// Middleware
	api := s.Group("/api")
	// api.Get("/public", handler.Public)
	// api.Get("/user", k.NeedsRole([]string{"user"}), handler.User)
	// api.Get("/admin", k.NeedsRole([]string{"admin"}), handler.Admin)
	// api.Get("/all", k.NeedsRole([]string{"user", "admin"}), handler.All)

	PostRouter(api, postService, k)

	s.Use(func(c *fiber.Ctx) error {
		c.Status(404)
		return c.JSON(&entities.ApiResponse{
			Code:    404,
			Type:    "NotFoundError",
			Message: fmt.Sprintf("Not Found: %s", c.Path()),
		}) // => 404 "Not Found"
	})
}
