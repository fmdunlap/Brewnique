package data

import (
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
	Category     RecipeCategory
	Subcategory  RecipeCategory
	Attributes   []*RecipeAttribute
	Version      int `json:"version"`
}

func ValidateRecipe(v *validator.Validator, recipe Recipe) {
	v.Check(len(recipe.Name) > 0, "name", "name is required")
	v.Check(len(recipe.Ingredients) > 0, "ingredients", "ingredients is required")
	v.Check(len(recipe.Instructions) > 0, "instructions", "instructions is required")
}
