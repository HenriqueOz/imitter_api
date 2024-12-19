package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"regexp"
	"time"
	"unicode"

	"github.com/golang-jwt/jwt/v5"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/constants"
)

func ValidateEmail(email string) error {
	if ok, _ := regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, email); !ok {
		return apperrors.ErrInvalidEmail
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < int(constants.PASSWORD_MIN_LENGTH) {
		return apperrors.ErrShortPassword
	}
	if len(password) > int(constants.PASSWORD_MAX_LENGTH) {
		return apperrors.ErrLongPassword
	}

	var upper, lower, special, number int
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			upper++
		case unicode.IsLower(char):
			lower++
		case unicode.IsNumber(char):
			number++
		case unicode.IsSymbol(char) || unicode.IsPunct(char):
			special++
		}
	}

	if upper >= 1 && lower >= 1 && special >= 1 && number >= 1 {
		return nil
	}
	return apperrors.ErrIvalidPassword
}

func ValidateUsername(username string) error {
	if len(username) < int(constants.USER_NAME_MIN_LENGTH) {
		return apperrors.ErrShortName
	}
	if len(username) > int(constants.USER_NAME_MAX_LENGTH) {
		return apperrors.ErrLongName
	}

	for _, char := range username {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) && char != '_' {
			return apperrors.ErrInvalidName
		}
	}

	return nil
}

func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func GenerateJwtToken(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "",
		"sub": userId,
		"aud": "",
		"exp": "",
		"nbf": time.Now().UnixMilli(),
		"iat": "",
		"jti": "",
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWTSECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
