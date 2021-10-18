package main

import (
	"log"

	"git.snrd.de/serverlist.tf/backend/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func root(c *fiber.Ctx) error {
	return c.SendString("You are viewing the serverlist.tf API")
}

func middleware(c *fiber.Ctx) error {
	return c.Next()
}

func setupApiRoutes(api fiber.Router) {
	api.Get("/servers", handlers.GetServers)
	api.Get("/gamemodes", handlers.GetGamemodes)
	api.Get("/regions", handlers.GetRegions)
	api.Get("/flag", handlers.GetFlag)

	serverApi := api.Group("/server", middleware)
	setupServerRoutes(serverApi)

	gamemodeApi := api.Group("/gamemode", middleware)
	setupGamemodeRoutes(gamemodeApi)

	regionApi := api.Group("/region", middleware)
	setupRegionRoutes(regionApi)
}

func setupServerRoutes(api fiber.Router) {
	api.Get("/", handlers.GetServer)
	api.Post("/", handlers.NewServer)
	api.Delete("/", handlers.DeleteServer)
}

func setupGamemodeRoutes(api fiber.Router) {
	api.Get("/", handlers.GetGamemode)
	api.Post("/", handlers.NewGamemode)
	api.Delete("/", handlers.DeleteGamemode)
}

func setupRegionRoutes(api fiber.Router) {
	api.Get("/", handlers.GetRegion)
	api.Post("/", handlers.NewRegion)
	api.Delete("/", handlers.DeleteRegion)
}

func defaultRoute(c *fiber.Ctx) error {
	return c.SendStatus(404)
}

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cache.New()) // We want our API to be cache 1m (default is 1m)

	app.Get("/", root)

	api := app.Group("/api/v1", middleware)
	setupApiRoutes(api)

	app.Use(defaultRoute)

	log.Fatal(app.Listen(":3000"))
}
