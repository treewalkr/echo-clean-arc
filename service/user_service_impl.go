package service

import (
	"echo-clean-arc/domain"
	"echo-clean-arc/repository"
)

// UserService represent the user's service contract
type UserServiceImpl struct {
	useRepository repository.UserRepository
}

// NewUserService creates a new user service.
func NewUserService(useRepository repository.UserRepository) UserService {
	return &UserServiceImpl{useRepository}
}

// FindAll finds all users.
func (s *UserServiceImpl) FindAll() ([]domain.User, error) {
	return s.useRepository.FindAll()
}

// FindById finds a user by its id.
func (s *UserServiceImpl) FindById(id int) (*domain.User, error) {
	return s.useRepository.FindById(id)
}

// Create creates a new user.
func (s *UserServiceImpl) Create(u *domain.User) error {
	return s.useRepository.Create(u)
}
