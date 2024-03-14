package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func canaryHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"hello": "world"})
}
