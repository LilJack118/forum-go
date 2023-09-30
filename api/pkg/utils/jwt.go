package utils

import (
	"errors"
	"forum/api/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type authJWT struct {
	secret      []byte
	access_exp  int
	refresh_exp int
}

type jwtClaims struct {
	Type string `json:"type,omitempty"`
	jwt.RegisteredClaims
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

func (a *authJWT) verifyToken(tokenString string, expected_type string) (*jwt.Token, *jwtClaims, error) {
	claims := &jwtClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return a.secret, errors.New("invalid token")
		}

		return a.secret, nil
	})

	if claims.Type != expected_type {
		return nil, nil, errors.New("invalid token type")
	}

	if err != nil {
		return nil, nil, err
	}

	return token, claims, nil
}

func (a *authJWT) GetUserID(access_token string) (string, error) {
	_, claims, err := a.verifyToken(access_token, "access")
	if err != nil {
		return "", err
	}

	return claims.RegisteredClaims.Subject, nil
}

func (a *authJWT) VerifyAccessToken(access_token string) (*jwt.Token, *jwtClaims, error) {
	return a.verifyToken(access_token, "access")
}

func (a *authJWT) VerifyRefreshToken(refresh_token string) (*jwt.Token, *jwtClaims, error) {
	return a.verifyToken(refresh_token, "refresh")
}

func (a *authJWT) createToken(claims *jwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(a.secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *authJWT) CreateAccessToken(sub string) (string, error) {
	claims := jwtClaims{
		Type: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   sub,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(a.access_exp))),
		},
	}

	return a.createToken(&claims)
}

func (a *authJWT) CreateRefreshToken(sub string) (string, error) {
	claims := jwtClaims{
		Type: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   sub,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(a.access_exp))),
		},
	}

	return a.createToken(&claims)
}

func (a *authJWT) CreateTokens(sub string) (string, string, error) {
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
