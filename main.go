package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/deanrtaylor1/type-utils/parser"
)

type FieldType struct {
	Type     string
	IsArray  bool
	IsCustom bool
}

type SchemaType struct {
	Name   string
	Fields map[string]FieldType
}

type SchemaListener struct {
	*parser.BaseHCLLikeDSLListener
	Schema        map[string]*SchemaType
	currentType   *SchemaType
	typeStack     []*SchemaType
	callbackQueue map[string][]func(*SchemaType)
}

func NewSchemaListener() *SchemaListener {
	return &SchemaListener{
		Schema:        make(map[string]*SchemaType),
		typeStack:     []*SchemaType{},
		callbackQueue: make(map[string][]func(*SchemaType)),
	}
}

func (l *SchemaListener) EnterBlock(ctx *parser.BlockContext) {
	typeName := ctx.IDENTIFIER().GetText()
	newType := &SchemaType{
		Name:   typeName,
		Fields: make(map[string]FieldType),
	}

	if l.currentType != nil {
		l.typeStack = append(l.typeStack, l.currentType)
	}

	l.currentType = newType
	l.Schema[typeName] = newType

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
	} else {
		l.currentType = nil
	}
}

func (l *SchemaListener) EnterAttribute(ctx *parser.AttributeContext) {
	if l.currentType != nil {
		var fieldName string
		isRepeated := false

		// Check if the first child is 'repeated'
		if ctx.GetChildCount() > 3 {
			firstChild := ctx.GetChild(0)
			if firstChildTerminal, ok := firstChild.(antlr.TerminalNode); ok {
				if firstChildTerminal.GetSymbol().GetText() == "repeated" {
					isRepeated = true
					// If repeated, the identifier is the second child
					if identNode, ok := ctx.GetChild(1).(antlr.TerminalNode); ok {
						fieldName = identNode.GetSymbol().GetText()
					}
				}
			}
		}

		// If not repeated, the identifier is the first child
		if !isRepeated {
			if identNode, ok := ctx.GetChild(0).(antlr.TerminalNode); ok {
				fieldName = identNode.GetSymbol().GetText()
			}
		}

		// Process the value
		var valueCtx parser.IValueContext
		if isRepeated {
			valueCtx = ctx.GetChild(3).(parser.IValueContext)
		} else {
			valueCtx = ctx.GetChild(2).(parser.IValueContext)
		}

		fieldType := l.processValue(valueCtx, isRepeated)
		l.currentType.Fields[fieldName] = fieldType
	}
}

func (l *SchemaListener) processValue(ctx parser.IValueContext, isRepeated bool) FieldType {
	switch v := ctx.(type) {
	case *parser.ValueContext:
		if v.STRING() != nil {
			return FieldType{Type: "string", IsArray: isRepeated, IsCustom: false}
		} else if v.NUMBER() != nil {
			return FieldType{Type: "number", IsArray: isRepeated, IsCustom: false}
		} else if v.BOOLEAN() != nil {
			return FieldType{Type: "boolean", IsArray: isRepeated, IsCustom: false}
		} else if v.Array() != nil {
			if v.Array().GetChildCount() > 2 { // [ value ]
				elementType := l.processValue(v.Array().GetChild(1).(parser.IValueContext), false)
				elementType.IsArray = true
				return elementType
			}
			return FieldType{Type: "any", IsArray: true, IsCustom: false}
		} else if v.IDENTIFIER() != nil {
			typeName := v.IDENTIFIER().GetSymbol().GetText()
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

func main() {
	input, _ := antlr.NewFileStream("test.hcl")
	lexer := parser.NewHCLLikeDSLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewHCLLikeDSLParser(stream)

	tree := p.File()

	listener := NewSchemaListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	walkFields(listener.Schema)
}

func walkFields(schema map[string]*SchemaType) {
	for typeName, schemaType := range schema {
		fmt.Printf("Type: %s\n", typeName)
		for fieldName, fieldType := range schemaType.Fields {
			arrayStr := ""
			if fieldType.IsArray {
				arrayStr = "[]"
			}
			customStr := ""
			if fieldType.IsCustom {
				customStr = " (custom type)"
			}
			fmt.Printf("  %s: %s%s%s\n", fieldName, arrayStr, fieldType.Type, customStr)
		}
		fmt.Println()
	}
}
