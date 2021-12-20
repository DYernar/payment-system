package repository

import (
	"database/sql"

	"payment.system.com/domain"
)

type PgRoleRepository struct {
	Db *sql.DB
}

func NewPgRoleRepo(Db *sql.DB) *PgRoleRepository {
	return &PgRoleRepository{
		Db: Db,
	}
}

func (pgRepo *PgRoleRepository) CreateRole(r *domain.Role) (*domain.Role, error) {
	return nil, nil
}

func (pgRepo *PgRoleRepository) DeleteRole(r *domain.Role) (*domain.Role, error) {
	return nil, nil
}

func (pgRepo *PgRoleRepository) UpdateRole(r *domain.Role) (*domain.Role, error) {
	return nil, nil
}

func (pgRepo *PgRoleRepository) AddRoleForUser(userId, roleId int64) error {
	return nil
}

func (pgRepo *PgRoleRepository) CheckUserHasRole(userId, roleId int64) (bool, error) {
	return false, nil
}
