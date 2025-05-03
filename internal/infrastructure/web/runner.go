package web

import (
	"context"

	"github.com/vitaodemolay/notifsystem/internal/infrastructure/container"
	"github.com/vitaodemolay/notifsystem/internal/infrastructure/logger"
	"github.com/vitaodemolay/notifsystem/internal/infrastructure/web/service"
)

func Run(ctx context.Context) error {
	// Initialize the web server

	// Set the port for the web server (TODO: make this configurable)
	port := ":8080"

	webServer, err := service.CreateWebServer(port)
	if err != nil {
		return err
	}

	// Set up the logger
	logger := logger.NewLogger()
	webServer.SetLogger(logger)

	logger.Info("Mounting Dependencies")

	// Get Controllers
	entryPointContainer, err := container.NewEntryPointContainer()
	if err != nil {
		return err
	}
	// Initialize routes
	webServer.InitalizeRoutes(entryPointContainer.GetControllers()...)

	// Start the server
	logger.Info("Starting server on port " + port)
	return webServer.Start()
}
