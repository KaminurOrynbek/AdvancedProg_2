package usecase

import (
	"adp_practice1/internal/domain"
	"adp_practice1/internal/repository/user_repo"
	"errors"
)

type UserUsecase struct {
	repo user_repo.UserRepository
}

func NewUserUsecase(repo user_repo.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uc *UserUsecase) CreateUser(name, email string) (*domain.User, error) {
	user := &domain.User{Name: name, Email: email}
	if name == "" || email == "" {
		return nil, errors.New("name and email are required")
	}
	err := uc.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserUsecase) GetUserByID(id uint64) (*domain.User, error) {
	return uc.repo.GetByID(id)
}
