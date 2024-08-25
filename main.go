package main

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/deanrtaylor1/type-utils/parser"
)

type FieldType struct {
	Type     string
	IsArray  bool
	IsCustom bool
	Fields   map[string]FieldType // For nested types
}

type SchemaType struct {
	Name   string
	Fields map[string]FieldType
}

type Config struct {
	OutputDir   string
	PackageName string
}

func main() {
	listener := NewSchemaListener()
	processedFiles := make(map[string]bool)

	var processFile func(string)
	processFile = func(filename string) {
		if processedFiles[filename] {
			return // Already processed this file
		}
		processedFiles[filename] = true

		input, _ := antlr.NewFileStream(filename)
		lexer := parser.NewHCLLikeDSLLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		p := parser.NewHCLLikeDSLParser(stream)

		tree := p.File()
		antlr.ParseTreeWalkerDefault.Walk(listener, tree)

		// Process imports
		for _, importStmt := range p.File().AllImportStatement() {
			importedFile := importStmt.STRING().GetSymbol().GetText()
			// Remove quotes from the string value
			importedFile = importedFile[1 : len(importedFile)-1]
			processFile(importedFile)
		}
	}

	processFile("main.hcl")

	// Print the extracted schema
	for typeName, schemaType := range listener.Schema {
		fmt.Printf("Type: %s\n", typeName)
		printFields(schemaType.Fields, 1, typeName == "DCLCONFIG", &listener.Config)
		fmt.Println()
	}

	// Print the config
	fmt.Printf("Config: %+v\n", listener.Config)

}

func printFields(fields map[string]FieldType, indent int, isDCLCONFIG bool, config *Config) {
	for fieldName, fieldType := range fields {
		if isDCLCONFIG {
			value := fieldType.Type
			switch fieldName {
			case "output_dir":
				config.OutputDir = value
			case "package_name":
				config.PackageName = value
			}
			fmt.Printf("%s%s = %s\n", strings.Repeat("  ", indent), fieldName, value)
		} else {
			// Process regular schema fields
			arrayStr := ""
			if fieldType.IsArray {
				arrayStr = "[]"
			}
			customStr := ""
			if fieldType.IsCustom {
				customStr = " (custom type)"
			}
			fmt.Printf("%s%s: %s%s%s\n", strings.Repeat("  ", indent), fieldName, arrayStr, fieldType.Type, customStr)
			if len(fieldType.Fields) > 0 {
				printFields(fieldType.Fields, indent+1, false, nil)
			}
		}
	}
}
