package authentication

import (
	"event-system-backend/pkg/handler"
	"event-system-backend/pkg/model/dto"
	"event-system-backend/pkg/model/dto/request"
	authenticationservice "event-system-backend/pkg/service/authentication"
	"event-system-backend/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	authenticationService authenticationservice.AuthenticationService
}

func NewAuthenticationController(authenticationService authenticationservice.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{authenticationService: authenticationService}
}

func SetupAuthenticationRoutes(r *gin.RouterGroup, ac *AuthenticationController) {
	authGroup := r.Group("auth")
	authGroup.POST("/login", ac.Login)
}

// Login handles user login
// @Summary User login
// @Description Login with username and password to obtain authentication token
// @Accept json
// @Produce json
// @Tags Authentication
// @Param body body request.LoginDto true "User login credentials"
// @Success 200 {object} dto.Response{data=response.LoginResponse} "Login successful"
// @Failure 400 {object} dto.Response "Bad request"
// @Failure 500 {object} dto.Response "Internal server error"
// @Router /auth/login [post]
func (ac *AuthenticationController) Login(c *gin.Context) {
	var body request.LoginDto
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(handler.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	if utils.IsEmptyString(body.Username) {
		c.Error(handler.NewError(http.StatusBadRequest, "Username must be filled"))
		return
	}

	if utils.IsEmptyString(body.Password) {
		c.Error(handler.NewError(http.StatusBadRequest, "Password must be filled"))
		return
	}

	auth, errAuth := ac.authenticationService.Login(body)
	if errAuth != nil {
		c.Error(handler.NewError(errAuth.Code, errAuth.Message))
		return
	}

	c.JSON(http.StatusOK, dto.BaseResponse(true, "OK", auth))
}
