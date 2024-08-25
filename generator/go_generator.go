package generator

import (
	"fmt"
	"os"
	"strings"

	"github.com/deanrtaylor1/type-utils/listener"
	"github.com/deanrtaylor1/type-utils/utils"
)

type GoSchemaGenerator struct {
	file   *os.File
	config *listener.Config
	schema map[string]*listener.SchemaType
}

func (g *GoSchemaGenerator) GetFile() *os.File {
	return g.file
}

func (g *GoSchemaGenerator) GetConfig() *listener.Config {
	return g.config
}

func (g *GoSchemaGenerator) GetSchema() map[string]*listener.SchemaType {
	return g.schema
}

func (g *GoSchemaGenerator) GeneratePackageDeclaration(config *listener.Config) (string, error) {
	return fmt.Sprintf("package %s", config.PackageName), nil
}

func (g *GoSchemaGenerator) GenerateTypeDefinition(typeName string, schemaType *listener.SchemaType) (string, error) {
	var b strings.Builder
	fmt.Fprintf(&b, "type %s struct {\n", typeName)
	for fieldName, fieldType := range schemaType.Fields {
		fieldDef, err := g.GenerateFieldDefinition(fieldName, &fieldType)
		if err != nil {
			return "", err
		}
		fmt.Fprintf(&b, "\t%s\n", fieldDef)
	}
	fmt.Fprintln(&b, "}")
	fmt.Println(b.String())
	return b.String(), nil
}

func (g *GoSchemaGenerator) GenerateFieldDefinition(fieldName string, fieldType *listener.FieldType) (string, error) {
	arrayStr := ""
	optionalStr := ""
	if fieldType.IsArray {
		arrayStr = "[]"
	}
	if fieldType.IsOptional {
		optionalStr = ",omitempty"
	}
	goType := g.ConvertType(fieldType.Type)
	return fmt.Sprintf("%s %s%s `json:\"%s%s\"`", utils.CapitalizeFirstLetter(fieldName), arrayStr, goType, fieldName, optionalStr), nil
}

func (g *GoSchemaGenerator) ConvertType(fieldType string) string {
	switch fieldType {
	case "string":
		return "string"
	case "number":
		return "float64"
	case "boolean":
		return "bool"
	default:
		return fieldType // For custom types, use the type name as is
	}
}
