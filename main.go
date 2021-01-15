package main

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	repositories "github.com/takunyan/go-learn/repositories"
	routes "github.com/takunyan/go-learn/routes"
)

func main() {
	// set firestore client
	repositories.SetFirestoreClient()

	// set Server things
	app := fiber.New()
	app.Use(logger.New())
	setupRoutes(app)

	err := app.Listen(":8000")

	if err != nil {
		panic(err)
	}
}

func setupRoutes(app *fiber.App) {

	// give response when at /
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "rootやで",
		})
	})

	// api group
	api := app.Group("/api")

	// give response when at /api
	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the api endpoint 😉",
		})
	})

	// connect todo routes
	routes.UserRoute(api.Group("/users"))
}
