package route

import (
	"net/http"
	"zumm/internal/middleware"

	"github.com/labstack/echo/v4"
)

// CanaryHandler is a simple healthcheck endpoint returning 200.
func CanaryHandler(c echo.Context) error {
	middleware.Logger.Debug("Requested received to canary endpoint")
	return c.NoContent(http.StatusOK)
}
