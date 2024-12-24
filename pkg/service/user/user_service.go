package user

import (
	"event-system-backend/pkg/handler"
	"event-system-backend/pkg/model/domain"
	"event-system-backend/pkg/model/dto/response"
)

type UserService interface {
	GetUserByIdOrUsernameOrEmail(unameOrEmail string) (*domain.User, *handler.CustomError)
	GetVendorById(id string) (*domain.User, *handler.CustomError)
	GetAllVendors(name string) ([]response.VendorsResponse, *handler.CustomError)
}
