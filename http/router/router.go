package router

import (
	"api/db"
	"api/models"
	"api/http/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetUpRoutes(server *fiber.App) {
		// Middleware
	api := server.Group("/api", logger.New())
	api.Get("/public", handler.Public)
	api.Get("/user", handler.User)
	api.Get("/admin", handler.Admin)
	api.Get("/all", handler.All)

	server.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
}


//AddBook
func AddBook(c *fiber.Ctx) error {
	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Db.Create(&book)

	return c.Status(200).JSON(book)
}

//AllBooks
func AllBooks(c *fiber.Ctx) error {
	books := []models.Book{}
	database.DB.Db.Find(&books)

	return c.Status(200).JSON(books)
}

//Book
func Book(c *fiber.Ctx) error {
	book := []models.Book{}
	title := new(models.Book)
	if err := c.BodyParser(title); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DB.Db.Where("title = ?", title.Title).Find(&book)
	return c.Status(200).JSON(book)
}

//Update
func Update(c *fiber.Ctx) error {
	book := []models.Book{}
	title := new(models.Book)
	if err := c.BodyParser(title); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Db.Model(&book).Where("title = ?", title.Title).Update("author", title.Author)

	return c.Status(400).JSON("updated")
}

//Delete
func Delete(c *fiber.Ctx) error {
	book := []models.Book{}
	title := new(models.Book)
	if err := c.BodyParser(title); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DB.Db.Where("title = ?", title.Title).Delete(&book)

	return c.Status(200).JSON("deleted")
}
