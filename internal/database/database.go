package database

import (
	"sort"
	"sync"

	"github.com/eristow/recipe_helper_backend/internal/recipe"
)

// In-memory datastore
type Datastore struct {
	m map[string]recipe.Recipe
	*sync.RWMutex
}

func NewDatastore() *Datastore {
	return &Datastore{
		m:       make(map[string]recipe.Recipe),
		RWMutex: &sync.RWMutex{},
	}
}

func (ds *Datastore) AddRecipe(key string, r *recipe.Recipe) {
	ds.Lock()
	defer ds.Unlock()

	ds.m[key] = *r
}

func (ds *Datastore) UpdateRecipe(key string, r *recipe.Recipe) {
	ds.Lock()
	defer ds.Unlock()

	ds.m[key] = *r
}

func (ds *Datastore) GetRecipeByName(key string) (*recipe.Recipe, bool) {
	ds.RLock()
	defer ds.RUnlock()

	recipe, exists := ds.m[key]

	return &recipe, exists
}

func (ds *Datastore) GetRecipeById(key string) (*recipe.Recipe, bool) {
	ds.RLock()
	defer ds.RUnlock()

	recipe, exists := ds.m[key]

	return &recipe, exists
}

func (ds *Datastore) ListRecipes() (recipes []recipe.Recipe) {
	ds.RLock()
	defer ds.RUnlock()

	for _, r := range ds.m {
		recipes = append(recipes, r)
	}

	sort.Slice(recipes, func(i, j int) bool {
		return recipes[i].Name < recipes[j].Name
	})

	return
}

func (ds *Datastore) DeleteRecipe(key string) {
	ds.RLock()
	defer ds.RUnlock()

	delete(ds.m, key)
}
