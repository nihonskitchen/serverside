package repositories

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/iterator"
)

// User struct the same as users collection in firestore without Favorites
type User struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
}

// UserWithFavorites struct the same as users collection in firestore
type UserWithFavorites struct {
	User
	Favorites []interface{} `json:"favorites"`
}

const (
	usersCollectionName string = "users"
)

// FindUserByID find user by id
func FindUserByID(ctx *fiber.Ctx, docID string) User {
	client := SetFirestoreClient()
	// 必ずこの関数の最後でCLOSEするようにする
	defer client.Close()

	// 値の取得
	collection := client.Collection(usersCollectionName)
	doc := collection.Doc(docID)
	field, err := doc.Get(context.Background())
	if err != nil {
		fmt.Errorf("error get data: %v", err)
		//TODO 現状ないものを取得した場合落ちる
	}
	var user User

	if field != nil {
		user = User{
			UID:  field.Data()["UID"].(string),
			Name: field.Data()["Name"].(string),
			//Favorites: field.Data()["Favorite"].([]interface{}),
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
	iter := client.Collection(usersCollectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		user := User{
			UID:  doc.Data()["UID"].(string),
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
	//ref, result, err := client.Collection(usersCollectionName).Add(ctx, map[string]interface{}{
	_, err := client.Collection(usersCollectionName).Doc(user.UID).Set(ctx, map[string]interface{}{
		"UID":  user.UID,
		"Name": user.Name,
	}, firestore.MergeAll)

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

	_, err := client.Collection(usersCollectionName).Doc(docID).Set(ctx, user)

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
	_, err := client.Collection(usersCollectionName).Doc(docID).Delete(ctx)
	if err != nil {
		fmt.Errorf("error get data: %v", err)
	}
	return nil
}

// FindFavoritesByUID find user by id
func FindFavoritesByUID(ctx *fiber.Ctx, docID string) UserWithFavorites {
	client := SetFirestoreClient()
	// 必ずこの関数の最後でCLOSEするようにする
	defer client.Close()

	// 値の取得
	collection := client.Collection(usersCollectionName)
	doc := collection.Doc(docID)
	field, err := doc.Get(context.Background())
	if err != nil {
		fmt.Errorf("error get data: %v", err)
		//TODO 現状ないものを取得した場合落ちる
	}
	var favorites UserWithFavorites

	if field != nil {
		favorites = UserWithFavorites{
			Favorites: field.Data()["Favorites"].([]interface{}),
			User: User{
				UID:  field.Data()["UID"].(string),
				Name: field.Data()["Name"].(string),
			},
		}
	}
	return favorites
}
