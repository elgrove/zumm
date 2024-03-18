package middleware

import (
	"zumm/internal/model"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// JWTokenSecretKey stores the secret key used to encrypt and decrypt JWTs.
// Obviously this is an egrigeous security risk, the key should be injected as an env var.
var JWTokenSecretKey = []byte("tottenhamhotspurfootballclub")

// JWTokenConfig provides the configuration for our JWT middleware using echojwt.
func JWTokenConfig() echojwt.Config {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(model.UserClaims)
		},
		SigningKey: JWTokenSecretKey,
	}
	return config
}

// JWTMiddleware provides an echo framework middleware which protects routes on which it
// is used with JWT authentication
func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(JWTokenConfig())
}
