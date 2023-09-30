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

	// // validate user data
	// if err := req.ValidateData(); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// // check if account with given email already exists
	// if req.UserExists() {
	// 	http.Error(w, fmt.Sprintf("User with email %s already exists", req.Email), http.StatusBadRequest)
	// 	return
	// }

	// user := models.UserFromRequest(&req)
	// // TODO save to DB

	// auth, err := auth.AuthJWT(r)
	// if err != nil {
	// 	log.Print(err)
	// 	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	// 	return
	// }

	// access_token, refresh_token, err := auth.CreateTokens(user.ID)

	// if err != nil {
	// 	log.Print(err)
	// 	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	// 	return
	// }

	// response := models.AuthResponse{
	// 	AccessToken:  access_token,
	// 	RefreshToken: refresh_token,
	// 	User: models.UserResponse{
	// 		ID:        user.ID,
	// 		FirstName: user.FirstName,
	// 		LastName:  user.LastName,
	// 		Email:     user.Email,
	// 	},
	// }

	if err := json.NewEncoder(w).Encode(createdUser); err != nil {
		log.Print(err)
		httpErrors.JSONError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}