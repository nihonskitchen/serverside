package controllers

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
	repositories "github.com/nihonskitchen/serverside/repositories"
)

// CreateRecipe is called by POST /api/recipes
func CreateRecipe(ctx *fiber.Ctx) error {
	params := new(struct {
		UserID       string        `json:"user_id"`
		RecipeName   string        `json:"recipe_name"`
		PictureURL   string        `json:"picture_url"`
		Time         string        `json:"time"`
		Likes        string        `json:"likes"`
		Dislikes     string        `json:"dislikes"`
		Prices       string        `json:"prices"`
		Servings     string        `json:"servings"`
		IsVisible    bool          `json:"is_visible"`
		OwnerComment string        `json:"owner_comment"`
		Ingredients  []interface{} `json:"ingredients"`
		Steps        []interface{} `json:"steps"`
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
		UserID:       params.UserID,
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
	docID, createdRecipe, err := repositories.SaveRecipe(targetRecipe)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":        "Created New Recipe",
		"created_recipe": createdRecipe,
		"doc_id":         docID,
	})
}

// GetAllRecipes is called by GET /api/recipes
func GetAllRecipes(ctx *fiber.Ctx) error {
	recipes := repositories.FindAllRecipes()
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "Got All Recipes",
		"data": fiber.Map{
			"recipes": recipes,
		},
	})
}

// GetAllRecipesByUID is called by GET /api/recipes/uid/:uid
func GetAllRecipesByUID(ctx *fiber.Ctx) error {
	UID := ctx.Params("uid")
	recipes := repositories.FindAllRecipesByUID(UID)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "Got All Recipes By UID",
		"data": fiber.Map{
			"recipes": recipes,
		},
	})
}

// GetAllRecipesByName is called by GET /api/recipes/name/:name
func GetAllRecipesByName(ctx *fiber.Ctx) error {
	Name := ctx.Params("name")
	recipes := repositories.FindAllRecipesByName(Name)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "Got All Recipes By Name",
		"data": fiber.Map{
			"recipes": recipes,
		},
	})
}

// GetRecipeByID is called by GET /api/recipes/:id
func GetRecipeByID(ctx *fiber.Ctx) error {
	//ドキュメントIDを渡す必要がある
	docID := ctx.Params("id")
	if docID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "No Document ID",
		})
	}
	recipe := repositories.FindRecipeByID(ctx, docID)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "Got recipe by ID",
		"data": fiber.Map{
			"recipe": recipe,
		},
	})
}
