package utils

import (
	"log"
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

	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
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
		"nbf":  jwt.NewNumericDate(time.Now()),
		"iat":  jwt.NewNumericDate(time.Now()),
		"jti":  jti,
	})

	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) *jwt.Token {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		log.Printf("Failed parsing token: %v\n", err)
		return nil
	}
	return token
}
