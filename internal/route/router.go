package route

import (
	"zumm/internal/model"

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
			return new(model.UserClaims)
		},
		SigningKey: JWTSecretKey,
	}
	jwtMiddleware := echojwt.WithConfig(jwtConfig)

	e.GET("/", CanaryHandler)
	e.GET("/user/create", UserCreateHandler)
	e.POST("/login", LoginHandler)
	e.POST("/discover", DiscoverHandler, jwtMiddleware)
	e.POST("/swipe", SwipeHandler, jwtMiddleware)

}
