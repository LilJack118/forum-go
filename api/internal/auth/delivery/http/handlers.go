package http

import (
	"encoding/json"
	"forum/api/internal/auth"
	"forum/api/internal/models"
	"forum/api/pkg/httpErrors"
	"forum/api/pkg/utils"
	"log"
	"net/http"
)

type authHandlers struct {
	useCase auth.AuthUseCase
}

func NewAuthHandlers(useCase auth.AuthUseCase) *authHandlers {
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

	response, err := h.useCase.GenerateTokens(createdUser)
	if err != nil {
		log.Print(err)
		httpErrors.JSONError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
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

	response, err := h.useCase.GenerateTokens(user)
	if err != nil {
		log.Print(err)
		httpErrors.JSONError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Print(err)
		httpErrors.JSONError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

type refreshTokenInput struct {
	RefreshToken string `json:"refresh_token"`
}

type accessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

func (h *authHandlers) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var input *refreshTokenInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	auth_jwt, _ := utils.AuthJWT()
	_, claims, err := auth_jwt.VerifyRefreshToken(input.RefreshToken)
	if err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	access_token, err := auth_jwt.CreateAccessToken(claims.RegisteredClaims.Subject)
	if err != nil {
		log.Print(err)
		httpErrors.JSONError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response := &accessTokenResponse{AccessToken: access_token}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Print(err)
		httpErrors.JSONError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
