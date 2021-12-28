package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) router() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowed)

	router.GET("/healthcheck", app.healthcheck)
	router.POST("/signup", app.SignupHandler)
	router.POST("/login", app.LoginHandler)
	router.POST("/validate", app.ValidateToken)
	router.GET("/user/:login", app.GetUser)

	return router
}
