package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nihonskitchen/serverside/controllers"
)

// UserRoute set User andpoints group
func UserRoute(route fiber.Router) {
	route.Get("", controllers.GetAllUsers)
	route.Get("/:id", controllers.GetUserByID)
	route.Post("", controllers.CreateUser)
	route.Put("/:id", controllers.PutUser)
	route.Delete("/:id", controllers.DeleteUser)
}
