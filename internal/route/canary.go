package route

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CanaryHandler is a simple healthcheck endpoint returning 200.
func CanaryHandler(c echo.Context) error {
	slog.Debug("Requested received to canary endpoint")
	return c.NoContent(http.StatusOK)
}
