// Code generated from HCLLikeDSL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // HCLLikeDSL
import "github.com/antlr4-go/antlr/v4"

// HCLLikeDSLListener is a complete listener for a parse tree produced by HCLLikeDSLParser.
type HCLLikeDSLListener interface {
	antlr.ParseTreeListener

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterImportStatement is called when entering the importStatement production.
	EnterImportStatement(c *ImportStatementContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitImportStatement is called when exiting the importStatement production.
	ExitImportStatement(c *ImportStatementContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)
}
