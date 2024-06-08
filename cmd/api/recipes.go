package main

import (
	"fmt"
	"net/http"
)

func (app *application) newRecipeHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name         string   `json:"name"`
		Ingredients  []string `json:"ingredients"`
		Instructions []string `json:"instructions"`
	}

	err := app.readJson(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "I'M A NEW RECIPE WITH NAME %s, INGREDIENTS %v, INSTRUCTIONS %v", input.Name, input.Ingredients, input.Instructions)
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

	w.WriteHeader(http.StatusOK)
	err = app.writeJson(w, http.StatusOK, map[string]any{
		"id":           fmt.Sprintf("%d", id),
		"name":         "My Recipe",
		"ingredients":  []string{"Eggs", "Milk", "Cheese"},
		"instructions": []string{"Boil eggs", "Add milk", "Add cheese"},
	}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
