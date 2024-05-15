package rest

import (
	"html/template"
	"net/http"
	"regexp"

	"github.com/eristow/recipe_helper_backend/internal/recipe"
)

const TemplatesDir = "../../templates/"

var templates = template.Must(template.ParseGlob(TemplatesDir + "*.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func renderTemplate(w http.ResponseWriter, tmpl string, recipe *recipe.Recipe) {
	err := templates.ExecuteTemplate(w, tmpl+".html", recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2]) // The name is the second subexpression.
	}
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/view/FrontPage", http.StatusFound)
}

func ViewHandler(w http.ResponseWriter, r *http.Request, name string) {
	recipeToView, err := recipe.LoadRecipe(name)
	if err != nil {
		http.Redirect(w, r, "/edit/"+name, http.StatusOK)
		return
	}
	renderTemplate(w, TemplatesDir+"view", recipeToView)
}

func EditHandler(w http.ResponseWriter, r *http.Request, name string) {
	recipeToEdit, err := recipe.LoadRecipe(name)
	if err != nil {
		recipeToEdit = &recipe.Recipe{Name: name}
	}
	renderTemplate(w, TemplatesDir+"edit", recipeToEdit)
}

func SaveHandler(w http.ResponseWriter, r *http.Request, name string) {
	// err := r.ParseForm()
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	ingredients := r.FormValue("ingredients")
	recipeToSave := &recipe.Recipe{Name: name, Ingredients: []byte(ingredients)}
	err := recipeToSave.SaveRecipe()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+name, http.StatusFound)
}
