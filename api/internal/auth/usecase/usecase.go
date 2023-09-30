package usecase

import (
	"fmt"
	"forum/api/internal/auth"
	"forum/api/internal/models"
)

type authUseCase struct {
	authRepo auth.AuthRepository
}

func NewAuthUseCase(authRepo auth.AuthRepository) *authUseCase {
	return &authUseCase{authRepo: authRepo}
}

func (u *authUseCase) Register(user *models.User) (*models.User, error) {

	existingUser, err := u.authRepo.GetUserByEmail(user.Email)
	if existingUser != nil || err == nil {
		return nil, fmt.Errorf("user with email %s already registered", user.Email)
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.PrepareCreate(); err != nil {
		return nil, err
	}

	// save to db
	if err := u.authRepo.CreateUser(user); err != nil {
		return nil, err
	}

	user.CleanPassword()

	return user, nil
}
