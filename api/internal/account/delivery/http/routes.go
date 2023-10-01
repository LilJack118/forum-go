package http

import (
	"forum/api/internal/account"

	"github.com/gorilla/mux"
)

func RegisterAccountRoutes(r *mux.Router, u account.AccountUseCase) {
	handler := NewAccountHandlers(u)

	r.HandleFunc("/account", handler.Get).Methods("GET")
}
