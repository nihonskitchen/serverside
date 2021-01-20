package repositories

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/iterator"
)

// User struct the same as user collection in firestore
type Recipe struct {
	RecipeName   string        `json:"recipe_name"`
	PictureURL   string        `json:"picture_url"`
	Time         string        `json:"time"`
	Likes        string        `json:"likes"`
	Dislikes     string        `json:"dislikes"`
	Prices       string        `json:"prices"`
	Servings     string        `json:"servings"`
	IsVisible    bool          `json:"is_visible"`
	OwnerComment string        `json:"owner_comment"`
	Ingredients  []interface{} `json:"ingredients"`
	Steps        []interface{} `json:"steps"`
}

// type Material struct {
// 	Name   string `json:"name"`
// 	Amount string `json:"amount"`
// 	Unit   string `json:"unit"`
// }

const (
	recipesCollectionName string = "recipes"
)

// SaveUser create new user
func SaveRecipe(recipe Recipe) (string, Recipe, error) {
	ctx := context.Background()
	client := SetFirestoreClient()
	defer client.Close()

	// Firestore登録用にUser型からMapに変換
	// 使用するならref, resultを受け取る
	//ref, result, err := client.Collection(recipesCollectionName).Add(ctx, map[string]interface{}{
	ref, _, err := client.Collection(recipesCollectionName).Add(ctx, map[string]interface{}{
		"RecipeName":   recipe.RecipeName,
		"PictureURL":   recipe.PictureURL,
		"Time":         recipe.Time,
		"Likes":        recipe.Likes,
		"Dislikes":     recipe.Dislikes,
		"Prices":       recipe.Prices,
		"Servings":     recipe.Servings,
		"IsVisible":    recipe.IsVisible,
		"OwnerComment": recipe.OwnerComment,
		"Ingredients":  recipe.Ingredients,
		"Steps":        recipe.Steps,
	})

	if err != nil {
		fmt.Errorf("error get data: %v", err)
	}

	return ref.ID, recipe, err
}

// FindAllRecipes get all recipes
func FindAllRecipes() []Recipe {
	ctx := context.Background()
	client := SetFirestoreClient()
	// 必ずこの関数の最後でCLOSEするようにする
	defer client.Close()

	var recipes []Recipe
	iter := client.Collection(recipesCollectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		recipe := Recipe{
			RecipeName:   doc.Data()["RecipeName"].(string),
			PictureURL:   doc.Data()["PictureURL"].(string),
			Time:         doc.Data()["Time"].(string),
			Likes:        doc.Data()["Likes"].(string),
			Dislikes:     doc.Data()["Dislikes"].(string),
			Prices:       doc.Data()["Prices"].(string),
			Servings:     doc.Data()["Servings"].(string),
			IsVisible:    doc.Data()["IsVisible"].(bool),
			OwnerComment: doc.Data()["OwnerComment"].(string),
			Ingredients:  doc.Data()["Ingredients"].([]interface{}),
			Steps:        doc.Data()["Steps"].([]interface{}),
		}
		recipes = append(recipes, recipe)
	}

	return recipes
}
