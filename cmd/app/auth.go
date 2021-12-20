package main

import (
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"payment.system.com/domain"
)

func (app *application) SignupHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	app.Logger.Println("SIGNUP CALLED")
	type User struct {
		Iin      string `json:"iin"`
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	var user User

	err := app.readJSON(w, r, &user)

	app.Logger.Printf("PARSED USER %v", user)

	if err != nil {
		app.badRequest(w, r, err.Error())
		return
	}

	if !app.IsValidIIN(user.Iin) {
		app.badRequest(w, r, domain.ErrorInvalidIin.Error())
		return
	}

	var newUser domain.User

	newUser.Iin = user.Iin
	newUser.Login = user.Login
	newUser.Password = user.Password

	createdUser, err := app.UserUsecases.CreateUser(&newUser)

	if err != nil {
		if errors.Is(err, domain.ErrorLoginAlreadyExists) {
			app.badRequest(w, r, err.Error())
			return
		}
		app.serverError(w, r, err)
		return
	}

	app.writeJson(w, http.StatusOK, envelope{"user": createdUser}, nil)
}

func (app *application) LoginHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {}

func (app *application) ValidateToken(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {}
