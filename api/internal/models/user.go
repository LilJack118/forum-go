package models

import (
	"forum/api/pkg/validator"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"user_id" db:"user_id" validate:"omitempty"`
	FirstName string    `json:"first_name" db:"first_name" validate:"required,lte=30"`
	LastName  string    `json:"last_name" db:"last_name" validate:"required,lte=30"`
	Email     string    `json:"email" db:"email" validate:"omitempty,lte=60,email"`
	Password  string    `json:"password" db:"password" validate:"omitempty,required,gte=6"`
}

func (user *User) Validate() error {
	return validator.ValidateStruct(user)
}

func (user *User) PrepareCreate() error {
	user.ID = uuid.New()

	return nil
}

type UserWithToken struct {
	User         *User  `json:"user"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}
