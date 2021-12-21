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
	query := `
		INSERT INTO users(iin, login, password) 
		VALUES ($1, $2, $3)
		RETURNING id, iin, login, password, created_at
	`

	args := []interface{}{user.Iin, user.Login, user.Password}

	var newUser domain.User
	err := pgRep.Db.QueryRow(query, args...).Scan(&newUser.Id, &newUser.Iin, &newUser.Login, &newUser.Password, &newUser.CreatedAt)

	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_iin_key\"" {
			return nil, domain.ErrorIinAlreadyExists
		}
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_login_key\"" {
			return nil, domain.ErrorLoginAlreadyExists
		}
		return nil, err
	}

	return &newUser, nil
}

func (pgRep *PgUserRepository) UpdateUser(user *domain.User) (*domain.User, error) {
	return nil, nil
}

func (pgRep *PgUserRepository) GetUserByLogin(login string) (*domain.User, error) {
	query := `
		SELECT id, iin, login, password, created_at
		FROM users
		WHERE login=$1
	`

	var user domain.User
	err := pgRep.Db.QueryRow(query, login).Scan(&user.Id, &user.Iin, &user.Login, &user.Password, &user.CreatedAt)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, domain.ErrorUserLoginDoesntExists
		}
		return nil, err
	}

	return &user, err
}
