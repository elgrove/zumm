// Package routes instantiates and defines the API endpoint router for this application.
// Endpoints requiring authentication have been secured with JWT using Echo's built-in middleware
package route

import (
	"zumm/internal/model"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var JWTokenSecretKey = []byte("tottenhamhotspurfootballclub")

func SetupRouter() *echo.Echo {
	e := echo.New()
	configureRoutes(e)
	return e
}

func JWTokenConfig() echojwt.Config {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(model.UserClaims)
		},
		SigningKey: JWTokenSecretKey,
	}
	return config
}

func configureRoutes(e *echo.Echo) {
	jwtMiddleware := echojwt.WithConfig(JWTokenConfig())

	e.GET("/", CanaryHandler)
	e.GET("/user/create", UserCreateHandler)
	e.POST("/login", LoginHandler)
	e.POST("/discover", DiscoverHandler, jwtMiddleware)
	e.POST("/swipe", SwipeHandler, jwtMiddleware)
}
