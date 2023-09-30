package http

import (
	"encoding/json"
	"forum/api/internal/auth"
	"forum/api/internal/models"
	"forum/api/pkg/httpErrors"
	"log"
	"net/http"
)

type authHandlers struct {
	useCase auth.AuthUseCase
}

func NewAuthHandler(useCase auth.AuthUseCase) *authHandlers {
	return &authHandlers{
		useCase: useCase,
	}
}

// handler functions

func (h *authHandlers) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdUser, err := h.useCase.Register(&user)
	if err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO generate jwt tokens
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdUser); err != nil {
		log.Print(err)
		httpErrors.JSONError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

type loginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *authHandlers) Login(w http.ResponseWriter, r *http.Request) {
	var input loginInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, code, err := h.useCase.Login(input.Email, input.Password)
	if err != nil {
		httpErrors.JSONError(w, err.Error(), code)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Print(err)
		httpErrors.JSONError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
