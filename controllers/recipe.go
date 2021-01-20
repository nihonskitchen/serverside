package controllers

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
	repositories "github.com/nihonskitchen/serverside/repositories"
)

// CreateRecipe is called by POST /api/recipes
func CreateRecipe(ctx *fiber.Ctx) error {
	params := new(struct {
		RecipeName   string                    `json:"recipe_name"`
		PictureURL   string                    `json:"picture_url"`
		Time         string                    `json:"time"`
		Likes        string                    `json:"likes"`
		Dislikes     string                    `json:"dislikes"`
		Prices       string                    `json:"prices"`
		Servings     string                    `json:"servings"`
		IsVisible    bool                      `json:"is_visible"`
		OwnerComment string                    `json:"owner_comment"`
		Ingredients  []repositories.Ingredient `json:"ingredients"`
		Steps        []string                  `json:"steps"`
	})
	err := ctx.BodyParser(&params)

	// if error
	if err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	// if len(params.ID) == 0 || len(params.Name) == 0 {
	// 	return ctx.Status(400).JSON(fiber.Map{
	// 		"success": false,
	// 		"error":   "ID or name not specified.",
	// 	})
	// }

	targetRecipe := repositories.Recipe{
		RecipeName:   params.RecipeName,
		PictureURL:   params.PictureURL,
		Time:         params.Time,
		Likes:        params.Likes,
		Dislikes:     params.Dislikes,
		Prices:       params.Prices,
		Servings:     params.Servings,
		IsVisible:    params.IsVisible,
		OwnerComment: params.OwnerComment,
		Ingredients:  params.Ingredients,
		Steps:        params.Steps}
	createdRecipe, err := repositories.SaveRecipe(targetRecipe)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":       "Created New Recipe",
		"createdRecipe": createdRecipe,
	})
}