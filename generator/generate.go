package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/deanrtaylor1/type-utils/config"
	"github.com/deanrtaylor1/type-utils/listener"
)

func clearDirectory(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()

	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}

	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}

	return nil
}

func getOutputDir(config *listener.Config, outputLang string) string {
	switch strings.ToLower(outputLang) {
	case "go":
		if config.GoOutputDir != "" {
			return config.GoOutputDir
		}
		fmt.Printf("No go_output_dir specified in config, please review your schemas. Aborting...\n")
		os.Exit(1)
	case "typescript", "ts":
		if config.TypeScriptOutputDir != "" {
			return config.TypeScriptOutputDir
		} else {
			fmt.Printf("No typescript_output_dir specified in config, please review your schemas. Aborting...\n")
			os.Exit(1)
		}
	case "javascript", "js":
		if config.JavascriptOutputDir != "" {
			return config.JavascriptOutputDir
		}
		fmt.Printf("No javascript_output_dir specified in config, please review your schemas. Aborting...\n")
		os.Exit(1)
	case "java":
		if config.JavaOutputDir != "" {
			return config.JavaOutputDir
		}
		fmt.Printf("No java_output_dir specified in config, please review your schemas. Aborting...\n")
		os.Exit(1)
	case "swift":
		if config.SwiftOutputDir != "" {
			return config.SwiftOutputDir
		}
		fmt.Printf("No swift_output_dir specified in config, please review your schemas. Aborting...\n")
		os.Exit(1)
	default:
		fmt.Printf("unsupported output language, aborting...: %s\n", outputLang)
		os.Exit(1)
	}
	return "unsupported"
}

func getFileType(outputLang string) string {
	switch strings.ToLower(outputLang) {
	case "go":
		return "go"
	case "typescript", "ts":
		return "ts"
	case "javascript", "js":
		return "js"
	case "java":
		return "java"
	default:
		panic("unsupported output language")
	}
}

func Generate(typeUtilsConfig config.TypeUtilConfiger, config *listener.Config, schema map[string]*listener.SchemaType) error {
	outputLang := typeUtilsConfig.GetLanguage()
	outputDir := fmt.Sprintf("%s/%s", getOutputDir(config, outputLang), config.PackageName)
	fmt.Println("Output directory: ", outputDir)
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	clearDirectory(outputDir)

	filename := filepath.Join(outputDir, config.PackageName+fmt.Sprintf(".%s", getFileType(outputLang)))

	// Open file in write mode, create if not exists, truncate if exists
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return fmt.Errorf("failed to open or create file: %v", err)
	}
	defer file.Close()

	generator, err := getGenerator(typeUtilsConfig, config, file, schema)
	if err != nil {
		return err
	}
	gen(generator)

	return nil
}

type SchemaGenerator interface {
	config.TypeUtilConfiger
	GeneratePackageDeclaration(config *listener.Config) (string, error)
	GenerateTypeDefinition(typeName string, schemaType *listener.SchemaType) (string, error)
	GenerateFieldDefinition(fieldName string, fieldType *listener.FieldType) (string, error)
	ConvertType(fieldType string) string
	GetFile() *os.File
	GetConfig() *listener.Config
	GetSchema() map[string]*listener.SchemaType
	Final() (string, error)
}

func gen(sg SchemaGenerator) error {
	packageDecl, err := sg.GeneratePackageDeclaration(sg.GetConfig())
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(sg.GetFile(), packageDecl)
	if err != nil {
		return fmt.Errorf("failed to write package declaration: %v", err)
	}

	for typeName, schemaType := range sg.GetSchema() {
		if typeName == "Type_Config" {
			continue
		}
		typedef, err := sg.GenerateTypeDefinition(typeName, schemaType)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintln(sg.GetFile(), typedef)
		if err != nil {
			return fmt.Errorf("failed to write type definition for %s: %v", typeName, err)
		}
	}
	final, err := sg.Final()
	if err != nil {
		return err
	}
	fmt.Fprintln(sg.GetFile(), final)
	return nil
}

func getGenerator(yamlConfig config.TypeUtilConfiger, config *listener.Config, file *os.File, schema map[string]*listener.SchemaType) (SchemaGenerator, error) {
	switch strings.ToLower(yamlConfig.GetLanguage()) {
	case "typescript", "ts":
		return &TypeScriptSchemaGenerator{
			file:             file,
			listenerConfig:   config,
			schema:           schema,
			TypeUtilConfiger: yamlConfig,
		}, nil
	case "js", "javascript":
		return &JavaScriptSchemaGenerator{
			file:             file,
			listenerConfig:   config,
			schema:           schema,
			TypeUtilConfiger: yamlConfig,
		}, nil
	case "go":
		return &GoSchemaGenerator{
			file:             file,
			listenerConfig:   config,
			schema:           schema,
			TypeUtilConfiger: yamlConfig,
		}, nil
	case "java":
		return &JavaSchemaGenerator{
			file:             file,
			schema:           schema,
			listenerConfig:   config,
			TypeUtilConfiger: yamlConfig,
		}, nil
	case "swift":
		return &SwiftSchemaGenerator{
			file:             file,
			listenerConfig:   config,
			schema:           schema,
			TypeUtilConfiger: yamlConfig,
		}, nil

	default:
		return nil, fmt.Errorf("no generator for output language: %s", yamlConfig.GetLanguage())
	}
}
