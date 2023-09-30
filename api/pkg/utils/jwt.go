package utils

import (
	"forum/api/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type authJWT struct {
	secret      []byte
	access_exp  int
	refresh_exp int
}

func AuthJWT() (*authJWT, error) {
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
		secret:      []byte(secret.(string)),
		access_exp:  int(access_exp.(int)),
		refresh_exp: int(refresh_exp.(int)),
	}

	return &auth, nil
}

func (a *authJWT) GetUserID() (uuid.UUID, error) {
	// returns user based on access token
	return uuid.New(), nil
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
