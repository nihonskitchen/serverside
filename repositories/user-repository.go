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

// Define firestore client & context
var client *firestore.Client
var ctx = context.Background()

//TODO 使えてない
type User struct {
	Name string `json:"name"`
}

// TODO まだ全ユーザー取れてない
func GetUser() map[string]interface{} {
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

	// 値の取得
	collection := client.Collection("users")
	doc := collection.Doc("user2")
	field, err := doc.Get(ctx)
	if err != nil {
		fmt.Errorf("error get data: %v", err)
	}
	user := field.Data()
	for key, value := range user {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}

	return user
}

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