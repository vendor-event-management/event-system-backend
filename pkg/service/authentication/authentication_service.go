package authentication

import (
	"event-system-backend/pkg/handler"
	"event-system-backend/pkg/model/dto"
)

type AuthenticationService interface {
	Login(data dto.LoginDto) (dto.LoginResponse, *handler.CustomError)
}
