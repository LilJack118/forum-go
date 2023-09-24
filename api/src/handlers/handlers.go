package handlers

import (
	"encoding/json"
	"fmt"
	"forum/api/src/models"
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

	fmt.Fprintf(w, "User: %+v", user)
}
