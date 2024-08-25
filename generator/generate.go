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

func Generate(outputLang string, config *listener.Config, schema map[string]*listener.SchemaType) error {
	outputDir := fmt.Sprintf("%s/%s", config.GoOutputDir, config.PackageName)
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	clearDirectory(outputDir)

	filename := filepath.Join(outputDir, config.PackageName+".go")

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
	return nil
}

func getGenerator(outputLang string, config *listener.Config, file *os.File, schema map[string]*listener.SchemaType) (SchemaGenerator, error) {
	switch strings.ToLower(outputLang) {
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

func CapitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
