package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Version        string `yaml:"version"`
	Language       string `yaml:"language"`
	SchemasDirName string `yaml:"schemasDirName,omitempty"`
}

// readConfig reads and parses the type-utils.yaml file
func Read(filename string) (*Config, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(buf, config)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	return config, nil
}
