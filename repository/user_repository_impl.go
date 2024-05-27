package repository

import (
	"echo-clean-arc/domain"

	"gorm.io/gorm"
)

// UserRepositoryImpl represent the repository for the user
type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository will create an object that represent the UserRepository interface
func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

// FindAll will return all user
func (r *UserRepositoryImpl) FindAll() ([]domain.User, error) {
	var users []domain.User
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// FindById will return a user by its id
func (r *UserRepositoryImpl) FindById(id int) (*domain.User, error) {
	var user domain.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// Create will create a new user
func (r *UserRepositoryImpl) Create(u *domain.User) error {
	result := r.db.Create(u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
