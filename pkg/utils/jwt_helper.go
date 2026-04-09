package utils

import (
	"os"
	"time"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessJWT(user m.ActiveUser) (string, *m.Error) {
	secretKey := os.Getenv("ASCII_JWT_SECRET_KEY")

	user.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(4 * time.Hour))),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)

	jwt, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", &m.Error{
			Error:   UNAUTHORIZED_ERR,
			Details: UNAUTHORIZED_ERR_DETAIL,
			Code:    UNAUTHORIZED_ERR_CODE,
		}
	}
	return jwt, nil
}

func GenerateRefreshJWT(user m.ActiveUser) (string, *m.Error) {
	secretKey := os.Getenv("ASCII_JWT_SECRET_KEY")

	user.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(168 * time.Hour))),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)

	jwt, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", &m.Error{
			Error:   UNAUTHORIZED_ERR,
			Details: UNAUTHORIZED_ERR_DETAIL,
			Code:    UNAUTHORIZED_ERR_CODE,
		}
	}
	return jwt, nil
}
