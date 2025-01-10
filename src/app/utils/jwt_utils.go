package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	id "github.com/google/uuid"
)

func GenerateJwtToken(uuid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "temp-issuer",
		"sub": uuid,
		"aud": "temp-aud",
		"exp": jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
		"nbf": jwt.NewNumericDate(time.Now()),
		"iat": jwt.NewNumericDate(time.Now()),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWTSECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateRefreshJwtToken(uuid string, accessToken string) (string, error) {
	jti := id.NewString()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":  "temp-issuer",
		"sub":  accessToken,
		"uuid": uuid,
		"aud":  "temp-aud",
		"exp":  jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		"nbf":  jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
		"iat":  jwt.NewNumericDate(time.Now()),
		"jti":  jti,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWTSECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
