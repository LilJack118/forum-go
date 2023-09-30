package auth

import "forum/api/internal/models"

type AuthRepository interface {
	CreateUser(user *models.User) error
}
