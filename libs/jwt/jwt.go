package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct{}

// returns jwt token in string format
func Generate(payload string, exp time.Time, JWTSecret string) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   payload,
		ExpiresAt: jwt.NewNumericDate(exp),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// returns payload in string format
func Parse(tokenString string, JWTSecret string) (string, error) {
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecret), nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", errors.New("invalid token")
	}

	return claims.Subject, nil
}

// returns a new jwt token in string format or an error if the old token is still valid for more than 30 seconds
func Refresh(tokenString string, exp time.Time, JWTSecret string) (string, error) {
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecret), nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", errors.New("invalid token")
	}

	if time.Until(claims.ExpiresAt.Time) > 30*time.Second {
		return "", errors.New("token is not expired yet")
	}

	return Generate(claims.Subject, exp, JWTSecret)
}
