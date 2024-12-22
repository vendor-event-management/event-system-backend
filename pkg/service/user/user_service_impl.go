package user

import (
	"event-system-backend/pkg/handler"
	"event-system-backend/pkg/model/domain"
	"event-system-backend/pkg/model/dto/response"
	repository "event-system-backend/pkg/repository/user"
	"net/http"

	"gorm.io/gorm"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{userRepository: userRepository}
}

func (e *UserServiceImpl) GetUserByIdOrUsernameOrEmail(unameOrEmail string) (*domain.User, *handler.CustomError) {
	user, err := e.userRepository.FindByIdOrUsernameOrEmail(unameOrEmail)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, handler.NewError(http.StatusNotFound, "User not found")
		}
		return nil, handler.NewError(http.StatusInternalServerError, err.Error())
	}

	return user, nil
}

func (e *UserServiceImpl) GetVendorById(id string) (*domain.User, *handler.CustomError) {
	vendor, err := e.userRepository.FindOneVendorById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, handler.NewError(http.StatusNotFound, "Vendor not found")
		}
		return nil, handler.NewError(http.StatusInternalServerError, err.Error())
	}

	return vendor, nil
}

func (e *UserServiceImpl) GetAllVendors(name string) ([]response.VendorsResponse, *handler.CustomError) {
	fetchedVendors, errVendors := e.userRepository.FindAllVendors(name)
	if errVendors != nil && errVendors != gorm.ErrRecordNotFound {
		return nil, handler.NewError(http.StatusInternalServerError, errVendors.Error())
	}

	return response.BuildVendorsResponseFromFetchedVendors(fetchedVendors), nil
}
