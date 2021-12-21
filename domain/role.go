package domain

import "errors"

var ErrorRoleExists = errors.New("role already exists in a database")
var ErrorRoleDoesntExists = errors.New("role doesn't exists")
var ErrorUserAlreadyHasSameRole = errors.New("user already has same role")

// we will have two roles admin and user
// also created domain for the roles table
var ROLE_ADMIN = "ROLE_ADMIN"
var ROLE_USER = "ROLE_USER"

type Role struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type RoleUsecases interface {
	CreateRole(r *Role) (*Role, error)
	DeleteRole(r *Role) (*Role, error)
	UpdateRole(r *Role) (*Role, error)
	AddRoleForUser(userId int64, role string) error
	CheckUserHasRole(userId int64, role string) (bool, error)
	GetRoleForUser(userId int64) (*Role, error)
}
type RoleRepository interface {
	CreateRole(r *Role) (*Role, error)
	DeleteRole(r *Role) (*Role, error)
	UpdateRole(r *Role) (*Role, error)
	AddRoleForUser(userId int64, role string) error
	CheckUserHasRole(userId int64, role string) (bool, error)
	GetRoleForUser(userId int64) (*Role, error)
}
