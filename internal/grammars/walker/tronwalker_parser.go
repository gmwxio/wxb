// Code generated from TronWalker.g4 by ANTLR 4.7.1. DO NOT EDIT.

package walker // TronWalker
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 63, 122,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 3, 2, 3, 2, 3, 2, 5, 2, 20, 10, 2, 3, 2, 3, 2, 7, 2, 24, 10,
	2, 12, 2, 14, 2, 27, 11, 2, 3, 2, 5, 2, 30, 10, 2, 3, 2, 3, 2, 3, 2, 3,
	3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 42, 10, 4, 12, 4, 14, 4, 45,
	11, 4, 3, 4, 5, 4, 48, 10, 4, 3, 4, 3, 4, 3, 4, 7, 4, 53, 10, 4, 12, 4,
	14, 4, 56, 11, 4, 3, 4, 5, 4, 59, 10, 4, 3, 4, 3, 4, 3, 4, 7, 4, 64, 10,
	4, 12, 4, 14, 4, 67, 11, 4, 3, 4, 5, 4, 70, 10, 4, 3, 4, 3, 4, 3, 4, 7,
	4, 75, 10, 4, 12, 4, 14, 4, 78, 11, 4, 3, 4, 5, 4, 81, 10, 4, 5, 4, 83,
	10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 90, 10, 5, 12, 5, 14, 5, 93,
	11, 5, 3, 5, 5, 5, 96, 10, 5, 3, 5, 3, 5, 3, 5, 7, 5, 101, 10, 5, 12, 5,
	14, 5, 104, 11, 5, 3, 5, 5, 5, 107, 10, 5, 5, 5, 109, 10, 5, 3, 6, 3, 6,
	3, 6, 5, 6, 114, 10, 6, 5, 6, 116, 10, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8,
	2, 2, 9, 2, 4, 6, 8, 10, 12, 14, 2, 4, 4, 2, 40, 40, 50, 50, 4, 2, 40,
	40, 47, 47, 2, 139, 2, 16, 3, 2, 2, 2, 4, 34, 3, 2, 2, 2, 6, 82, 3, 2,
	2, 2, 8, 108, 3, 2, 2, 2, 10, 115, 3, 2, 2, 2, 12, 117, 3, 2, 2, 2, 14,
	119, 3, 2, 2, 2, 16, 17, 7, 34, 2, 2, 17, 19, 7, 36, 2, 2, 18, 20, 5, 4,
	3, 2, 19, 18, 3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 29, 3, 2, 2, 2, 21, 25,
	7, 34, 2, 2, 22, 24, 5, 6, 4, 2, 23, 22, 3, 2, 2, 2, 24, 27, 3, 2, 2, 2,
	25, 23, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 28, 3, 2, 2, 2, 27, 25, 3,
	2, 2, 2, 28, 30, 7, 35, 2, 2, 29, 21, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30,
	31, 3, 2, 2, 2, 31, 32, 7, 35, 2, 2, 32, 33, 7, 2, 2, 3, 33, 3, 3, 2, 2,
	2, 34, 35, 7, 4, 2, 2, 35, 5, 3, 2, 2, 2, 36, 83, 7, 39, 2, 2, 37, 83,
	7, 38, 2, 2, 38, 47, 7, 42, 2, 2, 39, 43, 7, 34, 2, 2, 40, 42, 5, 8, 5,
	2, 41, 40, 3, 2, 2, 2, 42, 45, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44,
	3, 2, 2, 2, 44, 46, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 46, 48, 7, 35, 2, 2,
	47, 39, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 83, 3, 2, 2, 2, 49, 58, 7,
	43, 2, 2, 50, 54, 7, 34, 2, 2, 51, 53, 5, 10, 6, 2, 52, 51, 3, 2, 2, 2,
	53, 56, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 57, 3,
	2, 2, 2, 56, 54, 3, 2, 2, 2, 57, 59, 7, 35, 2, 2, 58, 50, 3, 2, 2, 2, 58,
	59, 3, 2, 2, 2, 59, 83, 3, 2, 2, 2, 60, 69, 7, 44, 2, 2, 61, 65, 7, 34,
	2, 2, 62, 64, 5, 12, 7, 2, 63, 62, 3, 2, 2, 2, 64, 67, 3, 2, 2, 2, 65,
	63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 68, 3, 2, 2, 2, 67, 65, 3, 2, 2,
	2, 68, 70, 7, 35, 2, 2, 69, 61, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 83,
	3, 2, 2, 2, 71, 80, 7, 41, 2, 2, 72, 76, 7, 34, 2, 2, 73, 75, 5, 14, 8,
	2, 74, 73, 3, 2, 2, 2, 75, 78, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 76, 77,
	3, 2, 2, 2, 77, 79, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 79, 81, 7, 35, 2, 2,
	80, 72, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 83, 3, 2, 2, 2, 82, 36, 3,
	2, 2, 2, 82, 37, 3, 2, 2, 2, 82, 38, 3, 2, 2, 2, 82, 49, 3, 2, 2, 2, 82,
	60, 3, 2, 2, 2, 82, 71, 3, 2, 2, 2, 83, 7, 3, 2, 2, 2, 84, 109, 7, 40,
	2, 2, 85, 109, 7, 47, 2, 2, 86, 95, 7, 42, 2, 2, 87, 91, 7, 34, 2, 2, 88,
	90, 5, 8, 5, 2, 89, 88, 3, 2, 2, 2, 90, 93, 3, 2, 2, 2, 91, 89, 3, 2, 2,
	2, 91, 92, 3, 2, 2, 2, 92, 94, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 94, 96,
	7, 35, 2, 2, 95, 87, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 109, 3, 2, 2,
	2, 97, 106, 7, 43, 2, 2, 98, 102, 7, 34, 2, 2, 99, 101, 5, 10, 6, 2, 100,
	99, 3, 2, 2, 2, 101, 104, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 102, 103, 3,
	2, 2, 2, 103, 105, 3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 105, 107, 7, 35, 2,
	2, 106, 98, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 109, 3, 2, 2, 2, 108,
	84, 3, 2, 2, 2, 108, 85, 3, 2, 2, 2, 108, 86, 3, 2, 2, 2, 108, 97, 3, 2,
	2, 2, 109, 9, 3, 2, 2, 2, 110, 116, 7, 40, 2, 2, 111, 113, 7, 52, 2, 2,
	112, 114, 7, 53, 2, 2, 113, 112, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114,
	116, 3, 2, 2, 2, 115, 110, 3, 2, 2, 2, 115, 111, 3, 2, 2, 2, 116, 11, 3,
	2, 2, 2, 117, 118, 9, 2, 2, 2, 118, 13, 3, 2, 2, 2, 119, 120, 9, 3, 2,
	2, 120, 15, 3, 2, 2, 2, 21, 19, 25, 29, 43, 47, 54, 58, 65, 69, 76, 80,
	82, 91, 95, 102, 106, 108, 113, 115,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "'syntax'", "", "", "", "", "", "'{'", "'}'", "", "'\"'", "'''",
	"", "':'", "'('", "')'", "'.'", "'['", "']'", "','", "'<'", "'>'", "'-'",
	"'+'",
}
var symbolicNames = []string{
	"", "PROTO3", "SYNTAX", "EQ_PRE", "ENDPRE", "PRE_WS", "PRE_COMMENT", "PRE_LINE_COMMENT",
	"LCUR", "RCUR", "EQ", "DQ", "SQ", "SEMI", "COLON", "LPAREN", "RPAREN",
	"DOT", "LSB", "RSB", "COMMA", "LCHEVR", "RCHEVR", "DASH", "PLUS", "STR",
	"ID", "INT", "FLT", "WS", "COMMENT", "LINE_COMMENT", "DOWN", "UP", "ROOT",
	"ERROR", "Import", "Package", "Option", "Extend", "Message", "Enum", "Service",
	"Oneof", "Map", "Field", "Repeated", "Reserved", "Rpc", "Keytype", "EnumValue",
	"EnumNum", "INF", "NAN", "MAX", "WEAK", "PUBLIC", "Returns", "Stream",
	"To", "TRUE", "FALSE",
}

var ruleNames = []string{
	"proto", "syntax", "tld", "msgBody", "enumBody", "svcBody", "extendBody",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type TronWalker struct {
	*antlr.BaseParser
}

func NewTronWalker(input antlr.TokenStream) *TronWalker {
	this := new(TronWalker)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "TronWalker.g4"

	return this
}

// TronWalker tokens.
const (
	TronWalkerEOF              = antlr.TokenEOF
	TronWalkerPROTO3           = 1
	TronWalkerSYNTAX           = 2
	TronWalkerEQ_PRE           = 3
	TronWalkerENDPRE           = 4
	TronWalkerPRE_WS           = 5
	TronWalkerPRE_COMMENT      = 6
	TronWalkerPRE_LINE_COMMENT = 7
	TronWalkerLCUR             = 8
	TronWalkerRCUR             = 9
	TronWalkerEQ               = 10
	TronWalkerDQ               = 11
	TronWalkerSQ               = 12
	TronWalkerSEMI             = 13
	TronWalkerCOLON            = 14
	TronWalkerLPAREN           = 15
	TronWalkerRPAREN           = 16
	TronWalkerDOT              = 17
	TronWalkerLSB              = 18
	TronWalkerRSB              = 19
	TronWalkerCOMMA            = 20
	TronWalkerLCHEVR           = 21
	TronWalkerRCHEVR           = 22
	TronWalkerDASH             = 23
	TronWalkerPLUS             = 24
	TronWalkerSTR              = 25
	TronWalkerID               = 26
	TronWalkerINT              = 27
	TronWalkerFLT              = 28
	TronWalkerWS               = 29
	TronWalkerCOMMENT          = 30
	TronWalkerLINE_COMMENT     = 31
	TronWalkerDOWN             = 32
	TronWalkerUP               = 33
	TronWalkerROOT             = 34
	TronWalkerERROR            = 35
	TronWalkerImport           = 36
	TronWalkerPackage          = 37
	TronWalkerOption           = 38
	TronWalkerExtend           = 39
	TronWalkerMessage          = 40
	TronWalkerEnum             = 41
	TronWalkerService          = 42
	TronWalkerOneof            = 43
	TronWalkerMap              = 44
	TronWalkerField            = 45
	TronWalkerRepeated         = 46
	TronWalkerReserved         = 47
	TronWalkerRpc              = 48
	TronWalkerKeytype          = 49
	TronWalkerEnumValue        = 50
	TronWalkerEnumNum          = 51
	TronWalkerINF              = 52
	TronWalkerNAN              = 53
	TronWalkerMAX              = 54
	TronWalkerWEAK             = 55
	TronWalkerPUBLIC           = 56
	TronWalkerReturns          = 57
	TronWalkerStream           = 58
	TronWalkerTo               = 59
	TronWalkerTRUE             = 60
	TronWalkerFALSE            = 61
)

// TronWalker rules.
const (
	TronWalkerRULE_proto      = 0
	TronWalkerRULE_syntax     = 1
	TronWalkerRULE_tld        = 2
	TronWalkerRULE_msgBody    = 3
	TronWalkerRULE_enumBody   = 4
	TronWalkerRULE_svcBody    = 5
	TronWalkerRULE_extendBody = 6
)

// IProtoContext is an interface to support dynamic dispatch.
type IProtoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProtoContext differentiates from other interfaces.
	IsProtoContext()
}

type ProtoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProtoContext() *ProtoContext {
	var p = new(ProtoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TronWalkerRULE_proto
	return p
}

func (*ProtoContext) IsProtoContext() {}

func NewProtoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProtoContext {
	var p = new(ProtoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TronWalkerRULE_proto

	return p
}

func (s *ProtoContext) GetParser() antlr.Parser { return s.parser }

func (s *ProtoContext) AllDOWN() []antlr.TerminalNode {
	return s.GetTokens(TronWalkerDOWN)
}

func (s *ProtoContext) DOWN(i int) antlr.TerminalNode {
	return s.GetToken(TronWalkerDOWN, i)
}

func (s *ProtoContext) ROOT() antlr.TerminalNode {
	return s.GetToken(TronWalkerROOT, 0)
}

func (s *ProtoContext) AllUP() []antlr.TerminalNode {
	return s.GetTokens(TronWalkerUP)
}

func (s *ProtoContext) UP(i int) antlr.TerminalNode {
	return s.GetToken(TronWalkerUP, i)
}

func (s *ProtoContext) EOF() antlr.TerminalNode {
	return s.GetToken(TronWalkerEOF, 0)
}

func (s *ProtoContext) Syntax() ISyntaxContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISyntaxContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISyntaxContext)
}

func (s *ProtoContext) AllTld() []ITldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITldContext)(nil)).Elem())
	var tst = make([]ITldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITldContext)
		}
	}

	return tst
}

func (s *ProtoContext) Tld(i int) ITldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITldContext)
}

func (s *ProtoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProtoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProtoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.EnterProto(s)
	}
}

func (s *ProtoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.ExitProto(s)
	}
}

func (p *TronWalker) Proto() (localctx IProtoContext) {
	localctx = NewProtoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TronWalkerRULE_proto)
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
		p.SetState(14)
		p.Match(TronWalkerDOWN)
	}
	{
		p.SetState(15)
		p.Match(TronWalkerROOT)
	}
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TronWalkerSYNTAX {
		{
			p.SetState(16)
			p.Syntax()
		}

	}
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TronWalkerDOWN {
		{
			p.SetState(19)
			p.Match(TronWalkerDOWN)
		}
		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(TronWalkerImport-36))|(1<<(TronWalkerPackage-36))|(1<<(TronWalkerExtend-36))|(1<<(TronWalkerMessage-36))|(1<<(TronWalkerEnum-36))|(1<<(TronWalkerService-36)))) != 0 {
			{
				p.SetState(20)
				p.Tld()
			}

			p.SetState(25)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(26)
			p.Match(TronWalkerUP)
		}

	}
	{
		p.SetState(29)
		p.Match(TronWalkerUP)
	}
	{
		p.SetState(30)
		p.Match(TronWalkerEOF)
	}

	return localctx
}

// ISyntaxContext is an interface to support dynamic dispatch.
type ISyntaxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSyntaxContext differentiates from other interfaces.
	IsSyntaxContext()
}

type SyntaxContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySyntaxContext() *SyntaxContext {
	var p = new(SyntaxContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TronWalkerRULE_syntax
	return p
}

func (*SyntaxContext) IsSyntaxContext() {}

func NewSyntaxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SyntaxContext {
	var p = new(SyntaxContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TronWalkerRULE_syntax

	return p
}

func (s *SyntaxContext) GetParser() antlr.Parser { return s.parser }

func (s *SyntaxContext) SYNTAX() antlr.TerminalNode {
	return s.GetToken(TronWalkerSYNTAX, 0)
}

func (s *SyntaxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SyntaxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SyntaxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.EnterSyntax(s)
	}
}

func (s *SyntaxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.ExitSyntax(s)
	}
}

func (p *TronWalker) Syntax() (localctx ISyntaxContext) {
	localctx = NewSyntaxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TronWalkerRULE_syntax)

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
		p.SetState(32)
		p.Match(TronWalkerSYNTAX)
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
	p.RuleIndex = TronWalkerRULE_tld
	return p
}

func (*TldContext) IsTldContext() {}

func NewTldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TldContext {
	var p = new(TldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TronWalkerRULE_tld

	return p
}

func (s *TldContext) GetParser() antlr.Parser { return s.parser }

func (s *TldContext) Package() antlr.TerminalNode {
	return s.GetToken(TronWalkerPackage, 0)
}

func (s *TldContext) Import() antlr.TerminalNode {
	return s.GetToken(TronWalkerImport, 0)
}

func (s *TldContext) Message() antlr.TerminalNode {
	return s.GetToken(TronWalkerMessage, 0)
}

func (s *TldContext) DOWN() antlr.TerminalNode {
	return s.GetToken(TronWalkerDOWN, 0)
}

func (s *TldContext) UP() antlr.TerminalNode {
	return s.GetToken(TronWalkerUP, 0)
}

func (s *TldContext) AllMsgBody() []IMsgBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMsgBodyContext)(nil)).Elem())
	var tst = make([]IMsgBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMsgBodyContext)
		}
	}

	return tst
}

func (s *TldContext) MsgBody(i int) IMsgBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMsgBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMsgBodyContext)
}

func (s *TldContext) Enum() antlr.TerminalNode {
	return s.GetToken(TronWalkerEnum, 0)
}

func (s *TldContext) AllEnumBody() []IEnumBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumBodyContext)(nil)).Elem())
	var tst = make([]IEnumBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumBodyContext)
		}
	}

	return tst
}

func (s *TldContext) EnumBody(i int) IEnumBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

func (s *TldContext) Service() antlr.TerminalNode {
	return s.GetToken(TronWalkerService, 0)
}

func (s *TldContext) AllSvcBody() []ISvcBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISvcBodyContext)(nil)).Elem())
	var tst = make([]ISvcBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISvcBodyContext)
		}
	}

	return tst
}

func (s *TldContext) SvcBody(i int) ISvcBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISvcBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISvcBodyContext)
}

func (s *TldContext) Extend() antlr.TerminalNode {
	return s.GetToken(TronWalkerExtend, 0)
}

func (s *TldContext) AllExtendBody() []IExtendBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExtendBodyContext)(nil)).Elem())
	var tst = make([]IExtendBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExtendBodyContext)
		}
	}

	return tst
}

func (s *TldContext) ExtendBody(i int) IExtendBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExtendBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExtendBodyContext)
}

func (s *TldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.EnterTld(s)
	}
}

func (s *TldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.ExitTld(s)
	}
}

func (p *TronWalker) Tld() (localctx ITldContext) {
	localctx = NewTldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TronWalkerRULE_tld)
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

	p.SetState(80)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TronWalkerPackage:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.Match(TronWalkerPackage)
		}

	case TronWalkerImport:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(35)
			p.Match(TronWalkerImport)
		}

	case TronWalkerMessage:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(36)
			p.Match(TronWalkerMessage)
		}
		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TronWalkerDOWN {
			{
				p.SetState(37)
				p.Match(TronWalkerDOWN)
			}
			p.SetState(41)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(TronWalkerOption-38))|(1<<(TronWalkerMessage-38))|(1<<(TronWalkerEnum-38))|(1<<(TronWalkerField-38)))) != 0 {
				{
					p.SetState(38)
					p.MsgBody()
				}

				p.SetState(43)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(44)
				p.Match(TronWalkerUP)
			}

		}

	case TronWalkerEnum:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(47)
			p.Match(TronWalkerEnum)
		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TronWalkerDOWN {
			{
				p.SetState(48)
				p.Match(TronWalkerDOWN)
			}
			p.SetState(52)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == TronWalkerOption || _la == TronWalkerEnumValue {
				{
					p.SetState(49)
					p.EnumBody()
				}

				p.SetState(54)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(55)
				p.Match(TronWalkerUP)
			}

		}

	case TronWalkerService:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(58)
			p.Match(TronWalkerService)
		}
		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TronWalkerDOWN {
			{
				p.SetState(59)
				p.Match(TronWalkerDOWN)
			}
			p.SetState(63)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == TronWalkerOption || _la == TronWalkerRpc {
				{
					p.SetState(60)
					p.SvcBody()
				}

				p.SetState(65)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(66)
				p.Match(TronWalkerUP)
			}

		}

	case TronWalkerExtend:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(69)
			p.Match(TronWalkerExtend)
		}
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TronWalkerDOWN {
			{
				p.SetState(70)
				p.Match(TronWalkerDOWN)
			}
			p.SetState(74)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == TronWalkerOption || _la == TronWalkerField {
				{
					p.SetState(71)
					p.ExtendBody()
				}

				p.SetState(76)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(77)
				p.Match(TronWalkerUP)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMsgBodyContext is an interface to support dynamic dispatch.
type IMsgBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMsgBodyContext differentiates from other interfaces.
	IsMsgBodyContext()
}

type MsgBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMsgBodyContext() *MsgBodyContext {
	var p = new(MsgBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TronWalkerRULE_msgBody
	return p
}

func (*MsgBodyContext) IsMsgBodyContext() {}

func NewMsgBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MsgBodyContext {
	var p = new(MsgBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TronWalkerRULE_msgBody

	return p
}

func (s *MsgBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *MsgBodyContext) Option() antlr.TerminalNode {
	return s.GetToken(TronWalkerOption, 0)
}

func (s *MsgBodyContext) Field() antlr.TerminalNode {
	return s.GetToken(TronWalkerField, 0)
}

func (s *MsgBodyContext) Message() antlr.TerminalNode {
	return s.GetToken(TronWalkerMessage, 0)
}

func (s *MsgBodyContext) DOWN() antlr.TerminalNode {
	return s.GetToken(TronWalkerDOWN, 0)
}

func (s *MsgBodyContext) UP() antlr.TerminalNode {
	return s.GetToken(TronWalkerUP, 0)
}

func (s *MsgBodyContext) AllMsgBody() []IMsgBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMsgBodyContext)(nil)).Elem())
	var tst = make([]IMsgBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMsgBodyContext)
		}
	}

	return tst
}

func (s *MsgBodyContext) MsgBody(i int) IMsgBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMsgBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMsgBodyContext)
}

func (s *MsgBodyContext) Enum() antlr.TerminalNode {
	return s.GetToken(TronWalkerEnum, 0)
}

func (s *MsgBodyContext) AllEnumBody() []IEnumBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumBodyContext)(nil)).Elem())
	var tst = make([]IEnumBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumBodyContext)
		}
	}

	return tst
}

func (s *MsgBodyContext) EnumBody(i int) IEnumBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

func (s *MsgBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MsgBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MsgBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.EnterMsgBody(s)
	}
}

func (s *MsgBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.ExitMsgBody(s)
	}
}

func (p *TronWalker) MsgBody() (localctx IMsgBodyContext) {
	localctx = NewMsgBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TronWalkerRULE_msgBody)
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

	p.SetState(106)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TronWalkerOption:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.Match(TronWalkerOption)
		}

	case TronWalkerField:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			p.Match(TronWalkerField)
		}

	case TronWalkerMessage:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(84)
			p.Match(TronWalkerMessage)
		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TronWalkerDOWN {
			{
				p.SetState(85)
				p.Match(TronWalkerDOWN)
			}
			p.SetState(89)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(TronWalkerOption-38))|(1<<(TronWalkerMessage-38))|(1<<(TronWalkerEnum-38))|(1<<(TronWalkerField-38)))) != 0 {
				{
					p.SetState(86)
					p.MsgBody()
				}

				p.SetState(91)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(92)
				p.Match(TronWalkerUP)
			}

		}

	case TronWalkerEnum:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(95)
			p.Match(TronWalkerEnum)
		}
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TronWalkerDOWN {
			{
				p.SetState(96)
				p.Match(TronWalkerDOWN)
			}
			p.SetState(100)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == TronWalkerOption || _la == TronWalkerEnumValue {
				{
					p.SetState(97)
					p.EnumBody()
				}

				p.SetState(102)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(103)
				p.Match(TronWalkerUP)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEnumBodyContext is an interface to support dynamic dispatch.
type IEnumBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumBodyContext differentiates from other interfaces.
	IsEnumBodyContext()
}

type EnumBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumBodyContext() *EnumBodyContext {
	var p = new(EnumBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TronWalkerRULE_enumBody
	return p
}

func (*EnumBodyContext) IsEnumBodyContext() {}

func NewEnumBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumBodyContext {
	var p = new(EnumBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TronWalkerRULE_enumBody

	return p
}

func (s *EnumBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumBodyContext) Option() antlr.TerminalNode {
	return s.GetToken(TronWalkerOption, 0)
}

func (s *EnumBodyContext) EnumValue() antlr.TerminalNode {
	return s.GetToken(TronWalkerEnumValue, 0)
}

func (s *EnumBodyContext) EnumNum() antlr.TerminalNode {
	return s.GetToken(TronWalkerEnumNum, 0)
}

func (s *EnumBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.EnterEnumBody(s)
	}
}

func (s *EnumBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.ExitEnumBody(s)
	}
}

func (p *TronWalker) EnumBody() (localctx IEnumBodyContext) {
	localctx = NewEnumBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TronWalkerRULE_enumBody)
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

	p.SetState(113)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TronWalkerOption:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.Match(TronWalkerOption)
		}

	case TronWalkerEnumValue:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(109)
			p.Match(TronWalkerEnumValue)
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TronWalkerEnumNum {
			{
				p.SetState(110)
				p.Match(TronWalkerEnumNum)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISvcBodyContext is an interface to support dynamic dispatch.
type ISvcBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSvcBodyContext differentiates from other interfaces.
	IsSvcBodyContext()
}

type SvcBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySvcBodyContext() *SvcBodyContext {
	var p = new(SvcBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TronWalkerRULE_svcBody
	return p
}

func (*SvcBodyContext) IsSvcBodyContext() {}

func NewSvcBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SvcBodyContext {
	var p = new(SvcBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TronWalkerRULE_svcBody

	return p
}

func (s *SvcBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *SvcBodyContext) Option() antlr.TerminalNode {
	return s.GetToken(TronWalkerOption, 0)
}

func (s *SvcBodyContext) Rpc() antlr.TerminalNode {
	return s.GetToken(TronWalkerRpc, 0)
}

func (s *SvcBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SvcBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SvcBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.EnterSvcBody(s)
	}
}

func (s *SvcBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.ExitSvcBody(s)
	}
}

func (p *TronWalker) SvcBody() (localctx ISvcBodyContext) {
	localctx = NewSvcBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TronWalkerRULE_svcBody)
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
		p.SetState(115)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TronWalkerOption || _la == TronWalkerRpc) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IExtendBodyContext is an interface to support dynamic dispatch.
type IExtendBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExtendBodyContext differentiates from other interfaces.
	IsExtendBodyContext()
}

type ExtendBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtendBodyContext() *ExtendBodyContext {
	var p = new(ExtendBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TronWalkerRULE_extendBody
	return p
}

func (*ExtendBodyContext) IsExtendBodyContext() {}

func NewExtendBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtendBodyContext {
	var p = new(ExtendBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TronWalkerRULE_extendBody

	return p
}

func (s *ExtendBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtendBodyContext) Option() antlr.TerminalNode {
	return s.GetToken(TronWalkerOption, 0)
}

func (s *ExtendBodyContext) Field() antlr.TerminalNode {
	return s.GetToken(TronWalkerField, 0)
}

func (s *ExtendBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtendBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtendBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.EnterExtendBody(s)
	}
}

func (s *ExtendBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.ExitExtendBody(s)
	}
}

func (p *TronWalker) ExtendBody() (localctx IExtendBodyContext) {
	localctx = NewExtendBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TronWalkerRULE_extendBody)
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
		p.SetState(117)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TronWalkerOption || _la == TronWalkerField) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
