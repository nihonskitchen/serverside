package repositories

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// Set firestore client
func SetFirestoreClient() *firestore.Client {
	//TODO  envに隠す
	sa := option.WithCredentialsFile("./nihonskitchen-firebase-adminsdk-yjuaq-eac2eb7580.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	return client
}
