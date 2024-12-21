package handler

import (
	"event-system-backend/pkg/model/dto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomError struct {
	Message string
	Code    int
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func NewError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			var customErr *CustomError
			if ok := err.(*CustomError); ok != nil {
				customErr = ok
			} else {
				customErr = &CustomError{
					Message: "Internal Server Error",
					Code:    http.StatusInternalServerError,
				}
			}

			c.JSON(customErr.Code, dto.BaseResponse(false, customErr.Message, nil))
		}
	}
}
