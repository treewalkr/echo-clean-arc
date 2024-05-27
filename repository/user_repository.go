package repository

import "echo-clean-arc/domain"

// UserRepository represent the user's repository contract
type UserRepository interface {
	FindAll() ([]domain.User, error)
	FindById(id int) (*domain.User, error)
	Create(u *domain.User) error
}
