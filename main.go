package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"

	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./nihonskitchen-firebase-adminsdk-yjuaq-eac2eb7580.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	fmt.Println("sucsess")
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	iter := client.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}

	defer client.Close()
}
