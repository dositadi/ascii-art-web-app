package middleware

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
)

func AuthenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Step one: get the jwtToken from the authorization header and check if its empty
		jwtToken := r.Header.Get("Authorization")

		if jwtToken == "" {
			http.Error(w, h.UNAUTHORIZED_ERR_DETAIL, http.StatusUnauthorized)
			return
		}

		// Step two: Trim the bearer off from the jwtToken and ensure that the result string is not the same as the former
		trimmedJWT := strings.TrimPrefix(jwtToken, "Bearer ")
		if trimmedJWT == jwtToken {
			http.Error(w, h.UNAUTHORIZED_ERR_DETAIL, http.StatusUnauthorized)
			return
		}

		// Step three: Declare the claims variable to hold user details
		var activeUser m.ActiveUser

		// Step five: parse the trimmed token with the claims variable and the function to fetch the secret key from the terminal
		token, err := jwt.ParseWithClaims(trimmedJWT, &activeUser, func(t *jwt.Token) (any, error) {
			if t.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("Signing method mismatch!")
			}
			return os.Getenv("ASCII_JWT_SECRET_KEY"), nil
		})
		if err != nil {
			http.Error(w, h.UNAUTHORIZED_ERR_DETAIL, http.StatusUnauthorized)
			return
		}

		// Step six: Check the type of the recieved claims and also ensure that the claims is valid
		if _, ok := token.Claims.(*m.ActiveUser); !ok && !token.Valid {
			http.Error(w, h.UNAUTHORIZED_ERR_DETAIL, http.StatusUnauthorized)
			return
		}

		// Step seven: create a context with value to pass the user_id forward to the next server
		ctx := context.WithValue(r.Context(), "user_id", activeUser.ID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetToken(r *http.Request) string {
	jwtToken := r.Header.Get("Authorization")

	if jwtToken != "" {
		token := strings.TrimPrefix(jwtToken, "Bearer ")
		return token
	}

	cookie, err := r.Cookie("access-token")
	if err == nil {
		return cookie.Value
	}
	return ""
}
