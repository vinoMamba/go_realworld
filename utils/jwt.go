package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vinoMamba.com/go_realworld/config"
)

func GenerateJWT(email, username string) (string, error) {
	iat := time.Now()
	exp := iat.Add(24 * time.Second)

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

func VerifyJWT(token string) (*jwt.MapClaims, bool, error) {
	var claim jwt.MapClaims
	t, err := jwt.ParseWithClaims(token, &claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetJWTSecret()), nil
	})
	if err != nil {
		return nil, false, err
	}
	if t.Valid {
		fmt.Printf("claim:%v\n", claim)
		return nil, true, nil
	} else {
		return &claim, false, nil
	}
}
