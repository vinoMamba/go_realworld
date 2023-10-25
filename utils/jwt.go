package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/vinoMamba.com/go_realworld/config"
)

func GenerateJWT(email, username string) (string, error) {
	jwtKey := []byte(config.GetJWTSecret())
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": map[string]string{
			"email":    email,
			"username": username,
		},
	})
	return t.SignedString(jwtKey)
}
