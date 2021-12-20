package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/julienschmidt/httprouter"
	"payment.system.com/domain"
	"payment.system.com/repository"
	"payment.system.com/usecases"
)

type application struct {
	config       config
	Logger       log.Logger
	UserUsecases domain.UserUsecases
	RoleUsecases domain.RoleUsecases
	redisConn    *redis.Client
}

func NewApplication(db *sql.DB, config *config) *application {

	userRepo := repository.NewPgUserRepository(db)

	roleRepo := repository.NewPgRoleRepo(db)

	return &application{
		Logger:       *log.New(os.Stdout, time.Now().String()+" : ", 0),
		config:       *config,
		UserUsecases: usecases.NewUserUsecases(userRepo),
		RoleUsecases: usecases.NewRoleUsecases(roleRepo),
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
