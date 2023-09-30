package usecase

import (
	"errors"
	"fmt"
	"forum/api/internal/auth"
	"forum/api/internal/models"
	"forum/api/pkg/utils"
	"net/http"
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

func (u *authUseCase) Login(email string, password string) (*models.User, int, error) {
	user, err := u.authRepo.GetUserByEmail(email)

	if user == nil || err != nil {
		return nil, http.StatusNotFound, fmt.Errorf("user with email %s doesn't exist", email)
	}

	if err := user.CheckPassword(password); err != nil {
		return nil, http.StatusUnauthorized, errors.New("invalid email or password")
	}

	user.CleanPassword()

	return user, 200, nil
}

func (u *authUseCase) GenerateTokens(user *models.User) (*models.UserWithTokens, error) {
	auth_jwt, err := utils.AuthJWT()
	if err != nil {
		return nil, err
	}

	access_token, refresh_token, err := auth_jwt.CreateTokens(user.ID.String())
	if err != nil {
		return nil, err
	}

	user_with_tokens := &models.UserWithTokens{
		AccessToken:  access_token,
		RefreshToken: refresh_token,
		User:         user,
	}

	return user_with_tokens, nil
}
