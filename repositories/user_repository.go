package repositories

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	//"google.golang.org/api/iterator"
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

	credentials, err := google.CredentialsFromJSON(ctx, []byte(os.Getenv("FIREBASE_KEYFILE_JSON")))
	if err != nil {
		log.Printf("error credentials from json: %v\n", err)
	}
	opt := option.WithCredentials(credentials)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalln(err)
	}

	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	defer client.Close()

	// log.Printf("repo before get client")
	// // 値の取得
	// collection := client.Collection("users2")

	// log.Printf("repo before get doc ref")
	// doc := collection.Doc("user3")

	// log.Printf("repo before doc.get")
	// field, err := doc.Get(ctx)
	// if err != nil {
	// 	fmt.Errorf("error get data: %v", err)
	// }

	// log.Printf("repo before field.data")
	// var user map[string]interface{}
	// if field != nil {
	// 	user = field.Data()
	// }
	// // for key, value := range user {
	// // 	fmt.Printf("key: %v, value: %v\n", key, value)
	// // }

	// log.Printf("repo after get user")

	//var users []map[string]
	iter := client.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		log.Printf("%+v", doc.Data())
	}

	return nil
}

// Set firestore client
func SetFirestoreClient() {
	//TODO  envに隠す
	// sa := option.WithCredentialsFile("./nihonskitchen-firebase-adminsdk-yjuaq-eac2eb7580.json")
	// app, err := firebase.NewApp(ctx, nil, sa)

	credentials, err := google.CredentialsFromJSON(ctx, []byte(os.Getenv("FIREBASE_KEYFILE_JSON")))
	if err != nil {
		log.Printf("error credentials from json: %v\n", err)
	}
	opt := option.WithCredentials(credentials)
	app, err := firebase.NewApp(ctx, nil, opt)
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
