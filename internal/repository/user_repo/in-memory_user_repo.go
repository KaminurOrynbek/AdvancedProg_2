package user_repo

import (
	"adp_practice1/internal/domain"
	"errors"
	"sync"
)

type InMemoryUserRepository struct {
	users  map[uint64]*domain.User
	mu     sync.Mutex
	nextID uint64
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users:  make(map[uint64]*domain.User),
		nextID: 1,
	}
}

func (r *InMemoryUserRepository) CreateUser(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = r.nextID
	r.nextID++

	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) GetByID(id uint64) (*domain.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}
