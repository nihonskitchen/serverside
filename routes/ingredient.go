package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nihonskitchen/serverside/controllers"
)

// IngredientRoute to barcode data
func IngredientRoute(route fiber.Router) {
	route.Get("", controllers.GetIngredients)
}
