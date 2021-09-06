package router

import (
	"api/keycloak"
	"api/pkg/entities"
	"api/pkg/post"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const POSTS_PATH = "/posts"

func PostRouter(r fiber.Router, s post.Service, k *keycloak.Keycloak) {
	r.Get(POSTS_PATH, getPosts(s))
	r.Get(POSTS_PATH+"/:post_id", getPost(s))
	r.Post(POSTS_PATH, k.NeedsRole([]string{"admin"}), addPost(s))
	r.Put(POSTS_PATH+"/:post_id", k.NeedsRole([]string{"admin"}), updatePost(s))
	r.Delete(POSTS_PATH+"/:post_id", k.NeedsRole([]string{"admin"}), removePost(s))
}

func addPost(s post.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var post entities.Post

		c.BodyParser(&post)

		if err := post.Validate(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&entities.ApiResponse{
				Code:    fiber.StatusBadRequest,
				Type:    "ValidationError",
				Message: strings.Title(err.Error()),
			})
		}

		result, dberr := s.InsertPost(&post)

		if dberr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&entities.ApiResponse{
				Code:    fiber.StatusInternalServerError,
				Type:    "DatabaseError",
				Message: strings.Title(dberr.Error()),
			})
		}
		return c.JSON(result)
	}
}

func updatePost(service post.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var post entities.Post
		postID := c.Params("post_id")
		if err := c.BodyParser(&post); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&entities.ApiResponse{
				Code:    fiber.StatusInternalServerError,
				Type:    "ParsingError",
				Message: strings.Title(err.Error()),
			})
		}
		post.ID = postID
		if err := post.Validate(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&entities.ApiResponse{
				Code:    fiber.StatusBadRequest,
				Type:    "ValidationError",
				Message: strings.Title(err.Error()),
			})
		}
		result, dberr := service.UpdatePost(&post)

		if dberr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&entities.ApiResponse{
				Code:    fiber.StatusInternalServerError,
				Type:    "DatabaseError",
				Message: strings.Title(dberr.Error()),
			})
		}
		return c.JSON(result)
	}
}

func removePost(service post.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var post entities.DeleteRequest
		err := c.BodyParser(&post)
		postID := c.Params("post_id")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&entities.ApiResponse{
				Code:    fiber.StatusInternalServerError,
				Type:    "ParseError",
				Message: strings.Title(err.Error()),
			})
		}
		dberr := service.RemovePost(postID)
		if dberr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&entities.ApiResponse{
				Code:    fiber.StatusInternalServerError,
				Type:    "DatabaseError",
				Message: strings.Title(dberr.Error()),
			})
		}
		return c.JSON(&entities.ApiResponse{
			Code:    fiber.StatusOK,
			Type:    "Success",
			Message: fmt.Sprintf("Post with ID %s was deleted", postID),
		})
	}
}

func getPost(s post.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		postID := c.Params("post_id")
		post, err := s.FetchPost(postID)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&entities.ApiResponse{
				Code:    fiber.StatusInternalServerError,
				Type:    "DatabaseError",
				Message: strings.Title(err.Error()),
			})
		}

		return c.JSON(&post)
	}
}

func getPosts(s post.Service) fiber.Handler {
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
