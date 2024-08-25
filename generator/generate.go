package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
	case "typescript":
		if config.TypeScriptOutputDir != "" {
			fmt.Println("typescript output dir: ", config.TypeScriptOutputDir)
			return config.TypeScriptOutputDir
		}
	default:
		panic("unsupported output language")
	}
	return "unsupported"
}

func getFileType(outputLang string) string {
	switch strings.ToLower(outputLang) {
	case "go":
		return "go"
	case "typescript":
		return "ts"
	default:
		panic("unsupported output language")
	}
}

func Generate(outputLang string, config *listener.Config, schema map[string]*listener.SchemaType) error {
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

	generator, err := getGenerator(outputLang, config, file, schema)
	if err != nil {
		return err
	}
	gen(generator)

	return nil
}

type SchemaGenerator interface {
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
		if typeName == "HCLCONFIG" {
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

func getGenerator(outputLang string, config *listener.Config, file *os.File, schema map[string]*listener.SchemaType) (SchemaGenerator, error) {
	switch strings.ToLower(outputLang) {
	case "typescript":
		return &TypeScriptSchemaGenerator{
			file:   file,
			config: config,
			schema: schema,
		}, nil
	case "ts":
		return &TypeScriptSchemaGenerator{
			file:   file,
			config: config,
			schema: schema,
		}, nil
	case "go":
		return &GoSchemaGenerator{
			file:   file,
			config: config,
			schema: schema,
		}, nil
	default:
		return nil, fmt.Errorf("no generator for output language: %s", outputLang)
	}
}
