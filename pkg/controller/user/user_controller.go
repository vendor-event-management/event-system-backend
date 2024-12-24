package user

import (
	"event-system-backend/pkg/handler"
	"event-system-backend/pkg/middleware"
	"event-system-backend/pkg/model/dto"
	userservice "event-system-backend/pkg/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService userservice.UserService
}

func NewUserController(userService userservice.UserService) *UserController {
	return &UserController{userService: userService}
}

func SetupUserRoutes(r *gin.RouterGroup, uc *UserController) {
	userGroup := r.Group("/user")
	userGroup.Use(middleware.AuthMiddleware)
	userGroup.GET("/vendors", uc.ShowAllVendors)
}

// ShowAllVendors shows all vendors with optional full name filter
// @Summary Get all vendors
// @Description Retrieve the list of all vendors, optionally filtered by full name
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <your-token-here>"
// @Param fullName query string false "Filter vendors by full name"
// @Security BearerAuth
// @Success 200 {object} dto.Response{data=[]response.VendorsResponse} "Vendors retrieved successfully"
// @Failure 400 {object} dto.Response "Bad request"
// @Failure 500 {object} dto.Response "Internal server error"
// @Router /user/vendors [get]
func (uc *UserController) ShowAllVendors(c *gin.Context) {
	fullName := c.DefaultQuery("fullName", "")

	vendors, err := uc.userService.GetAllVendors(fullName)
	if err != nil {
		c.Error(handler.NewError(err.Code, err.Message))
		return
	}

	c.JSON(http.StatusOK, dto.BaseResponse(true, "OK", vendors))
}
