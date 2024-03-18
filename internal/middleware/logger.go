package middleware

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

// StartLogger creates a structured logger for the application
func StartLogger() {
	lvl := new(slog.LevelVar)
	lvl.Set(slog.LevelDebug)
	Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}
