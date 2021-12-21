package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"payment.system.com/domain"
)

func (app *application) SignupHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	type User struct {
		Iin      string `json:"iin"`
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	var user User

	err := app.readJSON(w, r, &user)

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
		app.Logger.Printf("CREATE USER ERROR %v", err)
		if errors.Is(err, domain.ErrorLoginAlreadyExists) {
			app.badRequest(w, r, err.Error())
			return
		}
		if errors.Is(err, domain.ErrorIinAlreadyExists) {
			app.badRequest(w, r, err.Error())
			return
		}
		app.serverError(w, r, err)
		return
	}

	err = app.RoleUsecases.AddRoleForUser(createdUser.Id, domain.ROLE_USER)

	if err != nil {
		app.Logger.Printf("ADD ROLE ERROR %v", err)
		app.serverError(w, r, err)
		return
	}

	createdUser.Role, err = app.RoleUsecases.GetRoleForUser(createdUser.Id)

	if err != nil {
		app.Logger.Printf("GET ROLE FOR USER ERROR %v", err)
		app.serverError(w, r, err)
		return
	}

	app.writeJson(w, http.StatusOK, envelope{"user": createdUser}, nil)
}

func (app *application) insertToken(token string, login string) error {
	key := fmt.Sprintf("user:%v", login)

	return app.redisConn.Set(key, token, app.config.Server.RefreshTtl).Err()
}

// func (app *application) findToken(login string, token string) bool {
// 	key := fmt.Sprintf("user:%v", login)

// 	value, err := app.redisConn.Get(key).Result()
// 	if err != nil {
// 		return false
// 	}

// 	return token == value
// }

func (app *application) LoginHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 1. get user login and password from request body

	type User struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	var user User

	err := app.readJSON(w, r, &user)

	if err != nil {
		app.badRequest(w, r, err.Error())
		return
	}

	// 2. get user by login and check password

	userFromDb, err := app.UserUsecases.GetUserByLogin(user.Login)

	if err != nil {
		app.Logger.Printf("Get user by login err %v", err)
		if errors.Is(err, domain.ErrorUserLoginDoesntExists) {
			app.badRequest(w, r, err.Error())
			return
		}
		app.serverError(w, r, err)
		return
	}

	// 3. if password is wrong return corresponding error

	if user.Password != userFromDb.Password {
		app.Logger.Printf("Incorrect user password")
		app.badRequest(w, r, domain.ErrorIncorrectUserPassword.Error())
		return
	}

	// 4. create jwt token

	accessTokenExp := time.Now().Add(app.config.Server.AccessTtl).Unix()
	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["id"] = userFromDb.Id
	accessTokenClaims["login"] = userFromDb.Login
	accessTokenClaims["iin"] = userFromDb.Iin
	accessTokenClaims["iat"] = time.Now().Unix()
	accessTokenClaims["exp"] = accessTokenExp
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	accessSignedToken, err := accessToken.SignedString([]byte(app.config.Server.AccessSecret))
	if err != nil {
		app.serverError(w, r, err)
		app.Logger.Printf("Coundn't create token. Error: %v", err)
		return
	}

	refreshTokenExp := time.Now().Add(app.config.Server.RefreshTtl).Unix()
	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims["id"] = userFromDb.Id
	refreshTokenClaims["login"] = userFromDb.Login
	refreshTokenClaims["iin"] = userFromDb.Iin
	refreshTokenClaims["iat"] = time.Now().Unix()
	refreshTokenClaims["exp"] = refreshTokenExp
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	refreshSignedToken, err := refreshToken.SignedString([]byte(app.config.Server.RefreshSecret))
	if err != nil {
		app.serverError(w, r, err)
		app.Logger.Printf("Coundn't create token. Error: %v", err)
		return
	}

	// 5. save token in a redis
	if err := app.insertToken(refreshSignedToken, userFromDb.Login); err != nil {
		app.serverError(w, r, err)
		app.Logger.Printf("Coundn't create token. Error: %v", err)
		return
	}
	// 6. return token
	app.writeJson(w, http.StatusAccepted, envelope{"AccessToken": "Bearer " + accessSignedToken, "RefreshToken": "Bearer " + refreshSignedToken}, nil)
}

func (app *application) parseToken(token string, isAccess bool) (string, error) {
	JWTToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("failed to extract token metadata, unexpected signing method: %v", token.Header["alg"])
		}
		if isAccess {
			return []byte(app.config.Server.AccessSecret), nil
		}
		return []byte(app.config.Server.RefreshSecret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := JWTToken.Claims.(jwt.MapClaims)

	var login string

	if ok && JWTToken.Valid {
		login, ok = claims["login"].(string)
		if !ok {
			return "", fmt.Errorf("field id not found")
		}

		exp, ok := claims["exp"].(float64)
		if !ok {
			return "", fmt.Errorf("field exp not found")
		}

		expiredTime := time.Unix(int64(exp), 0)
		log.Printf("Expired: %v", expiredTime)
		if time.Now().After(expiredTime) {
			return "", fmt.Errorf("token expired")
		}
		return login, nil
	}

	return "", fmt.Errorf("invalid token")
}

func (app *application) ValidateToken(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	token := r.Header.Get("Authorization")

	if token == "" {
		app.unauthorizedRequest(w, r, errors.New("token should be provided in a header"))
		return
	}

	tParts := strings.Split(token, " ")

	if len(tParts) != 2 {
		app.unauthorizedRequest(w, r, errors.New("token should containt two parts"))
		return
	}

	if tParts[0] != "Bearer" {
		app.unauthorizedRequest(w, r, errors.New("first part of the auth header should be \"Bearer\""))
		return
	}

	login, err := app.parseToken(tParts[1], true)
	if err != nil {
		log.Printf("Parse access token error: %v", err)
		w.WriteHeader(http.StatusMovedPermanently)
		w.Header().Add("Location", "/login")
		return
	}

	user, err := app.UserUsecases.GetUserByLogin(login)
	if err != nil {
		app.Logger.Printf("Get user error %v", err)
		app.serverError(w, r, err)
		return
	}

	app.writeJson(w, http.StatusOK, envelope{"message": "user validated", "user": user}, nil)
}
