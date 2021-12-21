package main

import "net/http"

// badRequest is used to return bad requests
func (app *application) badRequest(w http.ResponseWriter, r *http.Request, err string) {
	app.writeJson(w, http.StatusBadRequest, envelope{"message": "bad request: " + err}, nil)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	app.writeJson(w, http.StatusNotFound, envelope{"message": "requested data not found!"}, nil)
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	app.writeJson(w, http.StatusInternalServerError, envelope{"message": "internal server error ", "error": err}, nil)
}

func (app *application) unauthorizedRequest(w http.ResponseWriter, r *http.Request, err error) {
	if err == nil {
		app.writeJson(w, http.StatusUnauthorized, envelope{"message": "user is not authorized"}, nil)
		return
	}
	app.writeJson(w, http.StatusUnauthorized, envelope{"message": err.Error()}, nil)
}

func (app *application) methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	app.writeJson(w, http.StatusMethodNotAllowed, envelope{"message": r.Method + " is not implemented"}, nil)
}
