package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	// DataStore specifies the type of data store to use
	DataStore string `yaml:"dataStore"`
}

func LoadConfig() (*Config, error) {
	// Define the default configuration file path
	defaultConfigPath := "../../internal/config/config.yaml"

	// Check if a custom configuration file path is provided via an environment variable
	configPath := os.Getenv("CONFIG_FILE_PATH")
	if configPath == "" {
		// Use the default configuration file path
		configPath = defaultConfigPath
	}

	// Resolve the absolute path to the configuration file
	absConfigPath, err := filepath.Abs(configPath)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(absConfigPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
