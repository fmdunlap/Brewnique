package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "ok",
		"environment": app.config.env,
		"version":     version,
	}

	err := app.writeJson(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Printf("healthcheckHandler: error writing json: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
