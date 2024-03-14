package routes

import (
	"github.com/labstack/echo/v4"
)

var JWTSecretKey = []byte("tottenhamhotspurfootballclub")

func SetupRouter() *echo.Echo {
	e := echo.New()

	configureMiddleware(e)
	configureRoutes(e)
	return e
}

func configureRoutes(e *echo.Echo) {
	e.GET("/", canaryHandler)
	e.GET("/user/create", userCreateHandler)
	e.POST("/login", loginHandler)
}

func configureMiddleware(e *echo.Echo) {
	// TODO you are here
	// add a /login route posting email and pass and returning nothing
	// then implement issueing jwt tokens
	// then implement middleware and secure hello world with jwt

}
