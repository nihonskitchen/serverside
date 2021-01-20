package repositories

import (
	"context"
	"fmt"
)

// User struct the same as user collection in firestore
type Recipe struct {
	RecipeName   string       `json:"recipe_name"`
	PictureURL   string       `json:"picture_url"`
	Time         string       `json:"time"`
	Likes        string       `json:"likes"`
	Dislikes     string       `json:"dislikes"`
	Prices       string       `json:"prices"`
	Servings     string       `json:"servings"`
	IsVisible    bool         `json:"is_visible"`
	OwnerComment string       `json:"owner_comment"`
	Ingredients  []Ingredient `json:"ingredients"`
	Steps        []string     `json:"steps"`
}

type Ingredient struct {
	Name   string `json:"name"`
	Amount string `json:"amount"`
	Unit   string `json:"unit"`
}

const (
	recipesCollectionName string = "recipes"
)

// SaveUser create new user
func SaveRecipe(recipe Recipe) (Recipe, error) {
	ctx := context.Background()
	client := SetFirestoreClient()
	defer client.Close()

	// Firestore登録用にUser型からMapに変換
	// 使用するならref, resultを受け取る
	//ref, result, err := client.Collection(recipesCollectionName).Add(ctx, map[string]interface{}{
	_, _, err := client.Collection(recipesCollectionName).Add(ctx, map[string]interface{}{
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

	return recipe, err
}
