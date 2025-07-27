package web

import (
	"context"

	"github.com/vitaodemolay/notifsystem/internal/infrastructure/configs"
	"github.com/vitaodemolay/notifsystem/internal/infrastructure/container"
	"github.com/vitaodemolay/notifsystem/internal/infrastructure/logger"
	"github.com/vitaodemolay/notifsystem/internal/infrastructure/web/service"
	loader "github.com/vitaodemolay/notifsystem/pkg/conf-loader"
)

func Run(ctx context.Context) error {
	//Load configurations
	configs, err := loader.LoadConfig[configs.Config]()
	if err != nil {
		return err
	}

	// Initialize the web server
	port := ":" + configs.GetPort()
	webServer, err := service.CreateWebServer(port)
	if err != nil {
		return err
	}

	// Set up the logger
	logger := logger.NewLogger()
	if configs.IsCustomLoggerEnabled() {
		webServer.SetLogger(logger)
	}

	logger.Info("Mounting Dependencies")

	// Initialize the infrastructure container
	infraContainer, err := container.NewInfraContainer(configs.GetDatabaseConnectionString())
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
	entryPointContainer, err := container.NewEntryPointContainer(
		applicationContainer,
		configs.GetIdentityProviderClientID(),
		configs.GetIdentityProviderRedirectURL(),
		configs.GetIdentityProviderTokenType(),
	)
	if err != nil {
		logger.Error("Failed to initialize entrypoint container: " + err.Error())
		return err
	}
	// Initialize routes
	webServer.InitalizeRoutes(entryPointContainer.GetControllers()...)

	// Start the server
	logger.Info("Starting server on host http://" + configs.GetHost() + port)
	return webServer.Start()
}
