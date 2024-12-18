package middlewares

import (
	"fmt"
	"net/http"
	"slices"
)

var skipUrls []string = []string{
	"/signup",
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if slices.Contains(skipUrls, r.URL.Path) {
			fmt.Printf("skip %s\n", r.URL.Path)
			next.ServeHTTP(w, r)
			return
		}

		if r.Header["Authorization"] == nil {

		}

		next.ServeHTTP(w, r)
	})
}
