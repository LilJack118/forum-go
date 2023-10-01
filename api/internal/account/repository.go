package account

import "forum/api/internal/models"

type AccountRepository interface {
	UpdateUserAccount(id string, fields *models.UserEditableFields) (int, error)
}
