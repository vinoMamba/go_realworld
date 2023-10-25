package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vinoMamba.com/go_realworld/config"
)

func GenerateJWT(email, username string) (string, error) {
	iat := time.Now()
	exp := iat.Add(24 * time.Hour)

	jwtKey := []byte(config.GetJWTSecret())
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": map[string]string{
			"email":    email,
			"username": username,
		},
		"iat": iat.Unix(),
		"exp": exp.Unix(),
	})
	return t.SignedString(jwtKey)
}
