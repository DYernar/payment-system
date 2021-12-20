package domain

import "errors"

var ErrorInvalidIin = errors.New("provided iin is invalid")
var ErrorLoginAlreadyExists = errors.New("login already exists")

type User struct {
	Id       int64  `json:"id"`
	Iin      string `json:"iin"`
	Login    string `json:"email"`
	Role     Role   `json:"role"`
	Password string `json:"-"`
}

type UserUsecases interface {
	CreateUser(user *User) (*User, error)
	UpdateUser(user *User) (*User, error)
}
type UserRepository interface {
	CreateUser(user *User) (*User, error)
	UpdateUser(user *User) (*User, error)
}
