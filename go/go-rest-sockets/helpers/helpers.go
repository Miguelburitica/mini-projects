package helpers

import (
	"net/http"
	"strings"

	"github.com/Miguelburitica/go-rest-server/models"
	"github.com/Miguelburitica/go-rest-server/server"
	"github.com/golang-jwt/jwt"
)

func SecurityBasic(s server.Server, w http.ResponseWriter, r *http.Request) (*models.AppClaims, bool, error) {
	tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
	token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Config().JWTSecret), nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return nil, false, err
	}

	if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
		return claims, ok, nil
	}

	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	return nil, false, err
}
