package usecase

import (
	"forum/api/internal/auth"
	"forum/api/internal/models"
)

type authUseCase struct {
	authRepo *auth.AuthRepository
}

func NewAuthUseCase(authRepo auth.AuthRepository) *authUseCase {
	return &authUseCase{authRepo: &authRepo}
}

func (u *authUseCase) Register(user *models.User) (*models.User, error) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.PrepareCreate(); err != nil {
		return nil, err
	}

	// hash password

	return user, nil
}
