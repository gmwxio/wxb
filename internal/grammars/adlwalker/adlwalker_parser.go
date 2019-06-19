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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 51, 191,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3, 22, 10, 3, 12, 3, 14, 3,
	25, 11, 3, 3, 3, 3, 3, 7, 3, 29, 10, 3, 12, 3, 14, 3, 32, 11, 3, 3, 3,
	7, 3, 35, 10, 3, 12, 3, 14, 3, 38, 11, 3, 3, 4, 7, 4, 41, 10, 4, 12, 4,
	14, 4, 44, 11, 4, 3, 4, 3, 4, 3, 4, 5, 4, 49, 10, 4, 3, 4, 7, 4, 52, 10,
	4, 12, 4, 14, 4, 55, 11, 4, 3, 4, 5, 4, 58, 10, 4, 3, 4, 7, 4, 61, 10,
	4, 12, 4, 14, 4, 64, 11, 4, 3, 4, 3, 4, 3, 4, 5, 4, 69, 10, 4, 3, 4, 7,
	4, 72, 10, 4, 12, 4, 14, 4, 75, 11, 4, 3, 4, 5, 4, 78, 10, 4, 3, 4, 7,
	4, 81, 10, 4, 12, 4, 14, 4, 84, 11, 4, 3, 4, 3, 4, 3, 4, 5, 4, 89, 10,
	4, 3, 4, 5, 4, 92, 10, 4, 3, 4, 7, 4, 95, 10, 4, 12, 4, 14, 4, 98, 11,
	4, 3, 4, 5, 4, 101, 10, 4, 3, 4, 7, 4, 104, 10, 4, 12, 4, 14, 4, 107, 11,
	4, 3, 4, 3, 4, 3, 4, 5, 4, 112, 10, 4, 3, 4, 5, 4, 115, 10, 4, 3, 4, 7,
	4, 118, 10, 4, 12, 4, 14, 4, 121, 11, 4, 3, 4, 5, 4, 124, 10, 4, 3, 4,
	3, 4, 3, 4, 5, 4, 129, 10, 4, 3, 5, 7, 5, 132, 10, 5, 12, 5, 14, 5, 135,
	11, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 143, 10, 5, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 6, 6,
	170, 10, 6, 13, 6, 14, 6, 171, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 6, 6, 182, 10, 6, 13, 6, 14, 6, 183, 3, 6, 3, 6, 3, 6, 5, 6, 189,
	10, 6, 3, 6, 2, 2, 7, 2, 4, 6, 8, 10, 2, 2, 2, 222, 2, 12, 3, 2, 2, 2,
	4, 23, 3, 2, 2, 2, 6, 128, 3, 2, 2, 2, 8, 133, 3, 2, 2, 2, 10, 188, 3,
	2, 2, 2, 12, 13, 7, 26, 2, 2, 13, 14, 7, 30, 2, 2, 14, 15, 7, 26, 2, 2,
	15, 16, 5, 4, 3, 2, 16, 17, 7, 27, 2, 2, 17, 18, 7, 27, 2, 2, 18, 19, 7,
	2, 2, 3, 19, 3, 3, 2, 2, 2, 20, 22, 7, 33, 2, 2, 21, 20, 3, 2, 2, 2, 22,
	25, 3, 2, 2, 2, 23, 21, 3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 26, 3, 2, 2,
	2, 25, 23, 3, 2, 2, 2, 26, 30, 7, 31, 2, 2, 27, 29, 7, 32, 2, 2, 28, 27,
	3, 2, 2, 2, 29, 32, 3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2,
	31, 36, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2, 33, 35, 5, 6, 4, 2, 34, 33, 3,
	2, 2, 2, 35, 38, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37,
	5, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 39, 41, 7, 33, 2, 2, 40, 39, 3, 2, 2,
	2, 41, 44, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 45,
	3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 45, 57, 7, 34, 2, 2, 46, 48, 7, 26, 2,
	2, 47, 49, 7, 38, 2, 2, 48, 47, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 53,
	3, 2, 2, 2, 50, 52, 5, 8, 5, 2, 51, 50, 3, 2, 2, 2, 52, 55, 3, 2, 2, 2,
	53, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 56, 3, 2, 2, 2, 55, 53, 3,
	2, 2, 2, 56, 58, 7, 27, 2, 2, 57, 46, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58,
	129, 3, 2, 2, 2, 59, 61, 7, 33, 2, 2, 60, 59, 3, 2, 2, 2, 61, 64, 3, 2,
	2, 2, 62, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 65, 3, 2, 2, 2, 64, 62,
	3, 2, 2, 2, 65, 77, 7, 35, 2, 2, 66, 68, 7, 26, 2, 2, 67, 69, 7, 38, 2,
	2, 68, 67, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 73, 3, 2, 2, 2, 70, 72,
	5, 8, 5, 2, 71, 70, 3, 2, 2, 2, 72, 75, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2,
	73, 74, 3, 2, 2, 2, 74, 76, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 76, 78, 7,
	27, 2, 2, 77, 66, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 129, 3, 2, 2, 2,
	79, 81, 7, 33, 2, 2, 80, 79, 3, 2, 2, 2, 81, 84, 3, 2, 2, 2, 82, 80, 3,
	2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 85, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 85,
	100, 7, 37, 2, 2, 86, 88, 7, 26, 2, 2, 87, 89, 7, 38, 2, 2, 88, 87, 3,
	2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 91, 3, 2, 2, 2, 90, 92, 7, 39, 2, 2, 91,
	90, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 96, 3, 2, 2, 2, 93, 95, 5, 10,
	6, 2, 94, 93, 3, 2, 2, 2, 95, 98, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 96, 97,
	3, 2, 2, 2, 97, 99, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 99, 101, 7, 27, 2,
	2, 100, 86, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 129, 3, 2, 2, 2, 102,
	104, 7, 33, 2, 2, 103, 102, 3, 2, 2, 2, 104, 107, 3, 2, 2, 2, 105, 103,
	3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 108, 3, 2, 2, 2, 107, 105, 3, 2,
	2, 2, 108, 123, 7, 36, 2, 2, 109, 111, 7, 26, 2, 2, 110, 112, 7, 38, 2,
	2, 111, 110, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 114, 3, 2, 2, 2, 113,
	115, 7, 39, 2, 2, 114, 113, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 119,
	3, 2, 2, 2, 116, 118, 5, 10, 6, 2, 117, 116, 3, 2, 2, 2, 118, 121, 3, 2,
	2, 2, 119, 117, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 122, 3, 2, 2, 2,
	121, 119, 3, 2, 2, 2, 122, 124, 7, 27, 2, 2, 123, 109, 3, 2, 2, 2, 123,
	124, 3, 2, 2, 2, 124, 129, 3, 2, 2, 2, 125, 129, 7, 49, 2, 2, 126, 129,
	7, 50, 2, 2, 127, 129, 7, 51, 2, 2, 128, 42, 3, 2, 2, 2, 128, 62, 3, 2,
	2, 2, 128, 82, 3, 2, 2, 2, 128, 105, 3, 2, 2, 2, 128, 125, 3, 2, 2, 2,
	128, 126, 3, 2, 2, 2, 128, 127, 3, 2, 2, 2, 129, 7, 3, 2, 2, 2, 130, 132,
	7, 33, 2, 2, 131, 130, 3, 2, 2, 2, 132, 135, 3, 2, 2, 2, 133, 131, 3, 2,
	2, 2, 133, 134, 3, 2, 2, 2, 134, 136, 3, 2, 2, 2, 135, 133, 3, 2, 2, 2,
	136, 142, 7, 40, 2, 2, 137, 138, 7, 26, 2, 2, 138, 139, 7, 39, 2, 2, 139,
	140, 5, 10, 6, 2, 140, 141, 7, 27, 2, 2, 141, 143, 3, 2, 2, 2, 142, 137,
	3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 9, 3, 2, 2, 2, 144, 145, 7, 41,
	2, 2, 145, 146, 7, 26, 2, 2, 146, 147, 7, 42, 2, 2, 147, 189, 7, 27, 2,
	2, 148, 149, 7, 41, 2, 2, 149, 150, 7, 26, 2, 2, 150, 151, 7, 43, 2, 2,
	151, 189, 7, 27, 2, 2, 152, 153, 7, 41, 2, 2, 153, 154, 7, 26, 2, 2, 154,
	155, 7, 44, 2, 2, 155, 189, 7, 27, 2, 2, 156, 157, 7, 41, 2, 2, 157, 158,
	7, 26, 2, 2, 158, 159, 7, 45, 2, 2, 159, 189, 7, 27, 2, 2, 160, 161, 7,
	41, 2, 2, 161, 162, 7, 26, 2, 2, 162, 163, 7, 46, 2, 2, 163, 189, 7, 27,
	2, 2, 164, 165, 7, 41, 2, 2, 165, 166, 7, 26, 2, 2, 166, 167, 7, 47, 2,
	2, 167, 169, 7, 26, 2, 2, 168, 170, 5, 10, 6, 2, 169, 168, 3, 2, 2, 2,
	170, 171, 3, 2, 2, 2, 171, 169, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2, 172,
	173, 3, 2, 2, 2, 173, 174, 7, 27, 2, 2, 174, 175, 7, 27, 2, 2, 175, 189,
	3, 2, 2, 2, 176, 177, 7, 41, 2, 2, 177, 178, 7, 26, 2, 2, 178, 179, 7,
	48, 2, 2, 179, 181, 7, 26, 2, 2, 180, 182, 5, 10, 6, 2, 181, 180, 3, 2,
	2, 2, 182, 183, 3, 2, 2, 2, 183, 181, 3, 2, 2, 2, 183, 184, 3, 2, 2, 2,
	184, 185, 3, 2, 2, 2, 185, 186, 7, 27, 2, 2, 186, 187, 7, 27, 2, 2, 187,
	189, 3, 2, 2, 2, 188, 144, 3, 2, 2, 2, 188, 148, 3, 2, 2, 2, 188, 152,
	3, 2, 2, 2, 188, 156, 3, 2, 2, 2, 188, 160, 3, 2, 2, 2, 188, 164, 3, 2,
	2, 2, 188, 176, 3, 2, 2, 2, 189, 11, 3, 2, 2, 2, 29, 23, 30, 36, 42, 48,
	53, 57, 62, 68, 73, 77, 82, 88, 91, 96, 100, 105, 111, 114, 119, 123, 128,
	133, 142, 171, 183, 188,
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
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ADLWalkerAnnotation {
		{
			p.SetState(18)
			p.Match(ADLWalkerAnnotation)
		}

		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(24)
		p.Match(ADLWalkerModule)
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ADLWalkerImport {
		{
			p.SetState(25)
			p.Match(ADLWalkerImport)
		}

		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(ADLWalkerAnnotation-31))|(1<<(ADLWalkerStruct-31))|(1<<(ADLWalkerUnion-31))|(1<<(ADLWalkerNewtype-31))|(1<<(ADLWalkerType-31))|(1<<(ADLWalkerModuleAnno-31))|(1<<(ADLWalkerDeclAnno-31))|(1<<(ADLWalkerFieldAnno-31)))) != 0 {
		{
			p.SetState(31)
			p.Tld()
		}

		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

type PkgNodeContext struct {
	*TldContext
}

func NewPkgNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PkgNodeContext {
	var p = new(PkgNodeContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

func (s *PkgNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PkgNodeContext) Struct() antlr.TerminalNode {
	return s.GetToken(ADLWalkerStruct, 0)
}

func (s *PkgNodeContext) AllAnnotation() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerAnnotation)
}

func (s *PkgNodeContext) Annotation(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerAnnotation, i)
}

func (s *PkgNodeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *PkgNodeContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *PkgNodeContext) TypeParam() antlr.TerminalNode {
	return s.GetToken(ADLWalkerTypeParam, 0)
}

func (s *PkgNodeContext) AllNameBody() []INameBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameBodyContext)(nil)).Elem())
	var tst = make([]INameBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameBodyContext)
		}
	}

	return tst
}

func (s *PkgNodeContext) NameBody(i int) INameBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameBodyContext)
}

func (s *PkgNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterPkgNode(s)
	}
}

func (s *PkgNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitPkgNode(s)
	}
}

type ExtNodeContext struct {
	*TldContext
}

func NewExtNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExtNodeContext {
	var p = new(ExtNodeContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

func (s *ExtNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtNodeContext) ModuleAnno() antlr.TerminalNode {
	return s.GetToken(ADLWalkerModuleAnno, 0)
}

func (s *ExtNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterExtNode(s)
	}
}

func (s *ExtNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitExtNode(s)
	}
}

type FieldAnnoNodeContext struct {
	*TldContext
}

func NewFieldAnnoNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldAnnoNodeContext {
	var p = new(FieldAnnoNodeContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

func (s *FieldAnnoNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldAnnoNodeContext) FieldAnno() antlr.TerminalNode {
	return s.GetToken(ADLWalkerFieldAnno, 0)
}

func (s *FieldAnnoNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterFieldAnnoNode(s)
	}
}

func (s *FieldAnnoNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitFieldAnnoNode(s)
	}
}

type DeclAnnoNodeContext struct {
	*TldContext
}

func NewDeclAnnoNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclAnnoNodeContext {
	var p = new(DeclAnnoNodeContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

func (s *DeclAnnoNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclAnnoNodeContext) DeclAnno() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDeclAnno, 0)
}

func (s *DeclAnnoNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterDeclAnnoNode(s)
	}
}

func (s *DeclAnnoNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitDeclAnnoNode(s)
	}
}

type ImportNodeContext struct {
	*TldContext
}

func NewImportNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImportNodeContext {
	var p = new(ImportNodeContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

func (s *ImportNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportNodeContext) Union() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUnion, 0)
}

func (s *ImportNodeContext) AllAnnotation() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerAnnotation)
}

func (s *ImportNodeContext) Annotation(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerAnnotation, i)
}

func (s *ImportNodeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *ImportNodeContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *ImportNodeContext) TypeParam() antlr.TerminalNode {
	return s.GetToken(ADLWalkerTypeParam, 0)
}

func (s *ImportNodeContext) AllNameBody() []INameBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameBodyContext)(nil)).Elem())
	var tst = make([]INameBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameBodyContext)
		}
	}

	return tst
}

func (s *ImportNodeContext) NameBody(i int) INameBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameBodyContext)
}

func (s *ImportNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterImportNode(s)
	}
}

func (s *ImportNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitImportNode(s)
	}
}

type MsgNodeContext struct {
	*TldContext
}

func NewMsgNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MsgNodeContext {
	var p = new(MsgNodeContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

func (s *MsgNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MsgNodeContext) Type() antlr.TerminalNode {
	return s.GetToken(ADLWalkerType, 0)
}

func (s *MsgNodeContext) AllAnnotation() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerAnnotation)
}

func (s *MsgNodeContext) Annotation(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerAnnotation, i)
}

func (s *MsgNodeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *MsgNodeContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *MsgNodeContext) TypeParam() antlr.TerminalNode {
	return s.GetToken(ADLWalkerTypeParam, 0)
}

func (s *MsgNodeContext) TypeExpr() antlr.TerminalNode {
	return s.GetToken(ADLWalkerTypeExpr, 0)
}

func (s *MsgNodeContext) AllJsonVal() []IJsonValContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IJsonValContext)(nil)).Elem())
	var tst = make([]IJsonValContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonValContext)
		}
	}

	return tst
}

func (s *MsgNodeContext) JsonVal(i int) IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *MsgNodeContext) Newtype() antlr.TerminalNode {
	return s.GetToken(ADLWalkerNewtype, 0)
}

func (s *MsgNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterMsgNode(s)
	}
}

func (s *MsgNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitMsgNode(s)
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

	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPkgNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLWalkerAnnotation {
			{
				p.SetState(37)
				p.Match(ADLWalkerAnnotation)
			}

			p.SetState(42)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(43)
			p.Match(ADLWalkerStruct)
		}
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerDOWN {
			{
				p.SetState(44)
				p.Match(ADLWalkerDOWN)
			}
			p.SetState(46)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ADLWalkerTypeParam {
				{
					p.SetState(45)
					p.Match(ADLWalkerTypeParam)
				}

			}
			p.SetState(51)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLWalkerAnnotation || _la == ADLWalkerField {
				{
					p.SetState(48)
					p.NameBody()
				}

				p.SetState(53)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(54)
				p.Match(ADLWalkerUP)
			}

		}

	case 2:
		localctx = NewImportNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLWalkerAnnotation {
			{
				p.SetState(57)
				p.Match(ADLWalkerAnnotation)
			}

			p.SetState(62)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(63)
			p.Match(ADLWalkerUnion)
		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerDOWN {
			{
				p.SetState(64)
				p.Match(ADLWalkerDOWN)
			}
			p.SetState(66)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ADLWalkerTypeParam {
				{
					p.SetState(65)
					p.Match(ADLWalkerTypeParam)
				}

			}
			p.SetState(71)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLWalkerAnnotation || _la == ADLWalkerField {
				{
					p.SetState(68)
					p.NameBody()
				}

				p.SetState(73)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(74)
				p.Match(ADLWalkerUP)
			}

		}

	case 3:
		localctx = NewMsgNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLWalkerAnnotation {
			{
				p.SetState(77)
				p.Match(ADLWalkerAnnotation)
			}

			p.SetState(82)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(83)
			p.Match(ADLWalkerType)
		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerDOWN {
			{
				p.SetState(84)
				p.Match(ADLWalkerDOWN)
			}
			p.SetState(86)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ADLWalkerTypeParam {
				{
					p.SetState(85)
					p.Match(ADLWalkerTypeParam)
				}

			}
			p.SetState(89)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ADLWalkerTypeExpr {
				{
					p.SetState(88)
					p.Match(ADLWalkerTypeExpr)
				}

			}
			p.SetState(94)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLWalkerJson {
				{
					p.SetState(91)
					p.JsonVal()
				}

				p.SetState(96)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(97)
				p.Match(ADLWalkerUP)
			}

		}

	case 4:
		localctx = NewMsgNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLWalkerAnnotation {
			{
				p.SetState(100)
				p.Match(ADLWalkerAnnotation)
			}

			p.SetState(105)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(106)
			p.Match(ADLWalkerNewtype)
		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLWalkerDOWN {
			{
				p.SetState(107)
				p.Match(ADLWalkerDOWN)
			}
			p.SetState(109)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ADLWalkerTypeParam {
				{
					p.SetState(108)
					p.Match(ADLWalkerTypeParam)
				}

			}
			p.SetState(112)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ADLWalkerTypeExpr {
				{
					p.SetState(111)
					p.Match(ADLWalkerTypeExpr)
				}

			}
			p.SetState(117)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ADLWalkerJson {
				{
					p.SetState(114)
					p.JsonVal()
				}

				p.SetState(119)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(120)
				p.Match(ADLWalkerUP)
			}

		}

	case 5:
		localctx = NewExtNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(123)
			p.Match(ADLWalkerModuleAnno)
		}

	case 6:
		localctx = NewDeclAnnoNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(124)
			p.Match(ADLWalkerDeclAnno)
		}

	case 7:
		localctx = NewFieldAnnoNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(125)
			p.Match(ADLWalkerFieldAnno)
		}

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

func (s *NameBodyContext) Field() antlr.TerminalNode {
	return s.GetToken(ADLWalkerField, 0)
}

func (s *NameBodyContext) AllAnnotation() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerAnnotation)
}

func (s *NameBodyContext) Annotation(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerAnnotation, i)
}

func (s *NameBodyContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, 0)
}

func (s *NameBodyContext) TypeExpr() antlr.TerminalNode {
	return s.GetToken(ADLWalkerTypeExpr, 0)
}

func (s *NameBodyContext) JsonVal() IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *NameBodyContext) UP() antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, 0)
}

func (s *NameBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterNameBody(s)
	}
}

func (s *NameBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitNameBody(s)
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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ADLWalkerAnnotation {
		{
			p.SetState(128)
			p.Match(ADLWalkerAnnotation)
		}

		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(134)
		p.Match(ADLWalkerField)
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ADLWalkerDOWN {
		{
			p.SetState(135)
			p.Match(ADLWalkerDOWN)
		}
		{
			p.SetState(136)
			p.Match(ADLWalkerTypeExpr)
		}
		{
			p.SetState(137)
			p.JsonVal()
		}
		{
			p.SetState(138)
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

func (s *JsonValContext) Json() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJson, 0)
}

func (s *JsonValContext) AllDOWN() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerDOWN)
}

func (s *JsonValContext) DOWN(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerDOWN, i)
}

func (s *JsonValContext) JsonStr() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJsonStr, 0)
}

func (s *JsonValContext) AllUP() []antlr.TerminalNode {
	return s.GetTokens(ADLWalkerUP)
}

func (s *JsonValContext) UP(i int) antlr.TerminalNode {
	return s.GetToken(ADLWalkerUP, i)
}

func (s *JsonValContext) JsonBool() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJsonBool, 0)
}

func (s *JsonValContext) JsonNull() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJsonNull, 0)
}

func (s *JsonValContext) JsonInt() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJsonInt, 0)
}

func (s *JsonValContext) JsonFloat() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJsonFloat, 0)
}

func (s *JsonValContext) JsonArray() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJsonArray, 0)
}

func (s *JsonValContext) AllJsonVal() []IJsonValContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IJsonValContext)(nil)).Elem())
	var tst = make([]IJsonValContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonValContext)
		}
	}

	return tst
}

func (s *JsonValContext) JsonVal(i int) IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *JsonValContext) JsonObj() antlr.TerminalNode {
	return s.GetToken(ADLWalkerJsonObj, 0)
}

func (s *JsonValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JsonValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.EnterJsonVal(s)
	}
}

func (s *JsonValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLWalkerListener); ok {
		listenerT.ExitJsonVal(s)
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

	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(142)
			p.Match(ADLWalkerJson)
		}
		{
			p.SetState(143)
			p.Match(ADLWalkerDOWN)
		}
		{
			p.SetState(144)
			p.Match(ADLWalkerJsonStr)
		}
		{
			p.SetState(145)
			p.Match(ADLWalkerUP)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(146)
			p.Match(ADLWalkerJson)
		}
		{
			p.SetState(147)
			p.Match(ADLWalkerDOWN)
		}
		{
			p.SetState(148)
			p.Match(ADLWalkerJsonBool)
		}
		{
			p.SetState(149)
			p.Match(ADLWalkerUP)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(150)
			p.Match(ADLWalkerJson)
		}
		{
			p.SetState(151)
			p.Match(ADLWalkerDOWN)
		}
		{
			p.SetState(152)
			p.Match(ADLWalkerJsonNull)
		}
		{
			p.SetState(153)
			p.Match(ADLWalkerUP)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(154)
			p.Match(ADLWalkerJson)
		}
		{
			p.SetState(155)
			p.Match(ADLWalkerDOWN)
		}
		{
			p.SetState(156)
			p.Match(ADLWalkerJsonInt)
		}
		{
			p.SetState(157)
			p.Match(ADLWalkerUP)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(158)
			p.Match(ADLWalkerJson)
		}
		{
			p.SetState(159)
			p.Match(ADLWalkerDOWN)
		}
		{
			p.SetState(160)
			p.Match(ADLWalkerJsonFloat)
		}
		{
			p.SetState(161)
			p.Match(ADLWalkerUP)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(162)
			p.Match(ADLWalkerJson)
		}
		{
			p.SetState(163)
			p.Match(ADLWalkerDOWN)
		}
		{
			p.SetState(164)
			p.Match(ADLWalkerJsonArray)
		}
		{
			p.SetState(165)
			p.Match(ADLWalkerDOWN)
		}
		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ADLWalkerJson {
			{
				p.SetState(166)
				p.JsonVal()
			}

			p.SetState(169)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(171)
			p.Match(ADLWalkerUP)
		}
		{
			p.SetState(172)
			p.Match(ADLWalkerUP)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(174)
			p.Match(ADLWalkerJson)
		}
		{
			p.SetState(175)
			p.Match(ADLWalkerDOWN)
		}
		{
			p.SetState(176)
			p.Match(ADLWalkerJsonObj)
		}
		{
			p.SetState(177)
			p.Match(ADLWalkerDOWN)
		}
		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ADLWalkerJson {
			{
				p.SetState(178)
				p.JsonVal()
			}

			p.SetState(181)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(183)
			p.Match(ADLWalkerUP)
		}
		{
			p.SetState(184)
			p.Match(ADLWalkerUP)
		}

	}

	return localctx
}
