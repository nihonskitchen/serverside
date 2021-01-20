package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nihonskitchen/serverside/controllers"
)

// UserRoute set User andpoints group
func RecipeRoute(route fiber.Router) {
	route.Get("", controllers.GetAllRecipes)
	// route.Get("/:id", controllers.GetUserByID)
	route.Post("", controllers.CreateRecipe)
	// route.Put("/:id", controllers.PutUser)
	// route.Delete("/:id", controllers.DeleteUser)
}
