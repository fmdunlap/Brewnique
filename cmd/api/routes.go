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
		AllowedMethods:   []string{"GET", "POST", "DELETE"},
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
		app.addRecipeRoutes(r)
		app.addUserRoutes(r)
		app.addCommentRoutes(r)
	})

	return r

}

func (app *application) addUserRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", app.listUsersHandler)
		r.Get("/{id}", app.getUserHandler)
		r.Get("/email/{email}", app.getUserByEmailHandler)
		r.Get("/username/{username}", app.getUserByUsernameHandler)
		r.Post("/", app.newUserHandler)
		r.Delete("/{id}", app.deleteUserHandler)
	})
}

func (app *application) addRecipeRoutes(r chi.Router) {
	r.Route("/recipes", func(r chi.Router) {
		r.Get("/", app.listRecipesHandler)
		r.Get("/{id}", app.getRecipeHandler)
		r.Post("/", app.newRecipeHandler)
		r.Delete("/{id}", app.deleteRecipeHandler)
		r.Get("/categories", app.listCategoriesHandler)
		r.Get("/categories/{id}", app.listSubcategoriesHandler)
		r.Get("/attributes", app.listAttributesHandler)
		r.Get("/attributes/{id}", app.listAttributeValuesHandler)
		r.Get("/tags", app.listTagsHandler)
	})
}

func (app *application) addCommentRoutes(r chi.Router) {
	r.Route("/comments", func(r chi.Router) {
		r.Get("/{id}", app.getCommentHandler)
		r.Get("/recipe/{id}", app.listRecipeCommentsHandler)
		r.Post("/", app.newCommentHandler)
	})
}
