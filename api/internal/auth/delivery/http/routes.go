package http

import (
	"forum/api/internal/auth"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(r *mux.Router, uc auth.AuthUseCase) {
	handler := NewAuthHandlers(uc)

	r.HandleFunc("/login", handler.Login).Methods("POST")
	r.HandleFunc("/register", handler.Register).Methods("POST")
	r.HandleFunc("/token/refresh", handler.RefreshToken).Methods("POST")
}
