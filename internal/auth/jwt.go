package auth

import (
	"event-system-backend/pkg/model/domain"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GetJWTSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET_KEY"))
}

func GenerateJWT(user domain.User) (string, error) {
	secretKey := GetJWTSecret()

	claims := jwt.MapClaims{
		"username": user.Username,
		"company":  user.CompanyName,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // expired after 24 hours
		"iat":      time.Now().Unix(),                     // issued at current time
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
