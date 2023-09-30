package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/mail"

	"github.com/google/uuid"
)

func UserFromRequest(req *RegisterRequest) User {
	u := User{uuid.New(), req.FirstName, req.LastName, req.Email, req.Password}
	return u
}

type RegisterRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (req *RegisterRequest) UserExists() bool {
	// TODO check if user with specified email already exists
	return false
}

func (req *RegisterRequest) ValidateData() error {
	// Check if data is valid
	var errorMap map[string]string = make(map[string]string)

	if req.FirstName == "" {
		errorMap["FirstName"] = "first name is required"
	}

	if req.LastName == "" {
		errorMap["LastName"] = "last name is required"
	}

	if _, err := mail.ParseAddress(req.Email); err != nil {
		errorMap["Email"] = fmt.Sprintf("email address %s is invalid", req.Email)
	}

	if len(req.Password) < 6 {
		errorMap["Password"] = "password must be at least 6 characters"
	}

	if len(errorMap) > 0 {
		errorString, _ := json.Marshal(errorMap)
		return errors.New(string(errorString))
	} else {
		return nil
	}
}

type User struct {
	ID        uuid.UUID `json:"user_id" db:"user_id" validate:"omitempty"`
	FirstName string    `json:"first_name" db:"first_name" validate:"required,lte=30"`
	LastName  string    `json:"last_name" db:"last_name" validate:"required,lte=30"`
	Email     string    `json:"email" db:"email" validate:"omitempty,lte=60,email"`
	Password  string    `json:"password" db:"password" validate:"omitempty,required,gte=6"`
}

type UserWithToken struct {
	User         *User  `json:"user"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"uuid"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
}

// Tokens Response

type AuthResponse struct {
	// used in register and login endpoints
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	User         UserResponse `json:"user"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}
