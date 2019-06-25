package adlproc

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	antlr "github.com/wxio/goantlr"
	parser "github.com/wxio/wxb/internal/grammars/adl"
)

type Json interface {
}
type JsonElems []Json
type JsonObjs map[string]Json
type JsonStr struct {
	MyToken `json:"-"`
	string
}
type JsonBool struct {
	MyToken `json:"-"`
	bool
}
type JsonNull struct {
	MyToken `json:"-"`
}
type JsonInt struct {
	MyToken `json:"-"`
	int64
}
type JsonFloat struct {
	MyToken `json:"-"`
	float64
}
type JsonArray struct {
	MyToken   `json:"-"`
	JsonElems []Json
}
type JsonObj struct {
	MyToken `json:"-"`
	JsonObjs
}

func (a JsonStr) MarshalJSON() ([]byte, error)   { return []byte(a.string), nil }
func (a JsonBool) MarshalJSON() ([]byte, error)  { return json.Marshal(a.bool) }
func (a JsonNull) MarshalJSON() ([]byte, error)  { return json.Marshal(nil) }
func (a JsonInt) MarshalJSON() ([]byte, error)   { return json.Marshal(a.int64) }
func (a JsonFloat) MarshalJSON() ([]byte, error) { return json.Marshal(a.float64) }
func (a JsonArray) MarshalJSON() ([]byte, error) { return json.Marshal(a.JsonElems) }
func (a JsonObj) MarshalJSON() ([]byte, error)   { return json.Marshal(a.JsonObjs) }

func UnmarshalJSON(by []byte) (Json, error) {
	errListener := &jsonErrListener{}
	is := antlr.NewInputStream(string(by))
	lexer := parser.NewadlLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errListener)
	//
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewADLParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.AddErrorListener(errListener)
	// p.BuildParseTrees = true
	ctx := p.JsonValue()
	fmt.Printf("ctv %T\n", ctx)
	if errListener.err != nil {
		return nil, errListener.err
	}
	v := &jsonVisitor{}
	js := ctx.Visit(v)
	// js := v.Visit(ctx, v).(Json)
	return js, v.err
}

type jsonErrListener struct {
	err error
}
type jsonVisitor struct {
	*antlr.BaseParseTreeVisitor
	// js  Json
	err error
}

var _ parser.ObjStatementContextVisitor = &jsonVisitor{}

func (v *jsonVisitor) VisitStringStatement(ctx parser.IStringStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	fmt.Printf("ctv %T\n", ctx)
	result = &JsonStr{MyToken: MyToken{Token: ctx.GetS(), TType: parser.ADLParserJsonStr}, string: ctx.GetS().GetText()}
	return
}
func (v *jsonVisitor) VisitTrueFalseNull(ctx parser.ITrueFalseNullContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	fmt.Printf("ctv %T\n", ctx)
	switch strings.ToLower(ctx.GetKw().GetText()) {
	case "true":
		result = &JsonBool{
			MyToken: MyToken{Token: ctx.GetKw(), TType: parser.ADLParserJsonBool},
			bool:    true,
		}
	case "false":
		result = &JsonBool{
			MyToken: MyToken{Token: ctx.GetKw(), TType: parser.ADLParserJsonBool},
			bool:    false,
		}
	case "null":
		result = &JsonNull{
			MyToken: MyToken{Token: ctx.GetKw(), TType: parser.ADLParserJsonNull},
		}
	default:
		v.err = fmt.Errorf("Expected: \"true|false|null\", Received: %s", ctx.GetKw().GetText())
	}
	return
}
func (v *jsonVisitor) VisitNumberStatement(ctx parser.INumberStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	fmt.Printf("ctv %T\n", ctx)
	i, err := strconv.ParseInt(ctx.GetN().GetText(), 10, 64)
	if err != nil {
		v.err = fmt.Errorf("Expected: <number>, Received: %s", ctx.GetN().GetText())
	} else {
		result = &JsonInt{
			MyToken: MyToken{Token: ctx.GetStart(), TType: parser.ADLParserJsonInt},
			int64:   i,
		}
	}
	return
}
func (v *jsonVisitor) VisitFloatStatement(ctx parser.IFloatStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	fmt.Printf("ctv %T\n", ctx)
	i, err := strconv.ParseFloat(ctx.GetF().GetText(), 64)
	if err != nil {
		v.err = fmt.Errorf("Expected: <float>, Received: %s", ctx.GetF().GetText())
	} else {
		result = &JsonFloat{
			MyToken: MyToken{Token: ctx.GetStart(), TType: parser.ADLParserJsonFloat},
			float64: i,
		}
	}
	return
}
func (v *jsonVisitor) VisitArrayStatement(ctx parser.IArrayStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	fmt.Printf("ctv %T\n", ctx)
	ja := &JsonArray{MyToken: MyToken{Token: ctx.GetStart(), TType: parser.ADLParserJsonArray}}
	for _, child := range ctx.GetChildren() {
		switch child := child.(type) {
		case antlr.TerminalNode:
			// delegate.VisitTerminal(child)
		case antlr.ErrorNode:
			v.err = fmt.Errorf("error node %v", child)
			return
		case antlr.RuleNode:
			r := child.Visit(delegate, args...)
			if r != nil {
				ja.JsonElems = append(ja.JsonElems, r.(Json))
			}
		default:
			// can this happen??
		}
	}
	result = ja
	return
}
func (v *jsonVisitor) VisitObjStatement(ctx parser.IObjStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	fmt.Printf("ctv %T\n", ctx)
	jo := &JsonObj{MyToken: MyToken{Token: ctx.GetStart(), TType: parser.ADLParserJsonObj}, JsonObjs: make(map[string]Json)}
	vs := ctx.GetV()
	for i, child := range ctx.GetK() {
		jo.JsonObjs[child.GetText()] = vs[i].Visit(delegate, args...)
	}
	result = jo
	return
}

func (d *jsonErrListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	d.err = fmt.Errorf("line " + strconv.Itoa(line) + ":" + strconv.Itoa(column) + " " + msg)
}

func (d *jsonErrListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	d.err = fmt.Errorf("ReportAmbiguity rec:%v dfs:%v start:%d stop:%d, exact:%v, ambigAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
}

func (d *jsonErrListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	d.err = fmt.Errorf("ReportAttemptingFullContext rec:%v dfs:%v start:%d stop:%d, conflictingAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
}

func (d *jsonErrListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	d.err = fmt.Errorf("ReportContextSensitivity rec:%v dfs:%v start:%d stop:%d, config:%v\n", recognizer, dfa, startIndex, stopIndex, configs)
}

// func (tbl *jsonBuildListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
// 	// t, ok := offendingSymbol.(antlr.Token)
// 	// if !ok && e != nil {
// 	// 	t = e.GetOffendingToken()
// 	// }
// 	tbl.err = fmt.Errorf("Expected %v Received %v", msg, offendingSymbol)
// }

// func (tbl *jsonBuildListener) ReportAmbiguity(
// 	recognizer antlr.Parser,
// 	dfa *antlr.DFA,
// 	startIndex, stopIndex int,
// 	exact bool,
// 	ambigAlts *antlr.BitSet,
// 	configs antlr.ATNConfigSet) {
// 	tbl.err = fmt.Errorf("ReportAmbiguity rec:%v dfs:%v start:%d stop:%d, exact:%v, ambigAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
// }

// func (tbl *jsonBuildListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
// 	tbl.err = fmt.Errorf("ReportAttemptingFullContext rec:%v dfs:%v start:%d stop:%d, conflictingAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
// }

// func (tbl *jsonBuildListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
// 	tbl.err = fmt.Errorf("ReportContextSensitivity rec:%v dfs:%v start:%d stop:%d, config:%v\n", recognizer, dfa, startIndex, stopIndex, configs)
// }
