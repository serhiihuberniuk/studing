package google_provider

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	CredentialsFile string `yaml:"credentials_file"`
}

func ReadConfig(pathConfig string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(pathConfig)
	if err != nil {
		return nil, fmt.Errorf("error occured while opening config file: %w", err)
	}
	defer file.Close()

	if err := yaml.NewDecoder(file).Decode(config); err != nil {
		return nil, fmt.Errorf("error while decoding config file: %w", err)
	}

	return config, nil
}
