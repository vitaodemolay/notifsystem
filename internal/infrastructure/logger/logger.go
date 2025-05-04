package logger

import (
	"log/slog"
	"time"

	"github.com/go-chi/httplog/v2"
)

func NewLogger() *httplog.Logger {
	logger := httplog.NewLogger("api", httplog.Options{
		LogLevel:        slog.LevelDebug,
		RequestHeaders:  false,
		ResponseHeaders: false,
		//JSON:            true,
		// MessageFieldName: "message",
		// LevelFieldName:   "severity",
		TimeFieldName: time.RFC3339,
		// Tags: map[string]string{
		// 	"version": "v1.0",
		// 	"env":     "dev",
		// },
		QuietDownRoutes: []string{
			"/",
			"/ping",
		},
		QuietDownPeriod: 10 * time.Second,
	})
	return logger
}
