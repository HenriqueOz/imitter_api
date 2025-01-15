package utilstest

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"sm.com/m/src/app/utils"
)

func Test_ResponseError(t *testing.T) {
	err := errors.New("test error")
	details := []string{"bunch", "of", "details"}

	expected := map[string]any{
		"error":   err.Error(),
		"details": details,
	}

	result := utils.ResponseError(err, details)

	assert.Equal(t, expected, result)
}

func Test_ResponseSuccess(t *testing.T) {
	data := []string{"bunch", "of", "data"}

	expected := map[string]any{
		"data": data,
	}

	result := utils.ResponseSuccess(data)

	assert.Equal(t, expected, result)
}
