package main

import (
	"brewnique.fdunlap.com/internal/data"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type CreateRecipeInput struct {
	Name          string   `json:"name"`
	AuthorId      int64    `json:"author_id"`
	Ingredients   []string `json:"ingredients"`
	Instructions  []string `json:"instructions"`
	CategoryId    int64    `json:"category"`
	SubcategoryId int64    `json:"subcategory"`
	Attributes    []int64  `json:"attributes"`
	Tags          []int64  `json:"tags"`
}

func (input *CreateRecipeInput) Validate() (map[string]string, error) {
	fieldErrors := make(map[string]string)

	if len(input.Name) == 0 {
		fieldErrors["name"] = "name is required"
	}
	if input.AuthorId == 0 {
		fieldErrors["author_id"] = "authorId is required"
	}
	if len(input.Ingredients) == 0 {
		fieldErrors["ingredients"] = "ingredients is required"
	}
	if len(input.Instructions) == 0 {
		fieldErrors["instructions"] = "instructions is required"
	}
	if input.CategoryId == 0 {
		fieldErrors["category_id"] = "categoryId is required"
	}

	if len(fieldErrors) > 0 {
		return fieldErrors, errors.New("validation failed")
	}

	return nil, nil
}

func (app *application) newRecipeHandler(w http.ResponseWriter, r *http.Request) {
	var input CreateRecipeInput

	err := app.readJson(w, r, &input)
	if err != nil {
		app.logError(r, err)
		app.badRequestResponse(w, r)
		return
	}

	fieldErrors, err := input.Validate()
	if err != nil {
		app.failedValidationResponse(w, r, fieldErrors)
		return
	}

	newRecipe, err := app.Services.Recipes.CreateRecipe(input.Name, input.AuthorId, input.Ingredients, input.Instructions)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	responseTags := make([]string, 0)
	for _, tag := range newRecipe.Tags {
		responseTags = append(responseTags, tag.Name)
	}

	response := data.RecipeAPIResponse{
		Id:           newRecipe.Id,
		Name:         newRecipe.Name,
		Author:       newRecipe.Author,
		Ingredients:  newRecipe.Ingredients,
		Instructions: newRecipe.Instructions,
		Category:     newRecipe.Category,
		Subcategory:  newRecipe.Subcategory,
		Attributes:   newRecipe.Attributes,
		Tags:         responseTags,
		CreatedAt:    newRecipe.CreatedAt,
		UpdatedAt:    newRecipe.UpdatedAt,
	}

	w.WriteHeader(http.StatusOK)
	err = app.writeJson(w, http.StatusOK, response, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listRecipesHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")

	fmt.Println(userId)

	if userId != "" {
		userIdInt, err := strconv.ParseInt(userId, 10, 64)
		if err != nil {
			app.badRequestResponse(w, r)
			return
		}

		recipes, err := app.Services.Recipes.ListRecipesByAuthorId(userIdInt)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		w.WriteHeader(http.StatusOK)
		app.writeJson(w, http.StatusOK, recipes, nil)
		return
	}

	recipes, err := app.Services.Recipes.ListRecipes()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	app.writeJson(w, http.StatusOK, recipes, nil)
}

func (app *application) getRecipeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
	}

	recipe, err := app.Services.Recipes.GetRecipe(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			app.notFoundResponse(w, r)
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = app.writeJson(w, http.StatusOK, recipe, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteRecipeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParam(r)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	err = app.Services.Recipes.DeleteRecipe(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	w.WriteHeader(http.StatusOK)
	app.writeJson(w, http.StatusOK, map[string]string{"status": "ok"}, nil)
}

func (app *application) listCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := app.Services.Recipes.ListCategories()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	app.writeJson(w, http.StatusOK, categories, nil)
}

func (app *application) listSubcategoriesHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
	}

	subcategories, err := app.Services.Recipes.ListSubcategories(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	app.writeJson(w, http.StatusOK, subcategories, nil)
}

func (app *application) listAttributesHandler(w http.ResponseWriter, r *http.Request) {
	attributes, err := app.Services.Recipes.GetAttributes()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	app.writeJson(w, http.StatusOK, attributes, nil)
}

func (app *application) listAttributeValuesHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
	}

	attributeValues, err := app.Services.Recipes.GetAttributeValues(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	app.writeJson(w, http.StatusOK, attributeValues, nil)
}

func (app *application) listTagsHandler(w http.ResponseWriter, r *http.Request) {
	tags, err := app.Services.Recipes.GetTags()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	app.writeJson(w, http.StatusOK, tags, nil)
}
