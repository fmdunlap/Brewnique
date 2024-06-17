package main

import (
	"database/sql"
	"errors"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func (app *application) newUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	err := app.readJson(w, r, &input)
	if err != nil {
		app.logError(r, err)
		app.badRequestResponse(w, r)
		return
	}

	newUser, err := app.Services.Users.CreateNewUser(input.Email, input.Username)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = app.writeJson(w, http.StatusOK, newUser, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
	}

	user, err := app.Services.Users.GetUser(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			app.notFoundResponse(w, r)
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = app.writeJson(w, http.StatusOK, user, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getUserByEmailHandler(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	if email == "" {
		app.badRequestResponse(w, r)
		return
	}

	user, err := app.Services.Users.GetUserByEmail(email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			app.notFoundResponse(w, r)
			return
		}
		log.Printf("error getting user by email %s: %s", email, err)
		app.serverErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = app.writeJson(w, http.StatusOK, user, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getUserByUsernameHandler(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	if username == "" {
		app.badRequestResponse(w, r)
		return
	}

	user, err := app.Services.Users.GetUserByUsername(username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			app.notFoundResponse(w, r)
			return
		}
		log.Printf("error getting user by username %s: %s", username, err)
		app.serverErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = app.writeJson(w, http.StatusOK, user, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// TODO: add pagination
func (app *application) listUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := app.Services.Users.ListUsers()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	app.writeJson(w, http.StatusOK, users, nil)
}

// TODO: Route protection
func (app *application) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParam(r)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	err = app.Services.Users.DeleteUser(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	w.WriteHeader(http.StatusOK)
	app.writeJson(w, http.StatusOK, map[string]string{"status": "ok"}, nil)
}
