package main

import "net/http"

func (app *application) logError(r *http.Request, err error) {
	app.logger.PrintError(err, &map[string]string{
		"request_method": r.Method,
		"request_url":    r.URL.String(),
	})
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	data := map[string]any{
		"status":      "error",
		"environment": app.config.env,
		"version":     version,
		"error":       message,
	}

	err := app.writeJson(w, status, data, nil)
	if err != nil {
		app.logError(r, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "Internal Server Error"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource was not found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request) {
	message := "bad request"
	app.errorResponse(w, r, http.StatusBadRequest, message)
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := "method not allowed"
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.errorResponse(w, r, http.StatusBadRequest, errors)

}
