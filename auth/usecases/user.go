package usecases

import "payment.system.com/domain"

type UserUsecases struct {
	Repo domain.UserRepository
}

func NewUserUsecases(pgRepo domain.UserRepository) *UserUsecases {
	return &UserUsecases{
		Repo: pgRepo,
	}
}
func (uc *UserUsecases) CreateUser(user *domain.User) (*domain.User, error) {
	return uc.Repo.CreateUser(user)
}

func (uc *UserUsecases) UpdateUser(user *domain.User) (*domain.User, error) {
	return uc.Repo.UpdateUser(user)
}

func (uc *UserUsecases) GetUserByLogin(login string) (*domain.User, error) {
	return uc.Repo.GetUserByLogin(login)
}
