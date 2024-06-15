package data

import (
	"fmt"
	"time"

	"brewnique.fdunlap.com/internal/validator"
)

type Recipe struct {
	Id           int64     `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	AuthorId     int64     `json:"author_id"`
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
	ListRecipesByAuthorId(userId int64) ([]Recipe, error)
}

type RecipeService struct {
	recipeProvider RecipeProvider
}

func NewRecipeService(recipeProvider RecipeProvider) *RecipeService {
	return &RecipeService{
		recipeProvider: recipeProvider,
	}
}

func (s *RecipeService) CreateRecipe(name string, authorId int64, ingredients []string, instructions []string) (Recipe, error) {
	if name == "" {
		return Recipe{}, fmt.Errorf("name is not set")
	}
	if authorId == 0 {
		return Recipe{}, fmt.Errorf("authorId is not set")
	}
	if len(ingredients) == 0 {
		return Recipe{}, fmt.Errorf("ingredients is not set")
	}
	if len(instructions) == 0 {
		return Recipe{}, fmt.Errorf("instructions is not set")
	}

	return s.recipeProvider.PutRecipe(Recipe{
		Name:         name,
		Ingredients:  ingredients,
		Instructions: instructions,
		AuthorId:     authorId,
	})
}

func (s *RecipeService) GetRecipe(id int64) (Recipe, error) {
	return s.recipeProvider.GetRecipe(id)
}

func (s *RecipeService) GetUserRecipes(userId int64) ([]Recipe, error) {
	return s.recipeProvider.ListRecipesByAuthorId(userId)
}

func (s *RecipeService) ListRecipes() ([]Recipe, error) {
	return s.recipeProvider.ListRecipes()
}

func (s *RecipeService) UpdateRecipe(recipe Recipe) (Recipe, error) {
	return s.recipeProvider.UpdateRecipe(recipe)
}

func (s *RecipeService) DeleteRecipe(id int64) error {
	return s.recipeProvider.DeleteRecipe(id)
}
