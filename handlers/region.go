package handlers

import "github.com/gofiber/fiber/v2"

func GetRegions(c *fiber.Ctx) error {
	return c.SendString("All regions")
}

func GetRegion(c *fiber.Ctx) error {
	return c.SendString("A single region")
}

func NewRegion(c *fiber.Ctx) error {
	return c.SendString("Add a region")
}

func DeleteRegion(c *fiber.Ctx) error {
	return c.SendString("Delete a region")
}
