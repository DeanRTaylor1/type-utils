package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type GitRepo struct {
	URL  string `yaml:"url"`
	Path string `yaml:"path"`
}

type Config struct {
	Version        string   `yaml:"version"`
	Language       string   `yaml:"language"`
	SchemasDirName string   `yaml:"schemasDirName,omitempty"`
	GitRepo        *GitRepo `yaml:"git_repo,omitempty"`
}

// Read reads and parses the type-utils.yaml file
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
