package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nihonskitchen/serverside/controllers"
)

func UserRoute(route fiber.Router) {
	log.Printf("Route")
	route.Get("", controllers.GetUsers)
}
