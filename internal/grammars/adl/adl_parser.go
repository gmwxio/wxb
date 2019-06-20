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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 51, 217,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 2, 3,
	3, 7, 3, 27, 10, 3, 12, 3, 14, 3, 30, 11, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7,
	3, 36, 10, 3, 12, 3, 14, 3, 39, 11, 3, 3, 3, 3, 3, 7, 3, 43, 10, 3, 12,
	3, 14, 3, 46, 11, 3, 3, 3, 7, 3, 49, 10, 3, 12, 3, 14, 3, 52, 11, 3, 3,
	3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 61, 10, 4, 12, 4, 14, 4, 64,
	11, 4, 3, 4, 3, 4, 5, 4, 68, 10, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5,
	5, 5, 76, 10, 5, 3, 6, 7, 6, 79, 10, 6, 12, 6, 14, 6, 82, 11, 6, 3, 6,
	3, 6, 3, 6, 5, 6, 87, 10, 6, 3, 6, 3, 6, 7, 6, 91, 10, 6, 12, 6, 14, 6,
	94, 11, 6, 3, 6, 3, 6, 3, 6, 7, 6, 99, 10, 6, 12, 6, 14, 6, 102, 11, 6,
	3, 6, 3, 6, 3, 6, 5, 6, 107, 10, 6, 3, 6, 3, 6, 3, 6, 5, 6, 112, 10, 6,
	3, 6, 3, 6, 5, 6, 116, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 5, 6, 138, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 144, 10, 7, 12, 7,
	14, 7, 147, 11, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 155, 10, 8,
	12, 8, 14, 8, 158, 11, 8, 3, 8, 3, 8, 3, 9, 3, 9, 5, 9, 164, 10, 9, 3,
	10, 7, 10, 167, 10, 10, 12, 10, 14, 10, 170, 11, 10, 3, 10, 3, 10, 5, 10,
	174, 10, 10, 3, 10, 3, 10, 3, 10, 5, 10, 179, 10, 10, 3, 10, 3, 10, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 191, 10, 11,
	12, 11, 14, 11, 194, 11, 11, 5, 11, 196, 10, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 207, 10, 11, 12, 11, 14,
	11, 210, 11, 11, 5, 11, 212, 10, 11, 3, 11, 5, 11, 215, 10, 11, 3, 11,
	2, 2, 12, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 2, 2, 2, 239, 2, 22, 3, 2,
	2, 2, 4, 28, 3, 2, 2, 2, 6, 56, 3, 2, 2, 2, 8, 75, 3, 2, 2, 2, 10, 137,
	3, 2, 2, 2, 12, 139, 3, 2, 2, 2, 14, 150, 3, 2, 2, 2, 16, 161, 3, 2, 2,
	2, 18, 168, 3, 2, 2, 2, 20, 214, 3, 2, 2, 2, 22, 23, 5, 4, 3, 2, 23, 24,
	7, 2, 2, 3, 24, 3, 3, 2, 2, 2, 25, 27, 5, 8, 5, 2, 26, 25, 3, 2, 2, 2,
	27, 30, 3, 2, 2, 2, 28, 26, 3, 2, 2, 2, 28, 29, 3, 2, 2, 2, 29, 31, 3,
	2, 2, 2, 30, 28, 3, 2, 2, 2, 31, 32, 7, 20, 2, 2, 32, 37, 7, 20, 2, 2,
	33, 34, 7, 13, 2, 2, 34, 36, 7, 20, 2, 2, 35, 33, 3, 2, 2, 2, 36, 39, 3,
	2, 2, 2, 37, 35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 40, 3, 2, 2, 2, 39,
	37, 3, 2, 2, 2, 40, 44, 7, 3, 2, 2, 41, 43, 5, 6, 4, 2, 42, 41, 3, 2, 2,
	2, 43, 46, 3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 50,
	3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 47, 49, 5, 10, 6, 2, 48, 47, 3, 2, 2, 2,
	49, 52, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 53, 3,
	2, 2, 2, 52, 50, 3, 2, 2, 2, 53, 54, 7, 4, 2, 2, 54, 55, 7, 10, 2, 2, 55,
	5, 3, 2, 2, 2, 56, 57, 7, 20, 2, 2, 57, 62, 7, 20, 2, 2, 58, 59, 7, 13,
	2, 2, 59, 61, 7, 20, 2, 2, 60, 58, 3, 2, 2, 2, 61, 64, 3, 2, 2, 2, 62,
	60, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 67, 3, 2, 2, 2, 64, 62, 3, 2, 2,
	2, 65, 66, 7, 13, 2, 2, 66, 68, 7, 17, 2, 2, 67, 65, 3, 2, 2, 2, 67, 68,
	3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 70, 7, 10, 2, 2, 70, 7, 3, 2, 2, 2,
	71, 72, 7, 18, 2, 2, 72, 73, 7, 20, 2, 2, 73, 76, 5, 20, 11, 2, 74, 76,
	7, 24, 2, 2, 75, 71, 3, 2, 2, 2, 75, 74, 3, 2, 2, 2, 76, 9, 3, 2, 2, 2,
	77, 79, 5, 8, 5, 2, 78, 77, 3, 2, 2, 2, 79, 82, 3, 2, 2, 2, 80, 78, 3,
	2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 83, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 83,
	84, 7, 20, 2, 2, 84, 86, 7, 20, 2, 2, 85, 87, 5, 12, 7, 2, 86, 85, 3, 2,
	2, 2, 86, 87, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 92, 7, 3, 2, 2, 89, 91,
	5, 18, 10, 2, 90, 89, 3, 2, 2, 2, 91, 94, 3, 2, 2, 2, 92, 90, 3, 2, 2,
	2, 92, 93, 3, 2, 2, 2, 93, 95, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 95, 96,
	7, 4, 2, 2, 96, 138, 7, 10, 2, 2, 97, 99, 5, 8, 5, 2, 98, 97, 3, 2, 2,
	2, 99, 102, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101,
	103, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 103, 104, 7, 20, 2, 2, 104, 106,
	7, 20, 2, 2, 105, 107, 5, 12, 7, 2, 106, 105, 3, 2, 2, 2, 106, 107, 3,
	2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 109, 7, 7, 2, 2, 109, 111, 7, 20, 2,
	2, 110, 112, 5, 14, 8, 2, 111, 110, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112,
	115, 3, 2, 2, 2, 113, 114, 7, 7, 2, 2, 114, 116, 5, 20, 11, 2, 115, 113,
	3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 138, 7, 10,
	2, 2, 118, 119, 7, 20, 2, 2, 119, 120, 7, 20, 2, 2, 120, 121, 5, 20, 11,
	2, 121, 122, 7, 10, 2, 2, 122, 138, 3, 2, 2, 2, 123, 124, 7, 20, 2, 2,
	124, 125, 7, 20, 2, 2, 125, 126, 7, 20, 2, 2, 126, 127, 5, 20, 11, 2, 127,
	128, 7, 10, 2, 2, 128, 138, 3, 2, 2, 2, 129, 130, 7, 20, 2, 2, 130, 131,
	7, 20, 2, 2, 131, 132, 7, 11, 2, 2, 132, 133, 7, 20, 2, 2, 133, 134, 7,
	20, 2, 2, 134, 135, 5, 20, 11, 2, 135, 136, 7, 10, 2, 2, 136, 138, 3, 2,
	2, 2, 137, 80, 3, 2, 2, 2, 137, 100, 3, 2, 2, 2, 137, 118, 3, 2, 2, 2,
	137, 123, 3, 2, 2, 2, 137, 129, 3, 2, 2, 2, 138, 11, 3, 2, 2, 2, 139, 140,
	7, 15, 2, 2, 140, 145, 7, 20, 2, 2, 141, 142, 7, 14, 2, 2, 142, 144, 7,
	20, 2, 2, 143, 141, 3, 2, 2, 2, 144, 147, 3, 2, 2, 2, 145, 143, 3, 2, 2,
	2, 145, 146, 3, 2, 2, 2, 146, 148, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 148,
	149, 7, 16, 2, 2, 149, 13, 3, 2, 2, 2, 150, 151, 7, 15, 2, 2, 151, 156,
	5, 16, 9, 2, 152, 153, 7, 14, 2, 2, 153, 155, 5, 16, 9, 2, 154, 152, 3,
	2, 2, 2, 155, 158, 3, 2, 2, 2, 156, 154, 3, 2, 2, 2, 156, 157, 3, 2, 2,
	2, 157, 159, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 159, 160, 7, 16, 2, 2, 160,
	15, 3, 2, 2, 2, 161, 163, 7, 20, 2, 2, 162, 164, 5, 14, 8, 2, 163, 162,
	3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 17, 3, 2, 2, 2, 165, 167, 5, 8,
	5, 2, 166, 165, 3, 2, 2, 2, 167, 170, 3, 2, 2, 2, 168, 166, 3, 2, 2, 2,
	168, 169, 3, 2, 2, 2, 169, 171, 3, 2, 2, 2, 170, 168, 3, 2, 2, 2, 171,
	173, 7, 20, 2, 2, 172, 174, 5, 14, 8, 2, 173, 172, 3, 2, 2, 2, 173, 174,
	3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 178, 7, 20, 2, 2, 176, 177, 7, 7,
	2, 2, 177, 179, 5, 20, 11, 2, 178, 176, 3, 2, 2, 2, 178, 179, 3, 2, 2,
	2, 179, 180, 3, 2, 2, 2, 180, 181, 7, 10, 2, 2, 181, 19, 3, 2, 2, 2, 182,
	215, 7, 19, 2, 2, 183, 215, 7, 20, 2, 2, 184, 215, 7, 21, 2, 2, 185, 215,
	7, 22, 2, 2, 186, 195, 7, 5, 2, 2, 187, 192, 5, 20, 11, 2, 188, 189, 7,
	14, 2, 2, 189, 191, 5, 20, 11, 2, 190, 188, 3, 2, 2, 2, 191, 194, 3, 2,
	2, 2, 192, 190, 3, 2, 2, 2, 192, 193, 3, 2, 2, 2, 193, 196, 3, 2, 2, 2,
	194, 192, 3, 2, 2, 2, 195, 187, 3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196,
	197, 3, 2, 2, 2, 197, 215, 7, 6, 2, 2, 198, 211, 7, 3, 2, 2, 199, 200,
	7, 19, 2, 2, 200, 201, 7, 12, 2, 2, 201, 208, 5, 20, 11, 2, 202, 203, 7,
	14, 2, 2, 203, 204, 7, 19, 2, 2, 204, 205, 7, 12, 2, 2, 205, 207, 5, 20,
	11, 2, 206, 202, 3, 2, 2, 2, 207, 210, 3, 2, 2, 2, 208, 206, 3, 2, 2, 2,
	208, 209, 3, 2, 2, 2, 209, 212, 3, 2, 2, 2, 210, 208, 3, 2, 2, 2, 211,
	199, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 215,
	7, 4, 2, 2, 214, 182, 3, 2, 2, 2, 214, 183, 3, 2, 2, 2, 214, 184, 3, 2,
	2, 2, 214, 185, 3, 2, 2, 2, 214, 186, 3, 2, 2, 2, 214, 198, 3, 2, 2, 2,
	215, 21, 3, 2, 2, 2, 28, 28, 37, 44, 50, 62, 67, 75, 80, 86, 92, 100, 106,
	111, 115, 137, 145, 156, 163, 168, 173, 178, 192, 195, 208, 211, 214,
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
	"WS", "LINE_DOC", "LINE_COMMENT", "DOWN", "UP", "ROOT", "ERROR", "ADL",
	"Module", "Import", "Annotation", "Struct", "Union", "Newtype", "Type",
	"TypeParam", "TypeExpr", "Field", "Json", "JsonStr", "JsonBool", "JsonNull",
	"JsonInt", "JsonFloat", "JsonArray", "JsonObj", "ModuleAnno", "DeclAnno",
	"FieldAnno",
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
	ADLParserADL          = 28
	ADLParserModule       = 29
	ADLParserImport       = 30
	ADLParserAnnotation   = 31
	ADLParserStruct       = 32
	ADLParserUnion        = 33
	ADLParserNewtype      = 34
	ADLParserType         = 35
	ADLParserTypeParam    = 36
	ADLParserTypeExpr     = 37
	ADLParserField        = 38
	ADLParserJson         = 39
	ADLParserJsonStr      = 40
	ADLParserJsonBool     = 41
	ADLParserJsonNull     = 42
	ADLParserJsonInt      = 43
	ADLParserJsonFloat    = 44
	ADLParserJsonArray    = 45
	ADLParserJsonObj      = 46
	ADLParserModuleAnno   = 47
	ADLParserDeclAnno     = 48
	ADLParserFieldAnno    = 49
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

func (s *AdlContext) EOF() antlr.TerminalNode {
	return s.GetToken(ADLParserEOF, 0)
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
		p.SetState(20)
		p.Module()
	}
	{
		p.SetState(21)
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
	kw   antlr.Token
	_ID  antlr.Token
	name []antlr.Token
}

func NewModuleStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModuleStatementContext {
	var p = new(ModuleStatementContext)

	p.ModuleContext = NewEmptyModuleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ModuleContext))

	return p
}

func (s *ModuleStatementContext) GetKw() antlr.Token { return s.kw }

func (s *ModuleStatementContext) Get_ID() antlr.Token { return s._ID }

func (s *ModuleStatementContext) SetKw(v antlr.Token) { s.kw = v }

func (s *ModuleStatementContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ModuleStatementContext) GetName() []antlr.Token { return s.name }

func (s *ModuleStatementContext) SetName(v []antlr.Token) { s.name = v }

func (s *ModuleStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleStatementContext) LCUR() antlr.TerminalNode {
	return s.GetToken(ADLParserLCUR, 0)
}

func (s *ModuleStatementContext) RCUR() antlr.TerminalNode {
	return s.GetToken(ADLParserRCUR, 0)
}

func (s *ModuleStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ADLParserSEMI, 0)
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

func (s *ModuleStatementContext) AllImports() []IImportsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportsContext)(nil)).Elem())
	var tst = make([]IImportsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportsContext)
		}
	}

	return tst
}

func (s *ModuleStatementContext) Imports(i int) IImportsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportsContext)
}

func (s *ModuleStatementContext) AllTop_level_statement() []ITop_level_statementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITop_level_statementContext)(nil)).Elem())
	var tst = make([]ITop_level_statementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITop_level_statementContext)
		}
	}

	return tst
}

func (s *ModuleStatementContext) Top_level_statement(i int) ITop_level_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITop_level_statementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITop_level_statementContext)
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

	var _alt int

	localctx = NewModuleStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ADLParserAT || _la == ADLParserLINE_DOC {
		{
			p.SetState(23)
			p.Annon()
		}

		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(29)

		var _m = p.Match(ADLParserID)

		localctx.(*ModuleStatementContext).kw = _m
	}
	{
		p.SetState(30)

		var _m = p.Match(ADLParserID)

		localctx.(*ModuleStatementContext)._ID = _m
	}
	localctx.(*ModuleStatementContext).name = append(localctx.(*ModuleStatementContext).name, localctx.(*ModuleStatementContext)._ID)
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ADLParserDOT {
		{
			p.SetState(31)
			p.Match(ADLParserDOT)
		}
		{
			p.SetState(32)

			var _m = p.Match(ADLParserID)

			localctx.(*ModuleStatementContext)._ID = _m
		}
		localctx.(*ModuleStatementContext).name = append(localctx.(*ModuleStatementContext).name, localctx.(*ModuleStatementContext)._ID)

		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(38)
		p.Match(ADLParserLCUR)
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(39)
				p.Imports()
			}

		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ADLParserAT)|(1<<ADLParserID)|(1<<ADLParserLINE_DOC))) != 0 {
		{
			p.SetState(45)
			p.Top_level_statement()
		}

		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(51)
		p.Match(ADLParserRCUR)
	}
	{
		p.SetState(52)
		p.Match(ADLParserSEMI)
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
	kw  antlr.Token
	_ID antlr.Token
	a   []antlr.Token
	s   antlr.Token
}

func NewImportStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImportStatementContext {
	var p = new(ImportStatementContext)

	p.ImportsContext = NewEmptyImportsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ImportsContext))

	return p
}

func (s *ImportStatementContext) GetKw() antlr.Token { return s.kw }

func (s *ImportStatementContext) Get_ID() antlr.Token { return s._ID }

func (s *ImportStatementContext) GetS() antlr.Token { return s.s }

func (s *ImportStatementContext) SetKw(v antlr.Token) { s.kw = v }

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
		p.SetState(54)

		var _m = p.Match(ADLParserID)

		localctx.(*ImportStatementContext).kw = _m
	}
	{
		p.SetState(55)

		var _m = p.Match(ADLParserID)

		localctx.(*ImportStatementContext)._ID = _m
	}
	localctx.(*ImportStatementContext).a = append(localctx.(*ImportStatementContext).a, localctx.(*ImportStatementContext)._ID)
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(56)
				p.Match(ADLParserDOT)
			}
			{
				p.SetState(57)

				var _m = p.Match(ADLParserID)

				localctx.(*ImportStatementContext)._ID = _m
			}
			localctx.(*ImportStatementContext).a = append(localctx.(*ImportStatementContext).a, localctx.(*ImportStatementContext)._ID)

		}
		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ADLParserDOT {
		{
			p.SetState(63)
			p.Match(ADLParserDOT)
		}
		{
			p.SetState(64)

			var _m = p.Match(ADLParserSTAR)

			localctx.(*ImportStatementContext).s = _m
		}

	}
	{
		p.SetState(67)
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
	a antlr.Token
}

func NewDocAnnoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DocAnnoContext {
	var p = new(DocAnnoContext)

	p.AnnonContext = NewEmptyAnnonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AnnonContext))

	return p
}

func (s *DocAnnoContext) GetA() antlr.Token { return s.a }

func (s *DocAnnoContext) SetA(v antlr.Token) { s.a = v }

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
	a antlr.Token
}

func NewLocalAnnoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LocalAnnoContext {
	var p = new(LocalAnnoContext)

	p.AnnonContext = NewEmptyAnnonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AnnonContext))

	return p
}

func (s *LocalAnnoContext) GetA() antlr.Token { return s.a }

func (s *LocalAnnoContext) SetA(v antlr.Token) { s.a = v }

func (s *LocalAnnoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalAnnoContext) AT() antlr.TerminalNode {
	return s.GetToken(ADLParserAT, 0)
}

func (s *LocalAnnoContext) JsonValue() IJsonValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValueContext)
}

func (s *LocalAnnoContext) ID() antlr.TerminalNode {
	return s.GetToken(ADLParserID, 0)
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

	p.SetState(73)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ADLParserAT:
		localctx = NewLocalAnnoContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(69)
			p.Match(ADLParserAT)
		}
		{
			p.SetState(70)

			var _m = p.Match(ADLParserID)

			localctx.(*LocalAnnoContext).a = _m
		}
		{
			p.SetState(71)
			p.JsonValue()
		}

	case ADLParserLINE_DOC:
		localctx = NewDocAnnoContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(72)

			var _m = p.Match(ADLParserLINE_DOC)

			localctx.(*DocAnnoContext).a = _m
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
	kw antlr.Token
	a  antlr.Token
}

func NewStructOrUnionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructOrUnionContext {
	var p = new(StructOrUnionContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

func (s *StructOrUnionContext) GetKw() antlr.Token { return s.kw }

func (s *StructOrUnionContext) GetA() antlr.Token { return s.a }

func (s *StructOrUnionContext) SetKw(v antlr.Token) { s.kw = v }

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
	kw antlr.Token
	a  antlr.Token
	b  antlr.Token
}

func NewDeclAnnotationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclAnnotationContext {
	var p = new(DeclAnnotationContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

func (s *DeclAnnotationContext) GetKw() antlr.Token { return s.kw }

func (s *DeclAnnotationContext) GetA() antlr.Token { return s.a }

func (s *DeclAnnotationContext) GetB() antlr.Token { return s.b }

func (s *DeclAnnotationContext) SetKw(v antlr.Token) { s.kw = v }

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
	kw antlr.Token
	a  antlr.Token
	b  antlr.Token
	c  antlr.Token
}

func NewFieldAnnotationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldAnnotationContext {
	var p = new(FieldAnnotationContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

func (s *FieldAnnotationContext) GetKw() antlr.Token { return s.kw }

func (s *FieldAnnotationContext) GetA() antlr.Token { return s.a }

func (s *FieldAnnotationContext) GetB() antlr.Token { return s.b }

func (s *FieldAnnotationContext) GetC() antlr.Token { return s.c }

func (s *FieldAnnotationContext) SetKw(v antlr.Token) { s.kw = v }

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
	kw antlr.Token
	a  antlr.Token
	b  antlr.Token
}

func NewTypeOrNewtypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeOrNewtypeContext {
	var p = new(TypeOrNewtypeContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

func (s *TypeOrNewtypeContext) GetKw() antlr.Token { return s.kw }

func (s *TypeOrNewtypeContext) GetA() antlr.Token { return s.a }

func (s *TypeOrNewtypeContext) GetB() antlr.Token { return s.b }

func (s *TypeOrNewtypeContext) SetKw(v antlr.Token) { s.kw = v }

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
	kw antlr.Token
	a  antlr.Token
}

func NewModuleAnnotationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModuleAnnotationContext {
	var p = new(ModuleAnnotationContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

func (s *ModuleAnnotationContext) GetKw() antlr.Token { return s.kw }

func (s *ModuleAnnotationContext) GetA() antlr.Token { return s.a }

func (s *ModuleAnnotationContext) SetKw(v antlr.Token) { s.kw = v }

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

	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStructOrUnionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLParserAT || _la == ADLParserLINE_DOC {
			{
				p.SetState(75)
				p.Annon()
			}

			p.SetState(80)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(81)

			var _m = p.Match(ADLParserID)

			localctx.(*StructOrUnionContext).kw = _m
		}
		{
			p.SetState(82)

			var _m = p.Match(ADLParserID)

			localctx.(*StructOrUnionContext).a = _m
		}
		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLParserLCHEVR {
			{
				p.SetState(83)
				p.TypeParam()
			}

		}
		{
			p.SetState(86)
			p.Match(ADLParserLCUR)
		}
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ADLParserAT)|(1<<ADLParserID)|(1<<ADLParserLINE_DOC))) != 0 {
			{
				p.SetState(87)
				p.SoruBody()
			}

			p.SetState(92)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(93)
			p.Match(ADLParserRCUR)
		}
		{
			p.SetState(94)
			p.Match(ADLParserSEMI)
		}

	case 2:
		localctx = NewTypeOrNewtypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLParserAT || _la == ADLParserLINE_DOC {
			{
				p.SetState(95)
				p.Annon()
			}

			p.SetState(100)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(101)

			var _m = p.Match(ADLParserID)

			localctx.(*TypeOrNewtypeContext).kw = _m
		}
		{
			p.SetState(102)

			var _m = p.Match(ADLParserID)

			localctx.(*TypeOrNewtypeContext).a = _m
		}
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLParserLCHEVR {
			{
				p.SetState(103)
				p.TypeParam()
			}

		}
		{
			p.SetState(106)
			p.Match(ADLParserEQ)
		}
		{
			p.SetState(107)

			var _m = p.Match(ADLParserID)

			localctx.(*TypeOrNewtypeContext).b = _m
		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLParserLCHEVR {
			{
				p.SetState(108)
				p.TypeExpr()
			}

		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLParserEQ {
			{
				p.SetState(111)
				p.Match(ADLParserEQ)
			}
			{
				p.SetState(112)
				p.JsonValue()
			}

		}
		{
			p.SetState(115)
			p.Match(ADLParserSEMI)
		}

	case 3:
		localctx = NewModuleAnnotationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(116)

			var _m = p.Match(ADLParserID)

			localctx.(*ModuleAnnotationContext).kw = _m
		}
		{
			p.SetState(117)

			var _m = p.Match(ADLParserID)

			localctx.(*ModuleAnnotationContext).a = _m
		}
		{
			p.SetState(118)
			p.JsonValue()
		}
		{
			p.SetState(119)
			p.Match(ADLParserSEMI)
		}

	case 4:
		localctx = NewDeclAnnotationContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(121)

			var _m = p.Match(ADLParserID)

			localctx.(*DeclAnnotationContext).kw = _m
		}
		{
			p.SetState(122)

			var _m = p.Match(ADLParserID)

			localctx.(*DeclAnnotationContext).a = _m
		}
		{
			p.SetState(123)

			var _m = p.Match(ADLParserID)

			localctx.(*DeclAnnotationContext).b = _m
		}
		{
			p.SetState(124)
			p.JsonValue()
		}
		{
			p.SetState(125)
			p.Match(ADLParserSEMI)
		}

	case 5:
		localctx = NewFieldAnnotationContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(127)

			var _m = p.Match(ADLParserID)

			localctx.(*FieldAnnotationContext).kw = _m
		}
		{
			p.SetState(128)

			var _m = p.Match(ADLParserID)

			localctx.(*FieldAnnotationContext).a = _m
		}
		{
			p.SetState(129)
			p.Match(ADLParserDCOLON)
		}
		{
			p.SetState(130)

			var _m = p.Match(ADLParserID)

			localctx.(*FieldAnnotationContext).b = _m
		}
		{
			p.SetState(131)

			var _m = p.Match(ADLParserID)

			localctx.(*FieldAnnotationContext).c = _m
		}
		{
			p.SetState(132)
			p.JsonValue()
		}
		{
			p.SetState(133)
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
		p.SetState(137)
		p.Match(ADLParserLCHEVR)
	}
	{
		p.SetState(138)

		var _m = p.Match(ADLParserID)

		localctx.(*TypeParameterContext)._ID = _m
	}
	localctx.(*TypeParameterContext).typep = append(localctx.(*TypeParameterContext).typep, localctx.(*TypeParameterContext)._ID)
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ADLParserCOMMA {
		{
			p.SetState(139)
			p.Match(ADLParserCOMMA)
		}
		{
			p.SetState(140)

			var _m = p.Match(ADLParserID)

			localctx.(*TypeParameterContext)._ID = _m
		}
		localctx.(*TypeParameterContext).typep = append(localctx.(*TypeParameterContext).typep, localctx.(*TypeParameterContext)._ID)

		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(146)
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
		p.SetState(148)
		p.Match(ADLParserLCHEVR)
	}
	{
		p.SetState(149)

		var _x = p.TypeExprElem()

		localctx.(*TypeExpressionContext)._typeExprElem = _x
	}
	localctx.(*TypeExpressionContext).typep = append(localctx.(*TypeExpressionContext).typep, localctx.(*TypeExpressionContext)._typeExprElem)
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ADLParserCOMMA {
		{
			p.SetState(150)
			p.Match(ADLParserCOMMA)
		}
		{
			p.SetState(151)

			var _x = p.TypeExprElem()

			localctx.(*TypeExpressionContext)._typeExprElem = _x
		}
		localctx.(*TypeExpressionContext).typep = append(localctx.(*TypeExpressionContext).typep, localctx.(*TypeExpressionContext)._typeExprElem)

		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(157)
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
		p.SetState(159)
		p.Match(ADLParserID)
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ADLParserLCHEVR {
		{
			p.SetState(160)
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
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ADLParserAT || _la == ADLParserLINE_DOC {
		{
			p.SetState(163)
			p.Annon()
		}

		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(169)

		var _m = p.Match(ADLParserID)

		localctx.(*FieldStatementContext).a = _m
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ADLParserLCHEVR {
		{
			p.SetState(170)
			p.TypeExpr()
		}

	}
	{
		p.SetState(173)

		var _m = p.Match(ADLParserID)

		localctx.(*FieldStatementContext).b = _m
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ADLParserEQ {
		{
			p.SetState(174)
			p.Match(ADLParserEQ)
		}
		{
			p.SetState(175)
			p.JsonValue()
		}

	}
	{
		p.SetState(178)
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
	kw antlr.Token
}

func NewTrueFalseNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueFalseNullContext {
	var p = new(TrueFalseNullContext)

	p.JsonValueContext = NewEmptyJsonValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValueContext))

	return p
}

func (s *TrueFalseNullContext) GetKw() antlr.Token { return s.kw }

func (s *TrueFalseNullContext) SetKw(v antlr.Token) { s.kw = v }

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

	p.SetState(212)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ADLParserSTR:
		localctx = NewStringStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)

			var _m = p.Match(ADLParserSTR)

			localctx.(*StringStatementContext).s = _m
		}

	case ADLParserID:
		localctx = NewTrueFalseNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(181)

			var _m = p.Match(ADLParserID)

			localctx.(*TrueFalseNullContext).kw = _m
		}

	case ADLParserINT:
		localctx = NewNumberStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(182)
			p.Match(ADLParserINT)
		}

	case ADLParserFLT:
		localctx = NewFloatStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(183)
			p.Match(ADLParserFLT)
		}

	case ADLParserLSQ:
		localctx = NewArrayStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(184)
			p.Match(ADLParserLSQ)
		}
		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ADLParserLCUR)|(1<<ADLParserLSQ)|(1<<ADLParserSTR)|(1<<ADLParserID)|(1<<ADLParserINT)|(1<<ADLParserFLT))) != 0 {
			{
				p.SetState(185)
				p.JsonValue()
			}
			p.SetState(190)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLParserCOMMA {
				{
					p.SetState(186)
					p.Match(ADLParserCOMMA)
				}
				{
					p.SetState(187)
					p.JsonValue()
				}

				p.SetState(192)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(195)
			p.Match(ADLParserRSQ)
		}

	case ADLParserLCUR:
		localctx = NewObjStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(196)
			p.Match(ADLParserLCUR)
		}
		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLParserSTR {
			{
				p.SetState(197)
				p.Match(ADLParserSTR)
			}
			{
				p.SetState(198)
				p.Match(ADLParserCOLON)
			}
			{
				p.SetState(199)
				p.JsonValue()
			}
			p.SetState(206)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLParserCOMMA {
				{
					p.SetState(200)
					p.Match(ADLParserCOMMA)
				}
				{
					p.SetState(201)
					p.Match(ADLParserSTR)
				}
				{
					p.SetState(202)
					p.Match(ADLParserCOLON)
				}
				{
					p.SetState(203)
					p.JsonValue()
				}

				p.SetState(208)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(211)
			p.Match(ADLParserRCUR)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
