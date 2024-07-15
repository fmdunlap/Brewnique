package data

import (
	"time"

	"brewnique.fdunlap.com/internal/validator"
)

type Recipe struct {
	Id            int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	AuthorId      int64
	Author        User
	Name          string
	Ingredients   []string
	Instructions  []string
	Category      string
	CategoryId    int64
	Subcategory   string
	SubcategoryId int64
	Attributes    []AttributeDTO
	Tags          []Tag
	Version       int
}

type RecipeAPIResponse struct {
	Id           int64          `json:"id"`
	Name         string         `json:"name"`
	Author       User           `json:"author"`
	Ingredients  []string       `json:"ingredients"`
	Instructions []string       `json:"instructions"`
	Category     string         `json:"category"`
	Subcategory  string         `json:"subcategory"`
	Attributes   []AttributeDTO `json:"attributes"`
	Tags         []string       `json:"tags"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type AttributeDTO struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type RecipeCategory struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	ParentId *int64 `json:"parent_id"`
}

type RecipeRating struct {
	Id       int64 `json:"id"`
	UserId   int64 `json:"user_id"`
	RecipeId int64 `json:"recipe_id"`
	Rating   int   `json:"rating"`
}

type RecipeTag struct {
	Id       int64  `json:"id"`
	RecipeId int64  `json:"recipe_id"`
	TagId    int64  `json:"tag_id"`
	Name     string `json:"name"`
}

func ValidateRecipe(v *validator.Validator, recipe Recipe) {
	v.Check(len(recipe.Name) > 0, "name", "name is required")
	v.Check(len(recipe.Ingredients) > 0, "ingredients", "ingredients is required")
	v.Check(len(recipe.Instructions) > 0, "instructions", "instructions is required")
}
