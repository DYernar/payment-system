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

func (uc *RoleUsecases) AddRoleForUser(userId int64, role string) error {
	return uc.Repo.AddRoleForUser(userId, role)
}

func (uc *RoleUsecases) CheckUserHasRole(userId int64, role string) (bool, error) {
	return uc.Repo.CheckUserHasRole(userId, role)
}

func (uc *RoleUsecases) GetRoleForUser(userId int64) (*domain.Role, error) {
	return uc.Repo.GetRoleForUser(userId)
}
