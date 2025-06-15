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
	connectionString := "host=localhost user=teste password=PassW0rd dbname=notifsystemdb port=5432 sslmode=disable" // Replace with your actual connection string

	webServer, err := service.CreateWebServer(port)
	if err != nil {
		return err
	}

	// Set up the logger
	logger := logger.NewLogger()
	// webServer.SetLogger(logger)

	logger.Info("Mounting Dependencies")

	// Initialize the infrastructure container
	infraContainer, err := container.NewInfraContainer(connectionString)
	if err != nil {
		logger.Error("Failed to initialize infrastructure container: " + err.Error())
		return err
	}

	// Initialize the application container
	applicationContainer, err := container.NewApplicationContainer(infraContainer)
	if err != nil {
		logger.Error("Failed to initialize application container: " + err.Error())
		return err
	}

	// Get Controllers
	entryPointContainer, err := container.NewEntryPointContainer(applicationContainer)
	if err != nil {
		logger.Error("Failed to initialize entrypoint container: " + err.Error())
		return err
	}
	// Initialize routes
	webServer.InitalizeRoutes(entryPointContainer.GetControllers()...)

	// Start the server
	logger.Info("Starting server on port " + port)
	return webServer.Start()
}
