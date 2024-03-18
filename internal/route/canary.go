package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CanaryHandler is a simple healthcheck endpoint returning 200.
func CanaryHandler(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
