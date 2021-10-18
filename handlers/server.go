package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetServers(c *fiber.Ctx) error {
	return c.SendString("All Servers")
}

func GetServer(c *fiber.Ctx) error {
	return c.SendString("A single Server")
}

func NewServer(c *fiber.Ctx) error {
	return c.SendString("Adds a new Server")
}

func DeleteServer(c *fiber.Ctx) error {
	return c.SendString("Deletes a Server")
}
