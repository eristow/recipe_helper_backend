// TODO: Fix the tests

package recipe

// import (
// 	"os"
// 	"reflect"
// 	"testing"
// )

// const ChocolateCakeFileLocation = RecipeDir + "chocolate_cake.txt"

// func TestSaveRecipe(t *testing.T) {
// 	// Create a recipe
// 	r := &Recipe{
// 		Name:        "chocolate_cake",
// 		Ingredients: []byte("Flour, Sugar, Cocoa Powder"),
// 	}

// 	// Save the recipe
// 	err := r.SaveRecipe()
// 	if err != nil {
// 		t.Errorf("SaveRecipe() returned an error: %v", err)
// 	}

// 	// Check if the file exists
// 	if _, err := os.Stat(ChocolateCakeFileLocation); os.IsNotExist(err) {
// 		t.Errorf("SaveRecipe() did not create the file")
// 	}

// 	// Check if the file content matches the expected values
// 	if _, err := os.ReadFile(ChocolateCakeFileLocation); err != nil {
// 		t.Errorf("SaveRecipe() did not write the expected content")
// 	}

// 	// Clean up
// 	os.Remove(ChocolateCakeFileLocation)
// }

// func TestLoadRecipe(t *testing.T) {
// 	name := "chocolate_cake"
// 	ingredients := []byte("Flour, Sugar, Cocoa Powder")

// 	// Create the file
// 	err := os.WriteFile(ChocolateCakeFileLocation, ingredients, 0600)
// 	if err != nil {
// 		t.Errorf("Failed to create the file: %v", err)
// 	}

// 	// Load the recipe
// 	r, err := LoadRecipe(name)
// 	if err != nil {
// 		t.Errorf("LoadRecipe() returned an error: %v", err)
// 	}

// 	// Check if r.Name matches the expected value
// 	if r.Name != name {
// 		t.Errorf("LoadRecipe() did not return the expected name")
// 	}

// 	// Check if r.Ingredients matches the expected value
// 	if !reflect.DeepEqual(r.Ingredients, ingredients) {
// 		t.Errorf("LoadRecipe() did not return the expected ingredients")
// 	}

// 	// Clean up
// 	os.Remove(ChocolateCakeFileLocation)

// 	// Test for non-existent file
// 	_, err = LoadRecipe("Non-existent Recipe")
// 	if err == nil {
// 		t.Errorf("LoadRecipe() did not return an error for a non-existent recipe")
// 	}
// }
