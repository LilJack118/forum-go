package models

import (
	"forum/api/pkg/validator"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID `json:"user_id" db:"user_id" validate:"omitempty"`
	FirstName string    `json:"first_name" db:"first_name" validate:"required,alphanum,lte=30"`
	LastName  string    `json:"last_name" db:"last_name" validate:"required,alphanum,lte=30"`
	Email     string    `json:"email" db:"email" validate:"omitempty,lte=60,email"`
	Password  string    `json:"password" db:"password" validate:"omitempty,required,gte=6"`
}

func (user *User) Validate() error {
	return validator.ValidateStruct(user)
}

func (user *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return nil
}

func (user *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}

	return nil
}

func (user *User) PrepareCreate() error {
	user.ID = uuid.New()

	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.ToLower(strings.TrimSpace(user.Email))

	if err := user.HashPassword(); err != nil {
		return err
	}

	return nil
}

type UserWithToken struct {
	User         *User  `json:"user"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}
