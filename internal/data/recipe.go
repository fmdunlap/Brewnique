package data

import (
	"time"

	"brewnique.fdunlap.com/internal/validator"
)

type Recipe struct {
	Id           int64     `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string    `json:"name"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	Version      int       `json:"version"`
}

func ValidateRecipe(v *validator.Validator, recipe Recipe) {
	v.Check(len(recipe.Name) > 0, "name", "name is required")
	v.Check(len(recipe.Ingredients) > 0, "ingredients", "ingredients is required")
	v.Check(len(recipe.Instructions) > 0, "instructions", "instructions is required")
}

type RecipeProvider interface {
	GetRecipe(
		d int64) (Recipe, error)
	ListRecipes() ([]Recipe, error)
	PutRecipe(recipe Recipe) (Recipe, error)
	UpdateRecipe(recipe Recipe) (Recipe, error)
	DeleteRecipe(id int64) error
}

type RecipeService struct {
	recipeProvider RecipeProvider
}

func NewRecipeService(recipeProvider RecipeProvider) *RecipeService {
	return &RecipeService{
		recipeProvider: recipeProvider,
	}
}

func (s *RecipeService) GetRecipe(id int64) (Recipe, error) {
	return s.recipeProvider.GetRecipe(id)
}

func (s *RecipeService) ListRecipes() ([]Recipe, error) {
	return s.recipeProvider.ListRecipes()
}

func (s *RecipeService) CreateRecipe(recipe Recipe) (Recipe, error) {
	return s.recipeProvider.PutRecipe(recipe)
}

func (s *RecipeService) UpdateRecipe(recipe Recipe) (Recipe, error) {
	return s.recipeProvider.UpdateRecipe(recipe)
}

func (s *RecipeService) DeleteRecipe(id int64) error {
	return s.recipeProvider.DeleteRecipe(id)
}
