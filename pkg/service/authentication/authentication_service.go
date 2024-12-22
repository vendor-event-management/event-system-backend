package authentication

import (
	"event-system-backend/pkg/handler"
	"event-system-backend/pkg/model/dto/request"
	"event-system-backend/pkg/model/dto/response"
)

type AuthenticationService interface {
	Login(data request.LoginDto) (response.LoginResponse, *handler.CustomError)
}
