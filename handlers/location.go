package handlers

import "github.com/gofiber/fiber/v2"

func GetFlag(c *fiber.Ctx) error {
	return c.SendString("germany")
}
