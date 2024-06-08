package data

import (
	"brewnique.fdunlap.com/internal/validator"
	"time"
)

type Recipe struct {
	ID           int64     `json:"id"`
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
