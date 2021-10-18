package router

import (
	d "api/db"
	"api/keycloak"
	"api/pkg/entities"
	"api/pkg/post"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const PostsPath = "/posts"

func SetUpRoutes(s *fiber.App, k *keycloak.Keycloak) {
	postCollection := d.DB

	postRepo := post.NewRepo(postCollection)

	// Public routes
	api := s.Group("/")
	postService := post.NewService(postRepo)
	PostRouter(api, postService, k)

	// Admin routes
	admin := s.Group("/admin")
	postAdminService := post.NewAdminService(postRepo)
	PostRouterAdmin(admin, postAdminService, k)

	s.Use(func(c *fiber.Ctx) error {
		c.Status(404)
		return c.JSON(&entities.ApiResponse{
			Code:    404,
			Type:    "NotFoundError",
			Message: fmt.Sprintf("Not Found: %s", c.Path()),
		}) // => 404 "Not Found"
	})
}
