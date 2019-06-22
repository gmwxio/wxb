// Code generated from DNAC_A_Walker.g4 by ANTLR 4.7.1. DO NOT EDIT.

package dnaca // DNAC_A_Walker
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 56, 148,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 7, 3,
	26, 10, 3, 12, 3, 14, 3, 29, 11, 3, 3, 3, 5, 3, 32, 10, 3, 3, 3, 7, 3,
	35, 10, 3, 12, 3, 14, 3, 38, 11, 3, 3, 3, 5, 3, 41, 10, 3, 3, 3, 3, 3,
	3, 3, 7, 3, 46, 10, 3, 12, 3, 14, 3, 49, 11, 3, 3, 3, 5, 3, 52, 10, 3,
	3, 3, 5, 3, 55, 10, 3, 3, 3, 7, 3, 58, 10, 3, 12, 3, 14, 3, 61, 11, 3,
	3, 3, 5, 3, 64, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	7, 3, 74, 10, 3, 12, 3, 14, 3, 77, 11, 3, 3, 3, 5, 3, 80, 10, 3, 3, 3,
	5, 3, 83, 10, 3, 3, 3, 5, 3, 86, 10, 3, 5, 3, 88, 10, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 5, 4, 95, 10, 4, 3, 5, 3, 5, 3, 5, 6, 5, 100, 10, 5, 13,
	5, 14, 5, 101, 3, 5, 3, 5, 5, 5, 106, 10, 5, 3, 6, 3, 6, 3, 6, 6, 6, 111,
	10, 6, 13, 6, 14, 6, 112, 3, 6, 3, 6, 5, 6, 117, 10, 6, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 6, 7, 127, 10, 7, 13, 7, 14, 7, 128, 3,
	7, 3, 7, 5, 7, 133, 10, 7, 3, 7, 3, 7, 3, 7, 6, 7, 138, 10, 7, 13, 7, 14,
	7, 139, 3, 7, 3, 7, 5, 7, 144, 10, 7, 5, 7, 146, 10, 7, 3, 7, 2, 2, 8,
	2, 4, 6, 8, 10, 12, 2, 2, 2, 172, 2, 14, 3, 2, 2, 2, 4, 87, 3, 2, 2, 2,
	6, 89, 3, 2, 2, 2, 8, 96, 3, 2, 2, 2, 10, 107, 3, 2, 2, 2, 12, 145, 3,
	2, 2, 2, 14, 15, 7, 26, 2, 2, 15, 16, 7, 30, 2, 2, 16, 17, 7, 26, 2, 2,
	17, 18, 5, 4, 3, 2, 18, 19, 7, 27, 2, 2, 19, 20, 7, 27, 2, 2, 20, 21, 7,
	2, 2, 3, 21, 3, 3, 2, 2, 2, 22, 40, 7, 54, 2, 2, 23, 27, 7, 26, 2, 2, 24,
	26, 5, 6, 4, 2, 25, 24, 3, 2, 2, 2, 26, 29, 3, 2, 2, 2, 27, 25, 3, 2, 2,
	2, 27, 28, 3, 2, 2, 2, 28, 31, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 30, 32,
	7, 38, 2, 2, 31, 30, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 36, 3, 2, 2, 2,
	33, 35, 5, 4, 3, 2, 34, 33, 3, 2, 2, 2, 35, 38, 3, 2, 2, 2, 36, 34, 3,
	2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 39, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 39,
	41, 7, 27, 2, 2, 40, 23, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 88, 3, 2,
	2, 2, 42, 63, 7, 37, 2, 2, 43, 47, 7, 26, 2, 2, 44, 46, 5, 6, 4, 2, 45,
	44, 3, 2, 2, 2, 46, 49, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2,
	2, 48, 51, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 50, 52, 7, 38, 2, 2, 51, 50,
	3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 54, 3, 2, 2, 2, 53, 55, 5, 8, 5, 2,
	54, 53, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 59, 3, 2, 2, 2, 56, 58, 5,
	12, 7, 2, 57, 56, 3, 2, 2, 2, 58, 61, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59,
	60, 3, 2, 2, 2, 60, 62, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 62, 64, 7, 27,
	2, 2, 63, 43, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 88, 3, 2, 2, 2, 65, 66,
	7, 55, 2, 2, 66, 67, 7, 26, 2, 2, 67, 68, 5, 12, 7, 2, 68, 69, 7, 27, 2,
	2, 69, 88, 3, 2, 2, 2, 70, 85, 7, 41, 2, 2, 71, 75, 7, 26, 2, 2, 72, 74,
	5, 6, 4, 2, 73, 72, 3, 2, 2, 2, 74, 77, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2,
	75, 76, 3, 2, 2, 2, 76, 79, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 78, 80, 5,
	8, 5, 2, 79, 78, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 82, 3, 2, 2, 2, 81,
	83, 5, 12, 7, 2, 82, 81, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 84, 3, 2,
	2, 2, 84, 86, 7, 27, 2, 2, 85, 71, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86,
	88, 3, 2, 2, 2, 87, 22, 3, 2, 2, 2, 87, 42, 3, 2, 2, 2, 87, 65, 3, 2, 2,
	2, 87, 70, 3, 2, 2, 2, 88, 5, 3, 2, 2, 2, 89, 94, 7, 33, 2, 2, 90, 91,
	7, 26, 2, 2, 91, 92, 5, 12, 7, 2, 92, 93, 7, 27, 2, 2, 93, 95, 3, 2, 2,
	2, 94, 90, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 7, 3, 2, 2, 2, 96, 105,
	7, 39, 2, 2, 97, 99, 7, 26, 2, 2, 98, 100, 5, 10, 6, 2, 99, 98, 3, 2, 2,
	2, 100, 101, 3, 2, 2, 2, 101, 99, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102,
	103, 3, 2, 2, 2, 103, 104, 7, 27, 2, 2, 104, 106, 3, 2, 2, 2, 105, 97,
	3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 9, 3, 2, 2, 2, 107, 116, 7, 40,
	2, 2, 108, 110, 7, 26, 2, 2, 109, 111, 5, 10, 6, 2, 110, 109, 3, 2, 2,
	2, 111, 112, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113,
	114, 3, 2, 2, 2, 114, 115, 7, 27, 2, 2, 115, 117, 3, 2, 2, 2, 116, 108,
	3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 11, 3, 2, 2, 2, 118, 146, 7, 43,
	2, 2, 119, 146, 7, 44, 2, 2, 120, 146, 7, 45, 2, 2, 121, 146, 7, 46, 2,
	2, 122, 146, 7, 47, 2, 2, 123, 132, 7, 48, 2, 2, 124, 126, 7, 26, 2, 2,
	125, 127, 5, 12, 7, 2, 126, 125, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128,
	126, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 131,
	7, 27, 2, 2, 131, 133, 3, 2, 2, 2, 132, 124, 3, 2, 2, 2, 132, 133, 3, 2,
	2, 2, 133, 146, 3, 2, 2, 2, 134, 143, 7, 49, 2, 2, 135, 137, 7, 26, 2,
	2, 136, 138, 5, 12, 7, 2, 137, 136, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139,
	137, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 142,
	7, 27, 2, 2, 142, 144, 3, 2, 2, 2, 143, 135, 3, 2, 2, 2, 143, 144, 3, 2,
	2, 2, 144, 146, 3, 2, 2, 2, 145, 118, 3, 2, 2, 2, 145, 119, 3, 2, 2, 2,
	145, 120, 3, 2, 2, 2, 145, 121, 3, 2, 2, 2, 145, 122, 3, 2, 2, 2, 145,
	123, 3, 2, 2, 2, 145, 134, 3, 2, 2, 2, 146, 13, 3, 2, 2, 2, 26, 27, 31,
	36, 40, 47, 51, 54, 59, 63, 75, 79, 82, 85, 87, 94, 101, 105, 112, 116,
	128, 132, 139, 143, 145,
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
	"TypeParam", "TypeExpr", "TypeExprElem", "Field", "Json", "JsonStr", "JsonBool",
	"JsonNull", "JsonInt", "JsonFloat", "JsonArray", "JsonObj", "ModuleAnno",
	"DeclAnno", "FieldAnno", "DNAC", "Name", "Exnotation", "Default",
}

var ruleNames = []string{
	"adl", "tld", "annotation", "typeExpr_", "typeExprElem_", "jsonVal",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type DNAC_A_Walker struct {
	*antlr.BaseParser
}

func NewDNAC_A_Walker(input antlr.TokenStream) *DNAC_A_Walker {
	this := new(DNAC_A_Walker)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "DNAC_A_Walker.g4"

	return this
}

// DNAC_A_Walker tokens.
const (
	DNAC_A_WalkerEOF          = antlr.TokenEOF
	DNAC_A_WalkerLCUR         = 1
	DNAC_A_WalkerRCUR         = 2
	DNAC_A_WalkerLSQ          = 3
	DNAC_A_WalkerRSQ          = 4
	DNAC_A_WalkerEQ           = 5
	DNAC_A_WalkerDQ           = 6
	DNAC_A_WalkerSQ           = 7
	DNAC_A_WalkerSEMI         = 8
	DNAC_A_WalkerDCOLON       = 9
	DNAC_A_WalkerCOLON        = 10
	DNAC_A_WalkerDOT          = 11
	DNAC_A_WalkerCOMMA        = 12
	DNAC_A_WalkerLCHEVR       = 13
	DNAC_A_WalkerRCHEVR       = 14
	DNAC_A_WalkerSTAR         = 15
	DNAC_A_WalkerAT           = 16
	DNAC_A_WalkerSTR          = 17
	DNAC_A_WalkerID           = 18
	DNAC_A_WalkerINT          = 19
	DNAC_A_WalkerFLT          = 20
	DNAC_A_WalkerWS           = 21
	DNAC_A_WalkerLINE_DOC     = 22
	DNAC_A_WalkerLINE_COMMENT = 23
	DNAC_A_WalkerDOWN         = 24
	DNAC_A_WalkerUP           = 25
	DNAC_A_WalkerROOT         = 26
	DNAC_A_WalkerERROR        = 27
	DNAC_A_WalkerADL          = 28
	DNAC_A_WalkerModule       = 29
	DNAC_A_WalkerImport       = 30
	DNAC_A_WalkerAnnotation   = 31
	DNAC_A_WalkerStruct       = 32
	DNAC_A_WalkerUnion        = 33
	DNAC_A_WalkerNewtype      = 34
	DNAC_A_WalkerType         = 35
	DNAC_A_WalkerTypeParam    = 36
	DNAC_A_WalkerTypeExpr     = 37
	DNAC_A_WalkerTypeExprElem = 38
	DNAC_A_WalkerField        = 39
	DNAC_A_WalkerJson         = 40
	DNAC_A_WalkerJsonStr      = 41
	DNAC_A_WalkerJsonBool     = 42
	DNAC_A_WalkerJsonNull     = 43
	DNAC_A_WalkerJsonInt      = 44
	DNAC_A_WalkerJsonFloat    = 45
	DNAC_A_WalkerJsonArray    = 46
	DNAC_A_WalkerJsonObj      = 47
	DNAC_A_WalkerModuleAnno   = 48
	DNAC_A_WalkerDeclAnno     = 49
	DNAC_A_WalkerFieldAnno    = 50
	DNAC_A_WalkerDNAC         = 51
	DNAC_A_WalkerName         = 52
	DNAC_A_WalkerExnotation   = 53
	DNAC_A_WalkerDefault      = 54
)

// DNAC_A_Walker rules.
const (
	DNAC_A_WalkerRULE_adl           = 0
	DNAC_A_WalkerRULE_tld           = 1
	DNAC_A_WalkerRULE_annotation    = 2
	DNAC_A_WalkerRULE_typeExpr_     = 3
	DNAC_A_WalkerRULE_typeExprElem_ = 4
	DNAC_A_WalkerRULE_jsonVal       = 5
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
	p.RuleIndex = DNAC_A_WalkerRULE_adl
	return p
}

func (*AdlContext) IsAdlContext() {}

func NewAdlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdlContext {
	var p = new(AdlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DNAC_A_WalkerRULE_adl

	return p
}

func (s *AdlContext) GetParser() antlr.Parser { return s.parser }

func (s *AdlContext) AllDOWN() []antlr.TerminalNode {
	return s.GetTokens(DNAC_A_WalkerDOWN)
}

func (s *AdlContext) DOWN(i int) antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerDOWN, i)
}

func (s *AdlContext) ADL() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerADL, 0)
}

func (s *AdlContext) Tld() ITldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITldContext)
}

func (s *AdlContext) AllUP() []antlr.TerminalNode {
	return s.GetTokens(DNAC_A_WalkerUP)
}

func (s *AdlContext) UP(i int) antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerUP, i)
}

func (s *AdlContext) EOF() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerEOF, 0)
}

func (s *AdlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.EnterAdl(s)
	}
}

func (s *AdlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.ExitAdl(s)
	}
}

func (s *AdlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DNAC_A_WalkerVisitor:
		return t.VisitAdl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DNAC_A_Walker) Adl() (localctx IAdlContext) {
	localctx = NewAdlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DNAC_A_WalkerRULE_adl)

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
		p.Match(DNAC_A_WalkerDOWN)
	}
	{
		p.SetState(13)
		p.Match(DNAC_A_WalkerADL)
	}
	{
		p.SetState(14)
		p.Match(DNAC_A_WalkerDOWN)
	}
	{
		p.SetState(15)
		p.Tld()
	}
	{
		p.SetState(16)
		p.Match(DNAC_A_WalkerUP)
	}
	{
		p.SetState(17)
		p.Match(DNAC_A_WalkerUP)
	}
	{
		p.SetState(18)
		p.Match(DNAC_A_WalkerEOF)
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
	p.RuleIndex = DNAC_A_WalkerRULE_tld
	return p
}

func (*TldContext) IsTldContext() {}

func NewTldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TldContext {
	var p = new(TldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DNAC_A_WalkerRULE_tld

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

type NameNodeContext struct {
	*TldContext
}

func NewNameNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NameNodeContext {
	var p = new(NameNodeContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

func (s *NameNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameNodeContext) Name() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerName, 0)
}

func (s *NameNodeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerDOWN, 0)
}

func (s *NameNodeContext) UP() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerUP, 0)
}

func (s *NameNodeContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *NameNodeContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *NameNodeContext) TypeParam() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerTypeParam, 0)
}

func (s *NameNodeContext) AllTld() []ITldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITldContext)(nil)).Elem())
	var tst = make([]ITldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITldContext)
		}
	}

	return tst
}

func (s *NameNodeContext) Tld(i int) ITldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITldContext)
}

func (s *NameNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.EnterNameNode(s)
	}
}

func (s *NameNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.ExitNameNode(s)
	}
}

func (s *NameNodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DNAC_A_WalkerVisitor:
		return t.VisitNameNode(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeNodeContext struct {
	*TldContext
}

func NewTypeNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeNodeContext {
	var p = new(TypeNodeContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

func (s *TypeNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNodeContext) Type() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerType, 0)
}

func (s *TypeNodeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerDOWN, 0)
}

func (s *TypeNodeContext) UP() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerUP, 0)
}

func (s *TypeNodeContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *TypeNodeContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *TypeNodeContext) TypeParam() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerTypeParam, 0)
}

func (s *TypeNodeContext) TypeExpr_() ITypeExpr_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeExpr_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeExpr_Context)
}

func (s *TypeNodeContext) AllJsonVal() []IJsonValContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IJsonValContext)(nil)).Elem())
	var tst = make([]IJsonValContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonValContext)
		}
	}

	return tst
}

func (s *TypeNodeContext) JsonVal(i int) IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *TypeNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.EnterTypeNode(s)
	}
}

func (s *TypeNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.ExitTypeNode(s)
	}
}

func (s *TypeNodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DNAC_A_WalkerVisitor:
		return t.VisitTypeNode(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExnotationNodeContext struct {
	*TldContext
}

func NewExnotationNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExnotationNodeContext {
	var p = new(ExnotationNodeContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

func (s *ExnotationNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExnotationNodeContext) Exnotation() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerExnotation, 0)
}

func (s *ExnotationNodeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerDOWN, 0)
}

func (s *ExnotationNodeContext) JsonVal() IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *ExnotationNodeContext) UP() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerUP, 0)
}

func (s *ExnotationNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.EnterExnotationNode(s)
	}
}

func (s *ExnotationNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.ExitExnotationNode(s)
	}
}

func (s *ExnotationNodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DNAC_A_WalkerVisitor:
		return t.VisitExnotationNode(s)

	default:
		return t.VisitChildren(s)
	}
}

type NameBodyNodeContext struct {
	*TldContext
}

func NewNameBodyNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NameBodyNodeContext {
	var p = new(NameBodyNodeContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

func (s *NameBodyNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameBodyNodeContext) Field() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerField, 0)
}

func (s *NameBodyNodeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerDOWN, 0)
}

func (s *NameBodyNodeContext) UP() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerUP, 0)
}

func (s *NameBodyNodeContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *NameBodyNodeContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *NameBodyNodeContext) TypeExpr_() ITypeExpr_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeExpr_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeExpr_Context)
}

func (s *NameBodyNodeContext) JsonVal() IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *NameBodyNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.EnterNameBodyNode(s)
	}
}

func (s *NameBodyNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.ExitNameBodyNode(s)
	}
}

func (s *NameBodyNodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DNAC_A_WalkerVisitor:
		return t.VisitNameBodyNode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DNAC_A_Walker) Tld() (localctx ITldContext) {
	localctx = NewTldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DNAC_A_WalkerRULE_tld)
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

	p.SetState(85)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DNAC_A_WalkerName:
		localctx = NewNameNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(20)
			p.Match(DNAC_A_WalkerName)
		}
		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DNAC_A_WalkerDOWN {
			{
				p.SetState(21)
				p.Match(DNAC_A_WalkerDOWN)
			}
			p.SetState(25)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == DNAC_A_WalkerAnnotation {
				{
					p.SetState(22)
					p.Annotation()
				}

				p.SetState(27)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(29)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == DNAC_A_WalkerTypeParam {
				{
					p.SetState(28)
					p.Match(DNAC_A_WalkerTypeParam)
				}

			}
			p.SetState(34)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(DNAC_A_WalkerType-35))|(1<<(DNAC_A_WalkerField-35))|(1<<(DNAC_A_WalkerName-35))|(1<<(DNAC_A_WalkerExnotation-35)))) != 0 {
				{
					p.SetState(31)
					p.Tld()
				}

				p.SetState(36)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(37)
				p.Match(DNAC_A_WalkerUP)
			}

		}

	case DNAC_A_WalkerType:
		localctx = NewTypeNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(40)
			p.Match(DNAC_A_WalkerType)
		}
		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DNAC_A_WalkerDOWN {
			{
				p.SetState(41)
				p.Match(DNAC_A_WalkerDOWN)
			}
			p.SetState(45)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == DNAC_A_WalkerAnnotation {
				{
					p.SetState(42)
					p.Annotation()
				}

				p.SetState(47)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(49)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == DNAC_A_WalkerTypeParam {
				{
					p.SetState(48)
					p.Match(DNAC_A_WalkerTypeParam)
				}

			}
			p.SetState(52)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == DNAC_A_WalkerTypeExpr {
				{
					p.SetState(51)
					p.TypeExpr_()
				}

			}
			p.SetState(57)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(DNAC_A_WalkerJsonStr-41))|(1<<(DNAC_A_WalkerJsonBool-41))|(1<<(DNAC_A_WalkerJsonNull-41))|(1<<(DNAC_A_WalkerJsonInt-41))|(1<<(DNAC_A_WalkerJsonFloat-41))|(1<<(DNAC_A_WalkerJsonArray-41))|(1<<(DNAC_A_WalkerJsonObj-41)))) != 0 {
				{
					p.SetState(54)
					p.JsonVal()
				}

				p.SetState(59)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(60)
				p.Match(DNAC_A_WalkerUP)
			}

		}

	case DNAC_A_WalkerExnotation:
		localctx = NewExnotationNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(63)
			p.Match(DNAC_A_WalkerExnotation)
		}
		{
			p.SetState(64)
			p.Match(DNAC_A_WalkerDOWN)
		}
		{
			p.SetState(65)
			p.JsonVal()
		}
		{
			p.SetState(66)
			p.Match(DNAC_A_WalkerUP)
		}

	case DNAC_A_WalkerField:
		localctx = NewNameBodyNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(68)
			p.Match(DNAC_A_WalkerField)
		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DNAC_A_WalkerDOWN {
			{
				p.SetState(69)
				p.Match(DNAC_A_WalkerDOWN)
			}
			p.SetState(73)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == DNAC_A_WalkerAnnotation {
				{
					p.SetState(70)
					p.Annotation()
				}

				p.SetState(75)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(77)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == DNAC_A_WalkerTypeExpr {
				{
					p.SetState(76)
					p.TypeExpr_()
				}

			}
			p.SetState(80)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(DNAC_A_WalkerJsonStr-41))|(1<<(DNAC_A_WalkerJsonBool-41))|(1<<(DNAC_A_WalkerJsonNull-41))|(1<<(DNAC_A_WalkerJsonInt-41))|(1<<(DNAC_A_WalkerJsonFloat-41))|(1<<(DNAC_A_WalkerJsonArray-41))|(1<<(DNAC_A_WalkerJsonObj-41)))) != 0 {
				{
					p.SetState(79)
					p.JsonVal()
				}

			}
			{
				p.SetState(82)
				p.Match(DNAC_A_WalkerUP)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAnnotationContext is an interface to support dynamic dispatch.
type IAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnnotationContext differentiates from other interfaces.
	IsAnnotationContext()
}

type AnnotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotationContext() *AnnotationContext {
	var p = new(AnnotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DNAC_A_WalkerRULE_annotation
	return p
}

func (*AnnotationContext) IsAnnotationContext() {}

func NewAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationContext {
	var p = new(AnnotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DNAC_A_WalkerRULE_annotation

	return p
}

func (s *AnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationContext) Annotation() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerAnnotation, 0)
}

func (s *AnnotationContext) DOWN() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerDOWN, 0)
}

func (s *AnnotationContext) JsonVal() IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *AnnotationContext) UP() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerUP, 0)
}

func (s *AnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.EnterAnnotation(s)
	}
}

func (s *AnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.ExitAnnotation(s)
	}
}

func (s *AnnotationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DNAC_A_WalkerVisitor:
		return t.VisitAnnotation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DNAC_A_Walker) Annotation() (localctx IAnnotationContext) {
	localctx = NewAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DNAC_A_WalkerRULE_annotation)
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
		p.SetState(87)
		p.Match(DNAC_A_WalkerAnnotation)
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DNAC_A_WalkerDOWN {
		{
			p.SetState(88)
			p.Match(DNAC_A_WalkerDOWN)
		}
		{
			p.SetState(89)
			p.JsonVal()
		}
		{
			p.SetState(90)
			p.Match(DNAC_A_WalkerUP)
		}

	}

	return localctx
}

// ITypeExpr_Context is an interface to support dynamic dispatch.
type ITypeExpr_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeExpr_Context differentiates from other interfaces.
	IsTypeExpr_Context()
}

type TypeExpr_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeExpr_Context() *TypeExpr_Context {
	var p = new(TypeExpr_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DNAC_A_WalkerRULE_typeExpr_
	return p
}

func (*TypeExpr_Context) IsTypeExpr_Context() {}

func NewTypeExpr_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeExpr_Context {
	var p = new(TypeExpr_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DNAC_A_WalkerRULE_typeExpr_

	return p
}

func (s *TypeExpr_Context) GetParser() antlr.Parser { return s.parser }

func (s *TypeExpr_Context) TypeExpr() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerTypeExpr, 0)
}

func (s *TypeExpr_Context) DOWN() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerDOWN, 0)
}

func (s *TypeExpr_Context) UP() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerUP, 0)
}

func (s *TypeExpr_Context) AllTypeExprElem_() []ITypeExprElem_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeExprElem_Context)(nil)).Elem())
	var tst = make([]ITypeExprElem_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeExprElem_Context)
		}
	}

	return tst
}

func (s *TypeExpr_Context) TypeExprElem_(i int) ITypeExprElem_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeExprElem_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeExprElem_Context)
}

func (s *TypeExpr_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExpr_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeExpr_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.EnterTypeExpr_(s)
	}
}

func (s *TypeExpr_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.ExitTypeExpr_(s)
	}
}

func (s *TypeExpr_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DNAC_A_WalkerVisitor:
		return t.VisitTypeExpr_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DNAC_A_Walker) TypeExpr_() (localctx ITypeExpr_Context) {
	localctx = NewTypeExpr_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DNAC_A_WalkerRULE_typeExpr_)
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
		p.SetState(94)
		p.Match(DNAC_A_WalkerTypeExpr)
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DNAC_A_WalkerDOWN {
		{
			p.SetState(95)
			p.Match(DNAC_A_WalkerDOWN)
		}
		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DNAC_A_WalkerTypeExprElem {
			{
				p.SetState(96)
				p.TypeExprElem_()
			}

			p.SetState(99)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(101)
			p.Match(DNAC_A_WalkerUP)
		}

	}

	return localctx
}

// ITypeExprElem_Context is an interface to support dynamic dispatch.
type ITypeExprElem_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeExprElem_Context differentiates from other interfaces.
	IsTypeExprElem_Context()
}

type TypeExprElem_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeExprElem_Context() *TypeExprElem_Context {
	var p = new(TypeExprElem_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DNAC_A_WalkerRULE_typeExprElem_
	return p
}

func (*TypeExprElem_Context) IsTypeExprElem_Context() {}

func NewTypeExprElem_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeExprElem_Context {
	var p = new(TypeExprElem_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DNAC_A_WalkerRULE_typeExprElem_

	return p
}

func (s *TypeExprElem_Context) GetParser() antlr.Parser { return s.parser }

func (s *TypeExprElem_Context) CopyFrom(ctx *TypeExprElem_Context) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeExprElem_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExprElem_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypeParamsContext struct {
	*TypeExprElem_Context
}

func NewTypeParamsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeParamsContext {
	var p = new(TypeParamsContext)

	p.TypeExprElem_Context = NewEmptyTypeExprElem_Context()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeExprElem_Context))

	return p
}

func (s *TypeParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeParamsContext) TypeExprElem() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerTypeExprElem, 0)
}

func (s *TypeParamsContext) DOWN() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerDOWN, 0)
}

func (s *TypeParamsContext) UP() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerUP, 0)
}

func (s *TypeParamsContext) AllTypeExprElem_() []ITypeExprElem_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeExprElem_Context)(nil)).Elem())
	var tst = make([]ITypeExprElem_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeExprElem_Context)
		}
	}

	return tst
}

func (s *TypeParamsContext) TypeExprElem_(i int) ITypeExprElem_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeExprElem_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeExprElem_Context)
}

func (s *TypeParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.EnterTypeParams(s)
	}
}

func (s *TypeParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.ExitTypeParams(s)
	}
}

func (s *TypeParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DNAC_A_WalkerVisitor:
		return t.VisitTypeParams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DNAC_A_Walker) TypeExprElem_() (localctx ITypeExprElem_Context) {
	localctx = NewTypeExprElem_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DNAC_A_WalkerRULE_typeExprElem_)
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

	localctx = NewTypeParamsContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(DNAC_A_WalkerTypeExprElem)
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DNAC_A_WalkerDOWN {
		{
			p.SetState(106)
			p.Match(DNAC_A_WalkerDOWN)
		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DNAC_A_WalkerTypeExprElem {
			{
				p.SetState(107)
				p.TypeExprElem_()
			}

			p.SetState(110)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(112)
			p.Match(DNAC_A_WalkerUP)
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
	p.RuleIndex = DNAC_A_WalkerRULE_jsonVal
	return p
}

func (*JsonValContext) IsJsonValContext() {}

func NewJsonValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonValContext {
	var p = new(JsonValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DNAC_A_WalkerRULE_jsonVal

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

func (s *JsonStrContext) JsonStr() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerJsonStr, 0)
}

func (s *JsonStrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.EnterJsonStr(s)
	}
}

func (s *JsonStrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.ExitJsonStr(s)
	}
}

func (s *JsonStrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DNAC_A_WalkerVisitor:
		return t.VisitJsonStr(s)

	default:
		return t.VisitChildren(s)
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

func (s *JsonArrayContext) JsonArray() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerJsonArray, 0)
}

func (s *JsonArrayContext) DOWN() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerDOWN, 0)
}

func (s *JsonArrayContext) UP() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerUP, 0)
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
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.EnterJsonArray(s)
	}
}

func (s *JsonArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.ExitJsonArray(s)
	}
}

func (s *JsonArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DNAC_A_WalkerVisitor:
		return t.VisitJsonArray(s)

	default:
		return t.VisitChildren(s)
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

func (s *JsonFloatContext) JsonFloat() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerJsonFloat, 0)
}

func (s *JsonFloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.EnterJsonFloat(s)
	}
}

func (s *JsonFloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.ExitJsonFloat(s)
	}
}

func (s *JsonFloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DNAC_A_WalkerVisitor:
		return t.VisitJsonFloat(s)

	default:
		return t.VisitChildren(s)
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

func (s *JsonObjContext) JsonObj() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerJsonObj, 0)
}

func (s *JsonObjContext) DOWN() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerDOWN, 0)
}

func (s *JsonObjContext) UP() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerUP, 0)
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
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.EnterJsonObj(s)
	}
}

func (s *JsonObjContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.ExitJsonObj(s)
	}
}

func (s *JsonObjContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DNAC_A_WalkerVisitor:
		return t.VisitJsonObj(s)

	default:
		return t.VisitChildren(s)
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

func (s *JsonBoolContext) JsonBool() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerJsonBool, 0)
}

func (s *JsonBoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.EnterJsonBool(s)
	}
}

func (s *JsonBoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.ExitJsonBool(s)
	}
}

func (s *JsonBoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DNAC_A_WalkerVisitor:
		return t.VisitJsonBool(s)

	default:
		return t.VisitChildren(s)
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

func (s *JsonIntContext) JsonInt() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerJsonInt, 0)
}

func (s *JsonIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.EnterJsonInt(s)
	}
}

func (s *JsonIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.ExitJsonInt(s)
	}
}

func (s *JsonIntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DNAC_A_WalkerVisitor:
		return t.VisitJsonInt(s)

	default:
		return t.VisitChildren(s)
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

func (s *JsonNullContext) JsonNull() antlr.TerminalNode {
	return s.GetToken(DNAC_A_WalkerJsonNull, 0)
}

func (s *JsonNullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.EnterJsonNull(s)
	}
}

func (s *JsonNullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_A_WalkerListener); ok {
		listenerT.ExitJsonNull(s)
	}
}

func (s *JsonNullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DNAC_A_WalkerVisitor:
		return t.VisitJsonNull(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DNAC_A_Walker) JsonVal() (localctx IJsonValContext) {
	localctx = NewJsonValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DNAC_A_WalkerRULE_jsonVal)
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

	p.SetState(143)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DNAC_A_WalkerJsonStr:
		localctx = NewJsonStrContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.Match(DNAC_A_WalkerJsonStr)
		}

	case DNAC_A_WalkerJsonBool:
		localctx = NewJsonBoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.Match(DNAC_A_WalkerJsonBool)
		}

	case DNAC_A_WalkerJsonNull:
		localctx = NewJsonNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(118)
			p.Match(DNAC_A_WalkerJsonNull)
		}

	case DNAC_A_WalkerJsonInt:
		localctx = NewJsonIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(119)
			p.Match(DNAC_A_WalkerJsonInt)
		}

	case DNAC_A_WalkerJsonFloat:
		localctx = NewJsonFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(120)
			p.Match(DNAC_A_WalkerJsonFloat)
		}

	case DNAC_A_WalkerJsonArray:
		localctx = NewJsonArrayContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(121)
			p.Match(DNAC_A_WalkerJsonArray)
		}
		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DNAC_A_WalkerDOWN {
			{
				p.SetState(122)
				p.Match(DNAC_A_WalkerDOWN)
			}
			p.SetState(124)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = (((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(DNAC_A_WalkerJsonStr-41))|(1<<(DNAC_A_WalkerJsonBool-41))|(1<<(DNAC_A_WalkerJsonNull-41))|(1<<(DNAC_A_WalkerJsonInt-41))|(1<<(DNAC_A_WalkerJsonFloat-41))|(1<<(DNAC_A_WalkerJsonArray-41))|(1<<(DNAC_A_WalkerJsonObj-41)))) != 0) {
				{
					p.SetState(123)
					p.JsonVal()
				}

				p.SetState(126)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(128)
				p.Match(DNAC_A_WalkerUP)
			}

		}

	case DNAC_A_WalkerJsonObj:
		localctx = NewJsonObjContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(132)
			p.Match(DNAC_A_WalkerJsonObj)
		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DNAC_A_WalkerDOWN {
			{
				p.SetState(133)
				p.Match(DNAC_A_WalkerDOWN)
			}
			p.SetState(135)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = (((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(DNAC_A_WalkerJsonStr-41))|(1<<(DNAC_A_WalkerJsonBool-41))|(1<<(DNAC_A_WalkerJsonNull-41))|(1<<(DNAC_A_WalkerJsonInt-41))|(1<<(DNAC_A_WalkerJsonFloat-41))|(1<<(DNAC_A_WalkerJsonArray-41))|(1<<(DNAC_A_WalkerJsonObj-41)))) != 0) {
				{
					p.SetState(134)
					p.JsonVal()
				}

				p.SetState(137)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(139)
				p.Match(DNAC_A_WalkerUP)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
