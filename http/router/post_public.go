package router

import (
	"api/keycloak"
	"api/pkg/entities"
	"api/pkg/post"
	"github.com/gofiber/fiber/v2"
	"strings"

)

func PostRouter(r fiber.Router, s post.PublicService, k *keycloak.Keycloak) {
	r.Get(PostsPath, getPosts(s))
	r.Get(PostsPath+"/:post_id", getPost(s))
}

func getPost(s post.PublicService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		postID := c.Params("post_id")
		p, err := s.FetchPost(postID)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&entities.ApiResponse{
				Code:    fiber.StatusInternalServerError,
				Type:    "DatabaseError",
				Message: strings.Title(err.Error()),
			})
		}

		return c.JSON(&p)
	}
}

func getPosts(s post.PublicService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		posts, err := s.FetchPosts()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&entities.ApiResponse{
				Code:    fiber.StatusInternalServerError,
				Type:    "DatabaseError",
				Message: strings.Title(err.Error()),
			})
		}

		publishedPosts := make([]entities.Post, 0)

		for _, p := range *posts {
			if p.Status == entities.PUBLISHED {
				publishedPosts = append(publishedPosts, p)
			}
		}

		return c.JSON(&publishedPosts)
	}
}
