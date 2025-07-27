package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
	confloader "github.com/vitaodemolay/notifsystem/pkg/conf-loader"
)

func Test_Config_Methods(t *testing.T) {
	configBase := &Config{
		ServerConfig: serverConfig{
			Port:                  "1000",
			Host:                  "host.test.com",
			CustomLoggerIsEnabled: true,
		},
		DatabaseConfig: databaseConfig{
			Host:     "localhost",
			Port:     "5432",
			User:     "userX",
			Password: "passwordX",
			DbName:   "dbnameX",
			SslMode:  "xyz",
		},
		IdentityProviderConfig: identityProviderConfig{
			ClientID:    "client_id",
			RedirectURL: "http://localhost/callback",
			TokenType:   "Tearer",
		},
	}

	t.Run("GetDatabaseConnectionString when sslMode is default value", func(t *testing.T) {
		// Arrange
		expected := "host=localhost user=userX password=passwordX dbname=dbnameX port=5432 sslmode=disable"
		_config := *configBase
		_config.DatabaseConfig.SslMode = ""

		// Act
		actual := _config.GetDatabaseConnectionString()

		// Assert
		assert.Equal(t, expected, actual, "Expected connection string to match")
	})

	t.Run("GetDatabaseConnectionString when sslMode is set", func(t *testing.T) {
		// Arrange
		expected := "host=localhost user=userX password=passwordX dbname=dbnameX port=5432 sslmode=xyz"
		_config := *configBase

		// Act
		actual := _config.GetDatabaseConnectionString()

		// Assert
		assert.Equal(t, expected, actual, "Expected connection string to match")
	})
	t.Run("GetPort when port is default value", func(t *testing.T) {
		// Arrange
		expected := "8080"
		_config := *configBase
		_config.ServerConfig.Port = ""

		// Act
		actual := _config.GetPort()

		// Assert
		assert.Equal(t, expected, actual, "Expected port to match")
	})
	t.Run("GetPort when port is set", func(t *testing.T) {
		// Arrange
		expected := "1000"
		_config := *configBase

		// Act
		actual := _config.GetPort()

		// Assert
		assert.Equal(t, expected, actual, "Expected port to match")
	})
	t.Run("GetHost when host is default value", func(t *testing.T) {
		// Arrange
		expected := "localhost"
		_config := *configBase
		_config.ServerConfig.Host = ""

		// Act
		actual := _config.GetHost()

		// Assert
		assert.Equal(t, expected, actual, "Expected host to match")
	})
	t.Run("GetHost when host is set", func(t *testing.T) {
		// Arrange
		expected := "host.test.com"
		_config := *configBase

		// Act
		actual := _config.GetHost()

		// Assert
		assert.Equal(t, expected, actual, "Expected host to match")
	})
	t.Run("IsCustomLoggerEnabled when custom logger is enabled", func(t *testing.T) {
		// Arrange
		expected := true
		_config := *configBase

		// Act
		actual := _config.IsCustomLoggerEnabled()

		// Assert
		assert.Equal(t, expected, actual, "Expected custom logger enabled to be true")
	})
	t.Run("IsCustomLoggerEnabled when custom logger is disabled", func(t *testing.T) {
		// Arrange
		expected := false
		_config := *configBase
		_config.ServerConfig.CustomLoggerIsEnabled = false

		// Act
		actual := _config.IsCustomLoggerEnabled()

		// Assert
		assert.Equal(t, expected, actual, "Expected custom logger enabled to be false")
	})
	t.Run("GetIdentityProviderClientID", func(t *testing.T) {
		// Arrange
		expected := "client_id"
		_config := *configBase

		// Act
		actual := _config.GetIdentityProviderClientID()

		// Assert
		assert.Equal(t, expected, actual, "Expected identity provider client ID to match")
	})
	t.Run("GetIdentityProviderRedirectURL", func(t *testing.T) {
		// Arrange
		expected := "http://localhost/callback"
		_config := *configBase

		// Act
		actual := _config.GetIdentityProviderRedirectURL()

		// Assert
		assert.Equal(t, expected, actual, "Expected identity provider redirect URL to match")
	})
	t.Run("GetIdentityProviderTokenTyp when tokenType is default", func(t *testing.T) {
		// Arrange
		expected := "Bearer"
		_config := *configBase
		_config.IdentityProviderConfig.TokenType = ""

		// Act
		actual := _config.GetIdentityProviderTokenType()

		// Assert
		assert.Equal(t, expected, actual, "Expected identity provider token type to be Bearer")
	})
	t.Run("GetIdentityProviderTokenType when tokenType is set", func(t *testing.T) {
		// Arrange
		expected := "Tearer"
		_config := *configBase

		// Act
		actual := _config.GetIdentityProviderTokenType()

		// Assert
		assert.Equal(t, expected, actual, "Expected identity provider token type to be Tearer")
	})
}

func Test_Config_With_Loader(t *testing.T) {
	configPath := "../../../resources/local-configs.yml"

	t.Run("should load config from file successfully", func(t *testing.T) {
		config, err := confloader.LoadConfigFromFile[Config](configPath)
		assert.NoError(t, err, "Expected no error when loading config from file")
		assert.NotNil(t, config, "Expected config to be loaded successfully")
	})
}
