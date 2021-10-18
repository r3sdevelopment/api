package router

import (
	"api/keycloak"
	"api/pkg/entities"
	"api/pkg/post"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func PostRouterAdmin(r fiber.Router, s post.AdminService, k *keycloak.Keycloak) {
	r.Get(PostsPath, k.NeedsRole([]string{"admin"}), getPostsAdmin(s))
	r.Get(PostsPath+"/:post_id", k.NeedsRole([]string{"admin"}), getPostAdmin(s))
	r.Post(PostsPath, k.NeedsRole([]string{"admin"}), addPostAdmin(s))
	r.Put(PostsPath+"/:post_id", k.NeedsRole([]string{"admin"}), updatePostAdmin(s))
	r.Delete(PostsPath+"/:post_id", k.NeedsRole([]string{"admin"}), removePostAdmin(s))
}

func addPostAdmin(s post.AdminService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var p entities.Post

		err := c.BodyParser(&p)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&entities.ApiResponse{
				Code:    fiber.StatusInternalServerError,
				Type:    "BodyParserError",
				Message: strings.Title(err.Error()),
			})
		}

		if userID, ok := c.Locals(keycloak.UserIdKey).(string); ok {
			p.UserId = userID
		}

		if err := p.Validate(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&entities.ApiResponse{
				Code:    fiber.StatusBadRequest,
				Type:    "ValidationError",
				Message: strings.Title(err.Error()),
			})
		}

		result, err := s.InsertPost(&p)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&entities.ApiResponse{
				Code:    fiber.StatusInternalServerError,
				Type:    "DatabaseError",
				Message: strings.Title(err.Error()),
			})
		}
		return c.JSON(result)
	}
}

func updatePostAdmin(service post.AdminService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var p entities.Post
		postID := c.Params("post_id")
		if err := c.BodyParser(&p); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&entities.ApiResponse{
				Code:    fiber.StatusInternalServerError,
				Type:    "ParsingError",
				Message: strings.Title(err.Error()),
			})
		}
		p.ID = postID
		if userID, ok := c.Locals(keycloak.UserIdKey).(string); ok {
			p.UserId = userID
		}
		if err := p.Validate(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&entities.ApiResponse{
				Code:    fiber.StatusBadRequest,
				Type:    "ValidationError",
				Message: strings.Title(err.Error()),
			})
		}
		result, err := service.UpdatePost(&p)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&entities.ApiResponse{
				Code:    fiber.StatusInternalServerError,
				Type:    "DatabaseError",
				Message: strings.Title(err.Error()),
			})
		}
		return c.JSON(result)
	}
}

func removePostAdmin(service post.AdminService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var p entities.DeleteRequest
		err := c.BodyParser(&p)
		postID := c.Params("post_id")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&entities.ApiResponse{
				Code:    fiber.StatusInternalServerError,
				Type:    "ParseError",
				Message: strings.Title(err.Error()),
			})
		}
		dbErr := service.RemovePost(postID)
		if dbErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&entities.ApiResponse{
				Code:    fiber.StatusInternalServerError,
				Type:    "DatabaseError",
				Message: strings.Title(dbErr.Error()),
			})
		}
		return c.JSON(&entities.ApiResponse{
			Code:    fiber.StatusOK,
			Type:    "Success",
			Message: fmt.Sprintf("Post with ID %s was deleted", postID),
		})
	}
}

func getPostAdmin(s post.AdminService) fiber.Handler {
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

func getPostsAdmin(s post.AdminService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		posts, err := s.FetchPosts()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&entities.ApiResponse{
				Code:    fiber.StatusInternalServerError,
				Type:    "DatabaseError",
				Message: strings.Title(err.Error()),
			})
		}

		return c.JSON(&posts)
	}
}
