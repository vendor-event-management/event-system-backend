package user

import (
	"event-system-backend/pkg/handler"
	"event-system-backend/pkg/model/domain"
)

type UserService interface {
	GetUserByUsernameOrEmail(unameOrEmail string) (*domain.User, *handler.CustomError)
	GetVendorById(id string) (*domain.User, *handler.CustomError)
}
