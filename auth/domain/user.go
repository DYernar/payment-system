package domain

import (
	"errors"
	"time"
)

var ErrorInvalidIin = errors.New("provided iin is invalid")
var ErrorLoginAlreadyExists = errors.New("login already exists")
var ErrorIinAlreadyExists = errors.New("iin already exists")
var ErrorIncorrectUserPassword = errors.New("user password is wrong!")
var ErrorUserLoginDoesntExists = errors.New("user login doesn't exists!")

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
}
type UserRepository interface {
	CreateUser(user *User) (*User, error)
	UpdateUser(user *User) (*User, error)
	GetUserByLogin(login string) (*User, error)
}
