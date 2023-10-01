package account

import "forum/api/internal/models"

type AccountUseCase interface {
	GetUserAccount(id string) (*models.User, int, error)
	UpdateUserAccount(id string, fields *models.UserEditableFields) (int, error)
	DeleteUserAccount(id string) (int, error)
}
