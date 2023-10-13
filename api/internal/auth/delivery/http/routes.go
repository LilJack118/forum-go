package http

import (
	"forum/api/internal/auth"

	"github.com/gorilla/mux"
)

func RegisterAuthNotProtectedRoutes(r *mux.Router, uc auth.AuthUseCase) {
	handler := NewAuthHandlers(uc)

	r.HandleFunc("/login", handler.Login).Methods("POST")
	r.HandleFunc("/register", handler.Register).Methods("POST")
	r.HandleFunc("/token/refresh", handler.RefreshToken).Methods("POST")
}
func RegisterAuthProtectedRoutes(r *mux.Router, uc auth.AuthUseCase) {
	// used to register routes that are protected with authentication middleware
	handler := NewAuthHandlers(uc)

	r.HandleFunc("/token/verify", handler.VerifyToken).Methods("GET")
}
