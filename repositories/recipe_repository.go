package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/iterator"
)

// User struct the same as user collection in firestore
type Recipe struct {
	UserID       string        `json:"user_id"`
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

type RecipeWithDocID struct {
	DocID string `json:"doc_id"`
	Recipe
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
		"UserID":       recipe.UserID,
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
func FindAllRecipes() []RecipeWithDocID {
	ctx := context.Background()
	client := SetFirestoreClient()
	// 必ずこの関数の最後でCLOSEするようにする
	defer client.Close()

	var recipes []RecipeWithDocID
	iter := client.Collection(recipesCollectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		recipe := RecipeWithDocID{
			DocID: doc.Ref.ID,
			Recipe: Recipe{
				UserID:       doc.Data()["UserID"].(string),
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
				// Ingredients:  doc.Data()["Ingredients"].([]Material{
				// 	Name:doc.Data()["Ingredients"]["Name"],
				// 	Amount:doc.Data()["Ingredients"]["Amount"],
				// 	Unit:doc.Data()["Ingredients"]["Unit"],
				// }),
				Steps: doc.Data()["Steps"].([]interface{}),
			},
		}
		recipes = append(recipes, recipe)
	}

	return recipes
}

// FindAllRecipesByUID get all recipes by uid
func FindAllRecipesByUID(UID string) []RecipeWithDocID {
	ctx := context.Background()
	client := SetFirestoreClient()
	// 必ずこの関数の最後でCLOSEするようにする
	defer client.Close()

	var recipes []RecipeWithDocID
	iter := client.Collection(recipesCollectionName).Where("UserID", "==", UID).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		recipe := RecipeWithDocID{
			DocID: doc.Ref.ID,
			Recipe: Recipe{
				UserID:       doc.Data()["UserID"].(string),
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
				// Ingredients:  doc.Data()["Ingredients"].([]Material{
				// 	Name:doc.Data()["Ingredients"]["Name"],
				// 	Amount:doc.Data()["Ingredients"]["Amount"],
				// 	Unit:doc.Data()["Ingredients"]["Unit"],
				// }),
				Steps: doc.Data()["Steps"].([]interface{}),
			},
		}
		recipes = append(recipes, recipe)
	}

	return recipes
}

// FindRecipeByID find recipe by id
func FindRecipeByID(ctx *fiber.Ctx, docID string) Recipe {
	client := SetFirestoreClient()
	defer client.Close()

	// 値の取得
	collection := client.Collection(recipesCollectionName)
	doc := collection.Doc(docID)
	docRef, err := doc.Get(context.Background())
	if err != nil {
		fmt.Errorf("error get data: %v", err)
	}
	var recipe Recipe
	//TODO 現状ないものを取得した場合落ちる
	if docRef != nil {
		recipe = Recipe{
			UserID:       docRef.Data()["UserID"].(string),
			RecipeName:   docRef.Data()["RecipeName"].(string),
			PictureURL:   docRef.Data()["PictureURL"].(string),
			Time:         docRef.Data()["Time"].(string),
			Likes:        docRef.Data()["Likes"].(string),
			Dislikes:     docRef.Data()["Dislikes"].(string),
			Prices:       docRef.Data()["Prices"].(string),
			Servings:     docRef.Data()["Servings"].(string),
			IsVisible:    docRef.Data()["IsVisible"].(bool),
			OwnerComment: docRef.Data()["OwnerComment"].(string),
			Ingredients:  docRef.Data()["Ingredients"].([]interface{}),
			Steps:        docRef.Data()["Steps"].([]interface{}),
		}
	}
	return recipe
}
