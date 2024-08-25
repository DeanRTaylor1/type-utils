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

// EnterHclconfig is called when production hclconfig is entered.
func (s *BaseHCLLikeDSLListener) EnterHclconfig(ctx *HclconfigContext) {}

// ExitHclconfig is called when production hclconfig is exited.
func (s *BaseHCLLikeDSLListener) ExitHclconfig(ctx *HclconfigContext) {}

// EnterConfigAttribute is called when production configAttribute is entered.
func (s *BaseHCLLikeDSLListener) EnterConfigAttribute(ctx *ConfigAttributeContext) {}

// ExitConfigAttribute is called when production configAttribute is exited.
func (s *BaseHCLLikeDSLListener) ExitConfigAttribute(ctx *ConfigAttributeContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseHCLLikeDSLListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseHCLLikeDSLListener) ExitBlock(ctx *BlockContext) {}

// EnterNestedBlock is called when production nestedBlock is entered.
func (s *BaseHCLLikeDSLListener) EnterNestedBlock(ctx *NestedBlockContext) {}

// ExitNestedBlock is called when production nestedBlock is exited.
func (s *BaseHCLLikeDSLListener) ExitNestedBlock(ctx *NestedBlockContext) {}

// EnterBlockBody is called when production blockBody is entered.
func (s *BaseHCLLikeDSLListener) EnterBlockBody(ctx *BlockBodyContext) {}

// ExitBlockBody is called when production blockBody is exited.
func (s *BaseHCLLikeDSLListener) ExitBlockBody(ctx *BlockBodyContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BaseHCLLikeDSLListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BaseHCLLikeDSLListener) ExitAttribute(ctx *AttributeContext) {}

// EnterValue is called when production value is entered.
func (s *BaseHCLLikeDSLListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseHCLLikeDSLListener) ExitValue(ctx *ValueContext) {}
