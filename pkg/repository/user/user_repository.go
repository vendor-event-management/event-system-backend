package repository

import (
	"event-system-backend/pkg/model/domain"
)

type UserRepository interface {
	FindByUsername(username string) (*domain.User, error)
}
