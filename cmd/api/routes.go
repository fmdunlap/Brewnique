package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/healthcheck", app.healthcheckHandler)
		r.Post("/recipe", app.newRecipeHandler)
		r.Get("/recipes", app.listRecipesHandler)
		r.Get("/recipe/{id}", app.getRecipeHandler)
	})

	return r

}
