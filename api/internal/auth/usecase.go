package auth

import (
	"forum/api/internal/models"
)

type AuthUseCase interface {
	Register(user *models.User) (*models.User, error)
	Login(email string, password string) (*models.User, int, error)
}
