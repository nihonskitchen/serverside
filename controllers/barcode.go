package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	repositories "github.com/nihonskitchen/serverside/repositories"
)

/*
// GetAllBarcodes is called by GET /api/barcode
func GetAllBarcodes(ctx *fiber.Ctx) error {
	barcodes := repositories.FindAllBarcodes()
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "成功やで",
		"data": fiber.Map{
			"barcodes": barcodes,
		},
	})
}
*/

//GetBarcode is called by GET /api/barcode/:jancode
func GetBarcode(ctx *fiber.Ctx) error {
	//log.Printf("controller before get barcode")

	barcode := repositories.FindBarcode(ctx)

	//log.Printf("controller after get barcode")

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "成功やで",
		"data": fiber.Map{
			"barcode": barcode,
		},
	})
}

// CreateBarcode is called by POST /api/barcode
func CreateBarcode(ctx *fiber.Ctx) error {
	params := new(struct {
		ID          string
		Barcode     string
		ProductName string
		Description string
		FrontPic    string
		BackPic     string
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

	if len(params.Barcode) == 0 || len(params.ProductName) == 0 {
		return ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Barcode or ProductName not specified.",
		})
	}

	targetBarcode := repositories.Barcode{Barcode: params.Barcode, ProductName: params.ProductName}
	createdBarcode, err := repositories.SaveBarcode(targetBarcode)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"ok":             true,
		"createdBarcode": createdBarcode,
	})
}
