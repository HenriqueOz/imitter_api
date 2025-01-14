package utilstest

import (
	"testing"

	"sm.com/m/src/app/utils"
)

func TestGenerateJwtToken_ErrSign(t *testing.T) {

	_, err := utils.GenerateJwtToken("test-uuid")

	if err == nil {
		t.Fatalf("Must return a error, got err = nil")
	}
}
