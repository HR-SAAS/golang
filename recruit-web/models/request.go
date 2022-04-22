package models

import "github.com/golang-jwt/jwt/v4"

type CustomClaims struct {
	ID       int64
	NickName string
	jwt.RegisteredClaims
}
