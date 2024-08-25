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
	"github.com/deanrtaylor1/type-utils/generator"
	"github.com/deanrtaylor1/type-utils/listener"
	"github.com/deanrtaylor1/type-utils/parser"
)

var DEBUG = true

func debug(format string, args ...interface{}) {
	if DEBUG {
		fmt.Printf(format, args...)
	}
}

func main() {
	// Read the configuration file
	config, err := config.Read(constants.CONFIG_FILE_NAME)
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	// Validate config
	if config.Version != "v1" {
		log.Fatalf("Unsupported version in config: %s", config.Version)
	}

	if config.Language == "" {
		log.Fatalf("Language not specified in config")
	}

	debug("Config loaded - Version: %s, Language: %s\n", config.Version, config.Language)

	var schemasDirName string
	if config.SchemasDirName == "" {
		schemasDirName = fmt.Sprintf("./%s", constants.DEFAULT_SCHEMA_DIR_NAME)
	} else {
		schemasDirName = fmt.Sprintf("./%s", config.SchemasDirName)
	}
	err = filepath.Walk(schemasDirName, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		debug("filename: %s", info.Name())
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".hcl") {
			fmt.Printf("Processing file: %s\n", path)
			err := processHCLFile(path, config.Language)
			if err != nil {
				return fmt.Errorf("error processing file %s: %v", path, err)
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error walking through schemas directory: %v", err)
	}
}

func processHCLFile(filePath string, language string) error {
	input, err := antlr.NewFileStream(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	lexer := parser.NewHCLLikeDSLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewHCLLikeDSLParser(stream)

	listener := listener.NewSchemaListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.File())

	fmt.Printf("\nProcessing: %s\n", filePath)
	fmt.Println("HCLCONFIG:")
	if hclconfig, ok := listener.Schema["HCLCONFIG"]; ok {
		for key, value := range hclconfig.Fields {
			fmt.Printf("  %s = %s\n", key, value.Type)
		}
	}
	fmt.Println()

	for typeName, schemaType := range listener.Schema {
		if typeName != "HCLCONFIG" {
			fmt.Printf("Type: %s\n", typeName)
			printFields(schemaType.Fields, 1)
			fmt.Println()
		}
	}

	err = generator.Generate(language, &listener.Config, listener.Schema)
	if err != nil {
		return fmt.Errorf("error creating types: %v", err)
	}

	fmt.Printf("Config: %+v\n", listener.Config)
	fmt.Println("------------------------------------")

	return nil
}

func printFields(fields map[string]listener.FieldType, indent int) {
	for fieldName, fieldType := range fields {
		arrayStr := ""
		optionalStr := ""
		if fieldType.IsArray {
			arrayStr = "[]"
		}
		if fieldType.IsOptional {
			optionalStr = "?"
		}
		customStr := ""
		if fieldType.IsCustom {
			customStr = " (custom type)"
		}
		fmt.Printf("%s%s%s: %s%s%s\n", strings.Repeat("  ", indent), fieldName, optionalStr, arrayStr, fieldType.Type, customStr)
		if len(fieldType.Fields) > 0 {
			printFields(fieldType.Fields, indent+1)
		}
	}
}
