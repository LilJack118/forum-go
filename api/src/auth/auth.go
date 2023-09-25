package auth

import (
	"forum/api/src/models"
	"net/http"

	"github.com/google/uuid"
)

type AuthJWT struct {
	Request *http.Request
}

func (a *AuthJWT) getUserID() (uuid.UUID, error) {
	// returns user based on access token
	return uuid.New(), nil
}

func (a *AuthJWT) GetUser() (*models.User, error) {

	id, err := a.getUserID()
	if err != nil {
		return nil, err
	}

	// TODO: get user from database
	user := models.User{ID: id, FirstName: "John", LastName: "Doe", Email: "doe@gmail.com", PasswordHash: "hdshshsfd"}

	// returns user based on access token
	return &user, nil
}

func (a *AuthJWT) VerifyAccessToken() bool {
	return true
}

func (a *AuthJWT) VerifyRefreshToken() bool {
	return true
}

func (a *AuthJWT) CreateAccessToken(id uuid.UUID) string {
	return "access_token"
}

func (a *AuthJWT) CreateRefreshToken(id uuid.UUID) string {
	return "refresh_token"
}
