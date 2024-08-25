package listener

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/deanrtaylor1/type-utils/parser"
	"github.com/deanrtaylor1/type-utils/utils"
)

type FieldType struct {
	Type       string
	IsArray    bool
	IsCustom   bool
	Fields     map[string]FieldType // For nested types
	IsOptional bool
}

type SchemaType struct {
	Name       string
	Fields     map[string]FieldType
	indent     int
	IsRepeated bool
	IsOptional bool
}

type Config struct {
	PackageName string
	FileName    string

	// Language-specific output directories
	GoOutputDir         string
	TypeScriptOutputDir string
	JavaOutputDir       string
	PythonOutputDir     string
	CsharpOutputDir     string
	RustOutputDir       string
	KotlinOutputDir     string
	SwiftOutputDir      string
	RubyOutputDir       string
	PhpOutputDir        string
	ScalaOutputDir      string
	DartOutputDir       string

	// Go-specific options
	GoModuleName   string
	GoGenerateTags []string // e.g., json, xml, yaml
	GoUsePointers  bool

	// TypeScript-specific options
	TsNamespace    string
	TsModuleSystem string // e.g., "CommonJS", "ES6"
	TsStrictMode   bool

	// Java-specific options
	JavaPackage     string
	JavaClassSuffix string // e.g., "DTO", "Model"
	JavaGenerateJPA bool

	// Python-specific options
	PythonPackageName      string
	PythonUseDataclasses   bool
	PythonGeneratePydantic bool

	// C#-specific options
	CsharpNamespace   string
	CsharpClassSuffix string
	CsharpNullable    bool

	// Rust-specific options
	RustModuleName   string
	RustDeriveTraits []string // e.g., Debug, Clone, Serialize

	// Kotlin-specific options
	KotlinPackage        string
	KotlinUseDataClasses bool

	// Swift-specific options
	SwiftModuleName      string
	SwiftGenerateCodable bool

	// Common options
	GenerateComments      bool
	GenerateValidation    bool
	GenerateToString      bool
	GenerateDocumentation bool
}

type SchemaListener struct {
	*parser.BaseHCLLikeDSLListener
	Schema        map[string]*SchemaType
	currentType   *SchemaType
	typeStack     []*SchemaType
	fieldStack    []string
	callbackQueue map[string][]func(*SchemaType)
	Config        Config
	isRepeated    bool
	isOptional    bool
}

func NewSchemaListener() *SchemaListener {
	return &SchemaListener{
		Schema:        make(map[string]*SchemaType),
		typeStack:     []*SchemaType{},
		fieldStack:    []string{},
		callbackQueue: make(map[string][]func(*SchemaType)),
		isRepeated:    false,
	}
}

func (l *SchemaListener) EnterHclconfig(ctx *parser.HclconfigContext) {
	l.currentType = &SchemaType{
		Name:   "HCLCONFIG",
		Fields: make(map[string]FieldType),
		indent: 1,
	}
	l.Schema["HCLCONFIG"] = l.currentType
}

func (l *SchemaListener) ExitHclconfig(ctx *parser.HclconfigContext) {
	l.currentType = nil
}

func (l *SchemaListener) EnterConfigAttribute(ctx *parser.ConfigAttributeContext) {
	if l.currentType != nil && l.currentType.Name == "HCLCONFIG" {
		key := ctx.IDENTIFIER().GetSymbol().GetText()
		value := ctx.STRING().GetSymbol().GetText()
		value = value[1 : len(value)-1] // Remove quotes

		l.currentType.Fields[key] = FieldType{Type: value, IsArray: false, IsCustom: false}

		switch key {
		// General options
		case "package_name":
			l.Config.PackageName = value
		case "file_name":
			l.Config.FileName = value

		// Language-specific output directories
		case "go_output_dir":
			l.Config.GoOutputDir = value
		case "typescript_output_dir":
			l.Config.TypeScriptOutputDir = value
		case "java_output_dir":
			l.Config.JavaOutputDir = value
		case "python_output_dir":
			l.Config.PythonOutputDir = value
		case "csharp_output_dir":
			l.Config.CsharpOutputDir = value
		case "rust_output_dir":
			l.Config.RustOutputDir = value
		case "kotlin_output_dir":
			l.Config.KotlinOutputDir = value
		case "swift_output_dir":
			l.Config.SwiftOutputDir = value
		case "ruby_output_dir":
			l.Config.RubyOutputDir = value
		case "php_output_dir":
			l.Config.PhpOutputDir = value
		case "scala_output_dir":
			l.Config.ScalaOutputDir = value
		case "dart_output_dir":
			l.Config.DartOutputDir = value

		// Go-specific options
		case "go_module_name":
			l.Config.GoModuleName = value
		case "go_generate_tags":
			l.Config.GoGenerateTags = strings.Split(value, ",")
		case "go_use_pointers":
			l.Config.GoUsePointers = parseBool(value)

		// TypeScript-specific options
		case "ts_namespace":
			l.Config.TsNamespace = value
		case "ts_module_system":
			l.Config.TsModuleSystem = value
		case "ts_strict_mode":
			l.Config.TsStrictMode = parseBool(value)

		// Java-specific options
		case "java_package":
			l.Config.JavaPackage = value
		case "java_class_suffix":
			l.Config.JavaClassSuffix = value
		case "java_generate_jpa":
			l.Config.JavaGenerateJPA = parseBool(value)

		// Python-specific options
		case "python_package_name":
			l.Config.PythonPackageName = value
		case "python_use_dataclasses":
			l.Config.PythonUseDataclasses = parseBool(value)
		case "python_generate_pydantic":
			l.Config.PythonGeneratePydantic = parseBool(value)

		// C#-specific options
		case "csharp_namespace":
			l.Config.CsharpNamespace = value
		case "csharp_class_suffix":
			l.Config.CsharpClassSuffix = value
		case "csharp_nullable":
			l.Config.CsharpNullable = parseBool(value)

		// Rust-specific options
		case "rust_module_name":
			l.Config.RustModuleName = value
		case "rust_derive_traits":
			l.Config.RustDeriveTraits = strings.Split(value, ",")

		// Kotlin-specific options
		case "kotlin_package":
			l.Config.KotlinPackage = value
		case "kotlin_use_data_classes":
			l.Config.KotlinUseDataClasses = parseBool(value)

		// Swift-specific options
		case "swift_module_name":
			l.Config.SwiftModuleName = value
		case "swift_generate_codable":
			l.Config.SwiftGenerateCodable = parseBool(value)

		// Common options
		case "generate_comments":
			l.Config.GenerateComments = parseBool(value)
		case "generate_validation":
			l.Config.GenerateValidation = parseBool(value)
		case "generate_to_string":
			l.Config.GenerateToString = parseBool(value)
		case "generate_documentation":
			l.Config.GenerateDocumentation = parseBool(value)
		}
	}
}

// Helper function to parse boolean values
func parseBool(value string) bool {
	return strings.ToLower(value) == "true" || value == "1"
}

func (l *SchemaListener) EnterBlock(ctx *parser.BlockContext) {
	typeName := utils.CapitalizeFirstLetter(ctx.IDENTIFIER().GetSymbol().GetText())
	newType := &SchemaType{
		Name:   typeName,
		Fields: make(map[string]FieldType),
		indent: 1,
	}

	// Handle repeated and optional flags
	isRepeated := strings.HasPrefix(ctx.GetText(), "repeated")
	isOptional := strings.HasPrefix(ctx.GetText(), "optional")
	if strings.HasPrefix(ctx.GetText(), "optional repeated") {
		isRepeated = true
		isOptional = true
	}

	if l.currentType != nil {
		l.typeStack = append(l.typeStack, l.currentType)
		if len(l.fieldStack) > 0 {
			lastField := l.fieldStack[len(l.fieldStack)-1]
			// Update the parent type's field to reference this new type
			if lastField != l.currentType.Name {
				l.currentType.Fields[strings.ToLower(lastField)] = FieldType{
					Type:       typeName,
					IsArray:    isRepeated,
					IsCustom:   true,
					IsOptional: isOptional,
					Fields:     newType.Fields,
				}
			}
		}
	}

	l.currentType = newType
	l.Schema[typeName] = newType

	// Add this field name to the stack
	l.fieldStack = append(l.fieldStack, typeName)

	// Process block body
	blockBody := ctx.BlockBody()
	if blockBody != nil {
		for _, child := range blockBody.GetChildren() {
			switch childCtx := child.(type) {
			case *parser.AttributeContext:
				l.processAttribute(childCtx, isRepeated, isOptional)
			case *parser.BlockContext:
				// This will be handled by a recursive call to EnterBlock
			default:
				fmt.Printf("Unknown type: %T, Text: %s\n", childCtx, childCtx.GetPayload())
			}
		}
	}

	// Store the isRepeated and isOptional flags in the current type
	l.currentType.IsRepeated = isRepeated
	l.currentType.IsOptional = isOptional
}

func (l *SchemaListener) ExitBlock(ctx *parser.BlockContext) {
	if len(l.typeStack) > 0 {
		parentType := l.typeStack[len(l.typeStack)-1]
		currentTypeName := l.currentType.Name

		// Update the parent type's field to reference this type
		if len(l.fieldStack) > 1 {
			fieldName := l.fieldStack[len(l.fieldStack)-1]
			// Ensure we're not adding a field with the same name as the type
			if strings.ToLower(fieldName) != strings.ToLower(parentType.Name) {
				// Get the flags from the current context
				isRepeated := strings.HasPrefix(ctx.GetText(), "repeated")
				isOptional := strings.HasPrefix(ctx.GetText(), "optional")
				if strings.HasPrefix(ctx.GetText(), "optional repeated") {
					isRepeated = true
					isOptional = true
				}

				parentType.Fields[strings.ToLower(fieldName)] = FieldType{
					Type:       currentTypeName,
					IsArray:    isRepeated,
					IsCustom:   true,
					IsOptional: isOptional,
					Fields:     l.currentType.Fields,
				}

			}
		}

		l.currentType = parentType
		l.typeStack = l.typeStack[:len(l.typeStack)-1]
	} else {
		l.currentType = nil
	}

	// Remove the last field from the stack
	if len(l.fieldStack) > 0 {
		l.fieldStack = l.fieldStack[:len(l.fieldStack)-1]
	}
}

func (l *SchemaListener) processAttribute(ctx *parser.AttributeContext, parentRepeated, parentOptional bool) {
	fieldName := ctx.IDENTIFIER().GetSymbol().GetText()
	isRepeated := parentRepeated || strings.HasPrefix(ctx.GetText(), "repeated")
	isOptional := parentOptional || strings.HasPrefix(ctx.GetText(), "optional")

	var fieldType FieldType
	if ctx.Value() != nil {
		fieldType = l.processValue(ctx.Value(), isRepeated, isOptional)
	} else {
		// If there's no value, set a default type
		fieldType = FieldType{
			Type:       "interface{}",
			IsArray:    isRepeated,
			IsOptional: isOptional,
		}
	}

	l.currentType.Fields[fieldName] = fieldType
}

func (l *SchemaListener) EnterAttribute(ctx *parser.AttributeContext) {
	if l.currentType != nil {
		fieldName := ctx.IDENTIFIER().GetSymbol().GetText()

		if ctx.GetChild(0).GetPayload().(antlr.Token).GetText() == "optional repeated" {
			l.isRepeated = true
			l.isOptional = true
		} else {
			l.isRepeated = ctx.GetChild(0).GetPayload().(antlr.Token).GetText() == "repeated"
			l.isOptional = ctx.GetChild(0).GetPayload().(antlr.Token).GetText() == "optional"
		}

		var fieldType FieldType
		if ctx.Value() != nil {
			fieldType = l.processValue(ctx.Value(), l.isRepeated, l.isOptional)
		}

		l.currentType.Fields[fieldName] = fieldType
	}
}

func (l *SchemaListener) processValue(ctx parser.IValueContext, isRepeated, isOptional bool) FieldType {
	switch v := ctx.(type) {
	case *parser.ValueContext:
		if v.STRING() != nil {
			value := v.STRING().GetText()
			value = value[1 : len(value)-1] // Remove quotes
			return FieldType{Type: value, IsArray: isRepeated, IsCustom: false, IsOptional: isOptional}
		} else if v.NUMBER() != nil {
			return FieldType{Type: "float64", IsArray: isRepeated, IsCustom: false, IsOptional: isOptional}
		} else if v.BOOLEAN() != nil {
			return FieldType{Type: "bool", IsArray: isRepeated, IsCustom: false, IsOptional: isOptional}
		} else if v.IDENTIFIER() != nil {
			typeName := v.IDENTIFIER().GetSymbol().GetText()
			return FieldType{Type: typeName, IsArray: isRepeated, IsCustom: true, IsOptional: isOptional}
		}
	}
	return FieldType{Type: "interface{}", IsArray: isRepeated, IsCustom: false, IsOptional: isOptional}
}
