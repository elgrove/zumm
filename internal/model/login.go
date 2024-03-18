package model

import (
	"github.com/golang-jwt/jwt/v5"
)

// LoginRequest holds the email and password the user wishes to log in with.
type LoginRequest struct {
	Email    string
	Password string
}

// LoginResponse holds the JWToken issued for a successful login.
type LoginResponse struct {
	Token string
}

// UserClaims is a type of custom JWT claims which includes information about the.
// user bearing the token
type UserClaims struct {
	jwt.RegisteredClaims
	User
}
