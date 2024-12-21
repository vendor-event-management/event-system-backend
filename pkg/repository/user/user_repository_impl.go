package repository

import (
	"event-system-backend/pkg/model/domain"

	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (u *userRepositoryImpl) FindByUsername(username string) (*domain.User, error) {
	var user domain.User
	if err := u.db.Table("users").Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
