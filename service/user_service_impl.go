package service

import (
	"echo-clean-arc/domain"

	"gorm.io/gorm"
)

type UserServiceImpl struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &UserServiceImpl{db}
}

func (s *UserServiceImpl) FindAll() ([]domain.User, error) {
	var users []domain.User
	result := s.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (s *UserServiceImpl) FindById(id int) (*domain.User, error) {
	var user domain.User
	result := s.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (s *UserServiceImpl) Create(u *domain.User) error {
	result := s.db.Create(u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
