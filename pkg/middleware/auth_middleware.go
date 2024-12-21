package middleware

import (
	"event-system-backend/internal/auth"
	"event-system-backend/pkg/handler"
	"event-system-backend/pkg/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if utils.IsEmptyString(authHeader) {
		c.Error(handler.NewError(http.StatusUnauthorized, "Authorization header is missing"))
		c.Abort()
		return
	}

	// check if header starts with Bearer
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.Error(handler.NewError(http.StatusUnauthorized, "Invalid authorization format"))
		c.Abort()
		return
	}

	jwtSecret := auth.GetJWTSecret()
	tokenString := parts[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		c.Error(handler.NewError(http.StatusUnauthorized, err.Error()))
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		username := claims["username"].(string)
		c.Set("username", username)
	} else {
		c.Error(handler.NewError(http.StatusUnauthorized, "Invalid token claims"))
		c.Abort()
		return
	}

	c.Next()
}
