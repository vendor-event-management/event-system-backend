package authentication

import (
	"event-system-backend/internal/auth"
	"event-system-backend/pkg/handler"
	"event-system-backend/pkg/model/dto/request"
	"event-system-backend/pkg/model/dto/response"
	"event-system-backend/pkg/service/user"
	"net/http"
)

type AuthenticationServiceImpl struct {
	userService user.UserService
}

func NewAuthenticationService(userService user.UserService) AuthenticationService {
	return &AuthenticationServiceImpl{userService: userService}
}

func (as *AuthenticationServiceImpl) Login(data request.LoginDto) (response.LoginResponse, *handler.CustomError) {
	var response response.LoginResponse

	user, errUser := as.userService.GetUserByIdOrUsernameOrEmail(data.Username)
	if errUser != nil {
		return response, handler.NewError(errUser.Code, errUser.Message)
	}

	isValid := auth.VerifyPassword(user.Password, data.Password)
	if !isValid {
		return response, handler.NewError(http.StatusUnauthorized, "Incorrect password")
	}

	token, errToken := auth.GenerateJWT(*user)
	if errToken != nil {
		return response, handler.NewError(http.StatusInternalServerError, errToken.Error())
	}

	response.Token = token

	return response, nil
}
