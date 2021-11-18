package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

type application struct {
	config config
	Logger log.Logger
}

func NewApplication(config *config) *application {
	return &application{
		Logger: *log.New(os.Stdout, time.Now().String()+" : ", 0),
		config: *config,
	}
}

func (app *application) run() {
	app.Logger.Printf(" Running at port %v\n", app.config.Server.Port)
	err := http.ListenAndServe(":"+strconv.Itoa(app.config.Server.Port), app.router())
	if err != nil {
		app.Logger.Fatalf(err.Error())
	}
}

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("HELLO WORLD"))
}
