package controllers

import (
	fiber "github.com/gofiber/fiber/v2"
	repositories "github.com/nihonskitchen/serverside/repositories"
)

// GetAllUsers is called by GET /api/users
func GetAllUsers(ctx *fiber.Ctx) error {
	users := repositories.FindAllUsers()
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "成功やで",
		"data": fiber.Map{
			"users": users,
		},
	})
}

// GetUserByID is called by GET /api/users/:id
func GetUserByID(ctx *fiber.Ctx) error {
	user := repositories.FindUserByID(ctx)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "成功やで",
		"data": fiber.Map{
			"users": user,
		},
	})
}

// //POST /api/users
// func CreateUser(ctx *fiber.Ctx) error {
// 	type Request struct {
// 		DishName string `json:"dish_name"`
// 	}

// 	var body Request

// 	err := ctx.BodyParser(&body)

// 	// if error
// 	if err != nil {
// 		fmt.Println(err)
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"success": false,
// 			"message": "Cannot parse JSON",
// 		})
// 	}

// 	// create a recipe variable
// 	recipe := &Recipe{
// 		DishID:   len(recipes) + 1,
// 		DishName: body.DishName,
// 	}

// 	// append in recipes
// 	recipes = append(recipes, recipe)

// 	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
// 		"success": true,
// 		"data": fiber.Map{
// 			"recipe": recipe,
// 		},
// 	})
// }

// func UpdateUser(ctx *fiber.Ctx) error {
// 	// find parameter
// 	paramID := c.Params("id")

// 	// convert parameter string to init
// 	id, err := strconv.Atoi(paramID)

// 	// if error when parsing string to int
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"success": false,
// 			"message": "Cannot parse Id",
// 		})
// 	}

// 	// request structure
// 	type Request struct {
// 		DishName *string `json:"dish_name"`
// 	}

// 	var body Request
// 	err = c.BodyParser(&body)

// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"success": false,
// 			"message": "Cannot parse JSON",
// 		})
// 	}

// 	var recipe *Recipe

// 	for _, r := range recipes {
// 		if r.DishID == id {
// 			recipe = r
// 			break
// 		}
// 	}

// 	if recipe.DishID == 0 {
// 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
// 			"success": false,
// 			"message": "Not found",
// 		})
// 	}

// 	if body.DishName != nil {
// 		recipe.DishName = *body.DishName
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"success": true,
// 		"data": fiber.Map{
// 			"recipe": recipe,
// 		},
// 	})
// }

// //DELETE /api/users/:id
// func DeleteUser(ctx *fiber.Ctx) error {
// 	// get param
// 	paramID := c.Params("id")

// 	// convert param string to int
// 	id, err := strconv.Atoi(paramID)

// 	// if parameter cannot parse
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"success": false,
// 			"message": "Cannot parse id",
// 		})
// 	}

// 	// find and delete recipe
// 	for i, recipe := range recipes {
// 		if recipe.DishID == id {

// 			recipes = append(recipes[:i], recipes[i+1:]...)

// 			return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
// 				"success": true,
// 				"message": "Deleted Successfully",
// 			})
// 		}
// 	}

// 	// if recipe not found
// 	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
// 		"success": false,
// 		"message": "Recipe not found",
// 	})
// }
