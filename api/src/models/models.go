package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/mail"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string
}

func UserFromRequest(ur *UserRequest) User {
	u := User{uuid.New(), ur.FirstName, ur.LastName, ur.Email, ur.Password}
	return u
}

type UserRequest struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (u *UserRequest) Exists() bool {
	// TODO check if user with specified email already exists
	return false
}

func (u *UserRequest) ValidateData() error {
	// Check if data is valid
	var errorMap map[string]string = make(map[string]string)

	if u.FirstName == "" {
		errorMap["FirstName"] = "first name is required"
	}

	if u.LastName == "" {
		errorMap["LastName"] = "last name is required"
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		errorMap["Email"] = fmt.Sprintf("email address %s is invalid", u.Email)
	}

	if len(u.Password) < 6 {
		errorMap["Password"] = "password must be at least 6 characters"
	}

	if len(errorMap) > 0 {
		errorString, _ := json.Marshal(errorMap)
		return errors.New(string(errorString))
	} else {
		return nil
	}
}
