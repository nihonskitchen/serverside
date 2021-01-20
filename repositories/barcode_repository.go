package repositories

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
)

// Barcode is the type of barcodes.
type Barcode struct {
	ID          string `json:"id"`
	Barcode     string `json:"barcode"`
	ProductName string `json:"product_name"`
	Description string `json:"description"`
	FrontPic    string `json:"front_picture"`
	BackPic     string `json:"back_picture"`
}

/*
const (
	collectionName string = "barcode"
)
*/

// FindBarcode find a single data of product by JANcode.
func FindBarcode(ctx *fiber.Ctx, docBarcode string) Barcode {

	client := SetFirestoreClient()
	// 必ずこの関数の最後でCLOSEするようにする
	defer client.Close()

	// ドキュメント名（jancode）を渡す
	// jancode := ctx.Params("jancode")

	// Firestore上のコレクション名
	collectionName := "barcode"

	// 値の取得
	collection := client.Collection(collectionName)
	doc := collection.Doc(docBarcode)
	field, err := doc.Get(context.Background())
	if err != nil {
		log.Printf("error get data: %v", err)
	}
	var barcode Barcode

	if field != nil {
		barcode = Barcode{
			ID:          field.Data()["id"].(string),
			Barcode:     field.Data()["barcode"].(string),
			ProductName: field.Data()["product_name"].(string),
			Description: field.Data()["description"].(string),
			FrontPic:    field.Data()["front_pic"].(string),
			BackPic:     field.Data()["back_pic"].(string),
		}
	}
	return barcode
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

// SaveBarcode create new user
func SaveBarcode(barcode Barcode) (Barcode, error) {
	ctx := context.Background()
	client := SetFirestoreClient()
	defer client.Close()

	// Firestore上のコレクション名
	collectionName := "barcode"

	// Firestore登録用にBarcode型からMapに変換
	// 使用するならref, resultを受け取る
	// Add()使用時はランダムな文字列がドキュメント識別子に設定される
	// ドキュメント識別子はDoc()で指定できる

	//ref, result, err := client.Collection(collectionName).Add(ctx, map[string]interface{}{
	//_, _, err := client.Collection(collectionName).Add(ctx, map[string]interface{}{
	_, err := client.Collection(collectionName).Doc(barcode.Barcode).Set(ctx, map[string]interface{}{
		"id":           barcode.ID,
		"barcode":      barcode.Barcode,
		"product_name": barcode.ProductName,
		"description":  barcode.Description,
		"front_pic":    barcode.FrontPic,
		"back_pic":     barcode.BackPic,
	})

	if err != nil {
		log.Printf("error get data: %v", err)
	}

	return barcode, err
}
