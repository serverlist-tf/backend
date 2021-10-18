package handlers

import "github.com/gofiber/fiber/v2"

func GetGamemodes(c *fiber.Ctx) error {
	return c.SendString("All gamemodes")
}

func GetGamemode(c *fiber.Ctx) error {
	return c.SendString("A single gamemode")
}

func NewGamemode(c *fiber.Ctx) error {
	return c.SendString("Add a gamemode")
}

func DeleteGamemode(c *fiber.Ctx) error {
	return c.SendString("Delete a gamemode")
}
