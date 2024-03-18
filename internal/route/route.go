// Package routes instantiates and defines the API endpoint router for this application.
// Endpoints requiring authentication have been secured with JWT using Echo's built-in middleware
package route

import (
	"zumm/internal/middleware"

	"github.com/labstack/echo/v4"
)

func SetupRouter() *echo.Echo {
	middleware.StartLogger()
	e := echo.New()
	configureRoutes(e)
	return e
}

func configureRoutes(e *echo.Echo) {
	jwtMiddleware := middleware.JWTMiddleware()
	e.GET("/", CanaryHandler)
	e.GET("/user/create", UserCreateHandler)
	e.POST("/login", LoginHandler)
	e.POST("/discover", DiscoverHandler, jwtMiddleware)
	e.POST("/swipe", SwipeHandler, jwtMiddleware)
}
