package middleware

import (
	"net/http"
	"strings"

	"github.com/Miguelburitica/go-rest-server/helpers"
	"github.com/Miguelburitica/go-rest-server/server"
)

var (
	NO_AUTH_NEEDED = []string{
		"login",
		"signup",
	}
)

func shouldCheckToken(route string) bool {
	for _, p := range NO_AUTH_NEEDED {
		if strings.Contains(route, p) {
			return false
		}
	}
	return true
}

func CheckAuthMiddleware(s server.Server) func (h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !shouldCheckToken(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}
			if _, _, err := helpers.SecurityBasic(s, w, r); err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			
			next.ServeHTTP(w, r)
		})
	}
}