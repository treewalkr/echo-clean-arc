package service

import (
	"echo-clean-arc/domain"
)

type UserService interface {
	FindAll() ([]domain.User, error)
	FindById(id int) (*domain.User, error)
	Create(u *domain.User) error
}
