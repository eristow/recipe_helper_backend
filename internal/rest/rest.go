package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/eristow/recipe_helper_backend/internal/database"
	"github.com/eristow/recipe_helper_backend/internal/recipe"
	"github.com/google/uuid"
)

var (
	listRecipeRe   = regexp.MustCompile(`^\/recipes\/*$`)
	getRecipeRe    = regexp.MustCompile(`^\/recipes\/([a-zA-Z0-9\-]+)\/?$`)
	createRecipeRe = regexp.MustCompile(`^\/recipes\/*$`)
)

type RootHandler struct{}
type RecipeHandler struct {
	store *database.Datastore
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
}

func HandleCors(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			enableCors(&w)
			w.WriteHeader(http.StatusOK)
		} else {
			h.ServeHTTP(w, r)
		}
	}
}

func NewRootHandler() *RootHandler {
	return &RootHandler{}
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
	// TODO: make this return a list of endpoints? or just use swagger/OpenAPI?
	enableCors(&w)
	log.Printf("Root: %s", r.Method)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the Recipe Helper Backend!"))
}

func (h *RecipeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	log.Printf("Recipes router: %s", r.Method)
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
	case r.Method == http.MethodPut && getRecipeRe.MatchString(r.URL.Path):
		h.Update(w, r)
		return
	case r.Method == http.MethodDelete && getRecipeRe.MatchString(r.URL.Path):
		h.Delete(w, r)
		return
	default:
		notFound(w, r)
		return
	}
}

func (h *RecipeHandler) List(w http.ResponseWriter, r *http.Request) {
	log.Println("List")
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
	log.Println("Get")
	recipeId := getRecipeNameIdFromUrl(r)
	log.Printf("Recipe ID: %s", recipeId)
	if recipeId == "" {
		notFound(w, r)
		return
	}

	recipe, exists := h.store.GetRecipeById(recipeId)

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
	log.Println("Create")
	var newRecipe recipe.Recipe
	if err := json.NewDecoder(r.Body).Decode(&newRecipe); err != nil {
		internalServerError(w, r)
		return
	}

	newRecipe.SetId(newRecipe.NewRecipeId())

	log.Printf("Adding new recipe: %+v", newRecipe)

	h.store.AddRecipe(newRecipe.Id.String(), &newRecipe)

	recipeJsonBytes, err := json.Marshal(newRecipe)
	if err != nil {
		internalServerError(w, r)
		return
	}

	log.Printf("Added new recipe: %+v", newRecipe)

	w.WriteHeader(http.StatusCreated)
	w.Write(recipeJsonBytes)
}

func (h *RecipeHandler) Update(w http.ResponseWriter, r *http.Request) {
	log.Println("Update")

	recipeId := getRecipeNameIdFromUrl(r)

	var updatedRecipe recipe.Recipe
	if err := json.NewDecoder(r.Body).Decode(&updatedRecipe); err != nil {
		internalServerError(w, r)
		return
	}

	id, err := uuid.Parse(recipeId)
	if err != nil {
		internalServerError(w, r)
		return
	}
	updatedRecipe.SetId(id)

	log.Printf("Updating recipe: %+v", updatedRecipe)

	h.store.UpdateRecipe(recipeId, &updatedRecipe)

	recipeJsonBytes, err := json.Marshal(updatedRecipe)
	if err != nil {
		internalServerError(w, r)
		return
	}

	log.Printf("Updated recipe: %+v", updatedRecipe)

	w.WriteHeader(http.StatusOK)
	w.Write(recipeJsonBytes)
}

func (h *RecipeHandler) Delete(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete")

	recipeId := getRecipeNameIdFromUrl(r)

	log.Printf("Deleting recipe: %s", recipeId)

	h.store.DeleteRecipe(recipeId)

	log.Printf("Deleted recipe: %s", recipeId)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted recipe: " + recipeId))
}

func internalServerError(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 internal server error"))
}

func notFound(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 not found"))
}
