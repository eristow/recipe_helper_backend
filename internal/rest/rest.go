package rest

import (
	"html/template"
	"net/http"

	"github.com/eristow/recipe_helper_backend/internal/recipe"
)

func renderTemplate(w http.ResponseWriter, tmpl string, recipe *recipe.Recipe) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, recipe)
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	recipeName := r.URL.Path[len("/view/"):]
	recipeToView, err := recipe.LoadRecipe(recipeName)
	if err != nil {
		http.Redirect(w, r, "/edit/"+recipeName, http.StatusFound)
		return
	}
	renderTemplate(w, "../../internal/rest/view", recipeToView)
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	recipeName := r.URL.Path[len("/edit/"):]
	recipeToEdit, err := recipe.LoadRecipe(recipeName)
	if err != nil {
		recipeToEdit = &recipe.Recipe{Name: recipeName}
	}
	renderTemplate(w, "../../internal/rest/edit", recipeToEdit)
}

// TODO: fix this. It saves as an empty file
func SaveHandler(w http.ResponseWriter, r *http.Request) {
	recipeName := r.URL.Path[len("/save/"):]
	ingredients := r.FormValue("ingredients")
	recipeToSave := &recipe.Recipe{Name: recipeName, Ingredients: []byte(ingredients)}
	recipeToSave.SaveRecipe()
	http.Redirect(w, r, "/view/"+recipeName, http.StatusFound)
}
