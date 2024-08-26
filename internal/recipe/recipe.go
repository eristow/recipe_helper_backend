package recipe

import (
	"github.com/google/uuid"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const RecipeDir = "../../data/recipes/"

type Recipe struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Ingredients []string  `json:"ingredients"`
	Steps       []string  `json:"steps"`
}

func NewRecipe(name string, ingredients []string, steps []string) *Recipe {
	return &Recipe{
		Id:          uuid.New(),
		Name:        name,
		Ingredients: ingredients,
		Steps:       steps,
	}
}

func (r *Recipe) GetNamePretty() string {
	caser := cases.Title(language.English)
	return caser.String(r.Name)
}

func (r *Recipe) NewRecipeId() uuid.UUID {
	return uuid.New()
}

func (r *Recipe) SetId(id uuid.UUID) {
	r.Id = id
}

func (r *Recipe) SetName(name string) {
	r.Name = name
}

func (r *Recipe) SetIngredients(ingredients []string) {
	r.Ingredients = ingredients
}

func (r *Recipe) SetSteps(steps []string) {
	r.Steps = steps
}
