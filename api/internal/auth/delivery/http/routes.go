package http

import (
	"forum/api/internal/auth"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, uc auth.AuthUseCase) {
	handler := NewAuthHandler(uc)

	r.HandleFunc("/register", handler.Register).Methods("POST")
}
