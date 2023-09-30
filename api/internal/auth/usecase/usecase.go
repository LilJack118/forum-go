package usecase

import "forum/api/internal/auth"

type authUseCase struct {
	authRepo *auth.AuthRepository
}

func NewAuthUseCase(authRepo auth.AuthRepository) *authUseCase {
	return &authUseCase{authRepo: &authRepo}
}
