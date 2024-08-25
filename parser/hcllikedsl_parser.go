// Code generated from HCLLikeDSL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // HCLLikeDSL
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type HCLLikeDSLParser struct {
	*antlr.BaseParser
}

var HCLLikeDSLParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func hcllikedslParserInit() {
	staticData := &HCLLikeDSLParserStaticData
	staticData.LiteralNames = []string{
		"", "'HCLCONFIG'", "'{'", "'}'", "'='", "'repeated'", "'optional'",
		"'optional repeated'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "STRING", "NUMBER", "BOOLEAN", "IDENTIFIER",
		"WS", "COMMENT",
	}
	staticData.RuleNames = []string{
		"file", "hclconfig", "configAttribute", "block", "nestedBlock", "blockBody",
		"attribute", "value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 13, 77, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 3, 0, 18, 8, 0, 1, 0, 5, 0,
		21, 8, 0, 10, 0, 12, 0, 24, 9, 0, 1, 1, 1, 1, 1, 1, 5, 1, 29, 8, 1, 10,
		1, 12, 1, 32, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 5, 5, 51, 8, 5, 10,
		5, 12, 5, 54, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3,
		6, 64, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 73, 8, 6,
		1, 7, 1, 7, 1, 7, 0, 0, 8, 0, 2, 4, 6, 8, 10, 12, 14, 0, 1, 1, 0, 8, 11,
		77, 0, 17, 1, 0, 0, 0, 2, 25, 1, 0, 0, 0, 4, 35, 1, 0, 0, 0, 6, 39, 1,
		0, 0, 0, 8, 44, 1, 0, 0, 0, 10, 52, 1, 0, 0, 0, 12, 72, 1, 0, 0, 0, 14,
		74, 1, 0, 0, 0, 16, 18, 3, 2, 1, 0, 17, 16, 1, 0, 0, 0, 17, 18, 1, 0, 0,
		0, 18, 22, 1, 0, 0, 0, 19, 21, 3, 6, 3, 0, 20, 19, 1, 0, 0, 0, 21, 24,
		1, 0, 0, 0, 22, 20, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 1, 1, 0, 0, 0,
		24, 22, 1, 0, 0, 0, 25, 26, 5, 1, 0, 0, 26, 30, 5, 2, 0, 0, 27, 29, 3,
		4, 2, 0, 28, 27, 1, 0, 0, 0, 29, 32, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 30,
		31, 1, 0, 0, 0, 31, 33, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 33, 34, 5, 3, 0,
		0, 34, 3, 1, 0, 0, 0, 35, 36, 5, 11, 0, 0, 36, 37, 5, 4, 0, 0, 37, 38,
		5, 8, 0, 0, 38, 5, 1, 0, 0, 0, 39, 40, 5, 11, 0, 0, 40, 41, 5, 2, 0, 0,
		41, 42, 3, 10, 5, 0, 42, 43, 5, 3, 0, 0, 43, 7, 1, 0, 0, 0, 44, 45, 5,
		2, 0, 0, 45, 46, 3, 10, 5, 0, 46, 47, 5, 3, 0, 0, 47, 9, 1, 0, 0, 0, 48,
		51, 3, 12, 6, 0, 49, 51, 3, 6, 3, 0, 50, 48, 1, 0, 0, 0, 50, 49, 1, 0,
		0, 0, 51, 54, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 11,
		1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 55, 56, 5, 11, 0, 0, 56, 57, 5, 4, 0, 0,
		57, 73, 3, 14, 7, 0, 58, 59, 5, 5, 0, 0, 59, 63, 5, 11, 0, 0, 60, 61, 5,
		4, 0, 0, 61, 64, 3, 14, 7, 0, 62, 64, 3, 8, 4, 0, 63, 60, 1, 0, 0, 0, 63,
		62, 1, 0, 0, 0, 64, 73, 1, 0, 0, 0, 65, 66, 5, 6, 0, 0, 66, 67, 5, 11,
		0, 0, 67, 68, 5, 4, 0, 0, 68, 73, 3, 14, 7, 0, 69, 70, 5, 7, 0, 0, 70,
		71, 5, 4, 0, 0, 71, 73, 3, 14, 7, 0, 72, 55, 1, 0, 0, 0, 72, 58, 1, 0,
		0, 0, 72, 65, 1, 0, 0, 0, 72, 69, 1, 0, 0, 0, 73, 13, 1, 0, 0, 0, 74, 75,
		7, 0, 0, 0, 75, 15, 1, 0, 0, 0, 7, 17, 22, 30, 50, 52, 63, 72,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// HCLLikeDSLParserInit initializes any static state used to implement HCLLikeDSLParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewHCLLikeDSLParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func HCLLikeDSLParserInit() {
	staticData := &HCLLikeDSLParserStaticData
	staticData.once.Do(hcllikedslParserInit)
}

// NewHCLLikeDSLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewHCLLikeDSLParser(input antlr.TokenStream) *HCLLikeDSLParser {
	HCLLikeDSLParserInit()
	this := new(HCLLikeDSLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &HCLLikeDSLParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "HCLLikeDSL.g4"

	return this
}

// HCLLikeDSLParser tokens.
const (
	HCLLikeDSLParserEOF        = antlr.TokenEOF
	HCLLikeDSLParserT__0       = 1
	HCLLikeDSLParserT__1       = 2
	HCLLikeDSLParserT__2       = 3
	HCLLikeDSLParserT__3       = 4
	HCLLikeDSLParserT__4       = 5
	HCLLikeDSLParserT__5       = 6
	HCLLikeDSLParserT__6       = 7
	HCLLikeDSLParserSTRING     = 8
	HCLLikeDSLParserNUMBER     = 9
	HCLLikeDSLParserBOOLEAN    = 10
	HCLLikeDSLParserIDENTIFIER = 11
	HCLLikeDSLParserWS         = 12
	HCLLikeDSLParserCOMMENT    = 13
)

// HCLLikeDSLParser rules.
const (
	HCLLikeDSLParserRULE_file            = 0
	HCLLikeDSLParserRULE_hclconfig       = 1
	HCLLikeDSLParserRULE_configAttribute = 2
	HCLLikeDSLParserRULE_block           = 3
	HCLLikeDSLParserRULE_nestedBlock     = 4
	HCLLikeDSLParserRULE_blockBody       = 5
	HCLLikeDSLParserRULE_attribute       = 6
	HCLLikeDSLParserRULE_value           = 7
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Hclconfig() IHclconfigContext
	AllBlock() []IBlockContext
	Block(i int) IBlockContext

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCLLikeDSLParserRULE_file
	return p
}

func InitEmptyFileContext(p *FileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCLLikeDSLParserRULE_file
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HCLLikeDSLParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) Hclconfig() IHclconfigContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHclconfigContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHclconfigContext)
}

func (s *FileContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *FileContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCLLikeDSLListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCLLikeDSLListener); ok {
		listenerT.ExitFile(s)
	}
}

func (p *HCLLikeDSLParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, HCLLikeDSLParserRULE_file)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == HCLLikeDSLParserT__0 {
		{
			p.SetState(16)
			p.Hclconfig()
		}

	}
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == HCLLikeDSLParserIDENTIFIER {
		{
			p.SetState(19)
			p.Block()
		}

		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHclconfigContext is an interface to support dynamic dispatch.
type IHclconfigContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllConfigAttribute() []IConfigAttributeContext
	ConfigAttribute(i int) IConfigAttributeContext

	// IsHclconfigContext differentiates from other interfaces.
	IsHclconfigContext()
}

type HclconfigContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHclconfigContext() *HclconfigContext {
	var p = new(HclconfigContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCLLikeDSLParserRULE_hclconfig
	return p
}

func InitEmptyHclconfigContext(p *HclconfigContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCLLikeDSLParserRULE_hclconfig
}

func (*HclconfigContext) IsHclconfigContext() {}

func NewHclconfigContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HclconfigContext {
	var p = new(HclconfigContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HCLLikeDSLParserRULE_hclconfig

	return p
}

func (s *HclconfigContext) GetParser() antlr.Parser { return s.parser }

func (s *HclconfigContext) AllConfigAttribute() []IConfigAttributeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConfigAttributeContext); ok {
			len++
		}
	}

	tst := make([]IConfigAttributeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConfigAttributeContext); ok {
			tst[i] = t.(IConfigAttributeContext)
			i++
		}
	}

	return tst
}

func (s *HclconfigContext) ConfigAttribute(i int) IConfigAttributeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConfigAttributeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConfigAttributeContext)
}

func (s *HclconfigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HclconfigContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HclconfigContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCLLikeDSLListener); ok {
		listenerT.EnterHclconfig(s)
	}
}

func (s *HclconfigContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCLLikeDSLListener); ok {
		listenerT.ExitHclconfig(s)
	}
}

func (p *HCLLikeDSLParser) Hclconfig() (localctx IHclconfigContext) {
	localctx = NewHclconfigContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, HCLLikeDSLParserRULE_hclconfig)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.Match(HCLLikeDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(26)
		p.Match(HCLLikeDSLParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == HCLLikeDSLParserIDENTIFIER {
		{
			p.SetState(27)
			p.ConfigAttribute()
		}

		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(33)
		p.Match(HCLLikeDSLParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConfigAttributeContext is an interface to support dynamic dispatch.
type IConfigAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsConfigAttributeContext differentiates from other interfaces.
	IsConfigAttributeContext()
}

type ConfigAttributeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConfigAttributeContext() *ConfigAttributeContext {
	var p = new(ConfigAttributeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCLLikeDSLParserRULE_configAttribute
	return p
}

func InitEmptyConfigAttributeContext(p *ConfigAttributeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCLLikeDSLParserRULE_configAttribute
}

func (*ConfigAttributeContext) IsConfigAttributeContext() {}

func NewConfigAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConfigAttributeContext {
	var p = new(ConfigAttributeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HCLLikeDSLParserRULE_configAttribute

	return p
}

func (s *ConfigAttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *ConfigAttributeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(HCLLikeDSLParserIDENTIFIER, 0)
}

func (s *ConfigAttributeContext) STRING() antlr.TerminalNode {
	return s.GetToken(HCLLikeDSLParserSTRING, 0)
}

func (s *ConfigAttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConfigAttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConfigAttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCLLikeDSLListener); ok {
		listenerT.EnterConfigAttribute(s)
	}
}

func (s *ConfigAttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCLLikeDSLListener); ok {
		listenerT.ExitConfigAttribute(s)
	}
}

func (p *HCLLikeDSLParser) ConfigAttribute() (localctx IConfigAttributeContext) {
	localctx = NewConfigAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, HCLLikeDSLParserRULE_configAttribute)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)
		p.Match(HCLLikeDSLParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(36)
		p.Match(HCLLikeDSLParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(37)
		p.Match(HCLLikeDSLParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	BlockBody() IBlockBodyContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCLLikeDSLParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCLLikeDSLParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HCLLikeDSLParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(HCLLikeDSLParserIDENTIFIER, 0)
}

func (s *BlockContext) BlockBody() IBlockBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockBodyContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCLLikeDSLListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCLLikeDSLListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *HCLLikeDSLParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, HCLLikeDSLParserRULE_block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(39)
		p.Match(HCLLikeDSLParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(40)
		p.Match(HCLLikeDSLParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(41)
		p.BlockBody()
	}
	{
		p.SetState(42)
		p.Match(HCLLikeDSLParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INestedBlockContext is an interface to support dynamic dispatch.
type INestedBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BlockBody() IBlockBodyContext

	// IsNestedBlockContext differentiates from other interfaces.
	IsNestedBlockContext()
}

type NestedBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNestedBlockContext() *NestedBlockContext {
	var p = new(NestedBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCLLikeDSLParserRULE_nestedBlock
	return p
}

func InitEmptyNestedBlockContext(p *NestedBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCLLikeDSLParserRULE_nestedBlock
}

func (*NestedBlockContext) IsNestedBlockContext() {}

func NewNestedBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NestedBlockContext {
	var p = new(NestedBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HCLLikeDSLParserRULE_nestedBlock

	return p
}

func (s *NestedBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *NestedBlockContext) BlockBody() IBlockBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockBodyContext)
}

func (s *NestedBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NestedBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCLLikeDSLListener); ok {
		listenerT.EnterNestedBlock(s)
	}
}

func (s *NestedBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCLLikeDSLListener); ok {
		listenerT.ExitNestedBlock(s)
	}
}

func (p *HCLLikeDSLParser) NestedBlock() (localctx INestedBlockContext) {
	localctx = NewNestedBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, HCLLikeDSLParserRULE_nestedBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Match(HCLLikeDSLParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(45)
		p.BlockBody()
	}
	{
		p.SetState(46)
		p.Match(HCLLikeDSLParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockBodyContext is an interface to support dynamic dispatch.
type IBlockBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAttribute() []IAttributeContext
	Attribute(i int) IAttributeContext
	AllBlock() []IBlockContext
	Block(i int) IBlockContext

	// IsBlockBodyContext differentiates from other interfaces.
	IsBlockBodyContext()
}

type BlockBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockBodyContext() *BlockBodyContext {
	var p = new(BlockBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCLLikeDSLParserRULE_blockBody
	return p
}

func InitEmptyBlockBodyContext(p *BlockBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCLLikeDSLParserRULE_blockBody
}

func (*BlockBodyContext) IsBlockBodyContext() {}

func NewBlockBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockBodyContext {
	var p = new(BlockBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HCLLikeDSLParserRULE_blockBody

	return p
}

func (s *BlockBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockBodyContext) AllAttribute() []IAttributeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeContext); ok {
			len++
		}
	}

	tst := make([]IAttributeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeContext); ok {
			tst[i] = t.(IAttributeContext)
			i++
		}
	}

	return tst
}

func (s *BlockBodyContext) Attribute(i int) IAttributeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *BlockBodyContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *BlockBodyContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *BlockBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCLLikeDSLListener); ok {
		listenerT.EnterBlockBody(s)
	}
}

func (s *BlockBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCLLikeDSLListener); ok {
		listenerT.ExitBlockBody(s)
	}
}

func (p *HCLLikeDSLParser) BlockBody() (localctx IBlockBodyContext) {
	localctx = NewBlockBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, HCLLikeDSLParserRULE_blockBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2272) != 0 {
		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(48)
				p.Attribute()
			}

		case 2:
			{
				p.SetState(49)
				p.Block()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAttributeContext is an interface to support dynamic dispatch.
type IAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Value() IValueContext
	NestedBlock() INestedBlockContext

	// IsAttributeContext differentiates from other interfaces.
	IsAttributeContext()
}

type AttributeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeContext() *AttributeContext {
	var p = new(AttributeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCLLikeDSLParserRULE_attribute
	return p
}

func InitEmptyAttributeContext(p *AttributeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCLLikeDSLParserRULE_attribute
}

func (*AttributeContext) IsAttributeContext() {}

func NewAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeContext {
	var p = new(AttributeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HCLLikeDSLParserRULE_attribute

	return p
}

func (s *AttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(HCLLikeDSLParserIDENTIFIER, 0)
}

func (s *AttributeContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *AttributeContext) NestedBlock() INestedBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INestedBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INestedBlockContext)
}

func (s *AttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCLLikeDSLListener); ok {
		listenerT.EnterAttribute(s)
	}
}

func (s *AttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCLLikeDSLListener); ok {
		listenerT.ExitAttribute(s)
	}
}

func (p *HCLLikeDSLParser) Attribute() (localctx IAttributeContext) {
	localctx = NewAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, HCLLikeDSLParserRULE_attribute)
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case HCLLikeDSLParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(55)
			p.Match(HCLLikeDSLParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(56)
			p.Match(HCLLikeDSLParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(57)
			p.Value()
		}

	case HCLLikeDSLParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(58)
			p.Match(HCLLikeDSLParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)
			p.Match(HCLLikeDSLParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case HCLLikeDSLParserT__3:
			{
				p.SetState(60)
				p.Match(HCLLikeDSLParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(61)
				p.Value()
			}

		case HCLLikeDSLParserT__1:
			{
				p.SetState(62)
				p.NestedBlock()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

	case HCLLikeDSLParserT__5:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(65)
			p.Match(HCLLikeDSLParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)
			p.Match(HCLLikeDSLParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.Match(HCLLikeDSLParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)
			p.Value()
		}

	case HCLLikeDSLParserT__6:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(69)
			p.Match(HCLLikeDSLParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.Match(HCLLikeDSLParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.Value()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCLLikeDSLParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCLLikeDSLParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HCLLikeDSLParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(HCLLikeDSLParserSTRING, 0)
}

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(HCLLikeDSLParserNUMBER, 0)
}

func (s *ValueContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(HCLLikeDSLParserBOOLEAN, 0)
}

func (s *ValueContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(HCLLikeDSLParserIDENTIFIER, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCLLikeDSLListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCLLikeDSLListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *HCLLikeDSLParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, HCLLikeDSLParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3840) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
