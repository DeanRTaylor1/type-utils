package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"

	"gopkg.in/yaml.v2"
)

type TypeUtilConfiger interface {
	GetVersion() string
	GetLanguage() string
	GetSchemasDirName() string
	GetGitRepo() GitRepoConfiger
}

type Configer interface {
	Read() Configer
	Validate() error
	GetConfig() TypeUtilConfiger
	GetFilename() string
}

type YamlConfiger struct {
	filename string
	config   TypeUtilConfiger
}

func NewYamlConfiger(filename string) *YamlConfiger {
	fmt.Printf("filename: %s\n", filename)
	return &YamlConfiger{filename: filename}
}

func (yc *YamlConfiger) Validate() error {
	version := yc.GetConfig().GetVersion()
	fmt.Printf("version: %s\n", version)
	if !slices.Contains(getSupportedVersions(), version) {
		return fmt.Errorf("unsupported version in config: %s", version)
	}
	language := yc.GetConfig().GetLanguage()
	if language == "" {
		return fmt.Errorf("language not specified in config")
	}
	return nil
}

func (yc *YamlConfiger) GetConfig() TypeUtilConfiger {
	return yc.config
}

func (yc *YamlConfiger) GetFilename() string {
	return yc.filename
}

func (yc *YamlConfiger) Read() Configer {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	filename := filepath.Join(cwd, yc.GetFilename())
	fmt.Printf("Attempting to read config file: %s\n", filename)

	info, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("Config file not found: %s", filename)
		}
		log.Fatalf("Error accessing config file: %v", err)
	}

	if info.IsDir() {
		log.Fatalf("Expected file, but found directory: %s", filename)
	}

	fmt.Printf("Found config file: %s\n", info.Name())

	buf, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	config := &Config{}
	err = yaml.Unmarshal(buf, config)
	if err != nil {
		log.Fatalf("Failed to parse YAML: %v", err)
	}

	fmt.Printf("Parsed config: %+v\n", config)
	yc.config = config
	return yc
}

func getSupportedVersions() []string {
	return []string{"v1"}
}
