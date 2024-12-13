package utils

import (
	"regexp"
	"unicode"

	"sm.com/m/src/app/constants"
)

func ValidateEmail(email string) bool {
	ok, _ := regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, email)
	return ok
}

func ValidatePassword(password string) bool {
	if len(password) < int(constants.PASSWORD_MIN_LENGTH) ||
		len(password) > int(constants.PASSWORD_MAX_LENGTH) {
		return false
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
	return upper >= 1 && lower >= 1 && special >= 1 && number >= 1
}

func ValidateUsername(username string) bool {

	if len(username) < int(constants.USER_NAME_MIN_LENGTH) ||
		len(username) > int(constants.USER_NAME_MAX_LENGTH) {
		return false
	}

	for _, char := range username {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) && char != '_' {
			return false
		}
	}

	return true
}

func HashPassword(password string) string {
	//TODO implement
	return ""
}
