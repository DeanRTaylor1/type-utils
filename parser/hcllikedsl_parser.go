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
		"", "'HCLCONFIG'", "'{'", "'}'", "'='", "'import'", "'repeated'", "'['",
		"','", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "STRING", "NUMBER", "BOOLEAN",
		"IDENTIFIER", "WS", "COMMENT",
	}
	staticData.RuleNames = []string{
		"file", "hclconfig", "configAttribute", "importStatement", "block",
		"blockBody", "attribute", "value", "array",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 15, 93, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 3, 0, 20, 8, 0,
		1, 0, 5, 0, 23, 8, 0, 10, 0, 12, 0, 26, 9, 0, 1, 0, 5, 0, 29, 8, 0, 10,
		0, 12, 0, 32, 9, 0, 1, 1, 1, 1, 1, 1, 5, 1, 37, 8, 1, 10, 1, 12, 1, 40,
		9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 5, 5, 58, 8, 5, 10, 5, 12, 5, 61, 9, 5, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 71, 8, 6, 3, 6, 73,
		8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 80, 8, 7, 1, 8, 1, 8, 1, 8, 1,
		8, 5, 8, 86, 8, 8, 10, 8, 12, 8, 89, 9, 8, 1, 8, 1, 8, 1, 8, 0, 0, 9, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 0, 0, 96, 0, 19, 1, 0, 0, 0, 2, 33, 1, 0, 0,
		0, 4, 43, 1, 0, 0, 0, 6, 47, 1, 0, 0, 0, 8, 50, 1, 0, 0, 0, 10, 59, 1,
		0, 0, 0, 12, 72, 1, 0, 0, 0, 14, 79, 1, 0, 0, 0, 16, 81, 1, 0, 0, 0, 18,
		20, 3, 2, 1, 0, 19, 18, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 24, 1, 0, 0,
		0, 21, 23, 3, 6, 3, 0, 22, 21, 1, 0, 0, 0, 23, 26, 1, 0, 0, 0, 24, 22,
		1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 30, 1, 0, 0, 0, 26, 24, 1, 0, 0, 0,
		27, 29, 3, 8, 4, 0, 28, 27, 1, 0, 0, 0, 29, 32, 1, 0, 0, 0, 30, 28, 1,
		0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 1, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 33,
		34, 5, 1, 0, 0, 34, 38, 5, 2, 0, 0, 35, 37, 3, 4, 2, 0, 36, 35, 1, 0, 0,
		0, 37, 40, 1, 0, 0, 0, 38, 36, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 41,
		1, 0, 0, 0, 40, 38, 1, 0, 0, 0, 41, 42, 5, 3, 0, 0, 42, 3, 1, 0, 0, 0,
		43, 44, 5, 13, 0, 0, 44, 45, 5, 4, 0, 0, 45, 46, 5, 10, 0, 0, 46, 5, 1,
		0, 0, 0, 47, 48, 5, 5, 0, 0, 48, 49, 5, 10, 0, 0, 49, 7, 1, 0, 0, 0, 50,
		51, 5, 13, 0, 0, 51, 52, 5, 2, 0, 0, 52, 53, 3, 10, 5, 0, 53, 54, 5, 3,
		0, 0, 54, 9, 1, 0, 0, 0, 55, 58, 3, 12, 6, 0, 56, 58, 3, 8, 4, 0, 57, 55,
		1, 0, 0, 0, 57, 56, 1, 0, 0, 0, 58, 61, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0,
		59, 60, 1, 0, 0, 0, 60, 11, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 62, 63, 5,
		13, 0, 0, 63, 64, 5, 4, 0, 0, 64, 73, 3, 14, 7, 0, 65, 66, 5, 6, 0, 0,
		66, 70, 5, 13, 0, 0, 67, 68, 5, 4, 0, 0, 68, 71, 3, 14, 7, 0, 69, 71, 3,
		8, 4, 0, 70, 67, 1, 0, 0, 0, 70, 69, 1, 0, 0, 0, 71, 73, 1, 0, 0, 0, 72,
		62, 1, 0, 0, 0, 72, 65, 1, 0, 0, 0, 73, 13, 1, 0, 0, 0, 74, 80, 5, 10,
		0, 0, 75, 80, 5, 11, 0, 0, 76, 80, 5, 12, 0, 0, 77, 80, 3, 16, 8, 0, 78,
		80, 5, 13, 0, 0, 79, 74, 1, 0, 0, 0, 79, 75, 1, 0, 0, 0, 79, 76, 1, 0,
		0, 0, 79, 77, 1, 0, 0, 0, 79, 78, 1, 0, 0, 0, 80, 15, 1, 0, 0, 0, 81, 82,
		5, 7, 0, 0, 82, 87, 3, 14, 7, 0, 83, 84, 5, 8, 0, 0, 84, 86, 3, 14, 7,
		0, 85, 83, 1, 0, 0, 0, 86, 89, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 88,
		1, 0, 0, 0, 88, 90, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 90, 91, 5, 9, 0, 0,
		91, 17, 1, 0, 0, 0, 10, 19, 24, 30, 38, 57, 59, 70, 72, 79, 87,
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
	HCLLikeDSLParserT__7       = 8
	HCLLikeDSLParserT__8       = 9
	HCLLikeDSLParserSTRING     = 10
	HCLLikeDSLParserNUMBER     = 11
	HCLLikeDSLParserBOOLEAN    = 12
	HCLLikeDSLParserIDENTIFIER = 13
	HCLLikeDSLParserWS         = 14
	HCLLikeDSLParserCOMMENT    = 15
)

// HCLLikeDSLParser rules.
const (
	HCLLikeDSLParserRULE_file            = 0
	HCLLikeDSLParserRULE_hclconfig       = 1
	HCLLikeDSLParserRULE_configAttribute = 2
	HCLLikeDSLParserRULE_importStatement = 3
	HCLLikeDSLParserRULE_block           = 4
	HCLLikeDSLParserRULE_blockBody       = 5
	HCLLikeDSLParserRULE_attribute       = 6
	HCLLikeDSLParserRULE_value           = 7
	HCLLikeDSLParserRULE_array           = 8
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Hclconfig() IHclconfigContext
	AllImportStatement() []IImportStatementContext
	ImportStatement(i int) IImportStatementContext
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

func (s *FileContext) AllImportStatement() []IImportStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImportStatementContext); ok {
			len++
		}
	}

	tst := make([]IImportStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImportStatementContext); ok {
			tst[i] = t.(IImportStatementContext)
			i++
		}
	}

	return tst
}

func (s *FileContext) ImportStatement(i int) IImportStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportStatementContext); ok {
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

	return t.(IImportStatementContext)
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
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == HCLLikeDSLParserT__0 {
		{
			p.SetState(18)
			p.Hclconfig()
		}

	}
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == HCLLikeDSLParserT__4 {
		{
			p.SetState(21)
			p.ImportStatement()
		}

		p.SetState(26)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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
			p.Block()
		}

		p.SetState(32)
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
		p.SetState(33)
		p.Match(HCLLikeDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(34)
		p.Match(HCLLikeDSLParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == HCLLikeDSLParserIDENTIFIER {
		{
			p.SetState(35)
			p.ConfigAttribute()
		}

		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(41)
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
		p.SetState(43)
		p.Match(HCLLikeDSLParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(44)
		p.Match(HCLLikeDSLParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(45)
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

// IImportStatementContext is an interface to support dynamic dispatch.
type IImportStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsImportStatementContext differentiates from other interfaces.
	IsImportStatementContext()
}

type ImportStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportStatementContext() *ImportStatementContext {
	var p = new(ImportStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCLLikeDSLParserRULE_importStatement
	return p
}

func InitEmptyImportStatementContext(p *ImportStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCLLikeDSLParserRULE_importStatement
}

func (*ImportStatementContext) IsImportStatementContext() {}

func NewImportStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStatementContext {
	var p = new(ImportStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HCLLikeDSLParserRULE_importStatement

	return p
}

func (s *ImportStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStatementContext) STRING() antlr.TerminalNode {
	return s.GetToken(HCLLikeDSLParserSTRING, 0)
}

func (s *ImportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCLLikeDSLListener); ok {
		listenerT.EnterImportStatement(s)
	}
}

func (s *ImportStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCLLikeDSLListener); ok {
		listenerT.ExitImportStatement(s)
	}
}

func (p *HCLLikeDSLParser) ImportStatement() (localctx IImportStatementContext) {
	localctx = NewImportStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, HCLLikeDSLParserRULE_importStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		p.Match(HCLLikeDSLParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(48)
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
	p.EnterRule(localctx, 8, HCLLikeDSLParserRULE_block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Match(HCLLikeDSLParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(51)
		p.Match(HCLLikeDSLParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(52)
		p.BlockBody()
	}
	{
		p.SetState(53)
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
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == HCLLikeDSLParserT__5 || _la == HCLLikeDSLParserIDENTIFIER {
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(55)
				p.Attribute()
			}

		case 2:
			{
				p.SetState(56)
				p.Block()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(61)
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
	Block() IBlockContext

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

func (s *AttributeContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
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
			p.SetState(62)
			p.Match(HCLLikeDSLParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)
			p.Match(HCLLikeDSLParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)
			p.Value()
		}

	case HCLLikeDSLParserT__5:
		p.EnterOuterAlt(localctx, 2)
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
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case HCLLikeDSLParserT__3:
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

		case HCLLikeDSLParserIDENTIFIER:
			{
				p.SetState(69)
				p.Block()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
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
	Array() IArrayContext
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

func (s *ValueContext) Array() IArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
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
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case HCLLikeDSLParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(74)
			p.Match(HCLLikeDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case HCLLikeDSLParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(75)
			p.Match(HCLLikeDSLParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case HCLLikeDSLParserBOOLEAN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(76)
			p.Match(HCLLikeDSLParserBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case HCLLikeDSLParserT__6:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(77)
			p.Array()
		}

	case HCLLikeDSLParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(78)
			p.Match(HCLLikeDSLParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllValue() []IValueContext
	Value(i int) IValueContext

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCLLikeDSLParserRULE_array
	return p
}

func InitEmptyArrayContext(p *ArrayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCLLikeDSLParserRULE_array
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HCLLikeDSLParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *ArrayContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
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

	return t.(IValueContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCLLikeDSLListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCLLikeDSLListener); ok {
		listenerT.ExitArray(s)
	}
}

func (p *HCLLikeDSLParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, HCLLikeDSLParserRULE_array)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Match(HCLLikeDSLParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.Value()
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == HCLLikeDSLParserT__7 {
		{
			p.SetState(83)
			p.Match(HCLLikeDSLParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(84)
			p.Value()
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(90)
		p.Match(HCLLikeDSLParserT__8)
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
