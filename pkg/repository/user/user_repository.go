package repository

import (
	"event-system-backend/pkg/model/domain"
)

type UserRepository interface {
	FindByIdOrUsernameOrEmail(username string) (*domain.User, error)
	FindOneVendorById(id string) (*domain.User, error)
	FindAllVendors(name string) ([]domain.User, error)
}
