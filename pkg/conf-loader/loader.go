package confloader

import (
	"log"
	"os"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

const (
	enviromentVariablePath = "CONFIG_PATH"
)

func LoadConfig[T any]() (*T, error) {
	configPath := os.Getenv(enviromentVariablePath)
	if configPath == "" {
		log.Println("Environment variable CONFIG_PATH is not set with config file location.")
		return nil, os.ErrNotExist
	}

	return LoadConfigFromFile[T](configPath)
}

func LoadConfigFromFile[T any](path string) (*T, error) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file. path: %v\n - error: %v\n", path, err)
		return nil, err
	}

	var config T
	if err := viper.Unmarshal(&config); err != nil {
		log.Printf("Error unmarshalling config file. path: %v\n - error: %v", path, err)
		return nil, err
	}

	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		log.Printf("Configuration is not valid - error: %v\n", err)
		return nil, err
	}

	return &config, nil
}
