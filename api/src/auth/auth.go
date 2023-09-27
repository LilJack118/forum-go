package auth

import (
	"forum/api/src/config"
	"forum/api/src/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type authJWT struct {
	request     *http.Request
	secret      []byte
	access_exp  int
	refresh_exp int
}

func AuthJWT(req *http.Request) (*authJWT, error) {
	secret, err := config.Config("SECRET_KEY", "string")
	if err != nil {
		return nil, err
	}
	access_exp, err := config.Config("ACCESS_TOKEN_EXP", "int")
	if err != nil {
		return nil, err
	}
	refresh_exp, err := config.Config("REFRESH_TOKEN_EXP", "int")
	if err != nil {
		return nil, err
	}

	auth := authJWT{
		request:     req,
		secret:      []byte(secret.(string)),
		access_exp:  int(access_exp.(int)),
		refresh_exp: int(refresh_exp.(int)),
	}

	return &auth, nil
}

func (a *authJWT) getUserID() (uuid.UUID, error) {
	// returns user based on access token
	return uuid.New(), nil
}

func (a *authJWT) GetUser() (*models.User, error) {

	id, err := a.getUserID()
	if err != nil {
		return nil, err
	}

	// TODO: get user from database
	user := models.User{ID: id, FirstName: "John", LastName: "Doe", Email: "doe@gmail.com", PasswordHash: "hdshshsfd"}

	// returns user based on access token
	return &user, nil
}

func (a *authJWT) VerifyAccessToken() bool {
	return true
}

func (a *authJWT) VerifyRefreshToken() bool {
	return true
}

func (a *authJWT) createToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(a.secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *authJWT) CreateAccessToken(sub uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"type": "access",
		"sub":  sub,
		"iat":  jwt.NewNumericDate(time.Now()),
		"exp":  jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(a.access_exp))),
	}

	return a.createToken(&claims)
}

func (a *authJWT) CreateRefreshToken(sub uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"type": "refresh",
		"sub":  sub,
		"iat":  jwt.NewNumericDate(time.Now()),
		"exp":  jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(a.refresh_exp))),
	}

	return a.createToken(&claims)
}

func (a *authJWT) CreateTokens(sub uuid.UUID) (string, string, error) {
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
