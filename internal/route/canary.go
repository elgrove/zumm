package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CanaryHandler(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
