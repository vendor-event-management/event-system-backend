package repository

import (
	"event-system-backend/pkg/model/domain"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (u *UserRepositoryImpl) FindByUsername(username string) (*domain.User, error) {
	var user domain.User
	if err := u.db.Table("users").Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepositoryImpl) FindOneVendorById(id string) (*domain.User, error) {
	var vendor domain.User
	if err := u.db.Table("users").Where("id = ? AND role = ?", id, domain.Vendor).First(&vendor).Error; err != nil {
		return nil, err
	}
	return &vendor, nil
}
