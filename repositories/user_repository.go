package repositories

import (
	"context"
	"fmt"
	//"google.golang.org/api/iterator"
)

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

	client := SetFirestoreClient()
	defer client.Close()

	// 値の取得
	collection := client.Collection("users")
	doc := collection.Doc("user2")
	field, err := doc.Get(context.Background())
	if err != nil {
		fmt.Errorf("error get data: %v", err)
	}
	var user map[string]interface{}
	if field != nil {
		user = field.Data()
	}
	// for key, value := range user {
	// 	fmt.Printf("key: %v, value: %v\n", key, value)
	// }

	return user
}
