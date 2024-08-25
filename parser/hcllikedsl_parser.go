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
		"file", "hclconfig", "configAttribute", "block", "blockBody", "attribute",
		"value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 13, 88, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 3, 0, 16, 8, 0, 1, 0, 5, 0, 19, 8, 0,
		10, 0, 12, 0, 22, 9, 0, 1, 1, 1, 1, 1, 1, 5, 1, 27, 8, 1, 10, 1, 12, 1,
		30, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 60, 8, 3, 1, 4, 1, 4, 5, 4, 64,
		8, 4, 10, 4, 12, 4, 67, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 84, 8, 5, 1, 6, 1,
		6, 1, 6, 0, 0, 7, 0, 2, 4, 6, 8, 10, 12, 0, 1, 1, 0, 8, 11, 91, 0, 15,
		1, 0, 0, 0, 2, 23, 1, 0, 0, 0, 4, 33, 1, 0, 0, 0, 6, 59, 1, 0, 0, 0, 8,
		65, 1, 0, 0, 0, 10, 83, 1, 0, 0, 0, 12, 85, 1, 0, 0, 0, 14, 16, 3, 2, 1,
		0, 15, 14, 1, 0, 0, 0, 15, 16, 1, 0, 0, 0, 16, 20, 1, 0, 0, 0, 17, 19,
		3, 6, 3, 0, 18, 17, 1, 0, 0, 0, 19, 22, 1, 0, 0, 0, 20, 18, 1, 0, 0, 0,
		20, 21, 1, 0, 0, 0, 21, 1, 1, 0, 0, 0, 22, 20, 1, 0, 0, 0, 23, 24, 5, 1,
		0, 0, 24, 28, 5, 2, 0, 0, 25, 27, 3, 4, 2, 0, 26, 25, 1, 0, 0, 0, 27, 30,
		1, 0, 0, 0, 28, 26, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29, 31, 1, 0, 0, 0,
		30, 28, 1, 0, 0, 0, 31, 32, 5, 3, 0, 0, 32, 3, 1, 0, 0, 0, 33, 34, 5, 11,
		0, 0, 34, 35, 5, 4, 0, 0, 35, 36, 5, 8, 0, 0, 36, 5, 1, 0, 0, 0, 37, 38,
		5, 11, 0, 0, 38, 39, 5, 2, 0, 0, 39, 40, 3, 8, 4, 0, 40, 41, 5, 3, 0, 0,
		41, 60, 1, 0, 0, 0, 42, 43, 5, 5, 0, 0, 43, 44, 5, 11, 0, 0, 44, 45, 5,
		2, 0, 0, 45, 46, 3, 8, 4, 0, 46, 47, 5, 3, 0, 0, 47, 60, 1, 0, 0, 0, 48,
		49, 5, 6, 0, 0, 49, 50, 5, 11, 0, 0, 50, 51, 5, 2, 0, 0, 51, 52, 3, 8,
		4, 0, 52, 53, 5, 3, 0, 0, 53, 60, 1, 0, 0, 0, 54, 55, 5, 7, 0, 0, 55, 56,
		5, 2, 0, 0, 56, 57, 3, 8, 4, 0, 57, 58, 5, 3, 0, 0, 58, 60, 1, 0, 0, 0,
		59, 37, 1, 0, 0, 0, 59, 42, 1, 0, 0, 0, 59, 48, 1, 0, 0, 0, 59, 54, 1,
		0, 0, 0, 60, 7, 1, 0, 0, 0, 61, 64, 3, 10, 5, 0, 62, 64, 3, 6, 3, 0, 63,
		61, 1, 0, 0, 0, 63, 62, 1, 0, 0, 0, 64, 67, 1, 0, 0, 0, 65, 63, 1, 0, 0,
		0, 65, 66, 1, 0, 0, 0, 66, 9, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 68, 69, 5,
		11, 0, 0, 69, 70, 5, 4, 0, 0, 70, 84, 3, 12, 6, 0, 71, 72, 5, 5, 0, 0,
		72, 73, 5, 11, 0, 0, 73, 74, 5, 4, 0, 0, 74, 84, 3, 12, 6, 0, 75, 76, 5,
		6, 0, 0, 76, 77, 5, 11, 0, 0, 77, 78, 5, 4, 0, 0, 78, 84, 3, 12, 6, 0,
		79, 80, 5, 7, 0, 0, 80, 81, 5, 11, 0, 0, 81, 82, 5, 4, 0, 0, 82, 84, 3,
		12, 6, 0, 83, 68, 1, 0, 0, 0, 83, 71, 1, 0, 0, 0, 83, 75, 1, 0, 0, 0, 83,
		79, 1, 0, 0, 0, 84, 11, 1, 0, 0, 0, 85, 86, 7, 0, 0, 0, 86, 13, 1, 0, 0,
		0, 7, 15, 20, 28, 59, 63, 65, 83,
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
	HCLLikeDSLParserRULE_blockBody       = 4
	HCLLikeDSLParserRULE_attribute       = 5
	HCLLikeDSLParserRULE_value           = 6
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
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == HCLLikeDSLParserT__0 {
		{
			p.SetState(14)
			p.Hclconfig()
		}

	}
	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2272) != 0 {
		{
			p.SetState(17)
			p.Block()
		}

		p.SetState(22)
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
		p.SetState(23)
		p.Match(HCLLikeDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(24)
		p.Match(HCLLikeDSLParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == HCLLikeDSLParserIDENTIFIER {
		{
			p.SetState(25)
			p.ConfigAttribute()
		}

		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(31)
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
		p.SetState(33)
		p.Match(HCLLikeDSLParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(34)
		p.Match(HCLLikeDSLParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(35)
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
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case HCLLikeDSLParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(37)
			p.Match(HCLLikeDSLParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(38)
			p.Match(HCLLikeDSLParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(39)
			p.BlockBody()
		}
		{
			p.SetState(40)
			p.Match(HCLLikeDSLParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case HCLLikeDSLParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(42)
			p.Match(HCLLikeDSLParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
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

	case HCLLikeDSLParserT__5:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(48)
			p.Match(HCLLikeDSLParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(49)
			p.Match(HCLLikeDSLParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.Match(HCLLikeDSLParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(51)
			p.BlockBody()
		}
		{
			p.SetState(52)
			p.Match(HCLLikeDSLParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case HCLLikeDSLParserT__6:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(54)
			p.Match(HCLLikeDSLParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(55)
			p.Match(HCLLikeDSLParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(56)
			p.BlockBody()
		}
		{
			p.SetState(57)
			p.Match(HCLLikeDSLParserT__2)
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
	p.EnterRule(localctx, 8, HCLLikeDSLParserRULE_blockBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2272) != 0 {
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(61)
				p.Attribute()
			}

		case 2:
			{
				p.SetState(62)
				p.Block()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(67)
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
	p.EnterRule(localctx, 10, HCLLikeDSLParserRULE_attribute)
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case HCLLikeDSLParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(68)
			p.Match(HCLLikeDSLParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(69)
			p.Match(HCLLikeDSLParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.Value()
		}

	case HCLLikeDSLParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.Match(HCLLikeDSLParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(72)
			p.Match(HCLLikeDSLParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(73)
			p.Match(HCLLikeDSLParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.Value()
		}

	case HCLLikeDSLParserT__5:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(75)
			p.Match(HCLLikeDSLParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(76)
			p.Match(HCLLikeDSLParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.Match(HCLLikeDSLParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)
			p.Value()
		}

	case HCLLikeDSLParserT__6:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(79)
			p.Match(HCLLikeDSLParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)
			p.Match(HCLLikeDSLParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(81)
			p.Match(HCLLikeDSLParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(82)
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
	p.EnterRule(localctx, 12, HCLLikeDSLParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
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
