package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"brewnique.fdunlap.com/internal/data"
	"brewnique.fdunlap.com/internal/validator"
)

func (app *application) newRecipeHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name         string   `json:"name"`
		AuthorId     int64    `json:"author_id"`
		Ingredients  []string `json:"ingredients"`
		Instructions []string `json:"instructions"`
	}

	err := app.readJson(w, r, &input)
	if err != nil {
		app.logError(r, err)
		app.badRequestResponse(w, r)
		return
	}

	recipe := data.Recipe{
		Name:         input.Name,
		AuthorId:     input.AuthorId,
		Ingredients:  input.Ingredients,
		Instructions: input.Instructions,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Version:      1,
	}

	v := validator.New()
	data.ValidateRecipe(v, recipe)

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	newRecipe, err := app.Services.Recipes.CreateRecipe(recipe.Name, recipe.AuthorId, recipe.Ingredients, recipe.Instructions)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = app.writeJson(w, http.StatusOK, newRecipe, nil)
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
