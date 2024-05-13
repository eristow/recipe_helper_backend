package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eristow/recipe_helper_backend/internal/recipe"
	"github.com/eristow/recipe_helper_backend/internal/rest"
)

func main() {
	http.HandleFunc("/view/", rest.ViewHandler)
	http.HandleFunc("/edit/", rest.EditHandler)
	http.HandleFunc("/save/", rest.SaveHandler)
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func oldMain() {
	fmt.Println("Hello, World!")

	recipe1 := recipe.Recipe{
		Name:        "pasta",
		Ingredients: []byte("Pasta, Tomato, Salt, Water"),
	}

	err := recipe1.SaveRecipe()
	if err != nil {
		fmt.Println(err)
	}

	recipe2, err := recipe.LoadRecipe("pasta")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(recipe2.Ingredients))
}
