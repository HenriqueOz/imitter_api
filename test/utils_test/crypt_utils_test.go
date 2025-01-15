package utilstest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"sm.com/m/src/app/utils"
)

func Test_CryptSha256(t *testing.T) {
	input := "Cool String"
	expected := "a89b5226763528f88a2a12f2804e80486d732bb2ac4f9347f58b856e1e1bd747"

	result := utils.HashSha256(input)

	assert.Equal(t, len(expected), len(result))
	assert.Equal(t, expected, result)
}
