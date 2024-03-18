package model

import (
	"github.com/golang-jwt/jwt/v5"
)

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	Token string
}

type UserClaims struct {
	jwt.RegisteredClaims
	User
}
