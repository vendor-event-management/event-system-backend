package repository

import (
	"event-system-backend/pkg/model/domain"
	"event-system-backend/pkg/utils"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (u *UserRepositoryImpl) FindByIdOrUsernameOrEmail(idOrUsername string) (*domain.User, error) {
	var user domain.User
	if err := u.db.Table("users").Where("id = ? OR username = ? OR email = ?", idOrUsername, idOrUsername, idOrUsername).First(&user).Error; err != nil {
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

func (u *UserRepositoryImpl) FindAllVendors(name string) ([]domain.User, error) {
	response := []domain.User{}

	baseQuery := u.db.Table("users").Where("role = ? AND deleted_at IS NULL", domain.Vendor)
	if !utils.IsEmptyString(name) {
		baseQuery = baseQuery.Where("LOWER(full_name) LIKE LOWER(?)", "%"+name+"%")
	}

	baseQuery = baseQuery.Order("full_name ASC")
	if err := baseQuery.Find(&response).Error; err != nil {
		return response, err
	}

	return response, nil
}
