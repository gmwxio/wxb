// Code generated from ADLParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // ADLParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/wxio/goantlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 33, 151,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	3, 2, 3, 2, 3, 2, 7, 2, 18, 10, 2, 12, 2, 14, 2, 21, 11, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 31, 10, 3, 12, 3, 14, 3, 34,
	11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 40, 10, 4, 12, 4, 14, 4, 43, 11, 4,
	3, 4, 3, 4, 5, 4, 47, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 53, 10, 4, 3,
	4, 3, 4, 7, 4, 57, 10, 4, 12, 4, 14, 4, 60, 11, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 5, 4, 67, 10, 4, 3, 4, 3, 4, 3, 4, 5, 4, 72, 10, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 93, 10, 4, 3, 5, 3, 5, 3, 5, 5, 5,
	98, 10, 5, 3, 5, 3, 5, 3, 5, 5, 5, 103, 10, 5, 7, 5, 105, 10, 5, 12, 5,
	14, 5, 108, 11, 5, 3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 114, 10, 6, 3, 6, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 127, 10, 7,
	12, 7, 14, 7, 130, 11, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 7, 7, 142, 10, 7, 12, 7, 14, 7, 145, 11, 7, 3, 7, 3, 7,
	5, 7, 149, 10, 7, 3, 7, 2, 2, 8, 2, 4, 6, 8, 10, 12, 2, 2, 2, 168, 2, 14,
	3, 2, 2, 2, 4, 26, 3, 2, 2, 2, 6, 92, 3, 2, 2, 2, 8, 94, 3, 2, 2, 2, 10,
	111, 3, 2, 2, 2, 12, 148, 3, 2, 2, 2, 14, 15, 5, 4, 3, 2, 15, 19, 7, 3,
	2, 2, 16, 18, 5, 6, 4, 2, 17, 16, 3, 2, 2, 2, 18, 21, 3, 2, 2, 2, 19, 17,
	3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 22, 3, 2, 2, 2, 21, 19, 3, 2, 2, 2,
	22, 23, 7, 4, 2, 2, 23, 24, 7, 10, 2, 2, 24, 25, 7, 2, 2, 3, 25, 3, 3,
	2, 2, 2, 26, 27, 7, 19, 2, 2, 27, 32, 7, 19, 2, 2, 28, 29, 7, 13, 2, 2,
	29, 31, 7, 19, 2, 2, 30, 28, 3, 2, 2, 2, 31, 34, 3, 2, 2, 2, 32, 30, 3,
	2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 5, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 35,
	36, 7, 19, 2, 2, 36, 41, 7, 19, 2, 2, 37, 38, 7, 13, 2, 2, 38, 40, 7, 19,
	2, 2, 39, 37, 3, 2, 2, 2, 40, 43, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 41, 42,
	3, 2, 2, 2, 42, 46, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 44, 45, 7, 13, 2, 2,
	45, 47, 7, 17, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 48, 3,
	2, 2, 2, 48, 93, 7, 10, 2, 2, 49, 50, 7, 19, 2, 2, 50, 52, 7, 19, 2, 2,
	51, 53, 5, 8, 5, 2, 52, 51, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 54, 3,
	2, 2, 2, 54, 58, 7, 3, 2, 2, 55, 57, 5, 10, 6, 2, 56, 55, 3, 2, 2, 2, 57,
	60, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 61, 3, 2, 2,
	2, 60, 58, 3, 2, 2, 2, 61, 62, 7, 4, 2, 2, 62, 93, 7, 10, 2, 2, 63, 64,
	7, 19, 2, 2, 64, 66, 7, 19, 2, 2, 65, 67, 5, 8, 5, 2, 66, 65, 3, 2, 2,
	2, 66, 67, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 69, 7, 7, 2, 2, 69, 71,
	7, 19, 2, 2, 70, 72, 5, 8, 5, 2, 71, 70, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2,
	72, 73, 3, 2, 2, 2, 73, 93, 7, 10, 2, 2, 74, 75, 7, 19, 2, 2, 75, 76, 7,
	19, 2, 2, 76, 77, 5, 12, 7, 2, 77, 78, 7, 10, 2, 2, 78, 93, 3, 2, 2, 2,
	79, 80, 7, 19, 2, 2, 80, 81, 7, 19, 2, 2, 81, 82, 7, 19, 2, 2, 82, 83,
	5, 12, 7, 2, 83, 84, 7, 10, 2, 2, 84, 93, 3, 2, 2, 2, 85, 86, 7, 19, 2,
	2, 86, 87, 7, 19, 2, 2, 87, 88, 7, 11, 2, 2, 88, 89, 7, 19, 2, 2, 89, 90,
	5, 12, 7, 2, 90, 91, 7, 10, 2, 2, 91, 93, 3, 2, 2, 2, 92, 35, 3, 2, 2,
	2, 92, 49, 3, 2, 2, 2, 92, 63, 3, 2, 2, 2, 92, 74, 3, 2, 2, 2, 92, 79,
	3, 2, 2, 2, 92, 85, 3, 2, 2, 2, 93, 7, 3, 2, 2, 2, 94, 95, 7, 15, 2, 2,
	95, 97, 7, 19, 2, 2, 96, 98, 5, 8, 5, 2, 97, 96, 3, 2, 2, 2, 97, 98, 3,
	2, 2, 2, 98, 106, 3, 2, 2, 2, 99, 100, 7, 14, 2, 2, 100, 102, 7, 19, 2,
	2, 101, 103, 5, 8, 5, 2, 102, 101, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103,
	105, 3, 2, 2, 2, 104, 99, 3, 2, 2, 2, 105, 108, 3, 2, 2, 2, 106, 104, 3,
	2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 109, 3, 2, 2, 2, 108, 106, 3, 2, 2,
	2, 109, 110, 7, 16, 2, 2, 110, 9, 3, 2, 2, 2, 111, 113, 7, 19, 2, 2, 112,
	114, 5, 8, 5, 2, 113, 112, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 115,
	3, 2, 2, 2, 115, 116, 7, 19, 2, 2, 116, 117, 7, 10, 2, 2, 117, 11, 3, 2,
	2, 2, 118, 149, 7, 18, 2, 2, 119, 149, 7, 19, 2, 2, 120, 149, 7, 20, 2,
	2, 121, 149, 7, 21, 2, 2, 122, 123, 7, 5, 2, 2, 123, 128, 5, 12, 7, 2,
	124, 125, 7, 14, 2, 2, 125, 127, 5, 12, 7, 2, 126, 124, 3, 2, 2, 2, 127,
	130, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 131,
	3, 2, 2, 2, 130, 128, 3, 2, 2, 2, 131, 132, 7, 6, 2, 2, 132, 149, 3, 2,
	2, 2, 133, 134, 7, 3, 2, 2, 134, 135, 7, 18, 2, 2, 135, 136, 7, 12, 2,
	2, 136, 143, 5, 12, 7, 2, 137, 138, 7, 14, 2, 2, 138, 139, 7, 18, 2, 2,
	139, 140, 7, 12, 2, 2, 140, 142, 5, 12, 7, 2, 141, 137, 3, 2, 2, 2, 142,
	145, 3, 2, 2, 2, 143, 141, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 146,
	3, 2, 2, 2, 145, 143, 3, 2, 2, 2, 146, 147, 7, 4, 2, 2, 147, 149, 3, 2,
	2, 2, 148, 118, 3, 2, 2, 2, 148, 119, 3, 2, 2, 2, 148, 120, 3, 2, 2, 2,
	148, 121, 3, 2, 2, 2, 148, 122, 3, 2, 2, 2, 148, 133, 3, 2, 2, 2, 149,
	13, 3, 2, 2, 2, 18, 19, 32, 41, 46, 52, 58, 66, 71, 92, 97, 102, 106, 113,
	128, 143, 148,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'{'", "'}'", "'['", "']'", "'='", "'\"'", "'''", "';'", "'::'", "':'",
	"'.'", "','", "'<'", "'>'", "'*'",
}
var symbolicNames = []string{
	"", "LCUR", "RCUR", "LSQ", "RSQ", "EQ", "DQ", "SQ", "SEMI", "DCOLON", "COLON",
	"DOT", "COMMA", "LCHEVR", "RCHEVR", "STAR", "STR", "ID", "INT", "FLT",
	"WS", "LINE_COMMENT", "DOWN", "UP", "ROOT", "ERROR", "Module", "Import",
	"Annotation", "Struct", "Union", "Newtype",
}

var ruleNames = []string{
	"adl", "module", "top_level_statement", "typeParam", "soruBody", "jsonValue",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ADLParser struct {
	*antlr.BaseParser
}

func NewADLParser(input antlr.TokenStream) *ADLParser {
	this := new(ADLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "ADLParser.g4"

	return this
}

// ADLParser tokens.
const (
	ADLParserEOF          = antlr.TokenEOF
	ADLParserLCUR         = 1
	ADLParserRCUR         = 2
	ADLParserLSQ          = 3
	ADLParserRSQ          = 4
	ADLParserEQ           = 5
	ADLParserDQ           = 6
	ADLParserSQ           = 7
	ADLParserSEMI         = 8
	ADLParserDCOLON       = 9
	ADLParserCOLON        = 10
	ADLParserDOT          = 11
	ADLParserCOMMA        = 12
	ADLParserLCHEVR       = 13
	ADLParserRCHEVR       = 14
	ADLParserSTAR         = 15
	ADLParserSTR          = 16
	ADLParserID           = 17
	ADLParserINT          = 18
	ADLParserFLT          = 19
	ADLParserWS           = 20
	ADLParserLINE_COMMENT = 21
	ADLParserDOWN         = 22
	ADLParserUP           = 23
	ADLParserROOT         = 24
	ADLParserERROR        = 25
	ADLParserModule       = 26
	ADLParserImport       = 27
	ADLParserAnnotation   = 28
	ADLParserStruct       = 29
	ADLParserUnion        = 30
	ADLParserNewtype      = 31
)

// ADLParser rules.
const (
	ADLParserRULE_adl                 = 0
	ADLParserRULE_module              = 1
	ADLParserRULE_top_level_statement = 2
	ADLParserRULE_typeParam           = 3
	ADLParserRULE_soruBody            = 4
	ADLParserRULE_jsonValue           = 5
)

// IAdlContext is an interface to support dynamic dispatch.
type IAdlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdlContext differentiates from other interfaces.
	IsAdlContext()
}

type AdlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdlContext() *AdlContext {
	var p = new(AdlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_adl
	return p
}

func (*AdlContext) IsAdlContext() {}

func NewAdlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdlContext {
	var p = new(AdlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_adl

	return p
}

func (s *AdlContext) GetParser() antlr.Parser { return s.parser }

func (s *AdlContext) Module() IModuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModuleContext)
}

func (s *AdlContext) LCUR() antlr.TerminalNode {
	return s.GetToken(ADLParserLCUR, 0)
}

func (s *AdlContext) RCUR() antlr.TerminalNode {
	return s.GetToken(ADLParserRCUR, 0)
}

func (s *AdlContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ADLParserSEMI, 0)
}

func (s *AdlContext) EOF() antlr.TerminalNode {
	return s.GetToken(ADLParserEOF, 0)
}

func (s *AdlContext) AllTop_level_statement() []ITop_level_statementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITop_level_statementContext)(nil)).Elem())
	var tst = make([]ITop_level_statementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITop_level_statementContext)
		}
	}

	return tst
}

func (s *AdlContext) Top_level_statement(i int) ITop_level_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITop_level_statementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITop_level_statementContext)
}

func (s *AdlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.EnterAdl(s)
	}
}

func (s *AdlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.ExitAdl(s)
	}
}

func (p *ADLParser) Adl() (localctx IAdlContext) {
	localctx = NewAdlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ADLParserRULE_adl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(12)
		p.Module()
	}
	{
		p.SetState(13)
		p.Match(ADLParserLCUR)
	}
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ADLParserID {
		{
			p.SetState(14)
			p.Top_level_statement()
		}

		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(20)
		p.Match(ADLParserRCUR)
	}
	{
		p.SetState(21)
		p.Match(ADLParserSEMI)
	}
	{
		p.SetState(22)
		p.Match(ADLParserEOF)
	}

	return localctx
}

// IModuleContext is an interface to support dynamic dispatch.
type IModuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModuleContext differentiates from other interfaces.
	IsModuleContext()
}

type ModuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleContext() *ModuleContext {
	var p = new(ModuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_module
	return p
}

func (*ModuleContext) IsModuleContext() {}

func NewModuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleContext {
	var p = new(ModuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_module

	return p
}

func (s *ModuleContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleContext) CopyFrom(ctx *ModuleContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ModuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ModuleStatementContext struct {
	*ModuleContext
	kw_mod antlr.Token
	_ID    antlr.Token
	name   []antlr.Token
}

func NewModuleStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModuleStatementContext {
	var p = new(ModuleStatementContext)

	p.ModuleContext = NewEmptyModuleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ModuleContext))

	return p
}

func (s *ModuleStatementContext) GetKw_mod() antlr.Token { return s.kw_mod }

func (s *ModuleStatementContext) Get_ID() antlr.Token { return s._ID }

func (s *ModuleStatementContext) SetKw_mod(v antlr.Token) { s.kw_mod = v }

func (s *ModuleStatementContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ModuleStatementContext) GetName() []antlr.Token { return s.name }

func (s *ModuleStatementContext) SetName(v []antlr.Token) { s.name = v }

func (s *ModuleStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleStatementContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ADLParserID)
}

func (s *ModuleStatementContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserID, i)
}

func (s *ModuleStatementContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(ADLParserDOT)
}

func (s *ModuleStatementContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserDOT, i)
}

func (s *ModuleStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.EnterModuleStatement(s)
	}
}

func (s *ModuleStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.ExitModuleStatement(s)
	}
}

func (p *ADLParser) Module() (localctx IModuleContext) {
	localctx = NewModuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ADLParserRULE_module)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewModuleStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)

		var _m = p.Match(ADLParserID)

		localctx.(*ModuleStatementContext).kw_mod = _m
	}
	{
		p.SetState(25)

		var _m = p.Match(ADLParserID)

		localctx.(*ModuleStatementContext)._ID = _m
	}
	localctx.(*ModuleStatementContext).name = append(localctx.(*ModuleStatementContext).name, localctx.(*ModuleStatementContext)._ID)
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ADLParserDOT {
		{
			p.SetState(26)
			p.Match(ADLParserDOT)
		}
		{
			p.SetState(27)

			var _m = p.Match(ADLParserID)

			localctx.(*ModuleStatementContext)._ID = _m
		}
		localctx.(*ModuleStatementContext).name = append(localctx.(*ModuleStatementContext).name, localctx.(*ModuleStatementContext)._ID)

		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITop_level_statementContext is an interface to support dynamic dispatch.
type ITop_level_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTop_level_statementContext differentiates from other interfaces.
	IsTop_level_statementContext()
}

type Top_level_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTop_level_statementContext() *Top_level_statementContext {
	var p = new(Top_level_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_top_level_statement
	return p
}

func (*Top_level_statementContext) IsTop_level_statementContext() {}

func NewTop_level_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Top_level_statementContext {
	var p = new(Top_level_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_top_level_statement

	return p
}

func (s *Top_level_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Top_level_statementContext) CopyFrom(ctx *Top_level_statementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Top_level_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Top_level_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructOrUnionContext struct {
	*Top_level_statementContext
	kw_soru antlr.Token
	a       antlr.Token
}

func NewStructOrUnionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructOrUnionContext {
	var p = new(StructOrUnionContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

func (s *StructOrUnionContext) GetKw_soru() antlr.Token { return s.kw_soru }

func (s *StructOrUnionContext) GetA() antlr.Token { return s.a }

func (s *StructOrUnionContext) SetKw_soru(v antlr.Token) { s.kw_soru = v }

func (s *StructOrUnionContext) SetA(v antlr.Token) { s.a = v }

func (s *StructOrUnionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructOrUnionContext) LCUR() antlr.TerminalNode {
	return s.GetToken(ADLParserLCUR, 0)
}

func (s *StructOrUnionContext) RCUR() antlr.TerminalNode {
	return s.GetToken(ADLParserRCUR, 0)
}

func (s *StructOrUnionContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ADLParserSEMI, 0)
}

func (s *StructOrUnionContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ADLParserID)
}

func (s *StructOrUnionContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserID, i)
}

func (s *StructOrUnionContext) TypeParam() ITypeParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeParamContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeParamContext)
}

func (s *StructOrUnionContext) AllSoruBody() []ISoruBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISoruBodyContext)(nil)).Elem())
	var tst = make([]ISoruBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISoruBodyContext)
		}
	}

	return tst
}

func (s *StructOrUnionContext) SoruBody(i int) ISoruBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISoruBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISoruBodyContext)
}

func (s *StructOrUnionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.EnterStructOrUnion(s)
	}
}

func (s *StructOrUnionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.ExitStructOrUnion(s)
	}
}

type DeclAnnotationContext struct {
	*Top_level_statementContext
	kw_anno antlr.Token
	a       antlr.Token
	b       antlr.Token
}

func NewDeclAnnotationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclAnnotationContext {
	var p = new(DeclAnnotationContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

func (s *DeclAnnotationContext) GetKw_anno() antlr.Token { return s.kw_anno }

func (s *DeclAnnotationContext) GetA() antlr.Token { return s.a }

func (s *DeclAnnotationContext) GetB() antlr.Token { return s.b }

func (s *DeclAnnotationContext) SetKw_anno(v antlr.Token) { s.kw_anno = v }

func (s *DeclAnnotationContext) SetA(v antlr.Token) { s.a = v }

func (s *DeclAnnotationContext) SetB(v antlr.Token) { s.b = v }

func (s *DeclAnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclAnnotationContext) JsonValue() IJsonValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValueContext)
}

func (s *DeclAnnotationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ADLParserSEMI, 0)
}

func (s *DeclAnnotationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ADLParserID)
}

func (s *DeclAnnotationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserID, i)
}

func (s *DeclAnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.EnterDeclAnnotation(s)
	}
}

func (s *DeclAnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.ExitDeclAnnotation(s)
	}
}

type FieldAnnotationContext struct {
	*Top_level_statementContext
	kw_anno antlr.Token
	a       antlr.Token
	b       antlr.Token
}

func NewFieldAnnotationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldAnnotationContext {
	var p = new(FieldAnnotationContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

func (s *FieldAnnotationContext) GetKw_anno() antlr.Token { return s.kw_anno }

func (s *FieldAnnotationContext) GetA() antlr.Token { return s.a }

func (s *FieldAnnotationContext) GetB() antlr.Token { return s.b }

func (s *FieldAnnotationContext) SetKw_anno(v antlr.Token) { s.kw_anno = v }

func (s *FieldAnnotationContext) SetA(v antlr.Token) { s.a = v }

func (s *FieldAnnotationContext) SetB(v antlr.Token) { s.b = v }

func (s *FieldAnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldAnnotationContext) DCOLON() antlr.TerminalNode {
	return s.GetToken(ADLParserDCOLON, 0)
}

func (s *FieldAnnotationContext) JsonValue() IJsonValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValueContext)
}

func (s *FieldAnnotationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ADLParserSEMI, 0)
}

func (s *FieldAnnotationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ADLParserID)
}

func (s *FieldAnnotationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserID, i)
}

func (s *FieldAnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.EnterFieldAnnotation(s)
	}
}

func (s *FieldAnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.ExitFieldAnnotation(s)
	}
}

type TypeOrNewtypeContext struct {
	*Top_level_statementContext
	kw_tnew antlr.Token
	a       antlr.Token
	b       antlr.Token
}

func NewTypeOrNewtypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeOrNewtypeContext {
	var p = new(TypeOrNewtypeContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

func (s *TypeOrNewtypeContext) GetKw_tnew() antlr.Token { return s.kw_tnew }

func (s *TypeOrNewtypeContext) GetA() antlr.Token { return s.a }

func (s *TypeOrNewtypeContext) GetB() antlr.Token { return s.b }

func (s *TypeOrNewtypeContext) SetKw_tnew(v antlr.Token) { s.kw_tnew = v }

func (s *TypeOrNewtypeContext) SetA(v antlr.Token) { s.a = v }

func (s *TypeOrNewtypeContext) SetB(v antlr.Token) { s.b = v }

func (s *TypeOrNewtypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeOrNewtypeContext) EQ() antlr.TerminalNode {
	return s.GetToken(ADLParserEQ, 0)
}

func (s *TypeOrNewtypeContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ADLParserSEMI, 0)
}

func (s *TypeOrNewtypeContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ADLParserID)
}

func (s *TypeOrNewtypeContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserID, i)
}

func (s *TypeOrNewtypeContext) AllTypeParam() []ITypeParamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeParamContext)(nil)).Elem())
	var tst = make([]ITypeParamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeParamContext)
		}
	}

	return tst
}

func (s *TypeOrNewtypeContext) TypeParam(i int) ITypeParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeParamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeParamContext)
}

func (s *TypeOrNewtypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.EnterTypeOrNewtype(s)
	}
}

func (s *TypeOrNewtypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.ExitTypeOrNewtype(s)
	}
}

type ImportStatementContext struct {
	*Top_level_statementContext
	kw_impt antlr.Token
	_ID     antlr.Token
	a       []antlr.Token
	s       antlr.Token
}

func NewImportStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImportStatementContext {
	var p = new(ImportStatementContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

func (s *ImportStatementContext) GetKw_impt() antlr.Token { return s.kw_impt }

func (s *ImportStatementContext) Get_ID() antlr.Token { return s._ID }

func (s *ImportStatementContext) GetS() antlr.Token { return s.s }

func (s *ImportStatementContext) SetKw_impt(v antlr.Token) { s.kw_impt = v }

func (s *ImportStatementContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ImportStatementContext) SetS(v antlr.Token) { s.s = v }

func (s *ImportStatementContext) GetA() []antlr.Token { return s.a }

func (s *ImportStatementContext) SetA(v []antlr.Token) { s.a = v }

func (s *ImportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ADLParserSEMI, 0)
}

func (s *ImportStatementContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ADLParserID)
}

func (s *ImportStatementContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserID, i)
}

func (s *ImportStatementContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(ADLParserDOT)
}

func (s *ImportStatementContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserDOT, i)
}

func (s *ImportStatementContext) STAR() antlr.TerminalNode {
	return s.GetToken(ADLParserSTAR, 0)
}

func (s *ImportStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.EnterImportStatement(s)
	}
}

func (s *ImportStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.ExitImportStatement(s)
	}
}

type ModuleAnnotationContext struct {
	*Top_level_statementContext
	kw_anno antlr.Token
	a       antlr.Token
}

func NewModuleAnnotationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModuleAnnotationContext {
	var p = new(ModuleAnnotationContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

func (s *ModuleAnnotationContext) GetKw_anno() antlr.Token { return s.kw_anno }

func (s *ModuleAnnotationContext) GetA() antlr.Token { return s.a }

func (s *ModuleAnnotationContext) SetKw_anno(v antlr.Token) { s.kw_anno = v }

func (s *ModuleAnnotationContext) SetA(v antlr.Token) { s.a = v }

func (s *ModuleAnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleAnnotationContext) JsonValue() IJsonValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValueContext)
}

func (s *ModuleAnnotationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ADLParserSEMI, 0)
}

func (s *ModuleAnnotationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ADLParserID)
}

func (s *ModuleAnnotationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserID, i)
}

func (s *ModuleAnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.EnterModuleAnnotation(s)
	}
}

func (s *ModuleAnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.ExitModuleAnnotation(s)
	}
}

func (p *ADLParser) Top_level_statement() (localctx ITop_level_statementContext) {
	localctx = NewTop_level_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ADLParserRULE_top_level_statement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewImportStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)

			var _m = p.Match(ADLParserID)

			localctx.(*ImportStatementContext).kw_impt = _m
		}
		{
			p.SetState(34)

			var _m = p.Match(ADLParserID)

			localctx.(*ImportStatementContext)._ID = _m
		}
		localctx.(*ImportStatementContext).a = append(localctx.(*ImportStatementContext).a, localctx.(*ImportStatementContext)._ID)
		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(35)
					p.Match(ADLParserDOT)
				}
				{
					p.SetState(36)

					var _m = p.Match(ADLParserID)

					localctx.(*ImportStatementContext)._ID = _m
				}
				localctx.(*ImportStatementContext).a = append(localctx.(*ImportStatementContext).a, localctx.(*ImportStatementContext)._ID)

			}
			p.SetState(41)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLParserDOT {
			{
				p.SetState(42)
				p.Match(ADLParserDOT)
			}
			{
				p.SetState(43)

				var _m = p.Match(ADLParserSTAR)

				localctx.(*ImportStatementContext).s = _m
			}

		}
		{
			p.SetState(46)
			p.Match(ADLParserSEMI)
		}

	case 2:
		localctx = NewStructOrUnionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(47)

			var _m = p.Match(ADLParserID)

			localctx.(*StructOrUnionContext).kw_soru = _m
		}
		{
			p.SetState(48)

			var _m = p.Match(ADLParserID)

			localctx.(*StructOrUnionContext).a = _m
		}
		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLParserLCHEVR {
			{
				p.SetState(49)
				p.TypeParam()
			}

		}
		{
			p.SetState(52)
			p.Match(ADLParserLCUR)
		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLParserID {
			{
				p.SetState(53)
				p.SoruBody()
			}

			p.SetState(58)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(59)
			p.Match(ADLParserRCUR)
		}
		{
			p.SetState(60)
			p.Match(ADLParserSEMI)
		}

	case 3:
		localctx = NewTypeOrNewtypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(61)

			var _m = p.Match(ADLParserID)

			localctx.(*TypeOrNewtypeContext).kw_tnew = _m
		}
		{
			p.SetState(62)

			var _m = p.Match(ADLParserID)

			localctx.(*TypeOrNewtypeContext).a = _m
		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLParserLCHEVR {
			{
				p.SetState(63)
				p.TypeParam()
			}

		}
		{
			p.SetState(66)
			p.Match(ADLParserEQ)
		}
		{
			p.SetState(67)

			var _m = p.Match(ADLParserID)

			localctx.(*TypeOrNewtypeContext).b = _m
		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLParserLCHEVR {
			{
				p.SetState(68)
				p.TypeParam()
			}

		}
		{
			p.SetState(71)
			p.Match(ADLParserSEMI)
		}

	case 4:
		localctx = NewModuleAnnotationContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(72)

			var _m = p.Match(ADLParserID)

			localctx.(*ModuleAnnotationContext).kw_anno = _m
		}
		{
			p.SetState(73)

			var _m = p.Match(ADLParserID)

			localctx.(*ModuleAnnotationContext).a = _m
		}
		{
			p.SetState(74)
			p.JsonValue()
		}
		{
			p.SetState(75)
			p.Match(ADLParserSEMI)
		}

	case 5:
		localctx = NewDeclAnnotationContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(77)

			var _m = p.Match(ADLParserID)

			localctx.(*DeclAnnotationContext).kw_anno = _m
		}
		{
			p.SetState(78)

			var _m = p.Match(ADLParserID)

			localctx.(*DeclAnnotationContext).a = _m
		}
		{
			p.SetState(79)

			var _m = p.Match(ADLParserID)

			localctx.(*DeclAnnotationContext).b = _m
		}
		{
			p.SetState(80)
			p.JsonValue()
		}
		{
			p.SetState(81)
			p.Match(ADLParserSEMI)
		}

	case 6:
		localctx = NewFieldAnnotationContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(83)

			var _m = p.Match(ADLParserID)

			localctx.(*FieldAnnotationContext).kw_anno = _m
		}
		{
			p.SetState(84)

			var _m = p.Match(ADLParserID)

			localctx.(*FieldAnnotationContext).a = _m
		}
		{
			p.SetState(85)
			p.Match(ADLParserDCOLON)
		}
		{
			p.SetState(86)

			var _m = p.Match(ADLParserID)

			localctx.(*FieldAnnotationContext).b = _m
		}
		{
			p.SetState(87)
			p.JsonValue()
		}
		{
			p.SetState(88)
			p.Match(ADLParserSEMI)
		}

	}

	return localctx
}

// ITypeParamContext is an interface to support dynamic dispatch.
type ITypeParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetTypep returns the typep token list.
	GetTypep() []antlr.Token

	// SetTypep sets the typep token list.
	SetTypep([]antlr.Token)

	// IsTypeParamContext differentiates from other interfaces.
	IsTypeParamContext()
}

type TypeParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_ID    antlr.Token
	typep  []antlr.Token
}

func NewEmptyTypeParamContext() *TypeParamContext {
	var p = new(TypeParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_typeParam
	return p
}

func (*TypeParamContext) IsTypeParamContext() {}

func NewTypeParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeParamContext {
	var p = new(TypeParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_typeParam

	return p
}

func (s *TypeParamContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeParamContext) Get_ID() antlr.Token { return s._ID }

func (s *TypeParamContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *TypeParamContext) GetTypep() []antlr.Token { return s.typep }

func (s *TypeParamContext) SetTypep(v []antlr.Token) { s.typep = v }

func (s *TypeParamContext) LCHEVR() antlr.TerminalNode {
	return s.GetToken(ADLParserLCHEVR, 0)
}

func (s *TypeParamContext) RCHEVR() antlr.TerminalNode {
	return s.GetToken(ADLParserRCHEVR, 0)
}

func (s *TypeParamContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ADLParserID)
}

func (s *TypeParamContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserID, i)
}

func (s *TypeParamContext) AllTypeParam() []ITypeParamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeParamContext)(nil)).Elem())
	var tst = make([]ITypeParamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeParamContext)
		}
	}

	return tst
}

func (s *TypeParamContext) TypeParam(i int) ITypeParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeParamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeParamContext)
}

func (s *TypeParamContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ADLParserCOMMA)
}

func (s *TypeParamContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserCOMMA, i)
}

func (s *TypeParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.EnterTypeParam(s)
	}
}

func (s *TypeParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.ExitTypeParam(s)
	}
}

func (p *ADLParser) TypeParam() (localctx ITypeParamContext) {
	localctx = NewTypeParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ADLParserRULE_typeParam)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(ADLParserLCHEVR)
	}
	{
		p.SetState(93)

		var _m = p.Match(ADLParserID)

		localctx.(*TypeParamContext)._ID = _m
	}
	localctx.(*TypeParamContext).typep = append(localctx.(*TypeParamContext).typep, localctx.(*TypeParamContext)._ID)
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ADLParserLCHEVR {
		{
			p.SetState(94)
			p.TypeParam()
		}

	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ADLParserCOMMA {
		{
			p.SetState(97)
			p.Match(ADLParserCOMMA)
		}
		{
			p.SetState(98)

			var _m = p.Match(ADLParserID)

			localctx.(*TypeParamContext)._ID = _m
		}
		localctx.(*TypeParamContext).typep = append(localctx.(*TypeParamContext).typep, localctx.(*TypeParamContext)._ID)
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLParserLCHEVR {
			{
				p.SetState(99)
				p.TypeParam()
			}

		}

		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(107)
		p.Match(ADLParserRCHEVR)
	}

	return localctx
}

// ISoruBodyContext is an interface to support dynamic dispatch.
type ISoruBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSoruBodyContext differentiates from other interfaces.
	IsSoruBodyContext()
}

type SoruBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySoruBodyContext() *SoruBodyContext {
	var p = new(SoruBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_soruBody
	return p
}

func (*SoruBodyContext) IsSoruBodyContext() {}

func NewSoruBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SoruBodyContext {
	var p = new(SoruBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_soruBody

	return p
}

func (s *SoruBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *SoruBodyContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ADLParserID)
}

func (s *SoruBodyContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserID, i)
}

func (s *SoruBodyContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ADLParserSEMI, 0)
}

func (s *SoruBodyContext) TypeParam() ITypeParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeParamContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeParamContext)
}

func (s *SoruBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SoruBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SoruBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.EnterSoruBody(s)
	}
}

func (s *SoruBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.ExitSoruBody(s)
	}
}

func (p *ADLParser) SoruBody() (localctx ISoruBodyContext) {
	localctx = NewSoruBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ADLParserRULE_soruBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(ADLParserID)
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ADLParserLCHEVR {
		{
			p.SetState(110)
			p.TypeParam()
		}

	}
	{
		p.SetState(113)
		p.Match(ADLParserID)
	}
	{
		p.SetState(114)
		p.Match(ADLParserSEMI)
	}

	return localctx
}

// IJsonValueContext is an interface to support dynamic dispatch.
type IJsonValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJsonValueContext differentiates from other interfaces.
	IsJsonValueContext()
}

type JsonValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonValueContext() *JsonValueContext {
	var p = new(JsonValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_jsonValue
	return p
}

func (*JsonValueContext) IsJsonValueContext() {}

func NewJsonValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonValueContext {
	var p = new(JsonValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_jsonValue

	return p
}

func (s *JsonValueContext) GetParser() antlr.Parser { return s.parser }

func (s *JsonValueContext) CopyFrom(ctx *JsonValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *JsonValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TrueFalseNullContext struct {
	*JsonValueContext
	kw_tfn antlr.Token
}

func NewTrueFalseNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueFalseNullContext {
	var p = new(TrueFalseNullContext)

	p.JsonValueContext = NewEmptyJsonValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValueContext))

	return p
}

func (s *TrueFalseNullContext) GetKw_tfn() antlr.Token { return s.kw_tfn }

func (s *TrueFalseNullContext) SetKw_tfn(v antlr.Token) { s.kw_tfn = v }

func (s *TrueFalseNullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueFalseNullContext) ID() antlr.TerminalNode {
	return s.GetToken(ADLParserID, 0)
}

func (s *TrueFalseNullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.EnterTrueFalseNull(s)
	}
}

func (s *TrueFalseNullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.ExitTrueFalseNull(s)
	}
}

type ObjStatementContext struct {
	*JsonValueContext
}

func NewObjStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjStatementContext {
	var p = new(ObjStatementContext)

	p.JsonValueContext = NewEmptyJsonValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValueContext))

	return p
}

func (s *ObjStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjStatementContext) LCUR() antlr.TerminalNode {
	return s.GetToken(ADLParserLCUR, 0)
}

func (s *ObjStatementContext) AllSTR() []antlr.TerminalNode {
	return s.GetTokens(ADLParserSTR)
}

func (s *ObjStatementContext) STR(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserSTR, i)
}

func (s *ObjStatementContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(ADLParserCOLON)
}

func (s *ObjStatementContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserCOLON, i)
}

func (s *ObjStatementContext) AllJsonValue() []IJsonValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IJsonValueContext)(nil)).Elem())
	var tst = make([]IJsonValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonValueContext)
		}
	}

	return tst
}

func (s *ObjStatementContext) JsonValue(i int) IJsonValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonValueContext)
}

func (s *ObjStatementContext) RCUR() antlr.TerminalNode {
	return s.GetToken(ADLParserRCUR, 0)
}

func (s *ObjStatementContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ADLParserCOMMA)
}

func (s *ObjStatementContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserCOMMA, i)
}

func (s *ObjStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.EnterObjStatement(s)
	}
}

func (s *ObjStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.ExitObjStatement(s)
	}
}

type FloatStatementContext struct {
	*JsonValueContext
}

func NewFloatStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatStatementContext {
	var p = new(FloatStatementContext)

	p.JsonValueContext = NewEmptyJsonValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValueContext))

	return p
}

func (s *FloatStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatStatementContext) FLT() antlr.TerminalNode {
	return s.GetToken(ADLParserFLT, 0)
}

func (s *FloatStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.EnterFloatStatement(s)
	}
}

func (s *FloatStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.ExitFloatStatement(s)
	}
}

type ArrayStatementContext struct {
	*JsonValueContext
}

func NewArrayStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayStatementContext {
	var p = new(ArrayStatementContext)

	p.JsonValueContext = NewEmptyJsonValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValueContext))

	return p
}

func (s *ArrayStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayStatementContext) LSQ() antlr.TerminalNode {
	return s.GetToken(ADLParserLSQ, 0)
}

func (s *ArrayStatementContext) AllJsonValue() []IJsonValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IJsonValueContext)(nil)).Elem())
	var tst = make([]IJsonValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonValueContext)
		}
	}

	return tst
}

func (s *ArrayStatementContext) JsonValue(i int) IJsonValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonValueContext)
}

func (s *ArrayStatementContext) RSQ() antlr.TerminalNode {
	return s.GetToken(ADLParserRSQ, 0)
}

func (s *ArrayStatementContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ADLParserCOMMA)
}

func (s *ArrayStatementContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserCOMMA, i)
}

func (s *ArrayStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.EnterArrayStatement(s)
	}
}

func (s *ArrayStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.ExitArrayStatement(s)
	}
}

type NumberStatementContext struct {
	*JsonValueContext
}

func NewNumberStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberStatementContext {
	var p = new(NumberStatementContext)

	p.JsonValueContext = NewEmptyJsonValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValueContext))

	return p
}

func (s *NumberStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberStatementContext) INT() antlr.TerminalNode {
	return s.GetToken(ADLParserINT, 0)
}

func (s *NumberStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.EnterNumberStatement(s)
	}
}

func (s *NumberStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.ExitNumberStatement(s)
	}
}

type StringStatementContext struct {
	*JsonValueContext
	s antlr.Token
}

func NewStringStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringStatementContext {
	var p = new(StringStatementContext)

	p.JsonValueContext = NewEmptyJsonValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValueContext))

	return p
}

func (s *StringStatementContext) GetS() antlr.Token { return s.s }

func (s *StringStatementContext) SetS(v antlr.Token) { s.s = v }

func (s *StringStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringStatementContext) STR() antlr.TerminalNode {
	return s.GetToken(ADLParserSTR, 0)
}

func (s *StringStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.EnterStringStatement(s)
	}
}

func (s *StringStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.ExitStringStatement(s)
	}
}

func (p *ADLParser) JsonValue() (localctx IJsonValueContext) {
	localctx = NewJsonValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ADLParserRULE_jsonValue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(146)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ADLParserSTR:
		localctx = NewStringStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)

			var _m = p.Match(ADLParserSTR)

			localctx.(*StringStatementContext).s = _m
		}

	case ADLParserID:
		localctx = NewTrueFalseNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)

			var _m = p.Match(ADLParserID)

			localctx.(*TrueFalseNullContext).kw_tfn = _m
		}

	case ADLParserINT:
		localctx = NewNumberStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(118)
			p.Match(ADLParserINT)
		}

	case ADLParserFLT:
		localctx = NewFloatStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(119)
			p.Match(ADLParserFLT)
		}

	case ADLParserLSQ:
		localctx = NewArrayStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(120)
			p.Match(ADLParserLSQ)
		}
		{
			p.SetState(121)
			p.JsonValue()
		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLParserCOMMA {
			{
				p.SetState(122)
				p.Match(ADLParserCOMMA)
			}
			{
				p.SetState(123)
				p.JsonValue()
			}

			p.SetState(128)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(129)
			p.Match(ADLParserRSQ)
		}

	case ADLParserLCUR:
		localctx = NewObjStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(131)
			p.Match(ADLParserLCUR)
		}
		{
			p.SetState(132)
			p.Match(ADLParserSTR)
		}
		{
			p.SetState(133)
			p.Match(ADLParserCOLON)
		}
		{
			p.SetState(134)
			p.JsonValue()
		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLParserCOMMA {
			{
				p.SetState(135)
				p.Match(ADLParserCOMMA)
			}
			{
				p.SetState(136)
				p.Match(ADLParserSTR)
			}
			{
				p.SetState(137)
				p.Match(ADLParserCOLON)
			}
			{
				p.SetState(138)
				p.JsonValue()
			}

			p.SetState(143)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(144)
			p.Match(ADLParserRCUR)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
