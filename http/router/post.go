package router

import (
	"api/pkg/entities"
	"api/pkg/post"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const POSTS_PATH = "/posts"

func PostRouter(r fiber.Router, s post.Service) {
	r.Get(POSTS_PATH, getPosts(s))
	r.Get(POSTS_PATH+"/:post_id", getPost(s))
	r.Post(POSTS_PATH, addPost(s))
	r.Put(POSTS_PATH+"/:post_id", updatePost(s))
	r.Delete(POSTS_PATH+"/:post_id", removePost(s))
}

func addPost(s post.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var post entities.Post

		c.BodyParser(&post)

		if err := post.Validate(); err != nil {
			c.Status(400)
			return c.JSON(&entities.ApiResponse{
				Code:    400,
				Type:    "ValidationError",
				Message: strings.Title(err.Error()),
			})
		}

		result, dberr := s.InsertPost(&post)

		if dberr != nil {
			c.Status(400)
			return c.JSON(&entities.ApiResponse{
				Code:    400,
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
			c.Status(400)
			return c.JSON(&entities.ApiResponse{
				Code:    400,
				Type:    "ParsingError",
				Message: strings.Title(err.Error()),
			})
		}
		post.ID = postID
		if err := post.Validate(); err != nil {
			c.Status(400)
			return c.JSON(&entities.ApiResponse{
				Code:    400,
				Type:    "ValidationError",
				Message: strings.Title(err.Error()),
			})
		}
		result, dberr := service.UpdatePost(&post)

		if dberr != nil {
			c.Status(400)
			return c.JSON(&entities.ApiResponse{
				Code:    400,
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
			_ = c.JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}
		dberr := service.RemovePost(postID)
		if dberr != nil {
			c.Status(400)
			return c.JSON(&entities.ApiResponse{
				Code:    400,
				Type:    "DatabaseError",
				Message: strings.Title(dberr.Error()),
			})
		}
		return c.JSON(&entities.ApiResponse{
			Code:    200,
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
			c.Status(400)
			return c.JSON(&entities.ApiResponse{
				Code:    400,
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
			c.Status(400)
			return c.JSON(&entities.ApiResponse{
				Code:    400,
				Type:    "DatabaseError",
				Message: strings.Title(err.Error()),
			})
		}

		return c.JSON(&posts)
	}
}
