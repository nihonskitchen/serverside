package controllers

import (
	"github.com/gofiber/fiber/v2"
	repositories "github.com/nihonskitchen/serverside/repositories"
)

//TODO get all users
func GetUsers(c *fiber.Ctx) error {
	users := repositories.GetUser()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "成功やで",
		"data": fiber.Map{
			"users": users,
		},
	})
}
