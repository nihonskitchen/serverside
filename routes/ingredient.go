package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nihonskitchen/serverside/controllers"
)

// BarcodeRoute to barcode data
func BarcodeRoute(route fiber.Router) {
	log.Printf("BarcodeRoute")
	//route.Get("", controllers.GetAllBarcodes)
	route.Get("/:jancode", controllers.GetIngredientWithBarcode)
	route.Post("", controllers.CreateIngredient)
	//route.Put("/:jancode", controllers.UpdateBarcode)
	//route.Delete("/:jancode", controllers.DeleteBarcode)
}
