package rest

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/eristow/recipe_helper_backend/internal/database"
	"github.com/eristow/recipe_helper_backend/internal/recipe"
)

var (
	listRecipeRe   = regexp.MustCompile(`^\/recipes\/*$`)
	getRecipeRe    = regexp.MustCompile(`^\/recipes\/([a-zA-Z0-9]+)\/?$`)
	createRecipeRe = regexp.MustCompile(`^\/recipes\/*$`)
)

type RootHandler struct{}
type RecipeHandler struct {
	store *database.Datastore
}

func NewRecipeHandler(store *database.Datastore) *RecipeHandler {
	return &RecipeHandler{store: store}
}

func getRecipeNameIdFromUrl(r *http.Request) string {
	matches := getRecipeRe.FindStringSubmatch(r.URL.Path)
	if matches == nil || len(matches) < 2 {
		return ""
	}
	return matches[1]
}

func (h *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: make this return a list of endpoints?
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the Recipe Helper Backend!"))
}

func (h *RecipeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && listRecipeRe.MatchString(r.URL.Path):
		h.List(w, r)
		return
	case r.Method == http.MethodGet && getRecipeRe.MatchString(r.URL.Path):
		h.Get(w, r)
		return
	case r.Method == http.MethodPost && createRecipeRe.MatchString(r.URL.Path):
		h.Create(w, r)
		return
	default:
		notFound(w, r)
		return
	}
}

func (h *RecipeHandler) List(w http.ResponseWriter, r *http.Request) {
	recipes := h.store.ListRecipes()

	recipesJsonBytes, err := json.Marshal(recipes)
	if err != nil {
		internalServerError(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(recipesJsonBytes)
}

func (h *RecipeHandler) Get(w http.ResponseWriter, r *http.Request) {
	recipeName := getRecipeNameIdFromUrl(r)
	if recipeName == "" {
		notFound(w, r)
		return
	}

	recipe, exists := h.store.GetRecipe(recipeName)

	if !exists {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("recipe not found"))
		return
	}

	recipeJsonBytes, err := json.Marshal(recipe)
	if err != nil {
		internalServerError(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(recipeJsonBytes)
}

func (h *RecipeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newRecipe recipe.Recipe
	if err := json.NewDecoder(r.Body).Decode(&newRecipe); err != nil {
		internalServerError(w, r)
		return
	}

	h.store.AddRecipe(newRecipe.Name, &newRecipe)

	recipeJsonBytes, err := json.Marshal(newRecipe)
	if err != nil {
		internalServerError(w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(recipeJsonBytes)
}

func internalServerError(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 internal server error"))
}

func notFound(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 not found"))
}
