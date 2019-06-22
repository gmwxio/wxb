// Code generated from DNAC_B_Walker.g4 by ANTLR 4.7.1. DO NOT EDIT.

package dnacb // DNAC_B_Walker
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 56, 75, 4,
	2, 9, 2, 4, 3, 9, 3, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	3, 3, 3, 3, 3, 7, 3, 18, 10, 3, 12, 3, 14, 3, 21, 11, 3, 3, 3, 5, 3, 24,
	10, 3, 3, 3, 7, 3, 27, 10, 3, 12, 3, 14, 3, 30, 11, 3, 3, 3, 5, 3, 33,
	10, 3, 3, 3, 3, 3, 3, 3, 7, 3, 38, 10, 3, 12, 3, 14, 3, 41, 11, 3, 3, 3,
	5, 3, 44, 10, 3, 3, 3, 5, 3, 47, 10, 3, 3, 3, 5, 3, 50, 10, 3, 3, 3, 5,
	3, 53, 10, 3, 3, 3, 3, 3, 3, 3, 7, 3, 58, 10, 3, 12, 3, 14, 3, 61, 11,
	3, 3, 3, 5, 3, 64, 10, 3, 3, 3, 5, 3, 67, 10, 3, 3, 3, 5, 3, 70, 10, 3,
	3, 3, 5, 3, 73, 10, 3, 3, 3, 2, 2, 4, 2, 4, 2, 2, 2, 88, 2, 6, 3, 2, 2,
	2, 4, 72, 3, 2, 2, 2, 6, 7, 7, 26, 2, 2, 7, 8, 7, 53, 2, 2, 8, 9, 7, 26,
	2, 2, 9, 10, 5, 4, 3, 2, 10, 11, 7, 27, 2, 2, 11, 12, 7, 27, 2, 2, 12,
	13, 7, 2, 2, 3, 13, 3, 3, 2, 2, 2, 14, 32, 7, 54, 2, 2, 15, 19, 7, 26,
	2, 2, 16, 18, 7, 33, 2, 2, 17, 16, 3, 2, 2, 2, 18, 21, 3, 2, 2, 2, 19,
	17, 3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 23, 3, 2, 2, 2, 21, 19, 3, 2, 2,
	2, 22, 24, 7, 38, 2, 2, 23, 22, 3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 28,
	3, 2, 2, 2, 25, 27, 5, 4, 3, 2, 26, 25, 3, 2, 2, 2, 27, 30, 3, 2, 2, 2,
	28, 26, 3, 2, 2, 2, 28, 29, 3, 2, 2, 2, 29, 31, 3, 2, 2, 2, 30, 28, 3,
	2, 2, 2, 31, 33, 7, 27, 2, 2, 32, 15, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33,
	73, 3, 2, 2, 2, 34, 52, 7, 37, 2, 2, 35, 39, 7, 26, 2, 2, 36, 38, 7, 33,
	2, 2, 37, 36, 3, 2, 2, 2, 38, 41, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 39, 40,
	3, 2, 2, 2, 40, 43, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 42, 44, 7, 38, 2, 2,
	43, 42, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 46, 3, 2, 2, 2, 45, 47, 7,
	39, 2, 2, 46, 45, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 49, 3, 2, 2, 2, 48,
	50, 7, 56, 2, 2, 49, 48, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 51, 3, 2,
	2, 2, 51, 53, 7, 27, 2, 2, 52, 35, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53,
	73, 3, 2, 2, 2, 54, 69, 7, 41, 2, 2, 55, 59, 7, 26, 2, 2, 56, 58, 7, 33,
	2, 2, 57, 56, 3, 2, 2, 2, 58, 61, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60,
	3, 2, 2, 2, 60, 63, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 62, 64, 7, 39, 2, 2,
	63, 62, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 66, 3, 2, 2, 2, 65, 67, 7,
	56, 2, 2, 66, 65, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68,
	70, 7, 27, 2, 2, 69, 55, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 73, 3, 2,
	2, 2, 71, 73, 7, 55, 2, 2, 72, 14, 3, 2, 2, 2, 72, 34, 3, 2, 2, 2, 72,
	54, 3, 2, 2, 2, 72, 71, 3, 2, 2, 2, 73, 5, 3, 2, 2, 2, 16, 19, 23, 28,
	32, 39, 43, 46, 49, 52, 59, 63, 66, 69, 72,
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
	"dnac", "tld",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type DNAC_B_Walker struct {
	*antlr.BaseParser
}

func NewDNAC_B_Walker(input antlr.TokenStream) *DNAC_B_Walker {
	this := new(DNAC_B_Walker)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "DNAC_B_Walker.g4"

	return this
}

// DNAC_B_Walker tokens.
const (
	DNAC_B_WalkerEOF          = antlr.TokenEOF
	DNAC_B_WalkerLCUR         = 1
	DNAC_B_WalkerRCUR         = 2
	DNAC_B_WalkerLSQ          = 3
	DNAC_B_WalkerRSQ          = 4
	DNAC_B_WalkerEQ           = 5
	DNAC_B_WalkerDQ           = 6
	DNAC_B_WalkerSQ           = 7
	DNAC_B_WalkerSEMI         = 8
	DNAC_B_WalkerDCOLON       = 9
	DNAC_B_WalkerCOLON        = 10
	DNAC_B_WalkerDOT          = 11
	DNAC_B_WalkerCOMMA        = 12
	DNAC_B_WalkerLCHEVR       = 13
	DNAC_B_WalkerRCHEVR       = 14
	DNAC_B_WalkerSTAR         = 15
	DNAC_B_WalkerAT           = 16
	DNAC_B_WalkerSTR          = 17
	DNAC_B_WalkerID           = 18
	DNAC_B_WalkerINT          = 19
	DNAC_B_WalkerFLT          = 20
	DNAC_B_WalkerWS           = 21
	DNAC_B_WalkerLINE_DOC     = 22
	DNAC_B_WalkerLINE_COMMENT = 23
	DNAC_B_WalkerDOWN         = 24
	DNAC_B_WalkerUP           = 25
	DNAC_B_WalkerROOT         = 26
	DNAC_B_WalkerERROR        = 27
	DNAC_B_WalkerADL          = 28
	DNAC_B_WalkerModule       = 29
	DNAC_B_WalkerImport       = 30
	DNAC_B_WalkerAnnotation   = 31
	DNAC_B_WalkerStruct       = 32
	DNAC_B_WalkerUnion        = 33
	DNAC_B_WalkerNewtype      = 34
	DNAC_B_WalkerType         = 35
	DNAC_B_WalkerTypeParam    = 36
	DNAC_B_WalkerTypeExpr     = 37
	DNAC_B_WalkerTypeExprElem = 38
	DNAC_B_WalkerField        = 39
	DNAC_B_WalkerJson         = 40
	DNAC_B_WalkerJsonStr      = 41
	DNAC_B_WalkerJsonBool     = 42
	DNAC_B_WalkerJsonNull     = 43
	DNAC_B_WalkerJsonInt      = 44
	DNAC_B_WalkerJsonFloat    = 45
	DNAC_B_WalkerJsonArray    = 46
	DNAC_B_WalkerJsonObj      = 47
	DNAC_B_WalkerModuleAnno   = 48
	DNAC_B_WalkerDeclAnno     = 49
	DNAC_B_WalkerFieldAnno    = 50
	DNAC_B_WalkerDNAC         = 51
	DNAC_B_WalkerName         = 52
	DNAC_B_WalkerExnotation   = 53
	DNAC_B_WalkerDefault      = 54
)

// DNAC_B_Walker rules.
const (
	DNAC_B_WalkerRULE_dnac = 0
	DNAC_B_WalkerRULE_tld  = 1
)

// IDnacContext is an interface to support dynamic dispatch.
type IDnacContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDnacContext differentiates from other interfaces.
	IsDnacContext()
}

type DnacContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDnacContext() *DnacContext {
	var p = new(DnacContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DNAC_B_WalkerRULE_dnac
	return p
}

func (*DnacContext) IsDnacContext() {}

func NewDnacContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DnacContext {
	var p = new(DnacContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DNAC_B_WalkerRULE_dnac

	return p
}

func (s *DnacContext) GetParser() antlr.Parser { return s.parser }

func (s *DnacContext) AllDOWN() []antlr.TerminalNode {
	return s.GetTokens(DNAC_B_WalkerDOWN)
}

func (s *DnacContext) DOWN(i int) antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerDOWN, i)
}

func (s *DnacContext) DNAC() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerDNAC, 0)
}

func (s *DnacContext) Tld() ITldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITldContext)
}

func (s *DnacContext) AllUP() []antlr.TerminalNode {
	return s.GetTokens(DNAC_B_WalkerUP)
}

func (s *DnacContext) UP(i int) antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerUP, i)
}

func (s *DnacContext) EOF() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerEOF, 0)
}

func (s *DnacContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DnacContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DnacContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_B_WalkerListener); ok {
		listenerT.EnterDnac(s)
	}
}

func (s *DnacContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_B_WalkerListener); ok {
		listenerT.ExitDnac(s)
	}
}

func (s *DnacContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DNAC_B_WalkerVisitor:
		return t.VisitDnac(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DNAC_B_Walker) Dnac() (localctx IDnacContext) {
	localctx = NewDnacContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DNAC_B_WalkerRULE_dnac)

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
		p.SetState(4)
		p.Match(DNAC_B_WalkerDOWN)
	}
	{
		p.SetState(5)
		p.Match(DNAC_B_WalkerDNAC)
	}
	{
		p.SetState(6)
		p.Match(DNAC_B_WalkerDOWN)
	}
	{
		p.SetState(7)
		p.Tld()
	}
	{
		p.SetState(8)
		p.Match(DNAC_B_WalkerUP)
	}
	{
		p.SetState(9)
		p.Match(DNAC_B_WalkerUP)
	}
	{
		p.SetState(10)
		p.Match(DNAC_B_WalkerEOF)
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
	p.RuleIndex = DNAC_B_WalkerRULE_tld
	return p
}

func (*TldContext) IsTldContext() {}

func NewTldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TldContext {
	var p = new(TldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DNAC_B_WalkerRULE_tld

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
	return s.GetToken(DNAC_B_WalkerName, 0)
}

func (s *NameNodeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerDOWN, 0)
}

func (s *NameNodeContext) UP() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerUP, 0)
}

func (s *NameNodeContext) AllAnnotation() []antlr.TerminalNode {
	return s.GetTokens(DNAC_B_WalkerAnnotation)
}

func (s *NameNodeContext) Annotation(i int) antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerAnnotation, i)
}

func (s *NameNodeContext) TypeParam() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerTypeParam, 0)
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
	if listenerT, ok := listener.(DNAC_B_WalkerListener); ok {
		listenerT.EnterNameNode(s)
	}
}

func (s *NameNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_B_WalkerListener); ok {
		listenerT.ExitNameNode(s)
	}
}

func (s *NameNodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DNAC_B_WalkerVisitor:
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
	return s.GetToken(DNAC_B_WalkerType, 0)
}

func (s *TypeNodeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerDOWN, 0)
}

func (s *TypeNodeContext) UP() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerUP, 0)
}

func (s *TypeNodeContext) AllAnnotation() []antlr.TerminalNode {
	return s.GetTokens(DNAC_B_WalkerAnnotation)
}

func (s *TypeNodeContext) Annotation(i int) antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerAnnotation, i)
}

func (s *TypeNodeContext) TypeParam() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerTypeParam, 0)
}

func (s *TypeNodeContext) TypeExpr() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerTypeExpr, 0)
}

func (s *TypeNodeContext) Default() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerDefault, 0)
}

func (s *TypeNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_B_WalkerListener); ok {
		listenerT.EnterTypeNode(s)
	}
}

func (s *TypeNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_B_WalkerListener); ok {
		listenerT.ExitTypeNode(s)
	}
}

func (s *TypeNodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DNAC_B_WalkerVisitor:
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
	return s.GetToken(DNAC_B_WalkerExnotation, 0)
}

func (s *ExnotationNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_B_WalkerListener); ok {
		listenerT.EnterExnotationNode(s)
	}
}

func (s *ExnotationNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_B_WalkerListener); ok {
		listenerT.ExitExnotationNode(s)
	}
}

func (s *ExnotationNodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DNAC_B_WalkerVisitor:
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
	return s.GetToken(DNAC_B_WalkerField, 0)
}

func (s *NameBodyNodeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerDOWN, 0)
}

func (s *NameBodyNodeContext) UP() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerUP, 0)
}

func (s *NameBodyNodeContext) AllAnnotation() []antlr.TerminalNode {
	return s.GetTokens(DNAC_B_WalkerAnnotation)
}

func (s *NameBodyNodeContext) Annotation(i int) antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerAnnotation, i)
}

func (s *NameBodyNodeContext) TypeExpr() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerTypeExpr, 0)
}

func (s *NameBodyNodeContext) Default() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerDefault, 0)
}

func (s *NameBodyNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_B_WalkerListener); ok {
		listenerT.EnterNameBodyNode(s)
	}
}

func (s *NameBodyNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DNAC_B_WalkerListener); ok {
		listenerT.ExitNameBodyNode(s)
	}
}

func (s *NameBodyNodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DNAC_B_WalkerVisitor:
		return t.VisitNameBodyNode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DNAC_B_Walker) Tld() (localctx ITldContext) {
	localctx = NewTldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DNAC_B_WalkerRULE_tld)
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

	p.SetState(70)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DNAC_B_WalkerName:
		localctx = NewNameNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(12)
			p.Match(DNAC_B_WalkerName)
		}
		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DNAC_B_WalkerDOWN {
			{
				p.SetState(13)
				p.Match(DNAC_B_WalkerDOWN)
			}
			p.SetState(17)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == DNAC_B_WalkerAnnotation {
				{
					p.SetState(14)
					p.Match(DNAC_B_WalkerAnnotation)
				}

				p.SetState(19)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(21)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == DNAC_B_WalkerTypeParam {
				{
					p.SetState(20)
					p.Match(DNAC_B_WalkerTypeParam)
				}

			}
			p.SetState(26)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(DNAC_B_WalkerType-35))|(1<<(DNAC_B_WalkerField-35))|(1<<(DNAC_B_WalkerName-35))|(1<<(DNAC_B_WalkerExnotation-35)))) != 0 {
				{
					p.SetState(23)
					p.Tld()
				}

				p.SetState(28)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(29)
				p.Match(DNAC_B_WalkerUP)
			}

		}

	case DNAC_B_WalkerType:
		localctx = NewTypeNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(32)
			p.Match(DNAC_B_WalkerType)
		}
		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DNAC_B_WalkerDOWN {
			{
				p.SetState(33)
				p.Match(DNAC_B_WalkerDOWN)
			}
			p.SetState(37)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == DNAC_B_WalkerAnnotation {
				{
					p.SetState(34)
					p.Match(DNAC_B_WalkerAnnotation)
				}

				p.SetState(39)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(41)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == DNAC_B_WalkerTypeParam {
				{
					p.SetState(40)
					p.Match(DNAC_B_WalkerTypeParam)
				}

			}
			p.SetState(44)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == DNAC_B_WalkerTypeExpr {
				{
					p.SetState(43)
					p.Match(DNAC_B_WalkerTypeExpr)
				}

			}
			p.SetState(47)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == DNAC_B_WalkerDefault {
				{
					p.SetState(46)
					p.Match(DNAC_B_WalkerDefault)
				}

			}
			{
				p.SetState(49)
				p.Match(DNAC_B_WalkerUP)
			}

		}

	case DNAC_B_WalkerField:
		localctx = NewNameBodyNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(52)
			p.Match(DNAC_B_WalkerField)
		}
		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DNAC_B_WalkerDOWN {
			{
				p.SetState(53)
				p.Match(DNAC_B_WalkerDOWN)
			}
			p.SetState(57)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == DNAC_B_WalkerAnnotation {
				{
					p.SetState(54)
					p.Match(DNAC_B_WalkerAnnotation)
				}

				p.SetState(59)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(61)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == DNAC_B_WalkerTypeExpr {
				{
					p.SetState(60)
					p.Match(DNAC_B_WalkerTypeExpr)
				}

			}
			p.SetState(64)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == DNAC_B_WalkerDefault {
				{
					p.SetState(63)
					p.Match(DNAC_B_WalkerDefault)
				}

			}
			{
				p.SetState(66)
				p.Match(DNAC_B_WalkerUP)
			}

		}

	case DNAC_B_WalkerExnotation:
		localctx = NewExnotationNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(69)
			p.Match(DNAC_B_WalkerExnotation)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
