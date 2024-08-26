package config

import (
	"fmt"

	"github.com/deanrtaylor1/type-utils/constants"
)

type Config struct {
	Version              string   `yaml:"version"`
	Language             string   `yaml:"language"`
	SchemasDirName       string   `yaml:"schemas_dir_name,omitempty"`
	GitRepo              *GitRepo `yaml:"git_repo,omitempty"`
	GenerateConstructors bool     `yaml:"generate_constructors,omitempty"`
}

func (c *Config) GetVersion() string {
	return c.Version
}

func (c *Config) GetLanguage() string {
	return c.Language
}

func (c *Config) GetSchemasDirName() string {
	if c.SchemasDirName != "" {
		return fmt.Sprintf("./%s", c.SchemasDirName)
	}
	return fmt.Sprintf("./%s", constants.DEFAULT_SCHEMA_DIR_NAME)
}

func (c *Config) GetGitRepo() GitRepoConfiger {
	return c.GitRepo
}

func (c *Config) GetGenerateConstructors() bool {
	return c.GenerateConstructors
}
