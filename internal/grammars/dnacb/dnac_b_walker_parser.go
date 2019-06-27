// Generated from DNAC_B_Walker.g4 by ANTLR 4.7.

package dnacb // DNAC_B_Walker
import (
	"fmt"
	"reflect"
	"strconv"
)
import "github.com/wxio/goantlr"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 59, 116,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 7, 3, 20, 10, 3, 12, 3, 14, 3, 23, 11, 3,
	3, 3, 7, 3, 26, 10, 3, 12, 3, 14, 3, 29, 11, 3, 3, 3, 5, 3, 32, 10, 3,
	3, 4, 3, 4, 3, 4, 7, 4, 37, 10, 4, 12, 4, 14, 4, 40, 11, 4, 3, 4, 5, 4,
	43, 10, 4, 3, 4, 7, 4, 46, 10, 4, 12, 4, 14, 4, 49, 11, 4, 3, 4, 5, 4,
	52, 10, 4, 3, 4, 3, 4, 3, 4, 7, 4, 57, 10, 4, 12, 4, 14, 4, 60, 11, 4,
	3, 4, 3, 4, 5, 4, 64, 10, 4, 3, 4, 5, 4, 67, 10, 4, 3, 4, 3, 4, 3, 4, 7,
	4, 72, 10, 4, 12, 4, 14, 4, 75, 11, 4, 3, 4, 5, 4, 78, 10, 4, 3, 4, 5,
	4, 81, 10, 4, 3, 4, 3, 4, 3, 4, 7, 4, 86, 10, 4, 12, 4, 14, 4, 89, 11,
	4, 3, 4, 3, 4, 5, 4, 93, 10, 4, 3, 4, 3, 4, 3, 4, 7, 4, 98, 10, 4, 12,
	4, 14, 4, 101, 11, 4, 3, 4, 5, 4, 104, 10, 4, 3, 4, 5, 4, 107, 10, 4, 3,
	4, 3, 4, 3, 4, 5, 4, 112, 10, 4, 5, 4, 114, 10, 4, 3, 4, 2, 2, 5, 2, 4,
	6, 2, 2, 2, 136, 2, 8, 3, 2, 2, 2, 4, 16, 3, 2, 2, 2, 6, 113, 3, 2, 2,
	2, 8, 9, 7, 26, 2, 2, 9, 10, 7, 56, 2, 2, 10, 11, 7, 26, 2, 2, 11, 12,
	5, 4, 3, 2, 12, 13, 7, 27, 2, 2, 13, 14, 7, 27, 2, 2, 14, 15, 7, 2, 2,
	3, 15, 3, 3, 2, 2, 2, 16, 31, 7, 57, 2, 2, 17, 21, 7, 26, 2, 2, 18, 20,
	7, 34, 2, 2, 19, 18, 3, 2, 2, 2, 20, 23, 3, 2, 2, 2, 21, 19, 3, 2, 2, 2,
	21, 22, 3, 2, 2, 2, 22, 27, 3, 2, 2, 2, 23, 21, 3, 2, 2, 2, 24, 26, 5,
	6, 4, 2, 25, 24, 3, 2, 2, 2, 26, 29, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 27,
	28, 3, 2, 2, 2, 28, 30, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 30, 32, 7, 27,
	2, 2, 31, 17, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 5, 3, 2, 2, 2, 33, 51,
	7, 57, 2, 2, 34, 38, 7, 26, 2, 2, 35, 37, 7, 34, 2, 2, 36, 35, 3, 2, 2,
	2, 37, 40, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 42,
	3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 41, 43, 7, 41, 2, 2, 42, 41, 3, 2, 2, 2,
	42, 43, 3, 2, 2, 2, 43, 47, 3, 2, 2, 2, 44, 46, 5, 6, 4, 2, 45, 44, 3,
	2, 2, 2, 46, 49, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48,
	50, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 50, 52, 7, 27, 2, 2, 51, 34, 3, 2,
	2, 2, 51, 52, 3, 2, 2, 2, 52, 114, 3, 2, 2, 2, 53, 66, 7, 40, 2, 2, 54,
	58, 7, 26, 2, 2, 55, 57, 7, 34, 2, 2, 56, 55, 3, 2, 2, 2, 57, 60, 3, 2,
	2, 2, 58, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 61, 3, 2, 2, 2, 60, 58,
	3, 2, 2, 2, 61, 63, 7, 41, 2, 2, 62, 64, 7, 42, 2, 2, 63, 62, 3, 2, 2,
	2, 63, 64, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 67, 7, 27, 2, 2, 66, 54,
	3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 114, 3, 2, 2, 2, 68, 80, 7, 40, 2,
	2, 69, 73, 7, 26, 2, 2, 70, 72, 7, 34, 2, 2, 71, 70, 3, 2, 2, 2, 72, 75,
	3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 77, 3, 2, 2, 2,
	75, 73, 3, 2, 2, 2, 76, 78, 7, 59, 2, 2, 77, 76, 3, 2, 2, 2, 77, 78, 3,
	2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 81, 7, 27, 2, 2, 80, 69, 3, 2, 2, 2, 80,
	81, 3, 2, 2, 2, 81, 114, 3, 2, 2, 2, 82, 92, 7, 44, 2, 2, 83, 87, 7, 26,
	2, 2, 84, 86, 7, 34, 2, 2, 85, 84, 3, 2, 2, 2, 86, 89, 3, 2, 2, 2, 87,
	85, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 90, 3, 2, 2, 2, 89, 87, 3, 2, 2,
	2, 90, 91, 7, 42, 2, 2, 91, 93, 7, 27, 2, 2, 92, 83, 3, 2, 2, 2, 92, 93,
	3, 2, 2, 2, 93, 114, 3, 2, 2, 2, 94, 106, 7, 44, 2, 2, 95, 99, 7, 26, 2,
	2, 96, 98, 7, 34, 2, 2, 97, 96, 3, 2, 2, 2, 98, 101, 3, 2, 2, 2, 99, 97,
	3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 103, 3, 2, 2, 2, 101, 99, 3, 2, 2,
	2, 102, 104, 7, 59, 2, 2, 103, 102, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104,
	105, 3, 2, 2, 2, 105, 107, 7, 27, 2, 2, 106, 95, 3, 2, 2, 2, 106, 107,
	3, 2, 2, 2, 107, 114, 3, 2, 2, 2, 108, 111, 7, 58, 2, 2, 109, 110, 7, 26,
	2, 2, 110, 112, 7, 27, 2, 2, 111, 109, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2,
	112, 114, 3, 2, 2, 2, 113, 33, 3, 2, 2, 2, 113, 53, 3, 2, 2, 2, 113, 68,
	3, 2, 2, 2, 113, 82, 3, 2, 2, 2, 113, 94, 3, 2, 2, 2, 113, 108, 3, 2, 2,
	2, 114, 7, 3, 2, 2, 2, 22, 21, 27, 31, 38, 42, 47, 51, 58, 63, 66, 73,
	77, 80, 87, 92, 99, 103, 106, 111, 113,
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
	"Module", "ImportModule", "ImportScopedName", "Annotation", "AnnotationNotScoped",
	"AnnotationScoped", "Struct", "Union", "Newtype", "Type", "TypeParam",
	"TypeExpr", "TypeExprElem", "Field", "Json", "JsonStr", "JsonBool", "JsonNull",
	"JsonInt", "JsonFloat", "JsonArray", "JsonObj", "ModuleAnno", "DeclAnno",
	"FieldAnno", "DNAC", "Name", "Exnotation", "Default",
}

var ruleNames = []string{
	"dnac", "name", "tld",
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
	DNAC_B_WalkerEOF                 = antlr.TokenEOF
	DNAC_B_WalkerLCUR                = 1
	DNAC_B_WalkerRCUR                = 2
	DNAC_B_WalkerLSQ                 = 3
	DNAC_B_WalkerRSQ                 = 4
	DNAC_B_WalkerEQ                  = 5
	DNAC_B_WalkerDQ                  = 6
	DNAC_B_WalkerSQ                  = 7
	DNAC_B_WalkerSEMI                = 8
	DNAC_B_WalkerDCOLON              = 9
	DNAC_B_WalkerCOLON               = 10
	DNAC_B_WalkerDOT                 = 11
	DNAC_B_WalkerCOMMA               = 12
	DNAC_B_WalkerLCHEVR              = 13
	DNAC_B_WalkerRCHEVR              = 14
	DNAC_B_WalkerSTAR                = 15
	DNAC_B_WalkerAT                  = 16
	DNAC_B_WalkerSTR                 = 17
	DNAC_B_WalkerID                  = 18
	DNAC_B_WalkerINT                 = 19
	DNAC_B_WalkerFLT                 = 20
	DNAC_B_WalkerWS                  = 21
	DNAC_B_WalkerLINE_DOC            = 22
	DNAC_B_WalkerLINE_COMMENT        = 23
	DNAC_B_WalkerDOWN                = 24
	DNAC_B_WalkerUP                  = 25
	DNAC_B_WalkerROOT                = 26
	DNAC_B_WalkerERROR               = 27
	DNAC_B_WalkerADL                 = 28
	DNAC_B_WalkerModule              = 29
	DNAC_B_WalkerImportModule        = 30
	DNAC_B_WalkerImportScopedName    = 31
	DNAC_B_WalkerAnnotation          = 32
	DNAC_B_WalkerAnnotationNotScoped = 33
	DNAC_B_WalkerAnnotationScoped    = 34
	DNAC_B_WalkerStruct              = 35
	DNAC_B_WalkerUnion               = 36
	DNAC_B_WalkerNewtype             = 37
	DNAC_B_WalkerType                = 38
	DNAC_B_WalkerTypeParam           = 39
	DNAC_B_WalkerTypeExpr            = 40
	DNAC_B_WalkerTypeExprElem        = 41
	DNAC_B_WalkerField               = 42
	DNAC_B_WalkerJson                = 43
	DNAC_B_WalkerJsonStr             = 44
	DNAC_B_WalkerJsonBool            = 45
	DNAC_B_WalkerJsonNull            = 46
	DNAC_B_WalkerJsonInt             = 47
	DNAC_B_WalkerJsonFloat           = 48
	DNAC_B_WalkerJsonArray           = 49
	DNAC_B_WalkerJsonObj             = 50
	DNAC_B_WalkerModuleAnno          = 51
	DNAC_B_WalkerDeclAnno            = 52
	DNAC_B_WalkerFieldAnno           = 53
	DNAC_B_WalkerDNAC                = 54
	DNAC_B_WalkerName                = 55
	DNAC_B_WalkerExnotation          = 56
	DNAC_B_WalkerDefault             = 57
)

// DNAC_B_Walker rules.
const (
	DNAC_B_WalkerRULE_dnac = 0
	DNAC_B_WalkerRULE_name = 1
	DNAC_B_WalkerRULE_tld  = 2
)

type IDnacContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	Name() INameContext
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	DNAC() antlr.TerminalNode
	EOF() antlr.TerminalNode
	//tokenListGetterDecl
	AllDOWN() []antlr.TerminalNode
	AllUP() []antlr.TerminalNode
	//tokenListIndexedGetterDecl
	DOWN(i int) antlr.TerminalNode
	UP(i int) antlr.TerminalNode

	// IsDnacContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
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

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters
func (s *DnacContext) AllDOWN() []antlr.TerminalNode {
	return s.GetTokens(DNAC_B_WalkerDOWN)
}

func (s *DnacContext) DOWN(i int) antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerDOWN, i)
}

func (s *DnacContext) DNAC() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerDNAC, 0)
}

func (s *DnacContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*NameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
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

//provideCopyFrom
func (s *DnacContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DnacContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *DnacContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DnacEntryListener); ok {
		listenerT.EnterDnac(s)
	}
}

func (s *DnacContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DnacExitListener); ok {
		listenerT.ExitDnac(s)
	}
}

func (s *DnacContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*DNAC_B_WalkerHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.Dnac != nil {
		h.Dnac(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *DnacContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case DnacContextVisitor:
		return t.VisitDnac(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//extensionMembers

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
		p.SetState(6)
		p.Match(DNAC_B_WalkerDOWN)
	}
	{
		p.SetState(7)
		p.Match(DNAC_B_WalkerDNAC)
	}
	{
		p.SetState(8)
		p.Match(DNAC_B_WalkerDOWN)
	}
	{
		p.SetState(9)
		p.Name()
	}
	{
		p.SetState(10)
		p.Match(DNAC_B_WalkerUP)
	}
	{
		p.SetState(11)
		p.Match(DNAC_B_WalkerUP)
	}
	{
		p.SetState(12)
		p.Match(DNAC_B_WalkerEOF)
	}

	return localctx
}

type INameContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsNameContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DNAC_B_WalkerRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DNAC_B_WalkerRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *NameContext) CopyFrom(ctx *NameContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type NameNodeContext struct {
	*NameContext
}

func NewNameNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NameNodeContext {
	var p = new(NameNodeContext)

	p.NameContext = NewEmptyNameContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NameContext))

	return p
}

type INameNodeContext interface {
	//Current rule
	INameContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllTld() []ITldContext
	//  ruleListIndexedGetterDecl
	Tld(i int) ITldContext

	//  tokenGetterDecl
	Name() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	//  tokenListGetterDecl
	AllAnnotation() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	Annotation(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*NameNodeContext) IsNameNodeContext() {}

//AltLabelStructDecl tokenDecls

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *NameNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
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

func (s *NameNodeContext) AllTld() []ITldContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*TldContext)(nil)).Elem())
	var tst = make([]ITldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITldContext)
		}
	}

	return tst
}

func (s *NameNodeContext) Tld(i int) ITldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITldContext)
}

func (s *NameNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NameNodeEntryListener); ok {
		listenerT.EnterNameNode(s)
	}
}

func (s *NameNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NameNodeExitListener); ok {
		listenerT.ExitNameNode(s)
	}
}

func (s *NameNodeContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*DNAC_B_WalkerHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.NameNode != nil {
		h.NameNode(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *NameNodeContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case NameNodeContextVisitor:
		return t.VisitNameNode(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *DNAC_B_Walker) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DNAC_B_WalkerRULE_name)
	var //TokenTypeDecl
	_la int

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

	localctx = NewNameNodeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(14)
		p.Match(DNAC_B_WalkerName)
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DNAC_B_WalkerDOWN {
		{
			p.SetState(15)
			p.Match(DNAC_B_WalkerDOWN)
		}
		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == DNAC_B_WalkerAnnotation {
			{
				p.SetState(16)
				p.Match(DNAC_B_WalkerAnnotation)
			}

			p.SetState(21)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(DNAC_B_WalkerType-38))|(1<<(DNAC_B_WalkerField-38))|(1<<(DNAC_B_WalkerName-38))|(1<<(DNAC_B_WalkerExnotation-38)))) != 0 {
			{
				p.SetState(22)
				p.Tld()
			}

			p.SetState(27)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(28)
			p.Match(DNAC_B_WalkerUP)
		}

	}

	return localctx
}

type ITldContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsTldContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
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

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *TldContext) CopyFrom(ctx *TldContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type NameRuleContext struct {
	*TldContext
}

func NewNameRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NameRuleContext {
	var p = new(NameRuleContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

type INameRuleContext interface {
	//Current rule
	ITldContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllTld() []ITldContext
	//  ruleListIndexedGetterDecl
	Tld(i int) ITldContext

	//  tokenGetterDecl
	Name() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	TypeParam() antlr.TerminalNode
	//  tokenListGetterDecl
	AllAnnotation() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	Annotation(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*NameRuleContext) IsNameRuleContext() {}

//AltLabelStructDecl tokenDecls

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *NameRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *NameRuleContext) Name() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerName, 0)
}

func (s *NameRuleContext) DOWN() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerDOWN, 0)
}

func (s *NameRuleContext) UP() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerUP, 0)
}

func (s *NameRuleContext) AllAnnotation() []antlr.TerminalNode {
	return s.GetTokens(DNAC_B_WalkerAnnotation)
}

func (s *NameRuleContext) Annotation(i int) antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerAnnotation, i)
}

func (s *NameRuleContext) TypeParam() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerTypeParam, 0)
}

func (s *NameRuleContext) AllTld() []ITldContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*TldContext)(nil)).Elem())
	var tst = make([]ITldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITldContext)
		}
	}

	return tst
}

func (s *NameRuleContext) Tld(i int) ITldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITldContext)
}

func (s *NameRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NameRuleEntryListener); ok {
		listenerT.EnterNameRule(s)
	}
}

func (s *NameRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NameRuleExitListener); ok {
		listenerT.ExitNameRule(s)
	}
}

func (s *NameRuleContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*DNAC_B_WalkerHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.NameRule != nil {
		h.NameRule(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *NameRuleContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case NameRuleContextVisitor:
		return t.VisitNameRule(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

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

type ITypeNodeContext interface {
	//Current rule
	ITldContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	Type() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	TypeParam() antlr.TerminalNode
	UP() antlr.TerminalNode
	TypeExpr() antlr.TerminalNode
	Default() antlr.TerminalNode
	//  tokenListGetterDecl
	AllAnnotation() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	Annotation(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*TypeNodeContext) IsTypeNodeContext() {}

//AltLabelStructDecl tokenDecls

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *TypeNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *TypeNodeContext) Type() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerType, 0)
}

func (s *TypeNodeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerDOWN, 0)
}

func (s *TypeNodeContext) TypeParam() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerTypeParam, 0)
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

func (s *TypeNodeContext) TypeExpr() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerTypeExpr, 0)
}

func (s *TypeNodeContext) Default() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerDefault, 0)
}

func (s *TypeNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeNodeEntryListener); ok {
		listenerT.EnterTypeNode(s)
	}
}

func (s *TypeNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeNodeExitListener); ok {
		listenerT.ExitTypeNode(s)
	}
}

func (s *TypeNodeContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*DNAC_B_WalkerHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.TypeNode != nil {
		h.TypeNode(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *TypeNodeContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case TypeNodeContextVisitor:
		return t.VisitTypeNode(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

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

type IExnotationNodeContext interface {
	//Current rule
	ITldContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	Exnotation() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*ExnotationNodeContext) IsExnotationNodeContext() {}

//AltLabelStructDecl tokenDecls

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *ExnotationNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *ExnotationNodeContext) Exnotation() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerExnotation, 0)
}

func (s *ExnotationNodeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerDOWN, 0)
}

func (s *ExnotationNodeContext) UP() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerUP, 0)
}

func (s *ExnotationNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExnotationNodeEntryListener); ok {
		listenerT.EnterExnotationNode(s)
	}
}

func (s *ExnotationNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExnotationNodeExitListener); ok {
		listenerT.ExitExnotationNode(s)
	}
}

func (s *ExnotationNodeContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*DNAC_B_WalkerHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.ExnotationNode != nil {
		h.ExnotationNode(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *ExnotationNodeContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case ExnotationNodeContextVisitor:
		return t.VisitExnotationNode(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

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

type INameBodyNodeContext interface {
	//Current rule
	ITldContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	Field() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	TypeExpr() antlr.TerminalNode
	UP() antlr.TerminalNode
	Default() antlr.TerminalNode
	//  tokenListGetterDecl
	AllAnnotation() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	Annotation(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*NameBodyNodeContext) IsNameBodyNodeContext() {}

//AltLabelStructDecl tokenDecls

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *NameBodyNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *NameBodyNodeContext) Field() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerField, 0)
}

func (s *NameBodyNodeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerDOWN, 0)
}

func (s *NameBodyNodeContext) TypeExpr() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerTypeExpr, 0)
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

func (s *NameBodyNodeContext) Default() antlr.TerminalNode {
	return s.GetToken(DNAC_B_WalkerDefault, 0)
}

func (s *NameBodyNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NameBodyNodeEntryListener); ok {
		listenerT.EnterNameBodyNode(s)
	}
}

func (s *NameBodyNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NameBodyNodeExitListener); ok {
		listenerT.ExitNameBodyNode(s)
	}
}

func (s *NameBodyNodeContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*DNAC_B_WalkerHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.NameBodyNode != nil {
		h.NameBodyNode(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *NameBodyNodeContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case NameBodyNodeContextVisitor:
		return t.VisitNameBodyNode(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *DNAC_B_Walker) Tld() (localctx ITldContext) {
	localctx = NewTldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DNAC_B_WalkerRULE_tld)
	var //TokenTypeDecl
	_la int

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

	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNameRuleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(31)
			p.Match(DNAC_B_WalkerName)
		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DNAC_B_WalkerDOWN {
			{
				p.SetState(32)
				p.Match(DNAC_B_WalkerDOWN)
			}
			p.SetState(36)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == DNAC_B_WalkerAnnotation {
				{
					p.SetState(33)
					p.Match(DNAC_B_WalkerAnnotation)
				}

				p.SetState(38)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(40)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == DNAC_B_WalkerTypeParam {
				{
					p.SetState(39)
					p.Match(DNAC_B_WalkerTypeParam)
				}

			}
			p.SetState(45)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(DNAC_B_WalkerType-38))|(1<<(DNAC_B_WalkerField-38))|(1<<(DNAC_B_WalkerName-38))|(1<<(DNAC_B_WalkerExnotation-38)))) != 0 {
				{
					p.SetState(42)
					p.Tld()
				}

				p.SetState(47)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(48)
				p.Match(DNAC_B_WalkerUP)
			}

		}

	case 2:
		localctx = NewTypeNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(51)
			p.Match(DNAC_B_WalkerType)
		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DNAC_B_WalkerDOWN {
			{
				p.SetState(52)
				p.Match(DNAC_B_WalkerDOWN)
			}
			p.SetState(56)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == DNAC_B_WalkerAnnotation {
				{
					p.SetState(53)
					p.Match(DNAC_B_WalkerAnnotation)
				}

				p.SetState(58)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(59)
				p.Match(DNAC_B_WalkerTypeParam)
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
			{
				p.SetState(63)
				p.Match(DNAC_B_WalkerUP)
			}

		}

	case 3:
		localctx = NewTypeNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(66)
			p.Match(DNAC_B_WalkerType)
		}
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DNAC_B_WalkerDOWN {
			{
				p.SetState(67)
				p.Match(DNAC_B_WalkerDOWN)
			}
			p.SetState(71)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == DNAC_B_WalkerAnnotation {
				{
					p.SetState(68)
					p.Match(DNAC_B_WalkerAnnotation)
				}

				p.SetState(73)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(75)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == DNAC_B_WalkerDefault {
				{
					p.SetState(74)
					p.Match(DNAC_B_WalkerDefault)
				}

			}
			{
				p.SetState(77)
				p.Match(DNAC_B_WalkerUP)
			}

		}

	case 4:
		localctx = NewNameBodyNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(80)
			p.Match(DNAC_B_WalkerField)
		}
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DNAC_B_WalkerDOWN {
			{
				p.SetState(81)
				p.Match(DNAC_B_WalkerDOWN)
			}
			p.SetState(85)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == DNAC_B_WalkerAnnotation {
				{
					p.SetState(82)
					p.Match(DNAC_B_WalkerAnnotation)
				}

				p.SetState(87)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(88)
				p.Match(DNAC_B_WalkerTypeExpr)
			}
			{
				p.SetState(89)
				p.Match(DNAC_B_WalkerUP)
			}

		}

	case 5:
		localctx = NewNameBodyNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(92)
			p.Match(DNAC_B_WalkerField)
		}
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DNAC_B_WalkerDOWN {
			{
				p.SetState(93)
				p.Match(DNAC_B_WalkerDOWN)
			}
			p.SetState(97)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == DNAC_B_WalkerAnnotation {
				{
					p.SetState(94)
					p.Match(DNAC_B_WalkerAnnotation)
				}

				p.SetState(99)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(101)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == DNAC_B_WalkerDefault {
				{
					p.SetState(100)
					p.Match(DNAC_B_WalkerDefault)
				}

			}
			{
				p.SetState(103)
				p.Match(DNAC_B_WalkerUP)
			}

		}

	case 6:
		localctx = NewExnotationNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(106)
			p.Match(DNAC_B_WalkerExnotation)
		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DNAC_B_WalkerDOWN {
			{
				p.SetState(107)
				p.Match(DNAC_B_WalkerDOWN)
			}
			{
				p.SetState(108)
				p.Match(DNAC_B_WalkerUP)
			}

		}

	}

	return localctx
}
