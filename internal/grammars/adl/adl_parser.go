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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 35, 216,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 2, 7,
	2, 26, 10, 2, 12, 2, 14, 2, 29, 11, 2, 3, 2, 7, 2, 32, 10, 2, 12, 2, 14,
	2, 35, 11, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3, 42, 10, 3, 12, 3, 14,
	3, 45, 11, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 51, 10, 3, 12, 3, 14, 3, 54,
	11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 60, 10, 4, 12, 4, 14, 4, 63, 11, 4,
	3, 4, 3, 4, 5, 4, 67, 10, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5,
	75, 10, 5, 3, 6, 7, 6, 78, 10, 6, 12, 6, 14, 6, 81, 11, 6, 3, 6, 3, 6,
	3, 6, 5, 6, 86, 10, 6, 3, 6, 3, 6, 7, 6, 90, 10, 6, 12, 6, 14, 6, 93, 11,
	6, 3, 6, 3, 6, 3, 6, 7, 6, 98, 10, 6, 12, 6, 14, 6, 101, 11, 6, 3, 6, 3,
	6, 3, 6, 5, 6, 106, 10, 6, 3, 6, 3, 6, 3, 6, 5, 6, 111, 10, 6, 3, 6, 3,
	6, 5, 6, 115, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5,
	6, 137, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 143, 10, 7, 12, 7, 14, 7,
	146, 11, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 154, 10, 8, 12, 8,
	14, 8, 157, 11, 8, 3, 8, 3, 8, 3, 9, 3, 9, 5, 9, 163, 10, 9, 3, 10, 7,
	10, 166, 10, 10, 12, 10, 14, 10, 169, 11, 10, 3, 10, 3, 10, 5, 10, 173,
	10, 10, 3, 10, 3, 10, 3, 10, 5, 10, 178, 10, 10, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 190, 10, 11, 12, 11,
	14, 11, 193, 11, 11, 5, 11, 195, 10, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 206, 10, 11, 12, 11, 14, 11, 209,
	11, 11, 5, 11, 211, 10, 11, 3, 11, 5, 11, 214, 10, 11, 3, 11, 2, 2, 12,
	2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 2, 2, 2, 238, 2, 22, 3, 2, 2, 2, 4,
	43, 3, 2, 2, 2, 6, 55, 3, 2, 2, 2, 8, 74, 3, 2, 2, 2, 10, 136, 3, 2, 2,
	2, 12, 138, 3, 2, 2, 2, 14, 149, 3, 2, 2, 2, 16, 160, 3, 2, 2, 2, 18, 167,
	3, 2, 2, 2, 20, 213, 3, 2, 2, 2, 22, 23, 5, 4, 3, 2, 23, 27, 7, 3, 2, 2,
	24, 26, 5, 6, 4, 2, 25, 24, 3, 2, 2, 2, 26, 29, 3, 2, 2, 2, 27, 25, 3,
	2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 33, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 30,
	32, 5, 10, 6, 2, 31, 30, 3, 2, 2, 2, 32, 35, 3, 2, 2, 2, 33, 31, 3, 2,
	2, 2, 33, 34, 3, 2, 2, 2, 34, 36, 3, 2, 2, 2, 35, 33, 3, 2, 2, 2, 36, 37,
	7, 4, 2, 2, 37, 38, 7, 10, 2, 2, 38, 39, 7, 2, 2, 3, 39, 3, 3, 2, 2, 2,
	40, 42, 5, 8, 5, 2, 41, 40, 3, 2, 2, 2, 42, 45, 3, 2, 2, 2, 43, 41, 3,
	2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 46, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 46,
	47, 7, 20, 2, 2, 47, 52, 7, 20, 2, 2, 48, 49, 7, 13, 2, 2, 49, 51, 7, 20,
	2, 2, 50, 48, 3, 2, 2, 2, 51, 54, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 52, 53,
	3, 2, 2, 2, 53, 5, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 55, 56, 7, 20, 2, 2,
	56, 61, 7, 20, 2, 2, 57, 58, 7, 13, 2, 2, 58, 60, 7, 20, 2, 2, 59, 57,
	3, 2, 2, 2, 60, 63, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2,
	62, 66, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 64, 65, 7, 13, 2, 2, 65, 67, 7,
	17, 2, 2, 66, 64, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68,
	69, 7, 10, 2, 2, 69, 7, 3, 2, 2, 2, 70, 71, 7, 18, 2, 2, 71, 72, 7, 20,
	2, 2, 72, 75, 5, 20, 11, 2, 73, 75, 7, 24, 2, 2, 74, 70, 3, 2, 2, 2, 74,
	73, 3, 2, 2, 2, 75, 9, 3, 2, 2, 2, 76, 78, 5, 8, 5, 2, 77, 76, 3, 2, 2,
	2, 78, 81, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 82,
	3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 82, 83, 7, 20, 2, 2, 83, 85, 7, 20, 2,
	2, 84, 86, 5, 12, 7, 2, 85, 84, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 87,
	3, 2, 2, 2, 87, 91, 7, 3, 2, 2, 88, 90, 5, 18, 10, 2, 89, 88, 3, 2, 2,
	2, 90, 93, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 94,
	3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 94, 95, 7, 4, 2, 2, 95, 137, 7, 10, 2,
	2, 96, 98, 5, 8, 5, 2, 97, 96, 3, 2, 2, 2, 98, 101, 3, 2, 2, 2, 99, 97,
	3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 102, 3, 2, 2, 2, 101, 99, 3, 2, 2,
	2, 102, 103, 7, 20, 2, 2, 103, 105, 7, 20, 2, 2, 104, 106, 5, 12, 7, 2,
	105, 104, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107,
	108, 7, 7, 2, 2, 108, 110, 7, 20, 2, 2, 109, 111, 5, 14, 8, 2, 110, 109,
	3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 114, 3, 2, 2, 2, 112, 113, 7, 7,
	2, 2, 113, 115, 5, 20, 11, 2, 114, 112, 3, 2, 2, 2, 114, 115, 3, 2, 2,
	2, 115, 116, 3, 2, 2, 2, 116, 137, 7, 10, 2, 2, 117, 118, 7, 20, 2, 2,
	118, 119, 7, 20, 2, 2, 119, 120, 5, 20, 11, 2, 120, 121, 7, 10, 2, 2, 121,
	137, 3, 2, 2, 2, 122, 123, 7, 20, 2, 2, 123, 124, 7, 20, 2, 2, 124, 125,
	7, 20, 2, 2, 125, 126, 5, 20, 11, 2, 126, 127, 7, 10, 2, 2, 127, 137, 3,
	2, 2, 2, 128, 129, 7, 20, 2, 2, 129, 130, 7, 20, 2, 2, 130, 131, 7, 11,
	2, 2, 131, 132, 7, 20, 2, 2, 132, 133, 7, 20, 2, 2, 133, 134, 5, 20, 11,
	2, 134, 135, 7, 10, 2, 2, 135, 137, 3, 2, 2, 2, 136, 79, 3, 2, 2, 2, 136,
	99, 3, 2, 2, 2, 136, 117, 3, 2, 2, 2, 136, 122, 3, 2, 2, 2, 136, 128, 3,
	2, 2, 2, 137, 11, 3, 2, 2, 2, 138, 139, 7, 15, 2, 2, 139, 144, 7, 20, 2,
	2, 140, 141, 7, 14, 2, 2, 141, 143, 7, 20, 2, 2, 142, 140, 3, 2, 2, 2,
	143, 146, 3, 2, 2, 2, 144, 142, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145,
	147, 3, 2, 2, 2, 146, 144, 3, 2, 2, 2, 147, 148, 7, 16, 2, 2, 148, 13,
	3, 2, 2, 2, 149, 150, 7, 15, 2, 2, 150, 155, 5, 16, 9, 2, 151, 152, 7,
	14, 2, 2, 152, 154, 5, 16, 9, 2, 153, 151, 3, 2, 2, 2, 154, 157, 3, 2,
	2, 2, 155, 153, 3, 2, 2, 2, 155, 156, 3, 2, 2, 2, 156, 158, 3, 2, 2, 2,
	157, 155, 3, 2, 2, 2, 158, 159, 7, 16, 2, 2, 159, 15, 3, 2, 2, 2, 160,
	162, 7, 20, 2, 2, 161, 163, 5, 14, 8, 2, 162, 161, 3, 2, 2, 2, 162, 163,
	3, 2, 2, 2, 163, 17, 3, 2, 2, 2, 164, 166, 5, 8, 5, 2, 165, 164, 3, 2,
	2, 2, 166, 169, 3, 2, 2, 2, 167, 165, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2,
	168, 170, 3, 2, 2, 2, 169, 167, 3, 2, 2, 2, 170, 172, 7, 20, 2, 2, 171,
	173, 5, 14, 8, 2, 172, 171, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 173, 174,
	3, 2, 2, 2, 174, 177, 7, 20, 2, 2, 175, 176, 7, 7, 2, 2, 176, 178, 5, 20,
	11, 2, 177, 175, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2,
	179, 180, 7, 10, 2, 2, 180, 19, 3, 2, 2, 2, 181, 214, 7, 19, 2, 2, 182,
	214, 7, 20, 2, 2, 183, 214, 7, 21, 2, 2, 184, 214, 7, 22, 2, 2, 185, 194,
	7, 5, 2, 2, 186, 191, 5, 20, 11, 2, 187, 188, 7, 14, 2, 2, 188, 190, 5,
	20, 11, 2, 189, 187, 3, 2, 2, 2, 190, 193, 3, 2, 2, 2, 191, 189, 3, 2,
	2, 2, 191, 192, 3, 2, 2, 2, 192, 195, 3, 2, 2, 2, 193, 191, 3, 2, 2, 2,
	194, 186, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196,
	214, 7, 6, 2, 2, 197, 210, 7, 3, 2, 2, 198, 199, 7, 19, 2, 2, 199, 200,
	7, 12, 2, 2, 200, 207, 5, 20, 11, 2, 201, 202, 7, 14, 2, 2, 202, 203, 7,
	19, 2, 2, 203, 204, 7, 12, 2, 2, 204, 206, 5, 20, 11, 2, 205, 201, 3, 2,
	2, 2, 206, 209, 3, 2, 2, 2, 207, 205, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2,
	208, 211, 3, 2, 2, 2, 209, 207, 3, 2, 2, 2, 210, 198, 3, 2, 2, 2, 210,
	211, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 214, 7, 4, 2, 2, 213, 181,
	3, 2, 2, 2, 213, 182, 3, 2, 2, 2, 213, 183, 3, 2, 2, 2, 213, 184, 3, 2,
	2, 2, 213, 185, 3, 2, 2, 2, 213, 197, 3, 2, 2, 2, 214, 21, 3, 2, 2, 2,
	28, 27, 33, 43, 52, 61, 66, 74, 79, 85, 91, 99, 105, 110, 114, 136, 144,
	155, 162, 167, 172, 177, 191, 194, 207, 210, 213,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'{'", "'}'", "'['", "']'", "'='", "'\"'", "'''", "';'", "'::'", "':'",
	"'.'", "','", "'<'", "'>'", "'*'", "'@'",
}
var symbolicNames = []string{
	"", "LCUR", "RCUR", "LSQ", "RSQ", "EQ", "DQ", "SQ", "SEMI", "DCOLON", "COLON",
	"DOT", "COMMA", "LCHEVR", "RCHEVR", "STAR", "AT", "STR", "ID", "INT", "FLT",
	"WS", "LINE_DOC", "LINE_COMMENT", "DOWN", "UP", "ROOT", "ERROR", "Module",
	"Import", "Annotation", "Struct", "Union", "Newtype",
}

var ruleNames = []string{
	"adl", "module", "imports", "annon", "top_level_statement", "typeParam",
	"typeExpr", "typeExprElem", "soruBody", "jsonValue",
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
	ADLParserAT           = 16
	ADLParserSTR          = 17
	ADLParserID           = 18
	ADLParserINT          = 19
	ADLParserFLT          = 20
	ADLParserWS           = 21
	ADLParserLINE_DOC     = 22
	ADLParserLINE_COMMENT = 23
	ADLParserDOWN         = 24
	ADLParserUP           = 25
	ADLParserROOT         = 26
	ADLParserERROR        = 27
	ADLParserModule       = 28
	ADLParserImport       = 29
	ADLParserAnnotation   = 30
	ADLParserStruct       = 31
	ADLParserUnion        = 32
	ADLParserNewtype      = 33
)

// ADLParser rules.
const (
	ADLParserRULE_adl                 = 0
	ADLParserRULE_module              = 1
	ADLParserRULE_imports             = 2
	ADLParserRULE_annon               = 3
	ADLParserRULE_top_level_statement = 4
	ADLParserRULE_typeParam           = 5
	ADLParserRULE_typeExpr            = 6
	ADLParserRULE_typeExprElem        = 7
	ADLParserRULE_soruBody            = 8
	ADLParserRULE_jsonValue           = 9
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

func (s *AdlContext) AllImports() []IImportsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportsContext)(nil)).Elem())
	var tst = make([]IImportsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportsContext)
		}
	}

	return tst
}

func (s *AdlContext) Imports(i int) IImportsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportsContext)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)
		p.Module()
	}
	{
		p.SetState(21)
		p.Match(ADLParserLCUR)
	}
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(22)
				p.Imports()
			}

		}
		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ADLParserAT)|(1<<ADLParserID)|(1<<ADLParserLINE_DOC))) != 0 {
		{
			p.SetState(28)
			p.Top_level_statement()
		}

		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(34)
		p.Match(ADLParserRCUR)
	}
	{
		p.SetState(35)
		p.Match(ADLParserSEMI)
	}
	{
		p.SetState(36)
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

func (s *ModuleStatementContext) AllAnnon() []IAnnonContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnonContext)(nil)).Elem())
	var tst = make([]IAnnonContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnonContext)
		}
	}

	return tst
}

func (s *ModuleStatementContext) Annon(i int) IAnnonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnonContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnonContext)
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
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ADLParserAT || _la == ADLParserLINE_DOC {
		{
			p.SetState(38)
			p.Annon()
		}

		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(44)

		var _m = p.Match(ADLParserID)

		localctx.(*ModuleStatementContext).kw_mod = _m
	}
	{
		p.SetState(45)

		var _m = p.Match(ADLParserID)

		localctx.(*ModuleStatementContext)._ID = _m
	}
	localctx.(*ModuleStatementContext).name = append(localctx.(*ModuleStatementContext).name, localctx.(*ModuleStatementContext)._ID)
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ADLParserDOT {
		{
			p.SetState(46)
			p.Match(ADLParserDOT)
		}
		{
			p.SetState(47)

			var _m = p.Match(ADLParserID)

			localctx.(*ModuleStatementContext)._ID = _m
		}
		localctx.(*ModuleStatementContext).name = append(localctx.(*ModuleStatementContext).name, localctx.(*ModuleStatementContext)._ID)

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IImportsContext is an interface to support dynamic dispatch.
type IImportsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportsContext differentiates from other interfaces.
	IsImportsContext()
}

type ImportsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportsContext() *ImportsContext {
	var p = new(ImportsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_imports
	return p
}

func (*ImportsContext) IsImportsContext() {}

func NewImportsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportsContext {
	var p = new(ImportsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_imports

	return p
}

func (s *ImportsContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportsContext) CopyFrom(ctx *ImportsContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ImportsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ImportStatementContext struct {
	*ImportsContext
	kw_impt antlr.Token
	_ID     antlr.Token
	a       []antlr.Token
	s       antlr.Token
}

func NewImportStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImportStatementContext {
	var p = new(ImportStatementContext)

	p.ImportsContext = NewEmptyImportsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ImportsContext))

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

func (p *ADLParser) Imports() (localctx IImportsContext) {
	localctx = NewImportsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ADLParserRULE_imports)
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

	localctx = NewImportStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)

		var _m = p.Match(ADLParserID)

		localctx.(*ImportStatementContext).kw_impt = _m
	}
	{
		p.SetState(54)

		var _m = p.Match(ADLParserID)

		localctx.(*ImportStatementContext)._ID = _m
	}
	localctx.(*ImportStatementContext).a = append(localctx.(*ImportStatementContext).a, localctx.(*ImportStatementContext)._ID)
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(55)
				p.Match(ADLParserDOT)
			}
			{
				p.SetState(56)

				var _m = p.Match(ADLParserID)

				localctx.(*ImportStatementContext)._ID = _m
			}
			localctx.(*ImportStatementContext).a = append(localctx.(*ImportStatementContext).a, localctx.(*ImportStatementContext)._ID)

		}
		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ADLParserDOT {
		{
			p.SetState(62)
			p.Match(ADLParserDOT)
		}
		{
			p.SetState(63)

			var _m = p.Match(ADLParserSTAR)

			localctx.(*ImportStatementContext).s = _m
		}

	}
	{
		p.SetState(66)
		p.Match(ADLParserSEMI)
	}

	return localctx
}

// IAnnonContext is an interface to support dynamic dispatch.
type IAnnonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnnonContext differentiates from other interfaces.
	IsAnnonContext()
}

type AnnonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnonContext() *AnnonContext {
	var p = new(AnnonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_annon
	return p
}

func (*AnnonContext) IsAnnonContext() {}

func NewAnnonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnonContext {
	var p = new(AnnonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_annon

	return p
}

func (s *AnnonContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnonContext) CopyFrom(ctx *AnnonContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AnnonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DocAnnoContext struct {
	*AnnonContext
}

func NewDocAnnoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DocAnnoContext {
	var p = new(DocAnnoContext)

	p.AnnonContext = NewEmptyAnnonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AnnonContext))

	return p
}

func (s *DocAnnoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocAnnoContext) LINE_DOC() antlr.TerminalNode {
	return s.GetToken(ADLParserLINE_DOC, 0)
}

func (s *DocAnnoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.EnterDocAnno(s)
	}
}

func (s *DocAnnoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.ExitDocAnno(s)
	}
}

type LocalAnnoContext struct {
	*AnnonContext
}

func NewLocalAnnoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LocalAnnoContext {
	var p = new(LocalAnnoContext)

	p.AnnonContext = NewEmptyAnnonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AnnonContext))

	return p
}

func (s *LocalAnnoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalAnnoContext) AT() antlr.TerminalNode {
	return s.GetToken(ADLParserAT, 0)
}

func (s *LocalAnnoContext) ID() antlr.TerminalNode {
	return s.GetToken(ADLParserID, 0)
}

func (s *LocalAnnoContext) JsonValue() IJsonValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValueContext)
}

func (s *LocalAnnoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.EnterLocalAnno(s)
	}
}

func (s *LocalAnnoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.ExitLocalAnno(s)
	}
}

func (p *ADLParser) Annon() (localctx IAnnonContext) {
	localctx = NewAnnonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ADLParserRULE_annon)

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

	p.SetState(72)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ADLParserAT:
		localctx = NewLocalAnnoContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(68)
			p.Match(ADLParserAT)
		}
		{
			p.SetState(69)
			p.Match(ADLParserID)
		}
		{
			p.SetState(70)
			p.JsonValue()
		}

	case ADLParserLINE_DOC:
		localctx = NewDocAnnoContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.Match(ADLParserLINE_DOC)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *StructOrUnionContext) AllAnnon() []IAnnonContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnonContext)(nil)).Elem())
	var tst = make([]IAnnonContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnonContext)
		}
	}

	return tst
}

func (s *StructOrUnionContext) Annon(i int) IAnnonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnonContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnonContext)
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
	c       antlr.Token
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

func (s *FieldAnnotationContext) GetC() antlr.Token { return s.c }

func (s *FieldAnnotationContext) SetKw_anno(v antlr.Token) { s.kw_anno = v }

func (s *FieldAnnotationContext) SetA(v antlr.Token) { s.a = v }

func (s *FieldAnnotationContext) SetB(v antlr.Token) { s.b = v }

func (s *FieldAnnotationContext) SetC(v antlr.Token) { s.c = v }

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

func (s *TypeOrNewtypeContext) AllEQ() []antlr.TerminalNode {
	return s.GetTokens(ADLParserEQ)
}

func (s *TypeOrNewtypeContext) EQ(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserEQ, i)
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

func (s *TypeOrNewtypeContext) AllAnnon() []IAnnonContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnonContext)(nil)).Elem())
	var tst = make([]IAnnonContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnonContext)
		}
	}

	return tst
}

func (s *TypeOrNewtypeContext) Annon(i int) IAnnonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnonContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnonContext)
}

func (s *TypeOrNewtypeContext) TypeParam() ITypeParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeParamContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeParamContext)
}

func (s *TypeOrNewtypeContext) TypeExpr() ITypeExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *TypeOrNewtypeContext) JsonValue() IJsonValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValueContext)
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
	p.EnterRule(localctx, 8, ADLParserRULE_top_level_statement)
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

	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStructOrUnionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLParserAT || _la == ADLParserLINE_DOC {
			{
				p.SetState(74)
				p.Annon()
			}

			p.SetState(79)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(80)

			var _m = p.Match(ADLParserID)

			localctx.(*StructOrUnionContext).kw_soru = _m
		}
		{
			p.SetState(81)

			var _m = p.Match(ADLParserID)

			localctx.(*StructOrUnionContext).a = _m
		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLParserLCHEVR {
			{
				p.SetState(82)
				p.TypeParam()
			}

		}
		{
			p.SetState(85)
			p.Match(ADLParserLCUR)
		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ADLParserAT)|(1<<ADLParserID)|(1<<ADLParserLINE_DOC))) != 0 {
			{
				p.SetState(86)
				p.SoruBody()
			}

			p.SetState(91)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(92)
			p.Match(ADLParserRCUR)
		}
		{
			p.SetState(93)
			p.Match(ADLParserSEMI)
		}

	case 2:
		localctx = NewTypeOrNewtypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLParserAT || _la == ADLParserLINE_DOC {
			{
				p.SetState(94)
				p.Annon()
			}

			p.SetState(99)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(100)

			var _m = p.Match(ADLParserID)

			localctx.(*TypeOrNewtypeContext).kw_tnew = _m
		}
		{
			p.SetState(101)

			var _m = p.Match(ADLParserID)

			localctx.(*TypeOrNewtypeContext).a = _m
		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLParserLCHEVR {
			{
				p.SetState(102)
				p.TypeParam()
			}

		}
		{
			p.SetState(105)
			p.Match(ADLParserEQ)
		}
		{
			p.SetState(106)

			var _m = p.Match(ADLParserID)

			localctx.(*TypeOrNewtypeContext).b = _m
		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLParserLCHEVR {
			{
				p.SetState(107)
				p.TypeExpr()
			}

		}
		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLParserEQ {
			{
				p.SetState(110)
				p.Match(ADLParserEQ)
			}
			{
				p.SetState(111)
				p.JsonValue()
			}

		}
		{
			p.SetState(114)
			p.Match(ADLParserSEMI)
		}

	case 3:
		localctx = NewModuleAnnotationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(115)

			var _m = p.Match(ADLParserID)

			localctx.(*ModuleAnnotationContext).kw_anno = _m
		}
		{
			p.SetState(116)

			var _m = p.Match(ADLParserID)

			localctx.(*ModuleAnnotationContext).a = _m
		}
		{
			p.SetState(117)
			p.JsonValue()
		}
		{
			p.SetState(118)
			p.Match(ADLParserSEMI)
		}

	case 4:
		localctx = NewDeclAnnotationContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(120)

			var _m = p.Match(ADLParserID)

			localctx.(*DeclAnnotationContext).kw_anno = _m
		}
		{
			p.SetState(121)

			var _m = p.Match(ADLParserID)

			localctx.(*DeclAnnotationContext).a = _m
		}
		{
			p.SetState(122)

			var _m = p.Match(ADLParserID)

			localctx.(*DeclAnnotationContext).b = _m
		}
		{
			p.SetState(123)
			p.JsonValue()
		}
		{
			p.SetState(124)
			p.Match(ADLParserSEMI)
		}

	case 5:
		localctx = NewFieldAnnotationContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(126)

			var _m = p.Match(ADLParserID)

			localctx.(*FieldAnnotationContext).kw_anno = _m
		}
		{
			p.SetState(127)

			var _m = p.Match(ADLParserID)

			localctx.(*FieldAnnotationContext).a = _m
		}
		{
			p.SetState(128)
			p.Match(ADLParserDCOLON)
		}
		{
			p.SetState(129)

			var _m = p.Match(ADLParserID)

			localctx.(*FieldAnnotationContext).b = _m
		}
		{
			p.SetState(130)

			var _m = p.Match(ADLParserID)

			localctx.(*FieldAnnotationContext).c = _m
		}
		{
			p.SetState(131)
			p.JsonValue()
		}
		{
			p.SetState(132)
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

	// IsTypeParamContext differentiates from other interfaces.
	IsTypeParamContext()
}

type TypeParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *TypeParamContext) CopyFrom(ctx *TypeParamContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypeParameterContext struct {
	*TypeParamContext
	_ID   antlr.Token
	typep []antlr.Token
}

func NewTypeParameterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeParameterContext {
	var p = new(TypeParameterContext)

	p.TypeParamContext = NewEmptyTypeParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeParamContext))

	return p
}

func (s *TypeParameterContext) Get_ID() antlr.Token { return s._ID }

func (s *TypeParameterContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *TypeParameterContext) GetTypep() []antlr.Token { return s.typep }

func (s *TypeParameterContext) SetTypep(v []antlr.Token) { s.typep = v }

func (s *TypeParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeParameterContext) LCHEVR() antlr.TerminalNode {
	return s.GetToken(ADLParserLCHEVR, 0)
}

func (s *TypeParameterContext) RCHEVR() antlr.TerminalNode {
	return s.GetToken(ADLParserRCHEVR, 0)
}

func (s *TypeParameterContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ADLParserID)
}

func (s *TypeParameterContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserID, i)
}

func (s *TypeParameterContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ADLParserCOMMA)
}

func (s *TypeParameterContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserCOMMA, i)
}

func (s *TypeParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.EnterTypeParameter(s)
	}
}

func (s *TypeParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.ExitTypeParameter(s)
	}
}

func (p *ADLParser) TypeParam() (localctx ITypeParamContext) {
	localctx = NewTypeParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ADLParserRULE_typeParam)
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

	localctx = NewTypeParameterContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(ADLParserLCHEVR)
	}
	{
		p.SetState(137)

		var _m = p.Match(ADLParserID)

		localctx.(*TypeParameterContext)._ID = _m
	}
	localctx.(*TypeParameterContext).typep = append(localctx.(*TypeParameterContext).typep, localctx.(*TypeParameterContext)._ID)
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ADLParserCOMMA {
		{
			p.SetState(138)
			p.Match(ADLParserCOMMA)
		}
		{
			p.SetState(139)

			var _m = p.Match(ADLParserID)

			localctx.(*TypeParameterContext)._ID = _m
		}
		localctx.(*TypeParameterContext).typep = append(localctx.(*TypeParameterContext).typep, localctx.(*TypeParameterContext)._ID)

		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(145)
		p.Match(ADLParserRCHEVR)
	}

	return localctx
}

// ITypeExprContext is an interface to support dynamic dispatch.
type ITypeExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeExprContext differentiates from other interfaces.
	IsTypeExprContext()
}

type TypeExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeExprContext() *TypeExprContext {
	var p = new(TypeExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_typeExpr
	return p
}

func (*TypeExprContext) IsTypeExprContext() {}

func NewTypeExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeExprContext {
	var p = new(TypeExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_typeExpr

	return p
}

func (s *TypeExprContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeExprContext) CopyFrom(ctx *TypeExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypeExpressionContext struct {
	*TypeExprContext
	_typeExprElem ITypeExprElemContext
	typep         []ITypeExprElemContext
}

func NewTypeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeExpressionContext {
	var p = new(TypeExpressionContext)

	p.TypeExprContext = NewEmptyTypeExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeExprContext))

	return p
}

func (s *TypeExpressionContext) Get_typeExprElem() ITypeExprElemContext { return s._typeExprElem }

func (s *TypeExpressionContext) Set_typeExprElem(v ITypeExprElemContext) { s._typeExprElem = v }

func (s *TypeExpressionContext) GetTypep() []ITypeExprElemContext { return s.typep }

func (s *TypeExpressionContext) SetTypep(v []ITypeExprElemContext) { s.typep = v }

func (s *TypeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExpressionContext) LCHEVR() antlr.TerminalNode {
	return s.GetToken(ADLParserLCHEVR, 0)
}

func (s *TypeExpressionContext) RCHEVR() antlr.TerminalNode {
	return s.GetToken(ADLParserRCHEVR, 0)
}

func (s *TypeExpressionContext) AllTypeExprElem() []ITypeExprElemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeExprElemContext)(nil)).Elem())
	var tst = make([]ITypeExprElemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeExprElemContext)
		}
	}

	return tst
}

func (s *TypeExpressionContext) TypeExprElem(i int) ITypeExprElemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeExprElemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeExprElemContext)
}

func (s *TypeExpressionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ADLParserCOMMA)
}

func (s *TypeExpressionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserCOMMA, i)
}

func (s *TypeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.EnterTypeExpression(s)
	}
}

func (s *TypeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.ExitTypeExpression(s)
	}
}

func (p *ADLParser) TypeExpr() (localctx ITypeExprContext) {
	localctx = NewTypeExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ADLParserRULE_typeExpr)
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

	localctx = NewTypeExpressionContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.Match(ADLParserLCHEVR)
	}
	{
		p.SetState(148)

		var _x = p.TypeExprElem()

		localctx.(*TypeExpressionContext)._typeExprElem = _x
	}
	localctx.(*TypeExpressionContext).typep = append(localctx.(*TypeExpressionContext).typep, localctx.(*TypeExpressionContext)._typeExprElem)
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ADLParserCOMMA {
		{
			p.SetState(149)
			p.Match(ADLParserCOMMA)
		}
		{
			p.SetState(150)

			var _x = p.TypeExprElem()

			localctx.(*TypeExpressionContext)._typeExprElem = _x
		}
		localctx.(*TypeExpressionContext).typep = append(localctx.(*TypeExpressionContext).typep, localctx.(*TypeExpressionContext)._typeExprElem)

		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(156)
		p.Match(ADLParserRCHEVR)
	}

	return localctx
}

// ITypeExprElemContext is an interface to support dynamic dispatch.
type ITypeExprElemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeExprElemContext differentiates from other interfaces.
	IsTypeExprElemContext()
}

type TypeExprElemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeExprElemContext() *TypeExprElemContext {
	var p = new(TypeExprElemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_typeExprElem
	return p
}

func (*TypeExprElemContext) IsTypeExprElemContext() {}

func NewTypeExprElemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeExprElemContext {
	var p = new(TypeExprElemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_typeExprElem

	return p
}

func (s *TypeExprElemContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeExprElemContext) CopyFrom(ctx *TypeExprElemContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeExprElemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExprElemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypeExpressionElemContext struct {
	*TypeExprElemContext
}

func NewTypeExpressionElemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeExpressionElemContext {
	var p = new(TypeExpressionElemContext)

	p.TypeExprElemContext = NewEmptyTypeExprElemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeExprElemContext))

	return p
}

func (s *TypeExpressionElemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExpressionElemContext) ID() antlr.TerminalNode {
	return s.GetToken(ADLParserID, 0)
}

func (s *TypeExpressionElemContext) TypeExpr() ITypeExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *TypeExpressionElemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.EnterTypeExpressionElem(s)
	}
}

func (s *TypeExpressionElemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.ExitTypeExpressionElem(s)
	}
}

func (p *ADLParser) TypeExprElem() (localctx ITypeExprElemContext) {
	localctx = NewTypeExprElemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ADLParserRULE_typeExprElem)
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

	localctx = NewTypeExpressionElemContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.Match(ADLParserID)
	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ADLParserLCHEVR {
		{
			p.SetState(159)
			p.TypeExpr()
		}

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

func (s *SoruBodyContext) CopyFrom(ctx *SoruBodyContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SoruBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SoruBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FieldStatementContext struct {
	*SoruBodyContext
	a antlr.Token
	b antlr.Token
}

func NewFieldStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldStatementContext {
	var p = new(FieldStatementContext)

	p.SoruBodyContext = NewEmptySoruBodyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SoruBodyContext))

	return p
}

func (s *FieldStatementContext) GetA() antlr.Token { return s.a }

func (s *FieldStatementContext) GetB() antlr.Token { return s.b }

func (s *FieldStatementContext) SetA(v antlr.Token) { s.a = v }

func (s *FieldStatementContext) SetB(v antlr.Token) { s.b = v }

func (s *FieldStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ADLParserSEMI, 0)
}

func (s *FieldStatementContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ADLParserID)
}

func (s *FieldStatementContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserID, i)
}

func (s *FieldStatementContext) AllAnnon() []IAnnonContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnonContext)(nil)).Elem())
	var tst = make([]IAnnonContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnonContext)
		}
	}

	return tst
}

func (s *FieldStatementContext) Annon(i int) IAnnonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnonContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnonContext)
}

func (s *FieldStatementContext) TypeExpr() ITypeExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *FieldStatementContext) EQ() antlr.TerminalNode {
	return s.GetToken(ADLParserEQ, 0)
}

func (s *FieldStatementContext) JsonValue() IJsonValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValueContext)
}

func (s *FieldStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.EnterFieldStatement(s)
	}
}

func (s *FieldStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLParserListener); ok {
		listenerT.ExitFieldStatement(s)
	}
}

func (p *ADLParser) SoruBody() (localctx ISoruBodyContext) {
	localctx = NewSoruBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ADLParserRULE_soruBody)
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

	localctx = NewFieldStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ADLParserAT || _la == ADLParserLINE_DOC {
		{
			p.SetState(162)
			p.Annon()
		}

		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(168)

		var _m = p.Match(ADLParserID)

		localctx.(*FieldStatementContext).a = _m
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ADLParserLCHEVR {
		{
			p.SetState(169)
			p.TypeExpr()
		}

	}
	{
		p.SetState(172)

		var _m = p.Match(ADLParserID)

		localctx.(*FieldStatementContext).b = _m
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ADLParserEQ {
		{
			p.SetState(173)
			p.Match(ADLParserEQ)
		}
		{
			p.SetState(174)
			p.JsonValue()
		}

	}
	{
		p.SetState(177)
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

func (s *ObjStatementContext) RCUR() antlr.TerminalNode {
	return s.GetToken(ADLParserRCUR, 0)
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

func (s *ArrayStatementContext) RSQ() antlr.TerminalNode {
	return s.GetToken(ADLParserRSQ, 0)
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
	p.EnterRule(localctx, 18, ADLParserRULE_jsonValue)
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

	p.SetState(211)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ADLParserSTR:
		localctx = NewStringStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(179)

			var _m = p.Match(ADLParserSTR)

			localctx.(*StringStatementContext).s = _m
		}

	case ADLParserID:
		localctx = NewTrueFalseNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(180)

			var _m = p.Match(ADLParserID)

			localctx.(*TrueFalseNullContext).kw_tfn = _m
		}

	case ADLParserINT:
		localctx = NewNumberStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(181)
			p.Match(ADLParserINT)
		}

	case ADLParserFLT:
		localctx = NewFloatStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(182)
			p.Match(ADLParserFLT)
		}

	case ADLParserLSQ:
		localctx = NewArrayStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(183)
			p.Match(ADLParserLSQ)
		}
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ADLParserLCUR)|(1<<ADLParserLSQ)|(1<<ADLParserSTR)|(1<<ADLParserID)|(1<<ADLParserINT)|(1<<ADLParserFLT))) != 0 {
			{
				p.SetState(184)
				p.JsonValue()
			}
			p.SetState(189)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLParserCOMMA {
				{
					p.SetState(185)
					p.Match(ADLParserCOMMA)
				}
				{
					p.SetState(186)
					p.JsonValue()
				}

				p.SetState(191)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(194)
			p.Match(ADLParserRSQ)
		}

	case ADLParserLCUR:
		localctx = NewObjStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(195)
			p.Match(ADLParserLCUR)
		}
		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLParserSTR {
			{
				p.SetState(196)
				p.Match(ADLParserSTR)
			}
			{
				p.SetState(197)
				p.Match(ADLParserCOLON)
			}
			{
				p.SetState(198)
				p.JsonValue()
			}
			p.SetState(205)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLParserCOMMA {
				{
					p.SetState(199)
					p.Match(ADLParserCOMMA)
				}
				{
					p.SetState(200)
					p.Match(ADLParserSTR)
				}
				{
					p.SetState(201)
					p.Match(ADLParserCOLON)
				}
				{
					p.SetState(202)
					p.JsonValue()
				}

				p.SetState(207)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(210)
			p.Match(ADLParserRCUR)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
