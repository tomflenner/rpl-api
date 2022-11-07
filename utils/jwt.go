package utils

import (
	"time"

	"github.com/b4cktr4ck5r3/rpl-api/config"
	"github.com/b4cktr4ck5r3/rpl-api/models"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(payload string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &models.Claims{
		SteamId: payload,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	var tokenString string
	var err error
	tokenString, err = token.SignedString([]byte(config.Cfg.JwtSecret))

	return tokenString, err
}

func VerifyJWT(tokenString string) (bool, error) {
	claims := &models.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.JwtSecret), nil
	})

	if err != nil || !token.Valid {
		return false, err
	}

	return true, nil
}
