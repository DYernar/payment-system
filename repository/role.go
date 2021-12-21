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

func (pgRepo *PgRoleRepository) AddRoleForUser(userId int64, role string) error {
	query := `
		INSERT INTO roles(name, user_id)
		VALUES ($1, $2)
	`

	_, err := pgRepo.Db.Exec(query, role, userId)

	return err
}

func (pgRepo *PgRoleRepository) CheckUserHasRole(userId int64, role string) (bool, error) {
	return false, nil
}

func (pgRepo *PgRoleRepository) GetRoleForUser(userId int64) (*domain.Role, error) {
	query := `
		SELECT id, name
		FROM roles 
		WHERE user_id=$1
	`

	var role domain.Role

	err := pgRepo.Db.QueryRow(query, userId).Scan(&role.Id, &role.Name)

	if err != nil {
		return nil, err
	}
	return &role, nil
}
