package data

import (
	"encoding/json"
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

func (r *Recipe) MarshalApiResponse() ([]byte, error) {
	return json.Marshal(r.ToRecipeAPIResponse())
}

func (r *Recipe) ToRecipeAPIResponse() RecipeAPIResponse {
	tagStrings := make([]string, 0)
	for _, tag := range r.Tags {
		tagStrings = append(tagStrings, tag.Name)
	}
	return RecipeAPIResponse{
		Id:   r.Id,
		Name: r.Name,
		Author: User{
			Id:       r.AuthorId,
			Username: r.Author.Username,
			Email:    r.Author.Email,
		},
		Ingredients:  r.Ingredients,
		Instructions: r.Instructions,
		Category:     r.Category,
		Subcategory:  r.Subcategory,
		Attributes:   r.Attributes,
		Tags:         tagStrings,
		CreatedAt:    r.CreatedAt,
		UpdatedAt:    r.UpdatedAt,
	}
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

type Category struct {
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
