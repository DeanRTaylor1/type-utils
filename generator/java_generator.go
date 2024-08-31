package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/deanrtaylor1/type-utils/config"
	"github.com/deanrtaylor1/type-utils/constants"
	"github.com/deanrtaylor1/type-utils/listener"
	"github.com/deanrtaylor1/type-utils/utils"
)

type JavaSchemaGenerator struct {
	file           *os.File
	schema         map[string]*listener.SchemaType
	listenerConfig *listener.Config
	config.TypeUtilConfiger
}

func NewJavaSchemaGenerator(conf config.TypeUtilConfiger, schema map[string]*listener.SchemaType) *JavaSchemaGenerator {
	return &JavaSchemaGenerator{
		TypeUtilConfiger: conf,
		schema:           schema,
	}
}

func (j *JavaSchemaGenerator) GetSchema() map[string]*listener.SchemaType {
	return j.schema
}

func (j *JavaSchemaGenerator) GeneratePackageDeclaration(config *listener.Config) (string, error) {
	return config.PackageName, nil
}

func (j *JavaSchemaGenerator) GenerateTypeDefinition(typeName string, schemaType *listener.SchemaType) (string, error) {
	baseOutputDir := j.listenerConfig.JavaOutputDir
	typeOutputDir := filepath.Join(baseOutputDir, strings.ToLower(typeName))

	if err := os.MkdirAll(typeOutputDir, 0755); err != nil {
		return "", fmt.Errorf("error creating directory for type %s: %v", typeName, err)
	}

	basePackage, err := j.GeneratePackageDeclaration(j.listenerConfig)
	if err != nil {
		return "", fmt.Errorf("error generating base package declaration: %v", err)
	}

	packageDecl := fmt.Sprintf("package %s.%s;\n", basePackage, strings.ToLower(typeName))

	// Generate interface
	interfaceContent := j.generateInterface(typeName, schemaType, packageDecl)
	interfaceFileName := filepath.Join(typeOutputDir, fmt.Sprintf("I%s.java", typeName))
	if err := os.WriteFile(interfaceFileName, []byte(interfaceContent), 0644); err != nil {
		return "", fmt.Errorf("error writing interface file: %v", err)
	}

	// Generate class
	classContent := j.generateClass(typeName, schemaType, packageDecl)
	classFileName := filepath.Join(typeOutputDir, fmt.Sprintf("%s.java", typeName))
	if err := os.WriteFile(classFileName, []byte(classContent), 0644); err != nil {
		return "", fmt.Errorf("error writing class file: %v", err)
	}

	return "", nil
}

func (j *JavaSchemaGenerator) generateInterface(typeName string, schemaType *listener.SchemaType, packageDecl string) string {
	var b strings.Builder
	b.WriteString(packageDecl)
	b.WriteString("\n")
	b.WriteString("import java.util.List;\n")
	b.WriteString("import java.time.LocalDateTime;\n\n")

	fmt.Fprintf(&b, "public interface I%s {\n", typeName)
	for fieldName, fieldType := range schemaType.Fields {
		javaType := j.ConvertType(fieldType.Type)
		if fieldType.IsArray {
			javaType = "List<" + javaType + ">"
		}
		fmt.Fprintf(&b, "    %s get%s();\n", javaType, utils.CapitalizeFirstLetter(fieldName))
		fmt.Fprintf(&b, "    void set%s(%s %s);\n", utils.CapitalizeFirstLetter(fieldName), javaType, fieldName)
	}
	fmt.Fprintln(&b, "}")

	return b.String()
}

func (j *JavaSchemaGenerator) generateClass(typeName string, schemaType *listener.SchemaType, packageDecl string) string {
	var b strings.Builder
	b.WriteString(packageDecl)
	b.WriteString("\n")
	b.WriteString("import java.util.List;\n")
	b.WriteString("import java.time.LocalDateTime;\n\n")

	fmt.Fprintf(&b, "public class %s implements I%s {\n", typeName, typeName)

	// Fields
	for fieldName, fieldType := range schemaType.Fields {
		javaType := j.ConvertType(fieldType.Type)
		if fieldType.IsArray {
			javaType = "List<" + javaType + ">"
		}
		fmt.Fprintf(&b, "    private %s %s;\n", javaType, fieldName)
	}
	fmt.Fprintln(&b)

	// Constructor
	if j.GetGenerateConstructors() {
		fmt.Fprintf(&b, "    public %s(", typeName)
		var params []string
		for fieldName, fieldType := range schemaType.Fields {
			javaType := j.ConvertType(fieldType.Type)
			if fieldType.IsArray {
				javaType = "List<" + javaType + ">"
			}
			params = append(params, fmt.Sprintf("%s %s", javaType, fieldName))
		}
		fmt.Fprintf(&b, strings.Join(params, ", "))
		fmt.Fprintln(&b, ") {")
		for fieldName := range schemaType.Fields {
			fmt.Fprintf(&b, "        this.%s = %s;\n", fieldName, fieldName)
		}
		fmt.Fprintln(&b, "    }\n")
	}

	// Getters and Setters
	for fieldName, fieldType := range schemaType.Fields {
		javaType := j.ConvertType(fieldType.Type)
		if fieldType.IsArray {
			javaType = "List<" + javaType + ">"
		}
		capFieldName := utils.CapitalizeFirstLetter(fieldName)

		// Getter
		fmt.Fprintf(&b, "    @Override\n")
		fmt.Fprintf(&b, "    public %s get%s() {\n", javaType, capFieldName)
		fmt.Fprintf(&b, "        return this.%s;\n", fieldName)
		fmt.Fprintf(&b, "    }\n\n")

		// Setter
		fmt.Fprintf(&b, "    @Override\n")
		fmt.Fprintf(&b, "    public void set%s(%s %s) {\n", capFieldName, javaType, fieldName)
		fmt.Fprintf(&b, "        this.%s = %s;\n", fieldName, fieldName)
		fmt.Fprintf(&b, "    }\n\n")
	}

	fmt.Fprintln(&b, "}")

	return b.String()
}

func (j *JavaSchemaGenerator) ConvertType(fieldType string) string {
	switch fieldType {
	case constants.TypeString:
		return "String"
	case constants.TypeInt, constants.TypeInt32:
		return "Integer"
	case constants.TypeInt64:
		return "Long"
	case constants.TypeFloat:
		return "Float"
	case constants.TypeDouble:
		return "Double"
	case constants.TypeBoolean:
		return "Boolean"
	case constants.TypeDate:
		return "LocalDate"
	case constants.TypeTime:
		return "LocalTime"
	case constants.TypeDateTime, constants.TypeTimestamp:
		return "LocalDateTime"
	default:
		return fieldType
	}
}

func (j *JavaSchemaGenerator) Final() (string, error) {
	return "// Types generated by type-utils", nil
}

func (j *JavaSchemaGenerator) GenerateFieldDefinition(fieldName string, fieldType *listener.FieldType) (string, error) {
	// This method is not directly used in Java generation, but we'll keep it for consistency
	javaType := j.ConvertType(fieldType.Type)
	if fieldType.IsArray {
		javaType = "List<" + javaType + ">"
	}
	return fmt.Sprintf("%s %s", javaType, fieldName), nil
}

func (j *JavaSchemaGenerator) GetConfig() *listener.Config {
	return j.listenerConfig
}

func (j *JavaSchemaGenerator) GetFile() *os.File {
	return j.file
}
