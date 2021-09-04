package handler

import (
	"github.com/gofiber/fiber/v2"
)

func Public(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "public", "data": nil})
}

func User(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "user", "data": nil})
}

func Admin(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "admin", "data": nil})
}

func All(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "all", "data": nil})
}

// //AddBook
// func AddBook(c *fiber.Ctx) error {
// 	book := new(models.Book)
// 	if err := c.BodyParser(book); err != nil {
// 		return c.Status(400).JSON(err.Error())
// 	}

// 	d.DB.Db.Create(&book)

// 	return c.Status(200).JSON(book)
// }

// //AllBooks
// func AllBooks(c *fiber.Ctx) error {
// 	books := []models.Book{}
// 	d.DB.Db.Find(&books)

// 	return c.Status(200).JSON(books)
// }

// //Book
// func Book(c *fiber.Ctx) error {
// 	book := []models.Book{}
// 	title := new(models.Book)
// 	if err := c.BodyParser(title); err != nil {
// 		return c.Status(400).JSON(err.Error())
// 	}
// 	d.DB.Db.Where("title = ?", title.Title).Find(&book)
// 	return c.Status(200).JSON(book)
// }

// //Update
// func Update(c *fiber.Ctx) error {
// 	book := []models.Book{}
// 	title := new(models.Book)
// 	if err := c.BodyParser(title); err != nil {
// 		return c.Status(400).JSON(err.Error())
// 	}

// 	d.DB.Db.Model(&book).Where("title = ?", title.Title).Update("author", title.Author)

// 	return c.Status(400).JSON("updated")
// }

// //Delete
// func Delete(c *fiber.Ctx) error {
// 	book := []models.Book{}
// 	title := new(models.Book)
// 	if err := c.BodyParser(title); err != nil {
// 		return c.Status(400).JSON(err.Error())
// 	}
// 	d.DB.Db.Where("title = ?", title.Title).Delete(&book)

// 	return c.Status(200).JSON("deleted")
// }
