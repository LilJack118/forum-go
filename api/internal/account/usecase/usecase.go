package usecase

import "forum/api/internal/account"

type accountUseCase struct {
	repo account.AccountRepository
}

func NewAccountUseCase(repo account.AccountRepository) *accountUseCase {
	return &accountUseCase{repo}
}
