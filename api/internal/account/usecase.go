package account

import "forum/api/internal/models"

type AccountUseCase interface {
	GetUserAccount(id string) (*models.User, int, error)
}
