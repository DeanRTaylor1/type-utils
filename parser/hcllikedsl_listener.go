// Code generated from HCLLikeDSL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // HCLLikeDSL
import "github.com/antlr4-go/antlr/v4"

// HCLLikeDSLListener is a complete listener for a parse tree produced by HCLLikeDSLParser.
type HCLLikeDSLListener interface {
	antlr.ParseTreeListener

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterHclconfig is called when entering the hclconfig production.
	EnterHclconfig(c *HclconfigContext)

	// EnterConfigAttribute is called when entering the configAttribute production.
	EnterConfigAttribute(c *ConfigAttributeContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterNestedBlock is called when entering the nestedBlock production.
	EnterNestedBlock(c *NestedBlockContext)

	// EnterBlockBody is called when entering the blockBody production.
	EnterBlockBody(c *BlockBodyContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitHclconfig is called when exiting the hclconfig production.
	ExitHclconfig(c *HclconfigContext)

	// ExitConfigAttribute is called when exiting the configAttribute production.
	ExitConfigAttribute(c *ConfigAttributeContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitNestedBlock is called when exiting the nestedBlock production.
	ExitNestedBlock(c *NestedBlockContext)

	// ExitBlockBody is called when exiting the blockBody production.
	ExitBlockBody(c *BlockBodyContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
