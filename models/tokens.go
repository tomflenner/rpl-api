package models

import "github.com/golang-jwt/jwt"

type TokensPayload struct {
	SteamID string `json:"steamId"`
}

type Claims struct {
	SteamId string `json:"steamId"`
	jwt.StandardClaims
}
