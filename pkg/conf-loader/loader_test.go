package confloader

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestConfig struct {
	Name string `mapstructure:"name" validate:"required"`
	Port int    `mapstructure:"port" validate:"required,min=1"`
}

func writeTempConfigFile(t *testing.T, content string) string {
	tmpfile, err := os.CreateTemp("", "testconfig-*.yaml")
	assert.NoError(t, err)
	_, err = tmpfile.Write([]byte(content))
	assert.NoError(t, err)
	tmpfile.Close()
	return tmpfile.Name()
}

func Test_loadConfigFromFile_success(t *testing.T) {
	content := `
name: "myapp"
port: 8080
`
	path := writeTempConfigFile(t, content)
	defer os.Remove(path)

	cfg, err := LoadConfigFromFile[TestConfig](path)
	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, "myapp", cfg.Name)
	assert.Equal(t, 8080, cfg.Port)
}

func Test_loadConfigFromFile_fileNotFound(t *testing.T) {
	_, err := LoadConfigFromFile[TestConfig]("/non/existent/path.yaml")
	assert.Error(t, err)
}

func Test_loadConfigFromFile_invalidConfig(t *testing.T) {
	content := `
name: ""
port: 0
`
	path := writeTempConfigFile(t, content)
	defer os.Remove(path)

	_, err := LoadConfigFromFile[TestConfig](path)
	assert.Error(t, err)
}

func Test_loadConfigFromFile_unmarshalling_config(t *testing.T) {
	content := `
name: "myapp"
port: "1x"
`
	path := writeTempConfigFile(t, content)
	defer os.Remove(path)

	_, err := LoadConfigFromFile[TestConfig](path)
	assert.Error(t, err)
}

func Test_LoadConfig_success(t *testing.T) {
	content := `
name: "myapp"
port: 8080
`
	path := writeTempConfigFile(t, content)
	defer os.Remove(path)

	os.Setenv("CONFIG_PATH", path)
	defer os.Unsetenv("CONFIG_PATH")

	cfg, err := LoadConfig[TestConfig]()
	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, "myapp", cfg.Name)
	assert.Equal(t, 8080, cfg.Port)
}

func Test_LoadConfig_envNotSet(t *testing.T) {
	os.Unsetenv("CONFIG_PATH")
	_, err := LoadConfig[TestConfig]()
	assert.Error(t, err)
}
