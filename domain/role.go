package domain

import "errors"

var ErrorRoleExists = errors.New("role already exists in a database")
var ErrorRoleDoesntExists = errors.New("role doesn't exists")
var ErrorUserAlreadyHasSameRole = errors.New("user already has same role")

type Role struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type RoleUsecases interface {
	CreateRole(r *Role) (*Role, error)
	DeleteRole(r *Role) (*Role, error)
	UpdateRole(r *Role) (*Role, error)
	AddRoleForUser(userId, roleId int64) error
	CheckUserHasRole(userId, roleId int64) (bool, error)
}
type RoleRepository interface {
	CreateRole(r *Role) (*Role, error)
	DeleteRole(r *Role) (*Role, error)
	UpdateRole(r *Role) (*Role, error)
	AddRoleForUser(userId, roleId int64) error
	CheckUserHasRole(userId, roleId int64) (bool, error)
}
