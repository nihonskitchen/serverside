package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/takunyan/go-learn/controllers"
)

func UserRoute(route fiber.Router) {
	route.Get("", controllers.GetUsers)
}
