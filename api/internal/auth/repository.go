package auth

import "forum/api/internal/models"

type AuthRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	UpdateUser(id string, fields *models.UserEditableFields) (int, error)
	DeleteUser(id string) (int, error)
}
