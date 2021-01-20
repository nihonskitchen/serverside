package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	repositories "github.com/nihonskitchen/serverside/repositories"
)

/*
// GetAllIngredients is called by GET /api/barcode
func GetAllIngredients(ctx *fiber.Ctx) error {
	ingredients := repositories.FindAllIngredients()
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "成功やで",
		"data": fiber.Map{
			"ingredients": ingredients,
		},
	})
}
*/

//GetIngredientWithBarcode is called by GET /api/barcode/:jancode
func GetIngredientWithBarcode(ctx *fiber.Ctx) error {
	// ドキュメント名を渡す
	docBarcode := ctx.Params("jancode")

	if docBarcode == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "No Document of the JANcode",
		})
	}

	ingredient, isInDB := repositories.FindIngredientByBarcode(ctx, docBarcode)

	// データベースに対象バーコードが登録されていない場合
	if !isInDB {
		return ctx.Status(404).JSON(fiber.Map{
			"success":      false,
			"message":      "Document of the JANcode is NOT in Database",
			"barcode_data": ingredient.BarcodeData,
		})
	}

	// データベースに対象バーコードが登録されている場合
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "Got the Ingredient by JANcode",
		"data": fiber.Map{
			"ingredient": ingredient,
		},
	})
}

// CreateIngredient is called by POST /api/barcode
func CreateIngredient(ctx *fiber.Ctx) error {
	params := new(struct {
		IngredientID   string
		BarcodeData    string
		IngredientName string
		Description    string
		//FrontPic    string
		//BackPic     string
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

	if len(params.BarcodeData) == 0 || len(params.IngredientName) == 0 {
		return ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "BarcodeData or IngredientName not specified.",
		})
	}

	targetIngredient := repositories.Ingredient{BarcodeData: params.BarcodeData, IngredientName: params.IngredientName, Description: params.Description}
	createdIngredient, err := repositories.SaveIngredient(targetIngredient)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":           "Created New Ingredient Data with Barcode",
		"createdIngredient": createdIngredient,
	})
}
