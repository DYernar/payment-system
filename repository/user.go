package repository

import (
	"database/sql"

	"payment.system.com/domain"
)

type PgUserRepository struct {
	Db *sql.DB
}

func NewPgUserRepository(Db *sql.DB) *PgUserRepository {
	return &PgUserRepository{
		Db: Db,
	}
}

func (pgRep *PgUserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	return nil, nil
}

func (pgRep *PgUserRepository) UpdateUser(user *domain.User) (*domain.User, error) {
	return nil, nil
}
