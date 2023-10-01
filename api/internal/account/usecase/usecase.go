package usecase

import (
	"forum/api/internal/auth"
	"forum/api/internal/models"
	"net/http"
)

type accountUseCase struct {
	authrepo auth.AuthRepository
}

func NewAccountUseCase(authrepo auth.AuthRepository) *accountUseCase {
	return &accountUseCase{authrepo}
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
	return u.authrepo.UpdateUser(id, fields)
}

func (u *accountUseCase) DeleteUserAccount(id string) (int, error) {
	return u.authrepo.DeleteUser(id)
}
