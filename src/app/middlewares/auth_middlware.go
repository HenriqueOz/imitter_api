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

		token, err := jwt.Parse(splitTokenString[1], func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWTSECRET")), nil
		})

		if err != nil {
			fmt.Printf("error parsing token: %v\n", err)
			sendInvalidTokenError(w)
			return
		}

		sub, _ := token.Claims.GetSubject()
		r.Header.Add("id", sub)

		next.ServeHTTP(w, r)
	})
}

func sendInvalidTokenError(w http.ResponseWriter) {
	fmt.Printf("error %v\n", apperrors.ErrInvalidToken)
	utils.SendError(w, &utils.RequestError{
		StatusCode: 403,
		Err:        apperrors.ErrForbidden,
		Message:    apperrors.ErrInvalidToken.Error(),
	})
}
