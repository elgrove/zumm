package routes

import (
	"zumm/models"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var JWTSecretKey = []byte("tottenhamhotspurfootballclub")

func SetupRouter() *echo.Echo {
	e := echo.New()
	configureRoutes(e)
	return e
}

func configureRoutes(e *echo.Echo) {
	jwtConfig := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(models.UserClaims)
		},
		SigningKey: JWTSecretKey,
	}
	jwtMiddleware := echojwt.WithConfig(jwtConfig)

	e.GET("/", canaryHandler)
	e.GET("/user/create", userCreateHandler)
	e.POST("/login", loginHandler)
	e.POST("/discover", discoverHandler, jwtMiddleware)

}
