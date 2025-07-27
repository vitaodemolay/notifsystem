package configs

type serverConfig struct {
	Port                  string `mapstructure:"port"`
	Host                  string `mapstructure:"host"`
	CustomLoggerIsEnabled bool   `mapstructure:"custom_logger_enabled"`
}

type databaseConfig struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     string `mapstructure:"port" validate:"required"`
	User     string `mapstructure:"user" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
	DbName   string `mapstructure:"dbname" validate:"required"`
	SslMode  string `mapstructure:"sslmode"`
}

type identityProviderConfig struct {
	ClientID    string `mapstructure:"client_id" validate:"required"`
	RedirectURL string `mapstructure:"redirect_uri" validate:"required"`
	TokenType   string `mapstructure:"token_type"`
}

type Config struct {
	ServerConfig           serverConfig           `mapstructure:"server"`
	DatabaseConfig         databaseConfig         `mapstructure:"database" validate:"required"`
	IdentityProviderConfig identityProviderConfig `mapstructure:"identity_provider" validate:"required"`
}

func (c *Config) GetDatabaseConnectionString() string {
	if c.DatabaseConfig.SslMode == "" {
		c.DatabaseConfig.SslMode = "disable"
	}

	return "host=" + c.DatabaseConfig.Host +
		" user=" + c.DatabaseConfig.User +
		" password=" + c.DatabaseConfig.Password +
		" dbname=" + c.DatabaseConfig.DbName +
		" port=" + c.DatabaseConfig.Port +
		" sslmode=" + c.DatabaseConfig.SslMode
}

func (c *Config) GetPort() string {
	if c.ServerConfig.Port == "" {
		c.ServerConfig.Port = "8080"
	}
	return c.ServerConfig.Port
}

func (c *Config) GetHost() string {
	if c.ServerConfig.Host == "" {
		c.ServerConfig.Host = "localhost"
	}
	return c.ServerConfig.Host
}

func (c *Config) IsCustomLoggerEnabled() bool {
	return c.ServerConfig.CustomLoggerIsEnabled
}

func (c *Config) GetIdentityProviderClientID() string {
	return c.IdentityProviderConfig.ClientID
}

func (c *Config) GetIdentityProviderRedirectURL() string {
	return c.IdentityProviderConfig.RedirectURL
}

func (c *Config) GetIdentityProviderTokenType() string {
	if c.IdentityProviderConfig.TokenType == "" {
		c.IdentityProviderConfig.TokenType = "Bearer"
	}
	return c.IdentityProviderConfig.TokenType
}
