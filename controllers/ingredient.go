package controllers

import (
	"github.com/gofiber/fiber/v2"
	repositories "github.com/nihonskitchen/serverside/repositories"
)

//GetIngredients TODO get all ingredients
func GetIngredients(c *fiber.Ctx) error {
	ingredients := repositories.GetIngredient()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "成功やで",
		"data": fiber.Map{
			"ingredients": ingredients,
		},
	})
}
