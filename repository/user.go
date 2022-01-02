package repository

import (
	"context"
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
		SELECT u.id, u.iin, u.login, u.password, u.created_at, r.id, r.name
		FROM users as u, roles as r
		WHERE login=$1 AND u.id = r.user_id
	`

	var user domain.User
	user.Role = &domain.Role{}
	err := pgRep.Db.QueryRow(query, login).Scan(&user.Id, &user.Iin, &user.Login, &user.Password, &user.CreatedAt, &user.Role.Id, &user.Role.Name)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, domain.ErrorUserLoginDoesntExists
		}
		return nil, err
	}

	return &user, err
}

func (pgRep *PgUserRepository) GetAllUsers() ([]*domain.User, error) {
	query := `
		SELECT u.id, u.iin, u.login, u.password, u.created_at, r.name, r.id
		FROM users as u, roles as r
		WHERE u.id = r.user_id
	`
	resp, err := pgRep.Db.QueryContext(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer resp.Close()

	users := []*domain.User{}

	for resp.Next() {
		u := domain.User{
			Role: &domain.Role{},
		}
		resp.Scan(&u.Id, &u.Iin, &u.Login, &u.Password, &u.CreatedAt, &u.Role.Name, &u.Role.Id)
		users = append(users, &u)
	}

	return users, nil

}
