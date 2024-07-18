package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/eristow/recipe_helper_backend/internal/database"
	"github.com/eristow/recipe_helper_backend/internal/recipe"
	"github.com/eristow/recipe_helper_backend/internal/rest"
)

type slashFix struct {
	mux http.Handler
}

func (h *slashFix) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.Replace(r.URL.Path, "//", "/", -1)
	h.mux.ServeHTTP(w, r)
}

func main() {
	mux := http.NewServeMux()
	ds := database.NewDatastore()
	rootH := rest.NewRootHandler()
	recipeH := rest.NewRecipeHandler(ds)

	// Add test recipe
	pancakeRecipe := recipe.NewRecipe(
		"Pancakes",
		[]string{"Flour", "Eggs", "Milk", "Sugar"},
		[]string{"Mix ingredients", "Cook on pan"},
	)
	ds.AddRecipe("pancakes", pancakeRecipe)
	mux.Handle("/", rootH)
	mux.Handle("/recipes/", recipeH)
	mux.Handle("/recipes", recipeH)
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", &slashFix{mux}))
}
