package usecase

import (
	"forum/api/internal/account"
	"forum/api/internal/auth"
	"forum/api/internal/models"
	"net/http"
)

type accountUseCase struct {
	authrepo    auth.AuthRepository
	accountrepo account.AccountRepository
}

func NewAccountUseCase(authrepo auth.AuthRepository, accountrepo account.AccountRepository) *accountUseCase {
	return &accountUseCase{authrepo, accountrepo}
}

func (u *accountUseCase) GetUserAccount(id string) (*models.User, int, error) {

	user, err := u.authrepo.GetUserByID(id)
	if err != nil {
		return nil, http.StatusNotFound, err
	}
	user.CleanPassword()
	return user, 0, nil
}

func (u *accountUseCase) UpdateUserAccount(id string, fields *models.UserEditableFields) (int, error) {
	return u.accountrepo.UpdateUserAccount(id, fields)
}
