package repositories

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"

	//"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// // Define firestore client & context
// var client *firestore.Client
// var ctx = context.Background()

// Ingredient is the type of ingredients.
type Ingredient struct {
	ID          int    `json:"id"`
	Barcode     string `json:"barcode"`
	ProductName string `json:"product_name"`
	FrontPic    string `json:"front_picture"`
	BackPic     string `json:"back_picture"`
	Description string `json:"description"`
}

// TODO まだ全ユーザー取れてない

// GetIngredient returns a ingredient.
func GetIngredient() map[string]interface{} {

	// Set firestore client
	//TODO  envに隠す
	sa := option.WithCredentialsFile("./nihonskitchen-firebase-adminsdk-yjuaq-eac2eb7580.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	//TODO closeをどうするか？
	//defer client.Close()

	//var users []map[string]
	// iter := client.Collection("users").Documents(ctx)
	// for {
	// 	doc, err := iter.Next()
	// 	if err == iterator.Done {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatalf("Failed to iterate: %v", err)
	// 	}
	// 	fmt.Println(doc.Data())
	// }

	// Define firestore client & context
	var client *firestore.Client
	var ctx = context.Background()

	// 値の取得
	collection := client.Collection("barcode")
	doc := collection.Doc("ltkpKJchM3iiOAJTl5HK")
	field, err := doc.Get(ctx)
	if err != nil {
		fmt.Errorf("error get data: %v", err)
	}
	ingredient := field.Data()
	for key, value := range ingredient {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}

	return ingredient
}

/*
	// Set firestore client
func SetFirestoreClient() {
	//TODO  envに隠す
	sa := option.WithCredentialsFile("./nihonskitchen-firebase-adminsdk-yjuaq-eac2eb7580.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	//TODO closeをどうするか？
	//defer client.Close()
}
*/
