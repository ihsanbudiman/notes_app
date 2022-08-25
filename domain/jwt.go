package domain

import "github.com/golang-jwt/jwt/v4"

type TokenClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}
