package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func (app *application) readIdParam(r *http.Request) (int64, error) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id param")
	}

	return id, nil
}

func (app *application) errorResponse(w http.ResponseWriter, status int, err error) {
	data := map[string]string{
		"status":      "error",
		"environment": app.config.env,
		"version":     version,
		"error":       err.Error(),
	}

	err = app.writeJson(w, status, data, nil)
	if err != nil {
		app.logger.Printf("errorResponse: error writing json: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) readJson(w http.ResponseWriter, r *http.Request, data any) error {
	r.Body = http.MaxBytesReader(w, r.Body, app.config.maxBodySize)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&data)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains invalid JSON at character %d", syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return fmt.Errorf("body contains invalid JSON")

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Value != "" {
				return fmt.Errorf("body contains invalid JSON for field %s", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains invalid JSON at character %d", unmarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			return fmt.Errorf("body must not be empty")
		case strings.HasPrefix(err.Error(), "json: unknown field"):
			FieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unkown key: %s", FieldName)
		case err.Error() == "http: request body too large":
			return fmt.Errorf("body must not be larger than %d bytes", app.config.maxBodySize)
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		default:
			return err
		}
	}

	err = decoder.Decode(&struct{}{})
	if err != nil {
		return errors.New("body must contain a single JSON object")
	}

	return nil
}

func (app *application) badRequestResponse(w http.ResponseWriter, err error) {
	app.errorResponse(w, http.StatusBadRequest, err)
}

func (app *application) writeJson(w http.ResponseWriter, status int, data any, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header().Set(key, value[0])
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
