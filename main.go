package main

import (
	"log"
	"os"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	repositories "github.com/nihonskitchen/serverside/repositories"
	routes "github.com/nihonskitchen/serverside/routes"
)

func main() {
	// set firestore client
	repositories.SetFirestoreClient()

	// set Server things
	app := fiber.New()
	app.Use(logger.New())
	setupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Printf("Listening on port %s\n\n", port)

	err := app.Listen("nihons-kitchen-server.herokuapp.com:" + port)

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
