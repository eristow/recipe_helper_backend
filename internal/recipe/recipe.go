package recipe

import (
	"os"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const RecipeDir = "../../data/recipes/"

type Recipe struct {
	Name string
	// Ingredients []string
	// Steps       []string
	Ingredients []byte
}

func (r *Recipe) PrettyName() string {
	caser := cases.Title(language.English)
	return caser.String(r.Name)
}

func (r *Recipe) SaveRecipe() error {
	filename := r.Name + ".txt"
	return os.WriteFile(RecipeDir+filename, r.Ingredients, 0600)
}

func LoadRecipe(name string) (*Recipe, error) {
	filename := name + ".txt"
	body, err := os.ReadFile(RecipeDir + filename)
	if err != nil {
		return nil, err
	}
	return &Recipe{Name: name, Ingredients: body}, nil
}
