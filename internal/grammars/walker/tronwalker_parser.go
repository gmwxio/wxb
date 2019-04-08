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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 63, 161,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 2, 5, 2, 22, 10, 2, 3, 2, 3, 2,
	7, 2, 26, 10, 2, 12, 2, 14, 2, 29, 11, 2, 3, 2, 5, 2, 32, 10, 2, 3, 2,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 44, 10, 4,
	12, 4, 14, 4, 47, 11, 4, 3, 4, 5, 4, 50, 10, 4, 3, 4, 3, 4, 3, 4, 7, 4,
	55, 10, 4, 12, 4, 14, 4, 58, 11, 4, 3, 4, 5, 4, 61, 10, 4, 3, 4, 3, 4,
	3, 4, 7, 4, 66, 10, 4, 12, 4, 14, 4, 69, 11, 4, 3, 4, 5, 4, 72, 10, 4,
	3, 4, 3, 4, 3, 4, 7, 4, 77, 10, 4, 12, 4, 14, 4, 80, 11, 4, 3, 4, 5, 4,
	83, 10, 4, 5, 4, 85, 10, 4, 3, 5, 3, 5, 3, 5, 5, 5, 90, 10, 5, 3, 5, 5,
	5, 93, 10, 5, 3, 5, 3, 5, 3, 5, 7, 5, 98, 10, 5, 12, 5, 14, 5, 101, 11,
	5, 3, 5, 5, 5, 104, 10, 5, 3, 5, 3, 5, 3, 5, 7, 5, 109, 10, 5, 12, 5, 14,
	5, 112, 11, 5, 3, 5, 5, 5, 115, 10, 5, 3, 5, 3, 5, 3, 5, 7, 5, 120, 10,
	5, 12, 5, 14, 5, 123, 11, 5, 3, 5, 5, 5, 126, 10, 5, 5, 5, 128, 10, 5,
	3, 6, 3, 6, 3, 6, 5, 6, 133, 10, 6, 5, 6, 135, 10, 6, 3, 7, 3, 7, 3, 7,
	5, 7, 140, 10, 7, 3, 7, 5, 7, 143, 10, 7, 5, 7, 145, 10, 7, 3, 8, 3, 8,
	5, 8, 149, 10, 8, 3, 9, 3, 9, 3, 9, 5, 9, 154, 10, 9, 3, 9, 5, 9, 157,
	10, 9, 5, 9, 159, 10, 9, 3, 9, 2, 2, 10, 2, 4, 6, 8, 10, 12, 14, 16, 2,
	2, 2, 189, 2, 18, 3, 2, 2, 2, 4, 36, 3, 2, 2, 2, 6, 84, 3, 2, 2, 2, 8,
	127, 3, 2, 2, 2, 10, 134, 3, 2, 2, 2, 12, 144, 3, 2, 2, 2, 14, 148, 3,
	2, 2, 2, 16, 158, 3, 2, 2, 2, 18, 19, 7, 34, 2, 2, 19, 21, 7, 36, 2, 2,
	20, 22, 5, 4, 3, 2, 21, 20, 3, 2, 2, 2, 21, 22, 3, 2, 2, 2, 22, 31, 3,
	2, 2, 2, 23, 27, 7, 34, 2, 2, 24, 26, 5, 6, 4, 2, 25, 24, 3, 2, 2, 2, 26,
	29, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 30, 3, 2, 2,
	2, 29, 27, 3, 2, 2, 2, 30, 32, 7, 35, 2, 2, 31, 23, 3, 2, 2, 2, 31, 32,
	3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 34, 7, 35, 2, 2, 34, 35, 7, 2, 2, 3,
	35, 3, 3, 2, 2, 2, 36, 37, 7, 4, 2, 2, 37, 5, 3, 2, 2, 2, 38, 85, 7, 39,
	2, 2, 39, 85, 7, 38, 2, 2, 40, 49, 7, 42, 2, 2, 41, 45, 7, 34, 2, 2, 42,
	44, 5, 8, 5, 2, 43, 42, 3, 2, 2, 2, 44, 47, 3, 2, 2, 2, 45, 43, 3, 2, 2,
	2, 45, 46, 3, 2, 2, 2, 46, 48, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 48, 50,
	7, 35, 2, 2, 49, 41, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 85, 3, 2, 2, 2,
	51, 60, 7, 43, 2, 2, 52, 56, 7, 34, 2, 2, 53, 55, 5, 12, 7, 2, 54, 53,
	3, 2, 2, 2, 55, 58, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2,
	57, 59, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 59, 61, 7, 35, 2, 2, 60, 52, 3,
	2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 85, 3, 2, 2, 2, 62, 71, 7, 44, 2, 2, 63,
	67, 7, 34, 2, 2, 64, 66, 5, 14, 8, 2, 65, 64, 3, 2, 2, 2, 66, 69, 3, 2,
	2, 2, 67, 65, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 70, 3, 2, 2, 2, 69, 67,
	3, 2, 2, 2, 70, 72, 7, 35, 2, 2, 71, 63, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2,
	72, 85, 3, 2, 2, 2, 73, 82, 7, 41, 2, 2, 74, 78, 7, 34, 2, 2, 75, 77, 5,
	16, 9, 2, 76, 75, 3, 2, 2, 2, 77, 80, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 78,
	79, 3, 2, 2, 2, 79, 81, 3, 2, 2, 2, 80, 78, 3, 2, 2, 2, 81, 83, 7, 35,
	2, 2, 82, 74, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 85, 3, 2, 2, 2, 84, 38,
	3, 2, 2, 2, 84, 39, 3, 2, 2, 2, 84, 40, 3, 2, 2, 2, 84, 51, 3, 2, 2, 2,
	84, 62, 3, 2, 2, 2, 84, 73, 3, 2, 2, 2, 85, 7, 3, 2, 2, 2, 86, 128, 7,
	40, 2, 2, 87, 89, 7, 47, 2, 2, 88, 90, 7, 48, 2, 2, 89, 88, 3, 2, 2, 2,
	89, 90, 3, 2, 2, 2, 90, 92, 3, 2, 2, 2, 91, 93, 7, 40, 2, 2, 92, 91, 3,
	2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 128, 3, 2, 2, 2, 94, 103, 7, 42, 2, 2,
	95, 99, 7, 34, 2, 2, 96, 98, 5, 8, 5, 2, 97, 96, 3, 2, 2, 2, 98, 101, 3,
	2, 2, 2, 99, 97, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 102, 3, 2, 2, 2,
	101, 99, 3, 2, 2, 2, 102, 104, 7, 35, 2, 2, 103, 95, 3, 2, 2, 2, 103, 104,
	3, 2, 2, 2, 104, 128, 3, 2, 2, 2, 105, 114, 7, 43, 2, 2, 106, 110, 7, 34,
	2, 2, 107, 109, 5, 12, 7, 2, 108, 107, 3, 2, 2, 2, 109, 112, 3, 2, 2, 2,
	110, 108, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 113, 3, 2, 2, 2, 112,
	110, 3, 2, 2, 2, 113, 115, 7, 35, 2, 2, 114, 106, 3, 2, 2, 2, 114, 115,
	3, 2, 2, 2, 115, 128, 3, 2, 2, 2, 116, 125, 7, 45, 2, 2, 117, 121, 7, 34,
	2, 2, 118, 120, 5, 10, 6, 2, 119, 118, 3, 2, 2, 2, 120, 123, 3, 2, 2, 2,
	121, 119, 3, 2, 2, 2, 121, 122, 3, 2, 2, 2, 122, 124, 3, 2, 2, 2, 123,
	121, 3, 2, 2, 2, 124, 126, 7, 35, 2, 2, 125, 117, 3, 2, 2, 2, 125, 126,
	3, 2, 2, 2, 126, 128, 3, 2, 2, 2, 127, 86, 3, 2, 2, 2, 127, 87, 3, 2, 2,
	2, 127, 94, 3, 2, 2, 2, 127, 105, 3, 2, 2, 2, 127, 116, 3, 2, 2, 2, 128,
	9, 3, 2, 2, 2, 129, 135, 7, 40, 2, 2, 130, 132, 7, 47, 2, 2, 131, 133,
	7, 40, 2, 2, 132, 131, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 135, 3, 2,
	2, 2, 134, 129, 3, 2, 2, 2, 134, 130, 3, 2, 2, 2, 135, 11, 3, 2, 2, 2,
	136, 145, 7, 40, 2, 2, 137, 139, 7, 52, 2, 2, 138, 140, 7, 53, 2, 2, 139,
	138, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 142, 3, 2, 2, 2, 141, 143,
	7, 40, 2, 2, 142, 141, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 145, 3, 2,
	2, 2, 144, 136, 3, 2, 2, 2, 144, 137, 3, 2, 2, 2, 145, 13, 3, 2, 2, 2,
	146, 149, 7, 40, 2, 2, 147, 149, 7, 50, 2, 2, 148, 146, 3, 2, 2, 2, 148,
	147, 3, 2, 2, 2, 149, 15, 3, 2, 2, 2, 150, 159, 7, 40, 2, 2, 151, 153,
	7, 47, 2, 2, 152, 154, 7, 48, 2, 2, 153, 152, 3, 2, 2, 2, 153, 154, 3,
	2, 2, 2, 154, 156, 3, 2, 2, 2, 155, 157, 7, 40, 2, 2, 156, 155, 3, 2, 2,
	2, 156, 157, 3, 2, 2, 2, 157, 159, 3, 2, 2, 2, 158, 150, 3, 2, 2, 2, 158,
	151, 3, 2, 2, 2, 159, 17, 3, 2, 2, 2, 32, 21, 27, 31, 45, 49, 56, 60, 67,
	71, 78, 82, 84, 89, 92, 99, 103, 110, 114, 121, 125, 127, 132, 134, 139,
	142, 144, 148, 153, 156, 158,
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
	"Oneof", "Map", "Field", "Datastructure", "Reserved", "Rpc", "Keytype",
	"EnumValue", "EnumNum", "INF", "NAN", "MAX", "WEAK", "PUBLIC", "Returns",
	"Stream", "To", "TRUE", "FALSE",
}

var ruleNames = []string{
	"proto", "syntax", "tld", "msgBody", "oneofBody", "enumBody", "svcBody",
	"extendBody",
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
	TronWalkerDatastructure    = 46
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
	TronWalkerRULE_oneofBody  = 4
	TronWalkerRULE_enumBody   = 5
	TronWalkerRULE_svcBody    = 6
	TronWalkerRULE_extendBody = 7
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
		p.SetState(16)
		p.Match(TronWalkerDOWN)
	}
	{
		p.SetState(17)
		p.Match(TronWalkerROOT)
	}
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TronWalkerSYNTAX {
		{
			p.SetState(18)
			p.Syntax()
		}

	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TronWalkerDOWN {
		{
			p.SetState(21)
			p.Match(TronWalkerDOWN)
		}
		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(TronWalkerImport-36))|(1<<(TronWalkerPackage-36))|(1<<(TronWalkerExtend-36))|(1<<(TronWalkerMessage-36))|(1<<(TronWalkerEnum-36))|(1<<(TronWalkerService-36)))) != 0 {
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
			p.Match(TronWalkerUP)
		}

	}
	{
		p.SetState(31)
		p.Match(TronWalkerUP)
	}
	{
		p.SetState(32)
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

func (s *SyntaxContext) CopyFrom(ctx *SyntaxContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SyntaxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SyntaxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SyntaxNodeContext struct {
	*SyntaxContext
}

func NewSyntaxNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SyntaxNodeContext {
	var p = new(SyntaxNodeContext)

	p.SyntaxContext = NewEmptySyntaxContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SyntaxContext))

	return p
}

func (s *SyntaxNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SyntaxNodeContext) SYNTAX() antlr.TerminalNode {
	return s.GetToken(TronWalkerSYNTAX, 0)
}

func (s *SyntaxNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.EnterSyntaxNode(s)
	}
}

func (s *SyntaxNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.ExitSyntaxNode(s)
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

	localctx = NewSyntaxNodeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
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

func (s *PkgNodeContext) Package() antlr.TerminalNode {
	return s.GetToken(TronWalkerPackage, 0)
}

func (s *PkgNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.EnterPkgNode(s)
	}
}

func (s *PkgNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
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

func (s *ExtNodeContext) Extend() antlr.TerminalNode {
	return s.GetToken(TronWalkerExtend, 0)
}

func (s *ExtNodeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(TronWalkerDOWN, 0)
}

func (s *ExtNodeContext) UP() antlr.TerminalNode {
	return s.GetToken(TronWalkerUP, 0)
}

func (s *ExtNodeContext) AllExtendBody() []IExtendBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExtendBodyContext)(nil)).Elem())
	var tst = make([]IExtendBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExtendBodyContext)
		}
	}

	return tst
}

func (s *ExtNodeContext) ExtendBody(i int) IExtendBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExtendBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExtendBodyContext)
}

func (s *ExtNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.EnterExtNode(s)
	}
}

func (s *ExtNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.ExitExtNode(s)
	}
}

type EnumNodeContext struct {
	*TldContext
}

func NewEnumNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EnumNodeContext {
	var p = new(EnumNodeContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

func (s *EnumNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumNodeContext) Enum() antlr.TerminalNode {
	return s.GetToken(TronWalkerEnum, 0)
}

func (s *EnumNodeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(TronWalkerDOWN, 0)
}

func (s *EnumNodeContext) UP() antlr.TerminalNode {
	return s.GetToken(TronWalkerUP, 0)
}

func (s *EnumNodeContext) AllEnumBody() []IEnumBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumBodyContext)(nil)).Elem())
	var tst = make([]IEnumBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumBodyContext)
		}
	}

	return tst
}

func (s *EnumNodeContext) EnumBody(i int) IEnumBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

func (s *EnumNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.EnterEnumNode(s)
	}
}

func (s *EnumNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.ExitEnumNode(s)
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

func (s *ImportNodeContext) Import() antlr.TerminalNode {
	return s.GetToken(TronWalkerImport, 0)
}

func (s *ImportNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.EnterImportNode(s)
	}
}

func (s *ImportNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.ExitImportNode(s)
	}
}

type SvcNodeContext struct {
	*TldContext
}

func NewSvcNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SvcNodeContext {
	var p = new(SvcNodeContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

func (s *SvcNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SvcNodeContext) Service() antlr.TerminalNode {
	return s.GetToken(TronWalkerService, 0)
}

func (s *SvcNodeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(TronWalkerDOWN, 0)
}

func (s *SvcNodeContext) UP() antlr.TerminalNode {
	return s.GetToken(TronWalkerUP, 0)
}

func (s *SvcNodeContext) AllSvcBody() []ISvcBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISvcBodyContext)(nil)).Elem())
	var tst = make([]ISvcBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISvcBodyContext)
		}
	}

	return tst
}

func (s *SvcNodeContext) SvcBody(i int) ISvcBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISvcBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISvcBodyContext)
}

func (s *SvcNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.EnterSvcNode(s)
	}
}

func (s *SvcNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.ExitSvcNode(s)
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

func (s *MsgNodeContext) Message() antlr.TerminalNode {
	return s.GetToken(TronWalkerMessage, 0)
}

func (s *MsgNodeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(TronWalkerDOWN, 0)
}

func (s *MsgNodeContext) UP() antlr.TerminalNode {
	return s.GetToken(TronWalkerUP, 0)
}

func (s *MsgNodeContext) AllMsgBody() []IMsgBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMsgBodyContext)(nil)).Elem())
	var tst = make([]IMsgBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMsgBodyContext)
		}
	}

	return tst
}

func (s *MsgNodeContext) MsgBody(i int) IMsgBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMsgBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMsgBodyContext)
}

func (s *MsgNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.EnterMsgNode(s)
	}
}

func (s *MsgNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.ExitMsgNode(s)
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

	p.SetState(82)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TronWalkerPackage:
		localctx = NewPkgNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(36)
			p.Match(TronWalkerPackage)
		}

	case TronWalkerImport:
		localctx = NewImportNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.Match(TronWalkerImport)
		}

	case TronWalkerMessage:
		localctx = NewMsgNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(38)
			p.Match(TronWalkerMessage)
		}
		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TronWalkerDOWN {
			{
				p.SetState(39)
				p.Match(TronWalkerDOWN)
			}
			p.SetState(43)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(TronWalkerOption-38))|(1<<(TronWalkerMessage-38))|(1<<(TronWalkerEnum-38))|(1<<(TronWalkerOneof-38))|(1<<(TronWalkerField-38)))) != 0 {
				{
					p.SetState(40)
					p.MsgBody()
				}

				p.SetState(45)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(46)
				p.Match(TronWalkerUP)
			}

		}

	case TronWalkerEnum:
		localctx = NewEnumNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(49)
			p.Match(TronWalkerEnum)
		}
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TronWalkerDOWN {
			{
				p.SetState(50)
				p.Match(TronWalkerDOWN)
			}
			p.SetState(54)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == TronWalkerOption || _la == TronWalkerEnumValue {
				{
					p.SetState(51)
					p.EnumBody()
				}

				p.SetState(56)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(57)
				p.Match(TronWalkerUP)
			}

		}

	case TronWalkerService:
		localctx = NewSvcNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(60)
			p.Match(TronWalkerService)
		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TronWalkerDOWN {
			{
				p.SetState(61)
				p.Match(TronWalkerDOWN)
			}
			p.SetState(65)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == TronWalkerOption || _la == TronWalkerRpc {
				{
					p.SetState(62)
					p.SvcBody()
				}

				p.SetState(67)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(68)
				p.Match(TronWalkerUP)
			}

		}

	case TronWalkerExtend:
		localctx = NewExtNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(71)
			p.Match(TronWalkerExtend)
		}
		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TronWalkerDOWN {
			{
				p.SetState(72)
				p.Match(TronWalkerDOWN)
			}
			p.SetState(76)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == TronWalkerOption || _la == TronWalkerField {
				{
					p.SetState(73)
					p.ExtendBody()
				}

				p.SetState(78)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(79)
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

func (s *MsgBodyContext) CopyFrom(ctx *MsgBodyContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *MsgBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MsgBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MsgOptNodeContext struct {
	*MsgBodyContext
}

func NewMsgOptNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MsgOptNodeContext {
	var p = new(MsgOptNodeContext)

	p.MsgBodyContext = NewEmptyMsgBodyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MsgBodyContext))

	return p
}

func (s *MsgOptNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MsgOptNodeContext) Option() antlr.TerminalNode {
	return s.GetToken(TronWalkerOption, 0)
}

func (s *MsgOptNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.EnterMsgOptNode(s)
	}
}

func (s *MsgOptNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.ExitMsgOptNode(s)
	}
}

type MsgOneofNodeContext struct {
	*MsgBodyContext
}

func NewMsgOneofNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MsgOneofNodeContext {
	var p = new(MsgOneofNodeContext)

	p.MsgBodyContext = NewEmptyMsgBodyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MsgBodyContext))

	return p
}

func (s *MsgOneofNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MsgOneofNodeContext) Oneof() antlr.TerminalNode {
	return s.GetToken(TronWalkerOneof, 0)
}

func (s *MsgOneofNodeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(TronWalkerDOWN, 0)
}

func (s *MsgOneofNodeContext) UP() antlr.TerminalNode {
	return s.GetToken(TronWalkerUP, 0)
}

func (s *MsgOneofNodeContext) AllOneofBody() []IOneofBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOneofBodyContext)(nil)).Elem())
	var tst = make([]IOneofBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOneofBodyContext)
		}
	}

	return tst
}

func (s *MsgOneofNodeContext) OneofBody(i int) IOneofBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOneofBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOneofBodyContext)
}

func (s *MsgOneofNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.EnterMsgOneofNode(s)
	}
}

func (s *MsgOneofNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.ExitMsgOneofNode(s)
	}
}

type MsgMsgNodeContext struct {
	*MsgBodyContext
}

func NewMsgMsgNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MsgMsgNodeContext {
	var p = new(MsgMsgNodeContext)

	p.MsgBodyContext = NewEmptyMsgBodyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MsgBodyContext))

	return p
}

func (s *MsgMsgNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MsgMsgNodeContext) Message() antlr.TerminalNode {
	return s.GetToken(TronWalkerMessage, 0)
}

func (s *MsgMsgNodeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(TronWalkerDOWN, 0)
}

func (s *MsgMsgNodeContext) UP() antlr.TerminalNode {
	return s.GetToken(TronWalkerUP, 0)
}

func (s *MsgMsgNodeContext) AllMsgBody() []IMsgBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMsgBodyContext)(nil)).Elem())
	var tst = make([]IMsgBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMsgBodyContext)
		}
	}

	return tst
}

func (s *MsgMsgNodeContext) MsgBody(i int) IMsgBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMsgBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMsgBodyContext)
}

func (s *MsgMsgNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.EnterMsgMsgNode(s)
	}
}

func (s *MsgMsgNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.ExitMsgMsgNode(s)
	}
}

type MsgFldNodeContext struct {
	*MsgBodyContext
}

func NewMsgFldNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MsgFldNodeContext {
	var p = new(MsgFldNodeContext)

	p.MsgBodyContext = NewEmptyMsgBodyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MsgBodyContext))

	return p
}

func (s *MsgFldNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MsgFldNodeContext) Field() antlr.TerminalNode {
	return s.GetToken(TronWalkerField, 0)
}

func (s *MsgFldNodeContext) Datastructure() antlr.TerminalNode {
	return s.GetToken(TronWalkerDatastructure, 0)
}

func (s *MsgFldNodeContext) Option() antlr.TerminalNode {
	return s.GetToken(TronWalkerOption, 0)
}

func (s *MsgFldNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.EnterMsgFldNode(s)
	}
}

func (s *MsgFldNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.ExitMsgFldNode(s)
	}
}

type MsgEnumNodeContext struct {
	*MsgBodyContext
}

func NewMsgEnumNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MsgEnumNodeContext {
	var p = new(MsgEnumNodeContext)

	p.MsgBodyContext = NewEmptyMsgBodyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MsgBodyContext))

	return p
}

func (s *MsgEnumNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MsgEnumNodeContext) Enum() antlr.TerminalNode {
	return s.GetToken(TronWalkerEnum, 0)
}

func (s *MsgEnumNodeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(TronWalkerDOWN, 0)
}

func (s *MsgEnumNodeContext) UP() antlr.TerminalNode {
	return s.GetToken(TronWalkerUP, 0)
}

func (s *MsgEnumNodeContext) AllEnumBody() []IEnumBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumBodyContext)(nil)).Elem())
	var tst = make([]IEnumBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumBodyContext)
		}
	}

	return tst
}

func (s *MsgEnumNodeContext) EnumBody(i int) IEnumBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

func (s *MsgEnumNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.EnterMsgEnumNode(s)
	}
}

func (s *MsgEnumNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.ExitMsgEnumNode(s)
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

	p.SetState(125)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TronWalkerOption:
		localctx = NewMsgOptNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)
			p.Match(TronWalkerOption)
		}

	case TronWalkerField:
		localctx = NewMsgFldNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.Match(TronWalkerField)
		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TronWalkerDatastructure {
			{
				p.SetState(86)
				p.Match(TronWalkerDatastructure)
			}

		}
		p.SetState(90)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(89)
				p.Match(TronWalkerOption)
			}

		}

	case TronWalkerMessage:
		localctx = NewMsgMsgNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(92)
			p.Match(TronWalkerMessage)
		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TronWalkerDOWN {
			{
				p.SetState(93)
				p.Match(TronWalkerDOWN)
			}
			p.SetState(97)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(TronWalkerOption-38))|(1<<(TronWalkerMessage-38))|(1<<(TronWalkerEnum-38))|(1<<(TronWalkerOneof-38))|(1<<(TronWalkerField-38)))) != 0 {
				{
					p.SetState(94)
					p.MsgBody()
				}

				p.SetState(99)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(100)
				p.Match(TronWalkerUP)
			}

		}

	case TronWalkerEnum:
		localctx = NewMsgEnumNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(103)
			p.Match(TronWalkerEnum)
		}
		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TronWalkerDOWN {
			{
				p.SetState(104)
				p.Match(TronWalkerDOWN)
			}
			p.SetState(108)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == TronWalkerOption || _la == TronWalkerEnumValue {
				{
					p.SetState(105)
					p.EnumBody()
				}

				p.SetState(110)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(111)
				p.Match(TronWalkerUP)
			}

		}

	case TronWalkerOneof:
		localctx = NewMsgOneofNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(114)
			p.Match(TronWalkerOneof)
		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TronWalkerDOWN {
			{
				p.SetState(115)
				p.Match(TronWalkerDOWN)
			}
			p.SetState(119)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == TronWalkerOption || _la == TronWalkerField {
				{
					p.SetState(116)
					p.OneofBody()
				}

				p.SetState(121)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(122)
				p.Match(TronWalkerUP)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOneofBodyContext is an interface to support dynamic dispatch.
type IOneofBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOneofBodyContext differentiates from other interfaces.
	IsOneofBodyContext()
}

type OneofBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOneofBodyContext() *OneofBodyContext {
	var p = new(OneofBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TronWalkerRULE_oneofBody
	return p
}

func (*OneofBodyContext) IsOneofBodyContext() {}

func NewOneofBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofBodyContext {
	var p = new(OneofBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TronWalkerRULE_oneofBody

	return p
}

func (s *OneofBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *OneofBodyContext) CopyFrom(ctx *OneofBodyContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *OneofBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneofBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type OneofFldNodeContext struct {
	*OneofBodyContext
}

func NewOneofFldNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OneofFldNodeContext {
	var p = new(OneofFldNodeContext)

	p.OneofBodyContext = NewEmptyOneofBodyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OneofBodyContext))

	return p
}

func (s *OneofFldNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneofFldNodeContext) Field() antlr.TerminalNode {
	return s.GetToken(TronWalkerField, 0)
}

func (s *OneofFldNodeContext) Option() antlr.TerminalNode {
	return s.GetToken(TronWalkerOption, 0)
}

func (s *OneofFldNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.EnterOneofFldNode(s)
	}
}

func (s *OneofFldNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.ExitOneofFldNode(s)
	}
}

type OneofOptNodeContext struct {
	*OneofBodyContext
}

func NewOneofOptNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OneofOptNodeContext {
	var p = new(OneofOptNodeContext)

	p.OneofBodyContext = NewEmptyOneofBodyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OneofBodyContext))

	return p
}

func (s *OneofOptNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneofOptNodeContext) Option() antlr.TerminalNode {
	return s.GetToken(TronWalkerOption, 0)
}

func (s *OneofOptNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.EnterOneofOptNode(s)
	}
}

func (s *OneofOptNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.ExitOneofOptNode(s)
	}
}

func (p *TronWalker) OneofBody() (localctx IOneofBodyContext) {
	localctx = NewOneofBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TronWalkerRULE_oneofBody)

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

	p.SetState(132)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TronWalkerOption:
		localctx = NewOneofOptNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)
			p.Match(TronWalkerOption)
		}

	case TronWalkerField:
		localctx = NewOneofFldNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(128)
			p.Match(TronWalkerField)
		}
		p.SetState(130)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(129)
				p.Match(TronWalkerOption)
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

func (s *EnumBodyContext) CopyFrom(ctx *EnumBodyContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *EnumBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EnumOptNodeContext struct {
	*EnumBodyContext
}

func NewEnumOptNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EnumOptNodeContext {
	var p = new(EnumOptNodeContext)

	p.EnumBodyContext = NewEmptyEnumBodyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*EnumBodyContext))

	return p
}

func (s *EnumOptNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumOptNodeContext) Option() antlr.TerminalNode {
	return s.GetToken(TronWalkerOption, 0)
}

func (s *EnumOptNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.EnterEnumOptNode(s)
	}
}

func (s *EnumOptNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.ExitEnumOptNode(s)
	}
}

type EnumValNodeContext struct {
	*EnumBodyContext
}

func NewEnumValNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EnumValNodeContext {
	var p = new(EnumValNodeContext)

	p.EnumBodyContext = NewEmptyEnumBodyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*EnumBodyContext))

	return p
}

func (s *EnumValNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValNodeContext) EnumValue() antlr.TerminalNode {
	return s.GetToken(TronWalkerEnumValue, 0)
}

func (s *EnumValNodeContext) EnumNum() antlr.TerminalNode {
	return s.GetToken(TronWalkerEnumNum, 0)
}

func (s *EnumValNodeContext) Option() antlr.TerminalNode {
	return s.GetToken(TronWalkerOption, 0)
}

func (s *EnumValNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.EnterEnumValNode(s)
	}
}

func (s *EnumValNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.ExitEnumValNode(s)
	}
}

func (p *TronWalker) EnumBody() (localctx IEnumBodyContext) {
	localctx = NewEnumBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TronWalkerRULE_enumBody)
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

	p.SetState(142)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TronWalkerOption:
		localctx = NewEnumOptNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(134)
			p.Match(TronWalkerOption)
		}

	case TronWalkerEnumValue:
		localctx = NewEnumValNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(135)
			p.Match(TronWalkerEnumValue)
		}
		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TronWalkerEnumNum {
			{
				p.SetState(136)
				p.Match(TronWalkerEnumNum)
			}

		}
		p.SetState(140)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(139)
				p.Match(TronWalkerOption)
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

func (s *SvcBodyContext) CopyFrom(ctx *SvcBodyContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SvcBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SvcBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SvcRpcNodeContext struct {
	*SvcBodyContext
}

func NewSvcRpcNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SvcRpcNodeContext {
	var p = new(SvcRpcNodeContext)

	p.SvcBodyContext = NewEmptySvcBodyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SvcBodyContext))

	return p
}

func (s *SvcRpcNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SvcRpcNodeContext) Rpc() antlr.TerminalNode {
	return s.GetToken(TronWalkerRpc, 0)
}

func (s *SvcRpcNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.EnterSvcRpcNode(s)
	}
}

func (s *SvcRpcNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.ExitSvcRpcNode(s)
	}
}

type SvcOptNodeContext struct {
	*SvcBodyContext
}

func NewSvcOptNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SvcOptNodeContext {
	var p = new(SvcOptNodeContext)

	p.SvcBodyContext = NewEmptySvcBodyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SvcBodyContext))

	return p
}

func (s *SvcOptNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SvcOptNodeContext) Option() antlr.TerminalNode {
	return s.GetToken(TronWalkerOption, 0)
}

func (s *SvcOptNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.EnterSvcOptNode(s)
	}
}

func (s *SvcOptNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.ExitSvcOptNode(s)
	}
}

func (p *TronWalker) SvcBody() (localctx ISvcBodyContext) {
	localctx = NewSvcBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TronWalkerRULE_svcBody)

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
	case TronWalkerOption:
		localctx = NewSvcOptNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(144)
			p.Match(TronWalkerOption)
		}

	case TronWalkerRpc:
		localctx = NewSvcRpcNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(145)
			p.Match(TronWalkerRpc)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *ExtendBodyContext) CopyFrom(ctx *ExtendBodyContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExtendBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtendBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExtFldNodeContext struct {
	*ExtendBodyContext
}

func NewExtFldNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExtFldNodeContext {
	var p = new(ExtFldNodeContext)

	p.ExtendBodyContext = NewEmptyExtendBodyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExtendBodyContext))

	return p
}

func (s *ExtFldNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtFldNodeContext) Field() antlr.TerminalNode {
	return s.GetToken(TronWalkerField, 0)
}

func (s *ExtFldNodeContext) Datastructure() antlr.TerminalNode {
	return s.GetToken(TronWalkerDatastructure, 0)
}

func (s *ExtFldNodeContext) Option() antlr.TerminalNode {
	return s.GetToken(TronWalkerOption, 0)
}

func (s *ExtFldNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.EnterExtFldNode(s)
	}
}

func (s *ExtFldNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.ExitExtFldNode(s)
	}
}

type ExtOptNodeContext struct {
	*ExtendBodyContext
}

func NewExtOptNodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExtOptNodeContext {
	var p = new(ExtOptNodeContext)

	p.ExtendBodyContext = NewEmptyExtendBodyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExtendBodyContext))

	return p
}

func (s *ExtOptNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtOptNodeContext) Option() antlr.TerminalNode {
	return s.GetToken(TronWalkerOption, 0)
}

func (s *ExtOptNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.EnterExtOptNode(s)
	}
}

func (s *ExtOptNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TronWalkerListener); ok {
		listenerT.ExitExtOptNode(s)
	}
}

func (p *TronWalker) ExtendBody() (localctx IExtendBodyContext) {
	localctx = NewExtendBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TronWalkerRULE_extendBody)
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

	p.SetState(156)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TronWalkerOption:
		localctx = NewExtOptNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(148)
			p.Match(TronWalkerOption)
		}

	case TronWalkerField:
		localctx = NewExtFldNodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(149)
			p.Match(TronWalkerField)
		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TronWalkerDatastructure {
			{
				p.SetState(150)
				p.Match(TronWalkerDatastructure)
			}

		}
		p.SetState(154)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(153)
				p.Match(TronWalkerOption)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
