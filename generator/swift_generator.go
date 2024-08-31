package generator

import (
	"fmt"
	"os"
	"strings"

	"github.com/deanrtaylor1/type-utils/config"
	"github.com/deanrtaylor1/type-utils/constants"
	"github.com/deanrtaylor1/type-utils/listener"
)

type SwiftSchemaGenerator struct {
	file           *os.File
	listenerConfig *listener.Config
	schema         map[string]*listener.SchemaType
	types          []string
	config.TypeUtilConfiger
}

func (s *SwiftSchemaGenerator) GetFile() *os.File {
	return s.file
}

func (s *SwiftSchemaGenerator) GetConfig() *listener.Config {
	return s.listenerConfig
}

func (s *SwiftSchemaGenerator) GetSchema() map[string]*listener.SchemaType {
	return s.schema
}

func (s *SwiftSchemaGenerator) GeneratePackageDeclaration(config *listener.Config) (string, error) {
	return fmt.Sprintf("// Generated types for %s\n\nimport Foundation", config.PackageName), nil
}

func (s *SwiftSchemaGenerator) GenerateTypeDefinition(typeName string, schemaType *listener.SchemaType) (string, error) {
	var b strings.Builder

	fmt.Fprintf(&b, "struct %s: Codable {\n", typeName)
	for fieldName, fieldType := range schemaType.Fields {
		fieldDef, err := s.GenerateFieldDefinition(fieldName, &fieldType)
		if err != nil {
			return "", err
		}
		fmt.Fprintf(&b, "    %s\n", fieldDef)
	}
	fmt.Fprintln(&b, "}")

	s.types = append(s.types, typeName)
	return b.String(), nil
}

func (s *SwiftSchemaGenerator) GenerateFieldDefinition(fieldName string, fieldType *listener.FieldType) (string, error) {
	swiftType := s.ConvertType(fieldType.Type)
	if fieldType.IsArray {
		swiftType = fmt.Sprintf("[%s]", swiftType)
	}
	if fieldType.IsOptional {
		swiftType = fmt.Sprintf("%s?", swiftType)
	}
	return fmt.Sprintf("let %s: %s", fieldName, swiftType), nil
}

func (s *SwiftSchemaGenerator) ConvertType(fieldType string) string {
	switch fieldType {
	case constants.TypeString:
		return "String"
	case constants.TypeInt, constants.TypeInt32:
		return "Int"
	case constants.TypeInt64:
		return "Int64"
	case constants.TypeFloat:
		return "Float"
	case constants.TypeDouble:
		return "Double"
	case constants.TypeBoolean:
		return "Bool"
	case constants.TypeDate, constants.TypeTime, constants.TypeDateTime, constants.TypeTimestamp:
		return "Date"
	default:
		return getTypeForSwift(fieldType)
	}
}

func getTypeForSwift(str string) string {
	s := strings.Split(str, ".")
	return s[len(s)-1]
}

func (s *SwiftSchemaGenerator) Final() (string, error) {
	return "", nil
}

func (s *SwiftSchemaGenerator) GetGenerateConstructors() bool {
	return false
}
