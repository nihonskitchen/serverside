package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/iterator"
)

// User struct the same as user collection in firestore
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

const (
	collectionName string = "users"
)

// FindUserByID find user by id
func FindUserByID(ctx *fiber.Ctx, docID string) User {
	client := SetFirestoreClient()
	// 必ずこの関数の最後でCLOSEするようにする
	defer client.Close()

	// 値の取得
	collection := client.Collection(collectionName)
	doc := collection.Doc(docID)
	field, err := doc.Get(context.Background())
	if err != nil {
		fmt.Errorf("error get data: %v", err)
	}
	var user User
	//TODO 現状ないものを取得した場合落ちる
	if field != nil {
		user = User{
			ID:   field.Data()["ID"].(string),
			Name: field.Data()["Name"].(string),
		}
	}
	return user
}

// FindAllUsers get all users
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

// SaveUser create new user
func SaveUser(user User) (User, error) {
	ctx := context.Background()
	client := SetFirestoreClient()
	defer client.Close()

	// Firestore登録用にUser型からMapに変換
	// 使用するならref, resultを受け取る
	//ref, result, err := client.Collection(collectionName).Add(ctx, map[string]interface{}{
	_, _, err := client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":   user.ID,
		"Name": user.Name,
	})

	if err != nil {
		fmt.Errorf("error get data: %v", err)
	}

	return user, err
}

// UpdateUser update user
func UpdateUser(docID string, user User) (User, error) {
	ctx := context.Background()
	client := SetFirestoreClient()
	defer client.Close()

	_, err := client.Collection(collectionName).Doc(docID).Set(ctx, user)

	if err != nil {
		fmt.Errorf("error get data: %v", err)
	}

	return user, err
}

// DeleteUserByID delete user by id
func DeleteUserByID(docID string) error {
	ctx := context.Background()
	client := SetFirestoreClient()
	// 必ずこの関数の最後でCLOSEするようにする
	defer client.Close()

	// 値の取得
	_, err := client.Collection(collectionName).Doc(docID).Delete(ctx)
	if err != nil {
		fmt.Errorf("error get data: %v", err)
	}
	return nil
}
