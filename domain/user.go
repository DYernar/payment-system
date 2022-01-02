package domain

import (
	"errors"
	"time"
)

var ErrorInvalidIin = errors.New("provided iin is invalid")
var ErrorLoginAlreadyExists = errors.New("login already exists")
var ErrorIinAlreadyExists = errors.New("iin already exists")
var ErrorIncorrectUserPassword = errors.New("user password is wrong")
var ErrorUserLoginDoesntExists = errors.New("user login doesn't exists")
var ErrorTokenShouldHaveBearer = errors.New("first part of the token should be \"Bearer\"")
var ErrorTokenShouldHaveTwoParts = errors.New("token should contain two parts")
var ErrorHeaderShouldContainAuth = errors.New("header should contain auth token")
var ErrorYouDontHavePermissionToSeeThisData = errors.New("you dont have permissions to see this page")

type User struct {
	Id        int64     `json:"id"`
	Iin       string    `json:"iin"`
	Login     string    `json:"email"`
	Role      *Role     `json:"role"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type UserUsecases interface {
	CreateUser(user *User) (*User, error)
	UpdateUser(user *User) (*User, error)
	GetUserByLogin(login string) (*User, error)
	GetAllUsers() ([]*User, error)
}
type UserRepository interface {
	CreateUser(user *User) (*User, error)
	UpdateUser(user *User) (*User, error)
	GetUserByLogin(login string) (*User, error)
	GetAllUsers() ([]*User, error)
}
