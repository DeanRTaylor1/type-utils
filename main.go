package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/deanrtaylor1/type-utils/config"
	"github.com/deanrtaylor1/type-utils/constants"
	"github.com/deanrtaylor1/type-utils/fetcher"
	"github.com/deanrtaylor1/type-utils/generator"
	"github.com/deanrtaylor1/type-utils/listener"
	"github.com/deanrtaylor1/type-utils/parser"
	"github.com/deanrtaylor1/type-utils/utils"
)

func main() {
	// Read the configuration file
	configer := config.NewYamlConfiger(constants.CONFIG_FILE_NAME)
	err := configer.Read().Validate()
	if err != nil {
		log.Fatalf("invalid configuration settings: %v", err)
	}

	utils.Debug("Config loaded - Version: %s, Language: %s\n", configer.GetConfig().GetVersion(), configer.GetConfig().GetLanguage())

	schemas := make([]*listener.SchemaListener, 0)

	if configer.GetConfig().GetGitRepo().CanFetch() {
		// Process schemas from Git repository
		gitRepo := configer.GetConfig().GetGitRepo()
		schemas, err = fetcher.ParseSchemasFromGitRepo(gitRepo.GetUrl(), gitRepo.GetPath())
		if err != nil {
			log.Fatalf("Error processing schemas from Git repo: %v", err)
		}

	} else {
		// Process local schemas
		schemasDirName := configer.GetConfig().GetSchemasDirName()

		filepath.Walk(schemasDirName, getFileWalkerCallback(schemas))
	}

	for _, v := range schemas {
		err = generator.Generate(configer.GetConfig(), &v.Config, v.Schema)
		if err != nil {
			log.Fatalf("Error generating types from Git repo schemas: %v", err)
		}
	}
}

func processHCLFile(filePath string) (*listener.SchemaListener, error) {
	input, err := antlr.NewFileStream(filePath)
	if err != nil {
		return &listener.SchemaListener{}, fmt.Errorf("error reading file: %v", err)
	}
	lexer := parser.NewHCLLikeDSLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewHCLLikeDSLParser(stream)
	listener := listener.NewSchemaListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.File())

	return listener, nil
}

func getFileWalkerCallback(schemas []*listener.SchemaListener) func(path string, info os.FileInfo, err error) error {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".hcl") {
			schemaConfig, err := processHCLFile(path)
			if err != nil {
				return fmt.Errorf("error processing file %s: %v", path, err)
			}
			schemas = append(schemas, schemaConfig)
		}
		return nil
	}
}
