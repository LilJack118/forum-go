package http

import (
	"forum/api/internal/account"

	"github.com/gorilla/mux"
)

func RegisterAccountRoutes(r *mux.Router, u account.AccountUseCase) {
	handler := NewAccountHandlers(u)

	r.HandleFunc("/account/{id}", handler.Get).Methods("GET")
}
