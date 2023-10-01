package http

import "forum/api/internal/account"

type accountHandlers struct {
	u account.AccountUseCase
}

func NewAccountHandlers(u account.AccountUseCase) *accountHandlers {
	return &accountHandlers{u}
}
