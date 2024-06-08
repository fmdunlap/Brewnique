package main

import (
	"brewnique.fdunlap.com/internal/data"
	"brewnique.fdunlap.com/internal/validator"
	"fmt"
	"net/http"
	"time"
)

func (app *application) newRecipeHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name         string   `json:"name"`
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
		Ingredients:  input.Ingredients,
		Instructions: input.Instructions,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	v := validator.New()
	data.ValidateRecipe(v, recipe)

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "I'M A NEW RECIPE WITH NAME %s, INGREDIENTS %v, INSTRUCTIONS %v", recipe.Name, recipe.Ingredients, recipe.Instructions)
}

func (app *application) listRecipesHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "I'M A LIST OF RECIPES")
}

func (app *application) getRecipeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParam(r)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	recipe := data.Recipe{
		ID:           id,
		CreatedAt:    time.Now().Add(-12 * time.Hour),
		UpdatedAt:    time.Now(),
		Name:         "My Recipe",
		Ingredients:  []string{"Eggs", "Milk", "Cheese"},
		Instructions: []string{"Boil eggs", "Add milk", "Add cheese"},
	}

	w.WriteHeader(http.StatusOK)
	err = app.writeJson(w, http.StatusOK, recipe, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
