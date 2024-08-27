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

	// Add test recipes
	pancakeRecipe := recipe.NewRecipe(
		"Pancakes",
		[]string{"Flour", "Eggs", "Milk", "Sugar"},
		[]string{"Mix ingredients", "Cook on pan"},
	)
	pizzaRecipe := recipe.NewRecipe(
		"Pizza",
		[]string{"Flour", "Yeast", "Sugar", "Olive Oil", "Salt", "Tomato Sauce", "Mozzarella Cheese", "Toppings"},
		[]string{"Mix warm water, yeast, sugar. Cover and allow to rest for 5 min.", "Add flour, salt, olive oil. Mix until dough forms.", "Knead dough for 5 min.", "Cover and allow to rise for 1 hour.", "Preheat oven to 475F.", "Roll out dough.", "Add sauce, cheese, toppings.", "Bake for 13-15 min."},
	)

	ds.AddRecipe(pancakeRecipe.Id.String(), pancakeRecipe)
	ds.AddRecipe(pizzaRecipe.Id.String(), pizzaRecipe)

	mux.Handle("/", rest.HandleCors(rootH))
	mux.Handle("/recipes/", rest.HandleCors(recipeH))
	mux.Handle("/recipes", rest.HandleCors(recipeH))
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", &slashFix{mux}))
}
