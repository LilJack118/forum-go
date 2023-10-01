package http

import (
	"fmt"
	"forum/api/internal/account"

	"github.com/gorilla/mux"
)

func RegisterAccountRoutes(r *mux.Router, u account.AccountUseCase) {
	handler := NewAccountHandlers(u)

	fmt.Println(handler)
}
