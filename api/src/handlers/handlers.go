package handlers

import (
	"encoding/json"
	"fmt"
	"forum/api/src/auth"
	"forum/api/src/models"
	"log"
	"net/http"
)

// Auth endpoint handlers

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var userRequest models.UserRequest

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// validate user data
	if err := userRequest.ValidateData(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check if account with given email already exists
	if userRequest.Exists() {
		http.Error(w, fmt.Sprintf("User with email %s already exists", userRequest.Email), http.StatusBadRequest)
		return
	}

	user := models.UserFromRequest(&userRequest)
	// TODO save to DB

	auth := auth.AuthJWT{Request: r}
	access_token, refresh_token, err := auth.CreateTokens(user.ID)

	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.AuthResponse{
		AccessToken:  access_token,
		RefreshToken: refresh_token,
		User: models.UserResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		},
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
