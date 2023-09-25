package auth

import (
	"forum/api/src/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthJWT struct {
	Request     *http.Request
	secret      []byte `default:"supersecrethere"`
	access_exp  int    `default:"900"`
	refresh_exp int    `default:"43200"`
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

func (a *AuthJWT) createToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(a.secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *AuthJWT) CreateAccessToken(sub uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"type": "access",
		"sub":  sub,
		"iat":  jwt.NewNumericDate(time.Now()),
		"exp":  jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(a.access_exp))),
	}

	return a.createToken(&claims)
}

func (a *AuthJWT) CreateRefreshToken(sub uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"type": "refresh",
		"sub":  sub,
		"iat":  jwt.NewNumericDate(time.Now()),
		"exp":  jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(a.refresh_exp))),
	}

	return a.createToken(&claims)
}

func (a *AuthJWT) CreateTokens(sub uuid.UUID) (string, string, error) {
	access_token, err := a.CreateAccessToken(sub)
	if err != nil {
		return "", "", err
	}

	refresh_token, err := a.CreateRefreshToken(sub)
	if err != nil {
		return "", "", err
	}

	return access_token, refresh_token, nil
}
