// Code generated from HCLLikeDSL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // HCLLikeDSL
import "github.com/antlr4-go/antlr/v4"

// BaseHCLLikeDSLListener is a complete listener for a parse tree produced by HCLLikeDSLParser.
type BaseHCLLikeDSLListener struct{}

var _ HCLLikeDSLListener = &BaseHCLLikeDSLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHCLLikeDSLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHCLLikeDSLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHCLLikeDSLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHCLLikeDSLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile is called when production file is entered.
func (s *BaseHCLLikeDSLListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseHCLLikeDSLListener) ExitFile(ctx *FileContext) {}

// EnterImportStatement is called when production importStatement is entered.
func (s *BaseHCLLikeDSLListener) EnterImportStatement(ctx *ImportStatementContext) {}

// ExitImportStatement is called when production importStatement is exited.
func (s *BaseHCLLikeDSLListener) ExitImportStatement(ctx *ImportStatementContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseHCLLikeDSLListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseHCLLikeDSLListener) ExitBlock(ctx *BlockContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BaseHCLLikeDSLListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BaseHCLLikeDSLListener) ExitAttribute(ctx *AttributeContext) {}

// EnterValue is called when production value is entered.
func (s *BaseHCLLikeDSLListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseHCLLikeDSLListener) ExitValue(ctx *ValueContext) {}

// EnterArray is called when production array is entered.
func (s *BaseHCLLikeDSLListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseHCLLikeDSLListener) ExitArray(ctx *ArrayContext) {}
