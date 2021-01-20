package repositories

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
)

// Ingredient is the type of ingredients.
type Ingredient struct {
	IngredientID   string `json:"ingredient_id"`
	BarcodeData    string `json:"barcode_data"`
	IngredientName string `json:"product_name"`
	Description    string `json:"description"`
	//FrontPic    string `json:"front_picture"`
	//BackPic     string `json:"back_picture"`
}

/*
const (
	collectionName string = "ingredient"
)
*/

// FindIngredientByBarcode finds a single data of ingredient by barcode.
func FindIngredientByBarcode(ctx *fiber.Ctx, docBarcode string) Ingredient {

	client := SetFirestoreClient()
	// 必ずこの関数の最後でCLOSEするようにする
	defer client.Close()

	// Firestore上のコレクション名
	collectionName := "ingredient"

	// 値の取得
	collection := client.Collection(collectionName)
	doc := collection.Doc(docBarcode)
	field, err := doc.Get(context.Background())
	if err != nil {
		log.Printf("error get data: %v", err)
	}
	var ingredient Ingredient

	if field != nil {
		ingredient = Ingredient{
			IngredientID:   field.Data()["ingredient_id"].(string),
			BarcodeData:    field.Data()["barcode_data"].(string),
			IngredientName: field.Data()["ingredient_name"].(string),
			Description:    field.Data()["description"].(string),
			//FrontPic:    field.Data()["front_pic"].(string),
			//BackPic:     field.Data()["back_pic"].(string),
		}
	}
	return ingredient
}

/*
// FindAllBarcodes get all users
func FindAllBarcodes() []Barcode {
	ctx := context.Background()
	client := SetFirestoreClient()
	// 必ずこの関数の最後でCLOSEするようにする
	defer client.Close()

	var barcodes []Barcode
	iter := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		barcode := Barcode{
			ID:          doc.Data()["id"].(string),
			Barcode:     doc.Data()["barcode"].(string),
			ProductName: doc.Data()["product_name"].(string),
			Description: doc.Data()["description"].(string),
			FrontPic:    doc.Data()["front_pic"].(string),
			BackPic:     doc.Data()["back_pic"].(string),
		}
		barcodes = append(barcodes, barcode)
	}

	return barcodes
}
*/

// SaveIngredient create new user
func SaveIngredient(ingredient Ingredient) (Ingredient, error) {
	ctx := context.Background()
	client := SetFirestoreClient()
	defer client.Close()

	// Firestore上のコレクション名
	collectionName := "ingredient"

	// Firestore登録用にBarcode型からMapに変換
	// 使用するならref, resultを受け取る
	// Add()使用時はランダムな文字列がドキュメント識別子に設定される
	// ドキュメント識別子はDoc()で指定できる

	//ref, result, err := client.Collection(collectionName).Add(ctx, map[string]interface{}{
	//_, _, err := client.Collection(collectionName).Add(ctx, map[string]interface{}{
	_, err := client.Collection(collectionName).Doc(ingredient.BarcodeData).Set(ctx, map[string]interface{}{
		"ingredient_id":   ingredient.IngredientID,
		"barcode_data":    ingredient.BarcodeData,
		"ingredient_name": ingredient.IngredientName,
		"description":     ingredient.Description,
		//"front_pic":    ingredient.FrontPic,
		//"back_pic":     ingredient.BackPic,
	})

	if err != nil {
		log.Printf("error get data: %v", err)
	}

	return ingredient, err
}
