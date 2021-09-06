package router

import (
	"api/pkg/entities"
	"api/pkg/post"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const POSTS_PATH = "/posts"

func PostRouter(r fiber.Router, s post.Service) {
	r.Get(POSTS_PATH, getPosts(s))
	r.Post(POSTS_PATH, addPost(s))
	r.Put(POSTS_PATH+"/:id", updatePost(s))
	r.Delete(POSTS_PATH, removePost(s))
}

func addPost(s post.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var post entities.Post

		c.BodyParser(&post)

		if err := post.Validate(); err != nil {
			c.Status(400)
			return c.JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}

		result, dberr := s.InsertPost(&post)
		return c.JSON(&fiber.Map{
			"status": result,
			"error":  dberr,
		})
	}
}

func updatePost(service post.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var post entities.Post

		if err := c.BodyParser(&post); err != nil {
			c.Status(400)
			return c.JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}

		if postId, err := uuid.Parse(c.Params("id")); err == nil {
			fmt.Printf("postID", postId)
		}

		if err := post.Validate(); err != nil {
			c.Status(400)
			return c.JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}
		result, dberr := service.UpdatePost(&post)
		return c.JSON(&fiber.Map{
			"status": result,
			"error":  dberr,
		})
	}
}

func removePost(service post.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var post entities.DeleteRequest
		err := c.BodyParser(&post)
		postID := post.ID
		if err != nil {
			_ = c.JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}
		dberr := service.RemovePost(postID)
		if dberr != nil {
			_ = c.JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}
		return c.JSON(&fiber.Map{
			"status":  false,
			"message": "updated successfully",
		})
	}
}

func getPosts(s post.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := s.FetchPosts()
		var result fiber.Map
		if err != nil {
			result = fiber.Map{
				"status": false,
				"error":  err.Error(),
			}
		} else {
			result = fiber.Map{
				"status": true,
				"posts":  fetched,
			}
		}
		return c.JSON(&result)
	}
}
