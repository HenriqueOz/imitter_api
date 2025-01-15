package utilstest

import (
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"sm.com/m/src/app/utils"
	"sm.com/m/test"
)

func Test_GenerateJwtToken(t *testing.T) {
	setUp := func() {
		test.LoadEnv()

		assert.NotEmpty(t, os.Getenv("JWT_SECRET"), "JWT_SECRET should not be empty")
	}

	t.Run("Should return a valid token string", func(tt *testing.T) {
		setUp()

		uuid := "test-uuid"
		tokenString, err := utils.GenerateJwtToken(uuid)

		assert.Nil(tt, err, "Error should be nil")
		assert.NotZero(tt, tokenString, "TokenString should not be zero")
	})

	t.Run("Should parse and verify token claims", func(tt *testing.T) {
		setUp()

		uuid := "test-uuid"
		tokenString, _ := utils.GenerateJwtToken(uuid)

		claims := jwt.MapClaims{}
		token, err := utils.ParseTokenWithClaims(tokenString, &claims)

		assert.Nil(t, err, "Error should be nil")
		assert.True(t, token.Valid, "Token should be valid")
		assert.Equal(t, uuid, claims["sub"], "Token sub claim should match the input uuid")
		assert.Equal(t, 6, len(claims), "Token Must have all 6 claims")
	})

	t.Run("Should comapre the first token with a second generated with a different UUID", func(t *testing.T) {
		setUp()

		uuid := "test-uuid"
		tokenString, _ := utils.GenerateJwtToken(uuid)

		uuid2 := "another-uuid"
		tokenString2, err := utils.GenerateJwtToken(uuid2)

		assert.Nil(t, err, "Error should be nil")
		assert.NotEqual(t, tokenString, tokenString2, "Tokens from different UUIDs should not be equals")
	})
}

func Test_GenerateRefreshJwtToken(t *testing.T) {
	setUp := func() {
		test.LoadEnv()

		assert.NotEmpty(t, os.Getenv("JWT_SECRET"), "JWT_SECRET should not be empty")
	}

	t.Run("Should return a valid token string", func(tt *testing.T) {
		setUp()

		uuid := "test-uuid"
		accessToken := "access-token"
		tokenString, err := utils.GenerateRefreshJwtToken(uuid, accessToken)

		assert.Nil(tt, err, "Error should be nil")
		assert.NotZero(tt, tokenString, "TokenString should not be zero")
	})

	t.Run("Should parse and verify token claims", func(tt *testing.T) {
		setUp()

		uuid := "test-uuid"
		accessToken := "access-token"
		tokenString, _ := utils.GenerateRefreshJwtToken(uuid, accessToken)

		claims := jwt.MapClaims{}
		token, err := utils.ParseTokenWithClaims(tokenString, &claims)

		assert.Nil(tt, err, "Error should be nil")
		assert.True(tt, token.Valid, "Token should be valid")
		assert.Equal(tt, accessToken, claims["sub"], "Token sub claim should match the input accessToken")
		assert.Equal(tt, uuid, claims["uuid"], "Token uuid claim should match the input uuid")
		assert.NotEmpty(tt, claims["jti"], "JTI claim should not be empty")
		assert.Equal(tt, 8, len(claims), "Token Must have all 8 claims")
	})

	t.Run("Should comapre the first token with a second generated with a different UUID and Access token", func(tt *testing.T) {
		setUp()

		uuid := "test-uuid"
		accessToken := "access-token"
		tokenString, _ := utils.GenerateRefreshJwtToken(uuid, accessToken)

		uuid2 := "another-uuid"
		accessToken2 := "another-acess-token"
		tokenString2, _ := utils.GenerateRefreshJwtToken(uuid2, accessToken2)

		assert.NotEqual(tt, tokenString, tokenString2, "Tokens from different UUIDs should not be equals")
	})
}

func Test_SignedString(t *testing.T) {
	test.LoadEnv()

	assert.NotEmpty(t, os.Getenv("JWT_SECRET"), "JWT_SECRET should not be empty")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{})

	tokenString, err := utils.SignedString(token)

	assert.Nil(t, err, "Error should be nil")
	assert.NotEmpty(t, tokenString, "TokenString should not be empty")
}

func Test_ParseToken(t *testing.T) {
	setUp := func() {
		test.LoadEnv()

		assert.NotEmpty(t, os.Getenv("JWT_SECRET"), "JWT_SECRET should not be empty")
	}

	t.Run("Should parse a invalid token string and return a error", func(tt *testing.T) {
		setUp()

		invalidToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"nbf": jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
		})
		invalidTokenString, _ := utils.SignedString(invalidToken)

		invalidTokenParse, err := utils.ParseToken(invalidTokenString)

		assert.NotNil(tt, err, "Error should not be nil")
		assert.Nil(tt, invalidTokenParse, "Invalid token parsed should be nil")
	})

	t.Run("Should parse a valid token string and return a token", func(tt *testing.T) {
		setUp()

		validToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"nbf": jwt.NewNumericDate(time.Now()),
		})
		validTokenString, _ := utils.SignedString(validToken)

		validTokenParse, err := utils.ParseToken(validTokenString)

		assert.Nil(tt, err, "Error should be nil")
		assert.NotNil(tt, validTokenParse, "Valid token parsed should not be nil")
	})
}

func Test_ParseTokenWithClaims(t *testing.T) {

	setUp := func() {
		test.LoadEnv()

		assert.NotEmpty(t, os.Getenv("JWT_SECRET"), "JWT_SECRET should not be empty")
	}

	t.Run("Should parse a invalid token string with claims and return a error", func(tt *testing.T) {
		setUp()

		invalidToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"nbf": jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
		})
		invalidTokenString, _ := utils.SignedString(invalidToken)
		invalidTokenClaims := jwt.MapClaims{}

		invalidTokenParse, err := utils.ParseTokenWithClaims(invalidTokenString, &invalidTokenClaims)

		assert.NotNil(tt, err, "Error should not be nil")
		assert.NotZero(tt, invalidTokenClaims["nbf"], "nbf claim should not be empty")
		assert.Nil(tt, invalidTokenParse, "Invalid token parsed should be nil")
	})

	t.Run("Should parse a valid token string wtih claims and return a token", func(tt *testing.T) {
		setUp()

		validToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"nbf": jwt.NewNumericDate(time.Now()),
		})
		validTokenString, _ := utils.SignedString(validToken)
		validTokenClaims := jwt.MapClaims{}

		validTokenParse, err := utils.ParseTokenWithClaims(validTokenString, &validTokenClaims)

		assert.Nil(tt, err, "Error should be nil")
		assert.NotZero(tt, validTokenClaims["nbf"], "nbf claim should not be empty")
		assert.NotNil(tt, validTokenParse, "Valid token should not be nil")
	})
}
