package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func applyMiddleware(r chi.Router) chi.Router {
	r.Use(middleware.Logger)
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	return r
}

func (app *application) routes() chi.Router {
	r := applyMiddleware(chi.NewRouter())

	r.NotFound(app.notFoundResponse)
	r.MethodNotAllowed(app.methodNotAllowedResponse)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/healthcheck", app.healthcheckHandler)
		r.Post("/recipe", app.newRecipeHandler)
		r.Get("/recipes", app.listRecipesHandler)
		r.Get("/recipe/{id}", app.getRecipeHandler)
	})

	return r

}
