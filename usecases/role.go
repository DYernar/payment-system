package usecases

import "payment.system.com/domain"

type RoleUsecases struct {
	Repo domain.RoleRepository
}

func NewRoleUsecases(pgRepo domain.RoleRepository) *RoleUsecases {
	return &RoleUsecases{
		Repo: pgRepo,
	}
}

func (uc *RoleUsecases) CreateRole(r *domain.Role) (*domain.Role, error) {
	return uc.Repo.CreateRole(r)
}

func (uc *RoleUsecases) DeleteRole(r *domain.Role) (*domain.Role, error) {
	return uc.Repo.DeleteRole(r)
}

func (uc *RoleUsecases) UpdateRole(r *domain.Role) (*domain.Role, error) {
	return uc.Repo.UpdateRole(r)
}

func (uc *RoleUsecases) AddRoleForUser(userId, roleId int64) error {
	return uc.Repo.AddRoleForUser(userId, roleId)
}

func (uc *RoleUsecases) CheckUserHasRole(userId, roleId int64) (bool, error) {
	return uc.Repo.CheckUserHasRole(userId, roleId)
}
