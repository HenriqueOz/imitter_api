package utilstest

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/constants"
	"sm.com/m/src/app/utils"
)

type ValidateEmailInput struct {
	Email    string
	Expected error
}

var validateEmailInputs = []ValidateEmailInput{
	{
		Email: "random@email.com", Expected: nil,
	},
	{
		Email: "randomemail.com", Expected: apperrors.ErrInvalidEmail,
	},
	{
		Email: "random@emailcom", Expected: apperrors.ErrInvalidEmail,
	},
	{
		Email: "random.com", Expected: apperrors.ErrInvalidEmail,
	},
	{
		Email: "randomemailcom", Expected: apperrors.ErrInvalidEmail,
	},
	{
		Email: ".", Expected: apperrors.ErrInvalidEmail,
	},
	{
		Email: "@.", Expected: apperrors.ErrInvalidEmail,
	},
}

func Test_ValidateEmail(t *testing.T) {
	for _, input := range validateEmailInputs {
		err := utils.ValidateEmail(input.Email)

		assert.Equal(t, input.Expected, err, "Error should match de input expected field")
	}
}

type ValidatePasswordInput struct {
	Password string
	Expected error
}

var validatePasswordInputs = []ValidatePasswordInput{
	{
		Password: "", Expected: apperrors.ErrShortPassword,
	},
	{
		Password: "a", Expected: apperrors.ErrShortPassword,
	},
	{
		Password: strings.Repeat("a", int(constants.PASSWORD_MAX_LENGTH)+1), Expected: apperrors.ErrLongPassword,
	},
	{
		Password: "password", Expected: apperrors.ErrInvalidPassword,
	},
	{
		Password: "password12", Expected: apperrors.ErrInvalidPassword,
	},
	{
		Password: "Password12", Expected: apperrors.ErrInvalidPassword,
	},
	{
		Password: "Password@12", Expected: nil,
	},
}

func Test_ValidatePassword(t *testing.T) {
	for _, input := range validatePasswordInputs {
		err := utils.ValidatePassword(input.Password)

		assert.Equal(t, input.Expected, err, "Error should match de input expected field")
	}
}

type ValidateNameInput struct {
	Name     string
	Expected error
}

var validateNameInputs = []ValidateNameInput{
	{
		Name: "a", Expected: apperrors.ErrShortName,
	},
	{
		Name: strings.Repeat("a", int(constants.USER_NAME_MAX_LENGTH)+1), Expected: apperrors.ErrLongName,
	},
	{
		Name: "username@", Expected: apperrors.ErrInvalidName,
	},
	{
		Name: "username12", Expected: nil,
	},
	{
		Name: "!@#$%*()", Expected: apperrors.ErrInvalidName,
	},
	{
		Name: "user_name12", Expected: nil,
	},
}

func Test_ValidateName(t *testing.T) {
	for _, input := range validateNameInputs {
		err := utils.ValidateName(input.Name)

		assert.Equal(t, input.Expected, err, "Error should match de input expected field: name: "+input.Name)
	}
}
