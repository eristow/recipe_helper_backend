package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eristow/recipe_helper_backend/internal/database"
	"github.com/eristow/recipe_helper_backend/internal/recipe"
	"github.com/eristow/recipe_helper_backend/internal/rest"
)

func main() {
	mux := http.NewServeMux()
	ds := database.NewDatastore()
	recipeH := rest.NewRecipeHandler(ds)

	// Add test recipe
	pancakeRecipe := recipe.NewRecipe(
		"Pancakes",
		[]string{"Flour", "Eggs", "Milk", "Sugar"},
		[]string{"Mix ingredients", "Cook on pan"},
	)
	ds.AddRecipe("pancakes", pancakeRecipe)
	mux.Handle("/", recipeH)
	mux.Handle("/recipes/", recipeH)
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
