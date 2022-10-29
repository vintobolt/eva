package models

import "github.com/golang-jwt/jwt"

// JwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

type Token struct {
	Token string `json:"token"`
}
