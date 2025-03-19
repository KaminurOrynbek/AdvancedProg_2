package user_repo

import (
	"adp_practice1/internal/domain"
)

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetByID(id uint64) (*domain.User, error)
}
