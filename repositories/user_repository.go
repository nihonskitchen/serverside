package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/iterator"
	//"google.golang.org/api/iterator"
)

//TODO 使えてない
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

const (
	collectionName string = "users-test"
)

// FindUserByID find user by id
func FindUserByID(ctx *fiber.Ctx) User {
	client := SetFirestoreClient()
	// 必ずこの関数の最後でCLOSEするようにする
	defer client.Close()

	//ドキュメントIDを渡す必要がある
	id := ctx.Params("id")

	// 値の取得
	collection := client.Collection(collectionName)
	doc := collection.Doc(id)
	field, err := doc.Get(context.Background())
	if err != nil {
		fmt.Errorf("error get data: %v", err)
	}
	var user User
	if field != nil {
		user = User{
			ID:   field.Data()["ID"].(string),
			Name: field.Data()["Name"].(string),
		}
	}
	return user
}

// FindAllUser get all users
func FindAllUsers() []User {
	ctx := context.Background()
	client := SetFirestoreClient()
	// 必ずこの関数の最後でCLOSEするようにする
	defer client.Close()

	var users []User
	iter := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		user := User{
			ID:   doc.Data()["ID"].(string),
			Name: doc.Data()["Name"].(string),
		}
		users = append(users, user)
	}

	return users
}

// func SaveUser(ctx *fiber.Ctx) map[string]interface{} {
// 	client := SetFirestoreClient()
// 	defer client.Close()

// 	// 値の取得
// 	ref, result, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
// 		"ID":   User.Id,
// 		"Name": User.Name,
// 	})

// 	if err != nil {
// 		fmt.Errorf("error get data: %v", err)
// 	}

// 	return nil
// }
