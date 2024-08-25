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
	Name   string
	Fields map[string]FieldType
	indent int
}

type Config struct {
	OutputDir   string
	PackageName string
	FileName    string
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
		case "output_dir":
			l.Config.OutputDir = value
		case "package_name":
			l.Config.PackageName = value
		case "file_name":
			l.Config.FileName = value
		}
	}
}

func (l *SchemaListener) EnterBlock(ctx *parser.BlockContext) {
	typeName := utils.CapitalizeFirstLetter(ctx.IDENTIFIER().GetSymbol().GetText())
	newType := &SchemaType{
		Name:   typeName,
		Fields: make(map[string]FieldType),
	}

	// Handle repeated and optional flags
	isRepeated := ctx.GetText()[:8] == "repeated"
	isOptional := ctx.GetText()[:8] == "optional" || (len(ctx.GetText()) > 17 && ctx.GetText()[:17] == "optional repeated")

	if l.currentType != nil {
		l.typeStack = append(l.typeStack, l.currentType)
		if len(l.fieldStack) > 0 {
			lastField := l.fieldStack[len(l.fieldStack)-1]
			// Update the parent type's field to reference this new type
			fmt.Printf("type name %s\n", typeName)
			if lastField != l.currentType.Name {
				fmt.Printf("Adding field %s of type %s to parent type %s\n", lastField, typeName, l.currentType.Name)
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
		fmt.Printf("Processing block body for %s\n", typeName)
		for i, child := range blockBody.GetChildren() {
			fmt.Printf("Child %d: ", i)
			switch childCtx := child.(type) {
			case *parser.AttributeContext:
				fmt.Printf("Attribute: %s\n", childCtx.GetText())
				l.processAttribute(childCtx)
			case *parser.BlockContext:
				fmt.Printf("Nested Block: %s\n", childCtx.IDENTIFIER().GetSymbol().GetText())
				// This will be handled by a recursive call to EnterBlock
			default:
				fmt.Printf("Unknown type: %T, Text: %s\n", childCtx, childCtx.GetPayload())
			}
		}
	}
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
				parentType.Fields[strings.ToLower(fieldName)] = FieldType{
					Type:     currentTypeName,
					IsCustom: true,
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

func (l *SchemaListener) processAttribute(ctx *parser.AttributeContext) {
	fieldName := ctx.IDENTIFIER().GetSymbol().GetText()
	isRepeated := ctx.GetText()[:8] == "repeated"
	isOptional := ctx.GetText()[:8] == "optional" || (len(ctx.GetText()) > 17 && ctx.GetText()[:17] == "optional repeated")

	var fieldType FieldType
	if ctx.Value() != nil {
		fieldType = l.processValue(ctx.Value(), isRepeated, isOptional)
	} else {
		// If there's no value, set a default type
		fieldType = FieldType{
			Type:       "any",
			IsArray:    isRepeated,
			IsOptional: isOptional,
		}
	}

	l.currentType.Fields[fieldName] = fieldType
}

func (l *SchemaListener) EnterAttribute(ctx *parser.AttributeContext) {
	if l.currentType != nil {
		fieldName := ctx.IDENTIFIER().GetSymbol().GetText()
		l.isRepeated = ctx.GetChild(0).GetPayload().(antlr.Token).GetText() == "repeated"
		l.isOptional = ctx.GetChild(0).GetPayload().(antlr.Token).GetText() == "optional"
		if ctx.GetChild(0).GetPayload().(antlr.Token).GetText() == "optional repeated" {
			l.isRepeated = true
			l.isOptional = true
		}

		fmt.Printf("value: %v\n", ctx.Value())

		var fieldType FieldType
		if ctx.Value() != nil {
			fieldType = l.processValue(ctx.Value(), l.isRepeated, l.isOptional)
		}

		l.currentType.Fields[fieldName] = fieldType
	}
}

func (l *SchemaListener) processValue(ctx parser.IValueContext, isRepeated bool, isOptional bool) FieldType {
	switch v := ctx.(type) {
	case *parser.ValueContext:
		if v.STRING() != nil {
			value := v.STRING().GetText()
			value = value[1 : len(value)-1] // Remove quotes
			return FieldType{Type: value, IsArray: isRepeated, IsCustom: false, IsOptional: isOptional}
		} else if v.NUMBER() != nil {
			return FieldType{Type: "number", IsArray: isRepeated, IsCustom: false, IsOptional: isOptional}
		} else if v.BOOLEAN() != nil {
			return FieldType{Type: "boolean", IsArray: isRepeated, IsCustom: false, IsOptional: isOptional}
		} else if v.IDENTIFIER() != nil {
			typeName := v.IDENTIFIER().GetSymbol().GetText()
			return FieldType{Type: typeName, IsArray: isRepeated, IsCustom: true, IsOptional: isOptional}
		}
	}
	return FieldType{Type: "any", IsArray: isRepeated, IsCustom: false, IsOptional: isOptional}
}
