package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/utils"
)

var skipUrls []string = []string{
	"/signup", "/signin",
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if slices.Contains(skipUrls, r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}

		if r.Header["Authorization"] == nil {
			sendInvalidTokenError(w)
			return
		}

		var tokenString string = r.Header["Authorization"][0]
		splitTokenString := strings.Split(tokenString, " ")

		if len(splitTokenString) < 2 || strings.Compare(splitTokenString[0], "Bearer") != 0 {
			sendInvalidTokenError(w)
			return
		}

		if token := parseToken(w, splitTokenString[1]); token != nil {
			claims := token.Claims.(jwt.MapClaims)
			if r.URL.Path == "/refresh" {
				claims = getTokenClaimsUnverified(w, claims["sub"].(string))
			}
			r.Header.Add("Uuid", claims["sub"].(string))
		}

		next.ServeHTTP(w, r)
	})
}

func getTokenClaimsUnverified(w http.ResponseWriter, expiredToken string) jwt.MapClaims {
	token, _, err := new(jwt.Parser).ParseUnverified(expiredToken, jwt.MapClaims{})
	if err != nil {
		fmt.Printf("error parsing token: %v\n", err)
		sendInvalidTokenError(w)
		return nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if _, exists := claims["sub"].(string); !exists {
			sendInvalidTokenError(w)
			return nil
		}
		return claims
	} else {
		sendInvalidTokenError(w)
	}
	return nil
}

func parseToken(w http.ResponseWriter, tokenString string) *jwt.Token {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWTSECRET")), nil
	})

	if err != nil {
		fmt.Printf("error parsing token: %v\n", err)
		sendInvalidTokenError(w)
		return nil
	}
	return token
}

func sendInvalidTokenError(w http.ResponseWriter) {
	fmt.Printf("error %v\n", apperrors.ErrInvalidToken)
	utils.SendError(w, &utils.RequestError{
		StatusCode: 403,
		Err:        apperrors.ErrForbidden,
		Message:    apperrors.ErrInvalidToken.Error(),
	})
}
