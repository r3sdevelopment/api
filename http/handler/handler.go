package handler

import "github.com/gofiber/fiber/v2"

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
