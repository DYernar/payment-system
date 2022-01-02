package main

import (
	"context"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"payment.system.com/domain"
)

type userCtxKey string

const iinContextKey = userCtxKey("iin")

func (app *application) checkAuthorized(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token := r.Header.Get("Authorization")

		if token == "" {
			app.unauthorizedRequest(w, r, domain.ErrorHeaderShouldContainAuth)
			return
		}

		tParts := strings.Split(token, " ")

		if len(tParts) != 2 {
			app.unauthorizedRequest(w, r, domain.ErrorTokenShouldHaveTwoParts)
			return
		}

		if tParts[0] != "Bearer" {
			app.unauthorizedRequest(w, r, domain.ErrorTokenShouldHaveBearer)
			return
		}

		login, err := app.parseToken(tParts[1], true)

		if err != nil {
			app.unauthorizedRequest(w, r, domain.ErrorUserLoginDoesntExists)
			return
		}

		user, err := app.UserUsecases.GetUserByLogin(login)

		if err != nil {
			app.unauthorizedRequest(w, r, domain.ErrorUserLoginDoesntExists)
			return
		}

		ctx := context.WithValue(r.Context(), iinContextKey, user.Iin)
		r = r.WithContext(ctx)
		next(w, r, ps)
	}
}

func (app *application) isAdmin(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token := r.Header.Get("Authorization")

		if token == "" {
			app.unauthorizedRequest(w, r, domain.ErrorHeaderShouldContainAuth)
			return
		}

		tParts := strings.Split(token, " ")

		if len(tParts) != 2 {
			app.unauthorizedRequest(w, r, domain.ErrorTokenShouldHaveTwoParts)
			return
		}

		if tParts[0] != "Bearer" {
			app.unauthorizedRequest(w, r, domain.ErrorTokenShouldHaveBearer)
			return
		}

		login, err := app.parseToken(tParts[1], true)

		if err != nil {
			app.unauthorizedRequest(w, r, domain.ErrorUserLoginDoesntExists)
			return
		}

		user, err := app.UserUsecases.GetUserByLogin(login)

		if err != nil {
			app.unauthorizedRequest(w, r, domain.ErrorUserLoginDoesntExists)
			return
		}

		role, err := app.RoleUsecases.GetRoleForUser(user.Id)

		if err != nil {
			app.serverError(w, r, err)
			return
		}

		if role.Name != domain.ROLE_ADMIN {
			app.unauthorizedRequest(w, r, domain.ErrorYouDontHavePermissionToSeeThisData)
			return
		}

		ctx := context.WithValue(r.Context(), iinContextKey, user.Iin)
		r = r.WithContext(ctx)
		next(w, r, ps)
	}
}
