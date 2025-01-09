package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken(uuid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "",
		"sub": uuid,
		"aud": "",
		"exp": jwt.NewNumericDate(time.Now().Add(time.Second * 10)),
		"nbf": jwt.NewNumericDate(time.Now()),
		"iat": jwt.NewNumericDate(time.Now()),
		"jti": "",
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWTSECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateRefreshJwtToken(accessToken string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "",
		"sub": accessToken,
		"aud": "",
		"exp": jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 15)),
		"nbf": jwt.NewNumericDate(time.Now()),
		"iat": jwt.NewNumericDate(time.Now()),
		"jti": "",
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWTSECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
