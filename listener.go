package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/deanrtaylor1/type-utils/parser"
)

type SchemaListener struct {
	*parser.BaseHCLLikeDSLListener
	Schema        map[string]*SchemaType
	currentType   *SchemaType
	typeStack     []*SchemaType
	fieldStack    []string
	callbackQueue map[string][]func(*SchemaType)
	Config        Config
	isRepeated    bool
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

func (l *SchemaListener) EnterBlock(ctx *parser.BlockContext) {
	typeName := ctx.IDENTIFIER().GetSymbol().GetText()
	newType := &SchemaType{
		Name:   typeName,
		Fields: make(map[string]FieldType),
	}

	if l.currentType != nil {
		l.typeStack = append(l.typeStack, l.currentType)
		if len(l.fieldStack) > 0 {
			lastField := l.fieldStack[len(l.fieldStack)-1]
			fieldType := l.currentType.Fields[lastField]
			fieldType.Fields = newType.Fields
			fieldType.IsArray = l.isRepeated
			l.currentType.Fields[lastField] = fieldType
		}
	}

	l.currentType = newType
	l.Schema[typeName] = newType
	l.isRepeated = false

	if callbacks, exists := l.callbackQueue[typeName]; exists {
		for _, callback := range callbacks {
			callback(newType)
		}
		delete(l.callbackQueue, typeName)
	}
}

func (l *SchemaListener) ExitBlock(ctx *parser.BlockContext) {
	if len(l.typeStack) > 0 {
		l.currentType = l.typeStack[len(l.typeStack)-1]
		l.typeStack = l.typeStack[:len(l.typeStack)-1]
		if len(l.fieldStack) > 0 {
			l.fieldStack = l.fieldStack[:len(l.fieldStack)-1]
		}
	} else {
		l.currentType = nil
	}
	l.isRepeated = false
}

func (l *SchemaListener) EnterAttribute(ctx *parser.AttributeContext) {
	if l.currentType != nil {
		fieldName := ctx.IDENTIFIER().GetSymbol().GetText()
		l.isRepeated = ctx.GetChild(0).GetPayload().(antlr.Token).GetText() == "repeated"

		var fieldType FieldType
		if ctx.Value() != nil {
			fieldType = l.processValue(ctx.Value(), l.isRepeated)
		} else if ctx.Block() != nil {
			// Handle nested block
			nestedType := &SchemaType{
				Name:   fieldName,
				Fields: make(map[string]FieldType),
			}
			fieldType = FieldType{
				Type:     "object",
				IsArray:  l.isRepeated,
				IsCustom: true,
				Fields:   nestedType.Fields,
			}
			l.typeStack = append(l.typeStack, l.currentType)
			l.currentType = nestedType
			l.fieldStack = append(l.fieldStack, fieldName)
		}

		l.currentType.Fields[fieldName] = fieldType
	}
}

func (l *SchemaListener) processValue(ctx parser.IValueContext, isRepeated bool) FieldType {
	switch v := ctx.(type) {
	case *parser.ValueContext:
		if v.STRING() != nil {
			value := v.STRING().GetText()
			// Remove surrounding quotes
			value = value[1 : len(value)-1]
			return FieldType{Type: value, IsArray: isRepeated, IsCustom: false}
		} else if v.NUMBER() != nil {
			return FieldType{Type: v.NUMBER().GetText(), IsArray: isRepeated, IsCustom: false}
		} else if v.BOOLEAN() != nil {
			return FieldType{Type: v.BOOLEAN().GetText(), IsArray: isRepeated, IsCustom: false}
		} else if v.Array() != nil {
			if v.Array().GetChildCount() > 2 { // [ value ]
				elementType := l.processValue(v.Array().GetChild(1).(parser.IValueContext), false)
				elementType.IsArray = true
				return elementType
			}
			return FieldType{Type: "any", IsArray: true, IsCustom: false}
		} else if v.IDENTIFIER() != nil {
			typeName := v.IDENTIFIER().GetSymbol().GetText()
			if l.currentType != nil && l.currentType.Name == "DCLCONFIG" {
				// For DCLCONFIG, store the actual value
				return FieldType{Type: typeName, IsArray: isRepeated, IsCustom: false}
			}
			if _, exists := l.Schema[typeName]; !exists {
				l.callbackQueue[typeName] = append(l.callbackQueue[typeName], func(schemaType *SchemaType) {
					l.currentType.Fields[typeName] = FieldType{Type: typeName, IsArray: isRepeated, IsCustom: true}
				})
			}
			return FieldType{Type: typeName, IsArray: isRepeated, IsCustom: true}
		}
	}
	return FieldType{Type: "any", IsArray: isRepeated, IsCustom: false}
}

func (l *SchemaListener) EnterConfigAttribute(ctx *parser.ConfigAttributeContext) {
	if l.currentType.Name == "DCLCONFIG" {
		key := ctx.IDENTIFIER().GetSymbol().GetText()
		value := ctx.STRING().GetSymbol().GetText()
		value = value[1 : len(value)-1] // Remove quotes
		fmt.Printf("Config Attribute: %s = %s\n", key, value)

		switch key {
		case "output_dir":
			l.Config.OutputDir = value
		case "package_name":
			l.Config.PackageName = value
		}
	}
}
