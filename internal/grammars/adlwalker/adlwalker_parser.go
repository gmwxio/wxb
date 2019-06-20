// Code generated from ADLWalker.g4 by ANTLR 4.7.1. DO NOT EDIT.

package walker // ADLWalker
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 51, 202,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 7, 3, 24, 10, 3,
	12, 3, 14, 3, 27, 11, 3, 3, 3, 7, 3, 30, 10, 3, 12, 3, 14, 3, 33, 11, 3,
	3, 3, 7, 3, 36, 10, 3, 12, 3, 14, 3, 39, 11, 3, 3, 3, 5, 3, 42, 10, 3,
	3, 4, 3, 4, 3, 4, 7, 4, 47, 10, 4, 12, 4, 14, 4, 50, 11, 4, 3, 4, 5, 4,
	53, 10, 4, 3, 4, 7, 4, 56, 10, 4, 12, 4, 14, 4, 59, 11, 4, 3, 4, 5, 4,
	62, 10, 4, 3, 4, 3, 4, 3, 4, 7, 4, 67, 10, 4, 12, 4, 14, 4, 70, 11, 4,
	3, 4, 5, 4, 73, 10, 4, 3, 4, 7, 4, 76, 10, 4, 12, 4, 14, 4, 79, 11, 4,
	3, 4, 5, 4, 82, 10, 4, 3, 4, 3, 4, 3, 4, 7, 4, 87, 10, 4, 12, 4, 14, 4,
	90, 11, 4, 3, 4, 5, 4, 93, 10, 4, 3, 4, 5, 4, 96, 10, 4, 3, 4, 7, 4, 99,
	10, 4, 12, 4, 14, 4, 102, 11, 4, 3, 4, 5, 4, 105, 10, 4, 3, 4, 3, 4, 3,
	4, 7, 4, 110, 10, 4, 12, 4, 14, 4, 113, 11, 4, 3, 4, 5, 4, 116, 10, 4,
	3, 4, 5, 4, 119, 10, 4, 3, 4, 7, 4, 122, 10, 4, 12, 4, 14, 4, 125, 11,
	4, 3, 4, 5, 4, 128, 10, 4, 3, 4, 3, 4, 3, 4, 5, 4, 133, 10, 4, 3, 5, 3,
	5, 3, 5, 7, 5, 138, 10, 5, 12, 5, 14, 5, 141, 11, 5, 3, 5, 5, 5, 144, 10,
	5, 3, 5, 5, 5, 147, 10, 5, 3, 5, 5, 5, 150, 10, 5, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 6, 6, 177, 10,
	6, 13, 6, 14, 6, 178, 3, 6, 3, 6, 5, 6, 183, 10, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 6, 6, 191, 10, 6, 13, 6, 14, 6, 192, 3, 6, 3, 6, 5, 6, 197,
	10, 6, 3, 6, 5, 6, 200, 10, 6, 3, 6, 2, 2, 7, 2, 4, 6, 8, 10, 2, 2, 2,
	238, 2, 12, 3, 2, 2, 2, 4, 20, 3, 2, 2, 2, 6, 132, 3, 2, 2, 2, 8, 134,
	3, 2, 2, 2, 10, 199, 3, 2, 2, 2, 12, 13, 7, 26, 2, 2, 13, 14, 7, 30, 2,
	2, 14, 15, 7, 26, 2, 2, 15, 16, 5, 4, 3, 2, 16, 17, 7, 27, 2, 2, 17, 18,
	7, 27, 2, 2, 18, 19, 7, 2, 2, 3, 19, 3, 3, 2, 2, 2, 20, 41, 7, 31, 2, 2,
	21, 25, 7, 26, 2, 2, 22, 24, 7, 33, 2, 2, 23, 22, 3, 2, 2, 2, 24, 27, 3,
	2, 2, 2, 25, 23, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 31, 3, 2, 2, 2, 27,
	25, 3, 2, 2, 2, 28, 30, 7, 32, 2, 2, 29, 28, 3, 2, 2, 2, 30, 33, 3, 2,
	2, 2, 31, 29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 37, 3, 2, 2, 2, 33, 31,
	3, 2, 2, 2, 34, 36, 5, 6, 4, 2, 35, 34, 3, 2, 2, 2, 36, 39, 3, 2, 2, 2,
	37, 35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 40, 3, 2, 2, 2, 39, 37, 3,
	2, 2, 2, 40, 42, 7, 27, 2, 2, 41, 21, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42,
	5, 3, 2, 2, 2, 43, 61, 7, 34, 2, 2, 44, 48, 7, 26, 2, 2, 45, 47, 7, 33,
	2, 2, 46, 45, 3, 2, 2, 2, 47, 50, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49,
	3, 2, 2, 2, 49, 52, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 51, 53, 7, 38, 2, 2,
	52, 51, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 57, 3, 2, 2, 2, 54, 56, 5,
	8, 5, 2, 55, 54, 3, 2, 2, 2, 56, 59, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 57,
	58, 3, 2, 2, 2, 58, 60, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 60, 62, 7, 27,
	2, 2, 61, 44, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 133, 3, 2, 2, 2, 63,
	81, 7, 35, 2, 2, 64, 68, 7, 26, 2, 2, 65, 67, 7, 33, 2, 2, 66, 65, 3, 2,
	2, 2, 67, 70, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 72,
	3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 71, 73, 7, 38, 2, 2, 72, 71, 3, 2, 2, 2,
	72, 73, 3, 2, 2, 2, 73, 77, 3, 2, 2, 2, 74, 76, 5, 8, 5, 2, 75, 74, 3,
	2, 2, 2, 76, 79, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78,
	80, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 80, 82, 7, 27, 2, 2, 81, 64, 3, 2,
	2, 2, 81, 82, 3, 2, 2, 2, 82, 133, 3, 2, 2, 2, 83, 104, 7, 37, 2, 2, 84,
	88, 7, 26, 2, 2, 85, 87, 7, 33, 2, 2, 86, 85, 3, 2, 2, 2, 87, 90, 3, 2,
	2, 2, 88, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 92, 3, 2, 2, 2, 90, 88,
	3, 2, 2, 2, 91, 93, 7, 38, 2, 2, 92, 91, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2,
	93, 95, 3, 2, 2, 2, 94, 96, 7, 39, 2, 2, 95, 94, 3, 2, 2, 2, 95, 96, 3,
	2, 2, 2, 96, 100, 3, 2, 2, 2, 97, 99, 5, 10, 6, 2, 98, 97, 3, 2, 2, 2,
	99, 102, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 103,
	3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 103, 105, 7, 27, 2, 2, 104, 84, 3, 2,
	2, 2, 104, 105, 3, 2, 2, 2, 105, 133, 3, 2, 2, 2, 106, 127, 7, 36, 2, 2,
	107, 111, 7, 26, 2, 2, 108, 110, 7, 33, 2, 2, 109, 108, 3, 2, 2, 2, 110,
	113, 3, 2, 2, 2, 111, 109, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 115,
	3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 114, 116, 7, 38, 2, 2, 115, 114, 3, 2,
	2, 2, 115, 116, 3, 2, 2, 2, 116, 118, 3, 2, 2, 2, 117, 119, 7, 39, 2, 2,
	118, 117, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 123, 3, 2, 2, 2, 120,
	122, 5, 10, 6, 2, 121, 120, 3, 2, 2, 2, 122, 125, 3, 2, 2, 2, 123, 121,
	3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 126, 3, 2, 2, 2, 125, 123, 3, 2,
	2, 2, 126, 128, 7, 27, 2, 2, 127, 107, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2,
	128, 133, 3, 2, 2, 2, 129, 133, 7, 49, 2, 2, 130, 133, 7, 50, 2, 2, 131,
	133, 7, 51, 2, 2, 132, 43, 3, 2, 2, 2, 132, 63, 3, 2, 2, 2, 132, 83, 3,
	2, 2, 2, 132, 106, 3, 2, 2, 2, 132, 129, 3, 2, 2, 2, 132, 130, 3, 2, 2,
	2, 132, 131, 3, 2, 2, 2, 133, 7, 3, 2, 2, 2, 134, 149, 7, 40, 2, 2, 135,
	139, 7, 26, 2, 2, 136, 138, 7, 33, 2, 2, 137, 136, 3, 2, 2, 2, 138, 141,
	3, 2, 2, 2, 139, 137, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 143, 3, 2,
	2, 2, 141, 139, 3, 2, 2, 2, 142, 144, 7, 39, 2, 2, 143, 142, 3, 2, 2, 2,
	143, 144, 3, 2, 2, 2, 144, 146, 3, 2, 2, 2, 145, 147, 5, 10, 6, 2, 146,
	145, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2, 148, 150,
	7, 27, 2, 2, 149, 135, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 9, 3, 2,
	2, 2, 151, 152, 7, 41, 2, 2, 152, 153, 7, 26, 2, 2, 153, 154, 7, 42, 2,
	2, 154, 200, 7, 27, 2, 2, 155, 156, 7, 41, 2, 2, 156, 157, 7, 26, 2, 2,
	157, 158, 7, 43, 2, 2, 158, 200, 7, 27, 2, 2, 159, 160, 7, 41, 2, 2, 160,
	161, 7, 26, 2, 2, 161, 162, 7, 44, 2, 2, 162, 200, 7, 27, 2, 2, 163, 164,
	7, 41, 2, 2, 164, 165, 7, 26, 2, 2, 165, 166, 7, 45, 2, 2, 166, 200, 7,
	27, 2, 2, 167, 168, 7, 41, 2, 2, 168, 169, 7, 26, 2, 2, 169, 170, 7, 46,
	2, 2, 170, 200, 7, 27, 2, 2, 171, 172, 7, 41, 2, 2, 172, 173, 7, 26, 2,
	2, 173, 182, 7, 47, 2, 2, 174, 176, 7, 26, 2, 2, 175, 177, 5, 10, 6, 2,
	176, 175, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 176, 3, 2, 2, 2, 178,
	179, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 181, 7, 27, 2, 2, 181, 183,
	3, 2, 2, 2, 182, 174, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183, 184, 3, 2,
	2, 2, 184, 200, 7, 27, 2, 2, 185, 186, 7, 41, 2, 2, 186, 187, 7, 26, 2,
	2, 187, 196, 7, 48, 2, 2, 188, 190, 7, 26, 2, 2, 189, 191, 5, 10, 6, 2,
	190, 189, 3, 2, 2, 2, 191, 192, 3, 2, 2, 2, 192, 190, 3, 2, 2, 2, 192,
	193, 3, 2, 2, 2, 193, 194, 3, 2, 2, 2, 194, 195, 7, 27, 2, 2, 195, 197,
	3, 2, 2, 2, 196, 188, 3, 2, 2, 2, 196, 197, 3, 2, 2, 2, 197, 198, 3, 2,
	2, 2, 198, 200, 7, 27, 2, 2, 199, 151, 3, 2, 2, 2, 199, 155, 3, 2, 2, 2,
	199, 159, 3, 2, 2, 2, 199, 163, 3, 2, 2, 2, 199, 167, 3, 2, 2, 2, 199,
	171, 3, 2, 2, 2, 199, 185, 3, 2, 2, 2, 200, 11, 3, 2, 2, 2, 34, 25, 31,
	37, 41, 48, 52, 57, 61, 68, 72, 77, 81, 88, 92, 95, 100, 104, 111, 115,
	118, 123, 127, 132, 139, 143, 146, 149, 178, 182, 192, 196, 199,
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
	"adl", "module", "tld", "nameBody", "jsonVal",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ADLWalker struct {
	*antlr.BaseParser
}

func NewADLWalker(input antlr.TokenStream) *ADLWalker {
	this := new(ADLWalker)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "ADLWalker.g4"

	return this
}

// ADLWalker tokens.
const (
	ADLWalkerEOF          = antlr.TokenEOF
	ADLWalkerLCUR         = 1
	ADLWalkerRCUR         = 2
	ADLWalkerLSQ          = 3
	ADLWalkerRSQ          = 4
	ADLWalkerEQ           = 5
	ADLWalkerDQ           = 6
	ADLWalkerSQ           = 7
	ADLWalkerSEMI         = 8
	ADLWalkerDCOLON       = 9
	ADLWalkerCOLON        = 10
	ADLWalkerDOT          = 11
	ADLWalkerCOMMA        = 12
	ADLWalkerLCHEVR       = 13
	ADLWalkerRCHEVR       = 14
	ADLWalkerSTAR         = 15
	ADLWalkerAT           = 16
	ADLWalkerSTR          = 17
	ADLWalkerID           = 18
	ADLWalkerINT          = 19
	ADLWalkerFLT          = 20
	ADLWalkerWS           = 21
	ADLWalkerLINE_DOC     = 22
	ADLWalkerLINE_COMMENT = 23
	ADLWalkerDOWN         = 24
	ADLWalkerUP           = 25
	ADLWalkerROOT         = 26
	ADLWalkerERROR        = 27
	ADLWalkerADL          = 28
	ADLWalkerModule       = 29
	ADLWalkerImport       = 30
	ADLWalkerAnnotation   = 31
	ADLWalkerStruct       = 32
	ADLWalkerUnion        = 33
	ADLWalkerNewtype      = 34
	ADLWalkerType         = 35
	ADLWalkerTypeParam    = 36
	ADLWalkerTypeExpr     = 37
	ADLWalkerField        = 38
	ADLWalkerJson         = 39
	ADLWalkerJsonStr      = 40
	ADLWalkerJsonBool     = 41
	ADLWalkerJsonNull     = 42
	ADLWalkerJsonInt      = 43
	ADLWalkerJsonFloat    = 44
	ADLWalkerJsonArray    = 45
	ADLWalkerJsonObj      = 46
	ADLWalkerModuleAnno   = 47
	ADLWalkerDeclAnno     = 48
	ADLWalkerFieldAnno    = 49
)

// ADLWalker rules.
const (
	ADLWalkerRULE_adl      = 0
	ADLWalkerRULE_module   = 1
	ADLWalkerRULE_tld      = 2
	ADLWalkerRULE_nameBody = 3
	ADLWalkerRULE_jsonVal  = 4
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
	p.RuleIndex = ADLWalkerRULE_adl
	return p
}

func (*AdlContext) IsAdlContext() {}

func NewAdlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdlContext {
	var p = new(AdlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLWalkerRULE_adl

	return p
}

func (s *AdlContext) GetParser() antlr.Parser { return s.parser }

func (s *AdlContext) AllDOWN() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerDOWN)
}

func (s *AdlContext) DOWN(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, i)
}

func (s *AdlContext) ADL() antlr.TerminalNode {
	return s.GetToken(ADLWalkerADL, 0)
}

func (s *AdlContext) Module() IModuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModuleContext)
}

func (s *AdlContext) AllUP() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerUP)
}

func (s *AdlContext) UP(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, i)
}

func (s *AdlContext) EOF() antlr.TerminalNode {
	return s.GetToken(ADLWalkerEOF, 0)
}

func (s *AdlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterAdl(s)
	}
}

func (s *AdlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitAdl(s)
	}
}

func (p *ADLWalker) Adl() (localctx IAdlContext) {
	localctx = NewAdlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ADLWalkerRULE_adl)

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
		p.SetState(10)
		p.Match(ADLWalkerDOWN)
	}
	{
		p.SetState(11)
		p.Match(ADLWalkerADL)
	}
	{
		p.SetState(12)
		p.Match(ADLWalkerDOWN)
	}
	{
		p.SetState(13)
		p.Module()
	}
	{
		p.SetState(14)
		p.Match(ADLWalkerUP)
	}
	{
		p.SetState(15)
		p.Match(ADLWalkerUP)
	}
	{
		p.SetState(16)
		p.Match(ADLWalkerEOF)
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
	p.RuleIndex = ADLWalkerRULE_module
	return p
}

func (*ModuleContext) IsModuleContext() {}

func NewModuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleContext {
	var p = new(ModuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLWalkerRULE_module

	return p
}

func (s *ModuleContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleContext) Module() antlr.TerminalNode {
	return s.GetToken(ADLWalkerModule, 0)
}

func (s *ModuleContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *ModuleContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *ModuleContext) AllAnnotation() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerAnnotation)
}

func (s *ModuleContext) Annotation(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerAnnotation, i)
}

func (s *ModuleContext) AllImport() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerImport)
}

func (s *ModuleContext) Import(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerImport, i)
}

func (s *ModuleContext) AllTld() []ITldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITldContext)(nil)).Elem())
	var tst = make([]ITldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITldContext)
		}
	}

	return tst
}

func (s *ModuleContext) Tld(i int) ITldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITldContext)
}

func (s *ModuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterModule(s)
	}
}

func (s *ModuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitModule(s)
	}
}

func (p *ADLWalker) Module() (localctx IModuleContext) {
	localctx = NewModuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ADLWalkerRULE_module)
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
		p.SetState(18)
		p.Match(ADLWalkerModule)
	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ADLWalkerDOWN {
		{
			p.SetState(19)
			p.Match(ADLWalkerDOWN)
		}
		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLWalkerAnnotation {
			{
				p.SetState(20)
				p.Match(ADLWalkerAnnotation)
			}

			p.SetState(25)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLWalkerImport {
			{
				p.SetState(26)
				p.Match(ADLWalkerImport)
			}

			p.SetState(31)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ADLWalkerStruct-32))|(1<<(ADLWalkerUnion-32))|(1<<(ADLWalkerNewtype-32))|(1<<(ADLWalkerType-32))|(1<<(ADLWalkerModuleAnno-32))|(1<<(ADLWalkerDeclAnno-32))|(1<<(ADLWalkerFieldAnno-32)))) != 0 {
			{
				p.SetState(32)
				p.Tld()
			}

			p.SetState(37)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(38)
			p.Match(ADLWalkerUP)
		}

	}

	return localctx
}

// ITldContext is an interface to support dynamic dispatch.
type ITldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTldContext differentiates from other interfaces.
	IsTldContext()
}

type TldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTldContext() *TldContext {
	var p = new(TldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLWalkerRULE_tld
	return p
}

func (*TldContext) IsTldContext() {}

func NewTldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TldContext {
	var p = new(TldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLWalkerRULE_tld

	return p
}

func (s *TldContext) GetParser() antlr.Parser { return s.parser }

func (s *TldContext) CopyFrom(ctx *TldContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypeContext struct {
	*TldContext
}

func NewTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeContext {
	var p = new(TypeContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) Type() antlr.TerminalNode {
	return s.GetToken(ADLWalkerType, 0)
}

func (s *TypeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *TypeContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *TypeContext) AllAnnotation() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerAnnotation)
}

func (s *TypeContext) Annotation(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerAnnotation, i)
}

func (s *TypeContext) TypeParam() antlr.TerminalNode {
	return s.GetToken(ADLWalkerTypeParam, 0)
}

func (s *TypeContext) TypeExpr() antlr.TerminalNode {
	return s.GetToken(ADLWalkerTypeExpr, 0)
}

func (s *TypeContext) AllJsonVal() []IJsonValContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IJsonValContext)(nil)).Elem())
	var tst = make([]IJsonValContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonValContext)
		}
	}

	return tst
}

func (s *TypeContext) JsonVal(i int) IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitType(s)
	}
}

type NewtypeContext struct {
	*TldContext
}

func NewNewtypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NewtypeContext {
	var p = new(NewtypeContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

func (s *NewtypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewtypeContext) Newtype() antlr.TerminalNode {
	return s.GetToken(ADLWalkerNewtype, 0)
}

func (s *NewtypeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *NewtypeContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *NewtypeContext) AllAnnotation() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerAnnotation)
}

func (s *NewtypeContext) Annotation(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerAnnotation, i)
}

func (s *NewtypeContext) TypeParam() antlr.TerminalNode {
	return s.GetToken(ADLWalkerTypeParam, 0)
}

func (s *NewtypeContext) TypeExpr() antlr.TerminalNode {
	return s.GetToken(ADLWalkerTypeExpr, 0)
}

func (s *NewtypeContext) AllJsonVal() []IJsonValContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IJsonValContext)(nil)).Elem())
	var tst = make([]IJsonValContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonValContext)
		}
	}

	return tst
}

func (s *NewtypeContext) JsonVal(i int) IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *NewtypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterNewtype(s)
	}
}

func (s *NewtypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitNewtype(s)
	}
}

type ModAnnoContext struct {
	*TldContext
}

func NewModAnnoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModAnnoContext {
	var p = new(ModAnnoContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

func (s *ModAnnoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModAnnoContext) ModuleAnno() antlr.TerminalNode {
	return s.GetToken(ADLWalkerModuleAnno, 0)
}

func (s *ModAnnoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterModAnno(s)
	}
}

func (s *ModAnnoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitModAnno(s)
	}
}

type FieldAnnoContext struct {
	*TldContext
}

func NewFieldAnnoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldAnnoContext {
	var p = new(FieldAnnoContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

func (s *FieldAnnoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldAnnoContext) FieldAnno() antlr.TerminalNode {
	return s.GetToken(ADLWalkerFieldAnno, 0)
}

func (s *FieldAnnoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterFieldAnno(s)
	}
}

func (s *FieldAnnoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitFieldAnno(s)
	}
}

type UnionContext struct {
	*TldContext
}

func NewUnionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnionContext {
	var p = new(UnionContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

func (s *UnionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionContext) Union() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUnion, 0)
}

func (s *UnionContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *UnionContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *UnionContext) AllAnnotation() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerAnnotation)
}

func (s *UnionContext) Annotation(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerAnnotation, i)
}

func (s *UnionContext) TypeParam() antlr.TerminalNode {
	return s.GetToken(ADLWalkerTypeParam, 0)
}

func (s *UnionContext) AllNameBody() []INameBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameBodyContext)(nil)).Elem())
	var tst = make([]INameBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameBodyContext)
		}
	}

	return tst
}

func (s *UnionContext) NameBody(i int) INameBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameBodyContext)
}

func (s *UnionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterUnion(s)
	}
}

func (s *UnionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitUnion(s)
	}
}

type DeclAnnoContext struct {
	*TldContext
}

func NewDeclAnnoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclAnnoContext {
	var p = new(DeclAnnoContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

func (s *DeclAnnoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclAnnoContext) DeclAnno() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDeclAnno, 0)
}

func (s *DeclAnnoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterDeclAnno(s)
	}
}

func (s *DeclAnnoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitDeclAnno(s)
	}
}

type StructContext struct {
	*TldContext
}

func NewStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructContext {
	var p = new(StructContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

func (s *StructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructContext) Struct() antlr.TerminalNode {
	return s.GetToken(ADLWalkerStruct, 0)
}

func (s *StructContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *StructContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *StructContext) AllAnnotation() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerAnnotation)
}

func (s *StructContext) Annotation(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerAnnotation, i)
}

func (s *StructContext) TypeParam() antlr.TerminalNode {
	return s.GetToken(ADLWalkerTypeParam, 0)
}

func (s *StructContext) AllNameBody() []INameBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameBodyContext)(nil)).Elem())
	var tst = make([]INameBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameBodyContext)
		}
	}

	return tst
}

func (s *StructContext) NameBody(i int) INameBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameBodyContext)
}

func (s *StructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterStruct(s)
	}
}

func (s *StructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitStruct(s)
	}
}

func (p *ADLWalker) Tld() (localctx ITldContext) {
	localctx = NewTldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ADLWalkerRULE_tld)
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

	p.SetState(130)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ADLWalkerStruct:
		localctx = NewStructContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)
			p.Match(ADLWalkerStruct)
		}
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerDOWN {
			{
				p.SetState(42)
				p.Match(ADLWalkerDOWN)
			}
			p.SetState(46)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLWalkerAnnotation {
				{
					p.SetState(43)
					p.Match(ADLWalkerAnnotation)
				}

				p.SetState(48)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(50)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ADLWalkerTypeParam {
				{
					p.SetState(49)
					p.Match(ADLWalkerTypeParam)
				}

			}
			p.SetState(55)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLWalkerField {
				{
					p.SetState(52)
					p.NameBody()
				}

				p.SetState(57)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(58)
				p.Match(ADLWalkerUP)
			}

		}

	case ADLWalkerUnion:
		localctx = NewUnionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(61)
			p.Match(ADLWalkerUnion)
		}
		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerDOWN {
			{
				p.SetState(62)
				p.Match(ADLWalkerDOWN)
			}
			p.SetState(66)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLWalkerAnnotation {
				{
					p.SetState(63)
					p.Match(ADLWalkerAnnotation)
				}

				p.SetState(68)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(70)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ADLWalkerTypeParam {
				{
					p.SetState(69)
					p.Match(ADLWalkerTypeParam)
				}

			}
			p.SetState(75)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLWalkerField {
				{
					p.SetState(72)
					p.NameBody()
				}

				p.SetState(77)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(78)
				p.Match(ADLWalkerUP)
			}

		}

	case ADLWalkerType:
		localctx = NewTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(81)
			p.Match(ADLWalkerType)
		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerDOWN {
			{
				p.SetState(82)
				p.Match(ADLWalkerDOWN)
			}
			p.SetState(86)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLWalkerAnnotation {
				{
					p.SetState(83)
					p.Match(ADLWalkerAnnotation)
				}

				p.SetState(88)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(90)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ADLWalkerTypeParam {
				{
					p.SetState(89)
					p.Match(ADLWalkerTypeParam)
				}

			}
			p.SetState(93)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ADLWalkerTypeExpr {
				{
					p.SetState(92)
					p.Match(ADLWalkerTypeExpr)
				}

			}
			p.SetState(98)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLWalkerJson {
				{
					p.SetState(95)
					p.JsonVal()
				}

				p.SetState(100)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(101)
				p.Match(ADLWalkerUP)
			}

		}

	case ADLWalkerNewtype:
		localctx = NewNewtypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(104)
			p.Match(ADLWalkerNewtype)
		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerDOWN {
			{
				p.SetState(105)
				p.Match(ADLWalkerDOWN)
			}
			p.SetState(109)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLWalkerAnnotation {
				{
					p.SetState(106)
					p.Match(ADLWalkerAnnotation)
				}

				p.SetState(111)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(113)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ADLWalkerTypeParam {
				{
					p.SetState(112)
					p.Match(ADLWalkerTypeParam)
				}

			}
			p.SetState(116)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ADLWalkerTypeExpr {
				{
					p.SetState(115)
					p.Match(ADLWalkerTypeExpr)
				}

			}
			p.SetState(121)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLWalkerJson {
				{
					p.SetState(118)
					p.JsonVal()
				}

				p.SetState(123)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(124)
				p.Match(ADLWalkerUP)
			}

		}

	case ADLWalkerModuleAnno:
		localctx = NewModAnnoContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(127)
			p.Match(ADLWalkerModuleAnno)
		}

	case ADLWalkerDeclAnno:
		localctx = NewDeclAnnoContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(128)
			p.Match(ADLWalkerDeclAnno)
		}

	case ADLWalkerFieldAnno:
		localctx = NewFieldAnnoContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(129)
			p.Match(ADLWalkerFieldAnno)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INameBodyContext is an interface to support dynamic dispatch.
type INameBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameBodyContext differentiates from other interfaces.
	IsNameBodyContext()
}

type NameBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameBodyContext() *NameBodyContext {
	var p = new(NameBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLWalkerRULE_nameBody
	return p
}

func (*NameBodyContext) IsNameBodyContext() {}

func NewNameBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameBodyContext {
	var p = new(NameBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLWalkerRULE_nameBody

	return p
}

func (s *NameBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *NameBodyContext) CopyFrom(ctx *NameBodyContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *NameBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FieldContext struct {
	*NameBodyContext
}

func NewFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldContext {
	var p = new(FieldContext)

	p.NameBodyContext = NewEmptyNameBodyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NameBodyContext))

	return p
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) Field() antlr.TerminalNode {
	return s.GetToken(ADLWalkerField, 0)
}

func (s *FieldContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *FieldContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *FieldContext) AllAnnotation() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerAnnotation)
}

func (s *FieldContext) Annotation(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerAnnotation, i)
}

func (s *FieldContext) TypeExpr() antlr.TerminalNode {
	return s.GetToken(ADLWalkerTypeExpr, 0)
}

func (s *FieldContext) JsonVal() IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *ADLWalker) NameBody() (localctx INameBodyContext) {
	localctx = NewNameBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ADLWalkerRULE_nameBody)
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

	localctx = NewFieldContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Match(ADLWalkerField)
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ADLWalkerDOWN {
		{
			p.SetState(133)
			p.Match(ADLWalkerDOWN)
		}
		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLWalkerAnnotation {
			{
				p.SetState(134)
				p.Match(ADLWalkerAnnotation)
			}

			p.SetState(139)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerTypeExpr {
			{
				p.SetState(140)
				p.Match(ADLWalkerTypeExpr)
			}

		}
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerJson {
			{
				p.SetState(143)
				p.JsonVal()
			}

		}
		{
			p.SetState(146)
			p.Match(ADLWalkerUP)
		}

	}

	return localctx
}

// IJsonValContext is an interface to support dynamic dispatch.
type IJsonValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJsonValContext differentiates from other interfaces.
	IsJsonValContext()
}

type JsonValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonValContext() *JsonValContext {
	var p = new(JsonValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLWalkerRULE_jsonVal
	return p
}

func (*JsonValContext) IsJsonValContext() {}

func NewJsonValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonValContext {
	var p = new(JsonValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLWalkerRULE_jsonVal

	return p
}

func (s *JsonValContext) GetParser() antlr.Parser { return s.parser }

func (s *JsonValContext) CopyFrom(ctx *JsonValContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *JsonValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type JsonStrContext struct {
	*JsonValContext
}

func NewJsonStrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonStrContext {
	var p = new(JsonStrContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

func (s *JsonStrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonStrContext) Json() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJson, 0)
}

func (s *JsonStrContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *JsonStrContext) JsonStr() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJsonStr, 0)
}

func (s *JsonStrContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *JsonStrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterJsonStr(s)
	}
}

func (s *JsonStrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitJsonStr(s)
	}
}

type JsonArrayContext struct {
	*JsonValContext
}

func NewJsonArrayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonArrayContext {
	var p = new(JsonArrayContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

func (s *JsonArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonArrayContext) Json() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJson, 0)
}

func (s *JsonArrayContext) AllDOWN() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerDOWN)
}

func (s *JsonArrayContext) DOWN(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, i)
}

func (s *JsonArrayContext) JsonArray() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJsonArray, 0)
}

func (s *JsonArrayContext) AllUP() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerUP)
}

func (s *JsonArrayContext) UP(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, i)
}

func (s *JsonArrayContext) AllJsonVal() []IJsonValContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IJsonValContext)(nil)).Elem())
	var tst = make([]IJsonValContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonValContext)
		}
	}

	return tst
}

func (s *JsonArrayContext) JsonVal(i int) IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *JsonArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterJsonArray(s)
	}
}

func (s *JsonArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitJsonArray(s)
	}
}

type JsonFloatContext struct {
	*JsonValContext
}

func NewJsonFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonFloatContext {
	var p = new(JsonFloatContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

func (s *JsonFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonFloatContext) Json() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJson, 0)
}

func (s *JsonFloatContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *JsonFloatContext) JsonFloat() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJsonFloat, 0)
}

func (s *JsonFloatContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *JsonFloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterJsonFloat(s)
	}
}

func (s *JsonFloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitJsonFloat(s)
	}
}

type JsonObjContext struct {
	*JsonValContext
}

func NewJsonObjContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonObjContext {
	var p = new(JsonObjContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

func (s *JsonObjContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonObjContext) Json() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJson, 0)
}

func (s *JsonObjContext) AllDOWN() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerDOWN)
}

func (s *JsonObjContext) DOWN(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, i)
}

func (s *JsonObjContext) JsonObj() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJsonObj, 0)
}

func (s *JsonObjContext) AllUP() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerUP)
}

func (s *JsonObjContext) UP(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, i)
}

func (s *JsonObjContext) AllJsonVal() []IJsonValContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IJsonValContext)(nil)).Elem())
	var tst = make([]IJsonValContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonValContext)
		}
	}

	return tst
}

func (s *JsonObjContext) JsonVal(i int) IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *JsonObjContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterJsonObj(s)
	}
}

func (s *JsonObjContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitJsonObj(s)
	}
}

type JsonBoolContext struct {
	*JsonValContext
}

func NewJsonBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonBoolContext {
	var p = new(JsonBoolContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

func (s *JsonBoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonBoolContext) Json() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJson, 0)
}

func (s *JsonBoolContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *JsonBoolContext) JsonBool() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJsonBool, 0)
}

func (s *JsonBoolContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *JsonBoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterJsonBool(s)
	}
}

func (s *JsonBoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitJsonBool(s)
	}
}

type JsonIntContext struct {
	*JsonValContext
}

func NewJsonIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonIntContext {
	var p = new(JsonIntContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

func (s *JsonIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonIntContext) Json() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJson, 0)
}

func (s *JsonIntContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *JsonIntContext) JsonInt() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJsonInt, 0)
}

func (s *JsonIntContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *JsonIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterJsonInt(s)
	}
}

func (s *JsonIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitJsonInt(s)
	}
}

type JsonNullContext struct {
	*JsonValContext
}

func NewJsonNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonNullContext {
	var p = new(JsonNullContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

func (s *JsonNullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonNullContext) Json() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJson, 0)
}

func (s *JsonNullContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *JsonNullContext) JsonNull() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJsonNull, 0)
}

func (s *JsonNullContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *JsonNullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterJsonNull(s)
	}
}

func (s *JsonNullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitJsonNull(s)
	}
}

func (p *ADLWalker) JsonVal() (localctx IJsonValContext) {
	localctx = NewJsonValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ADLWalkerRULE_jsonVal)
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

	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		localctx = NewJsonStrContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(149)
			p.Match(ADLWalkerJson)
		}
		{
			p.SetState(150)
			p.Match(ADLWalkerDOWN)
		}
		{
			p.SetState(151)
			p.Match(ADLWalkerJsonStr)
		}
		{
			p.SetState(152)
			p.Match(ADLWalkerUP)
		}

	case 2:
		localctx = NewJsonBoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(153)
			p.Match(ADLWalkerJson)
		}
		{
			p.SetState(154)
			p.Match(ADLWalkerDOWN)
		}
		{
			p.SetState(155)
			p.Match(ADLWalkerJsonBool)
		}
		{
			p.SetState(156)
			p.Match(ADLWalkerUP)
		}

	case 3:
		localctx = NewJsonNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(157)
			p.Match(ADLWalkerJson)
		}
		{
			p.SetState(158)
			p.Match(ADLWalkerDOWN)
		}
		{
			p.SetState(159)
			p.Match(ADLWalkerJsonNull)
		}
		{
			p.SetState(160)
			p.Match(ADLWalkerUP)
		}

	case 4:
		localctx = NewJsonIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(161)
			p.Match(ADLWalkerJson)
		}
		{
			p.SetState(162)
			p.Match(ADLWalkerDOWN)
		}
		{
			p.SetState(163)
			p.Match(ADLWalkerJsonInt)
		}
		{
			p.SetState(164)
			p.Match(ADLWalkerUP)
		}

	case 5:
		localctx = NewJsonFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(165)
			p.Match(ADLWalkerJson)
		}
		{
			p.SetState(166)
			p.Match(ADLWalkerDOWN)
		}
		{
			p.SetState(167)
			p.Match(ADLWalkerJsonFloat)
		}
		{
			p.SetState(168)
			p.Match(ADLWalkerUP)
		}

	case 6:
		localctx = NewJsonArrayContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(169)
			p.Match(ADLWalkerJson)
		}
		{
			p.SetState(170)
			p.Match(ADLWalkerDOWN)
		}
		{
			p.SetState(171)
			p.Match(ADLWalkerJsonArray)
		}
		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerDOWN {
			{
				p.SetState(172)
				p.Match(ADLWalkerDOWN)
			}
			p.SetState(174)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == ADLWalkerJson {
				{
					p.SetState(173)
					p.JsonVal()
				}

				p.SetState(176)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(178)
				p.Match(ADLWalkerUP)
			}

		}
		{
			p.SetState(182)
			p.Match(ADLWalkerUP)
		}

	case 7:
		localctx = NewJsonObjContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(183)
			p.Match(ADLWalkerJson)
		}
		{
			p.SetState(184)
			p.Match(ADLWalkerDOWN)
		}
		{
			p.SetState(185)
			p.Match(ADLWalkerJsonObj)
		}
		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerDOWN {
			{
				p.SetState(186)
				p.Match(ADLWalkerDOWN)
			}
			p.SetState(188)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == ADLWalkerJson {
				{
					p.SetState(187)
					p.JsonVal()
				}

				p.SetState(190)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(192)
				p.Match(ADLWalkerUP)
			}

		}
		{
			p.SetState(196)
			p.Match(ADLWalkerUP)
		}

	}

	return localctx
}
