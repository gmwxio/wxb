package adlproc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	antlr "github.com/wxio/goantlr"
	"github.com/wxio/wxb/internal/ctree"
	parser "github.com/wxio/wxb/internal/grammars/adl"
	walker "github.com/wxio/wxb/internal/grammars/adlwalker"
)

type Json interface {
}
type JsonElems []Json
type JsonObjs map[string]Json
type JsonStr struct {
	// MyToken `json:"-"`
	string
}
type JsonBool struct {
	// MyToken `json:"-"`
	bool
}
type JsonNull struct {
	// MyToken `json:"-"`
}
type JsonInt struct {
	// MyToken `json:"-"`
	int64
}
type JsonFloat struct {
	// MyToken `json:"-"`
	float64
}
type JsonArray struct {
	// MyToken   `json:"-"`
	JsonElems []Json
}
type JsonObj struct {
	// MyToken `json:"-"`
	JsonObjs
}

func (a JsonStr) String() string { return a.string }
func (a JsonBool) String() string {
	if a.bool {
		return "true"
	}
	return "false"
}
func (a JsonNull) String() string  { return "null" }
func (a JsonInt) String() string   { return strconv.FormatInt(a.int64, 10) }
func (a JsonFloat) String() string { return fmt.Sprintf("%v", a.float64) }
func (a JsonArray) String() string {
	b := bytes.Buffer{}
	b.WriteString("[")
	for i, v := range a.JsonElems {
		if i != 0 {
			b.WriteString(",")
		}
		b.WriteString(fmt.Sprintf("%v", v))
	}
	b.WriteString("]")
	return b.String()
}
func (a JsonObj) String() string {
	b := bytes.Buffer{}
	b.WriteString("{")
	f := true
	for k, v := range a.JsonObjs {
		if f {
			f = false
		} else {
			b.WriteString(",")
		}
		b.WriteString(k)
		b.WriteString(":")
		b.WriteString(fmt.Sprintf("%v", v))
	}
	b.WriteString("}")
	return b.String()
}

func (a JsonStr) MarshalJSON() ([]byte, error)   { return []byte(`"` + a.string + `"`), nil }
func (a JsonBool) MarshalJSON() ([]byte, error)  { return json.Marshal(a.bool) }
func (a JsonNull) MarshalJSON() ([]byte, error)  { return json.Marshal(nil) }
func (a JsonInt) MarshalJSON() ([]byte, error)   { return json.Marshal(a.int64) }
func (a JsonFloat) MarshalJSON() ([]byte, error) { return json.Marshal(a.float64) }
func (a JsonArray) MarshalJSON() ([]byte, error) { return json.Marshal(a.JsonElems) }
func (a JsonObj) MarshalJSON() ([]byte, error)   { return json.Marshal(a.JsonObjs) }

func UnmarshalJSON(by []byte) (ctree.Tree, Json, error) {
	errListener := &jsonErrListener{}
	lexer := parser.NewadlLexer(antlr.NewInputStream(string(by)))
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
	if errListener.err != nil {
		return nil, nil, errListener.err
	}
	vi := &jsonVisitor{}
	js := ctx.Visit(vi)
	tr := vi.bldr.Build()
	return tr, js, vi.err
}

func WalkJSON(tr ctree.Tree, list antlr.ParseTreeListener) error {
	var tttype *TTType
	var tts antlr.TokenStream = ctree.NewTreeTokenSource(tr, tttype)
	p := walker.NewADLWalker(tts)
	// debugTreeToken(tts, p)
	p.SetTokenStream(tts)
	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	errListener := &jsonErrListener{}
	p.AddErrorListener(errListener)
	p.BuildParseTrees = true
	jv := p.Json()
	antlr.ParseTreeWalkerDefault.Walk(list, jv)
	return errListener.err
}

func VisitJSON(tr ctree.Tree, vi antlr.ParseTreeVisitor) (Json, error) {
	var tttype *TTType
	var tts antlr.TokenStream = ctree.NewTreeTokenSource(tr, tttype)
	p := walker.NewADLWalker(tts)
	// debugTreeToken(tts, p)
	p.SetTokenStream(tts)
	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	errListener := &jsonErrListener{}
	p.AddErrorListener(errListener)
	ctx := p.Json()
	obj := ctx.Visit(vi)
	return obj, errListener.err
}

func debugTreeToken(tts antlr.TokenStream, p antlr.Recognizer) {
	i := 1
	for {
		to := tts.Get(i)
		if -1 == to.GetTokenType() {
			break
		}
		fmt.Printf("%d ", to.GetTokenType())
		fmt.Printf("%d : %v\t\t%v\n", i,
			p.GetSymbolicNames()[to.GetTokenType()],
			to.GetLine(),
		)
		i++
	}
}

type jsonErrListener struct {
	err error
}
type jsonVisitor struct {
	*antlr.BaseParseTreeVisitor
	bldr ctree.WalkableBuilder
	// js  Json
	err error
}

var _ parser.ObjStatementContextVisitor = &jsonVisitor{}

func (v *jsonVisitor) VisitStringStatement(ctx parser.IStringStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	fmt.Printf("ctv %T %s\n", ctx, ctx.GetS().GetText())
	s := ctx.GetS().GetText()
	result = JsonStr{
		// MyToken: MyToken{Token: ctx.GetS(), TType: parser.ADLParserJsonStr},
		string: s[1 : len(s)-1]}
	v.bldr.AddNode(ctx.GetStart(), parser.ADLParserJsonStr, result)
	return
}
func (v *jsonVisitor) VisitTrueFalseNull(ctx parser.ITrueFalseNullContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	fmt.Printf("ctv %T\n", ctx)
	switch strings.ToLower(ctx.GetKw().GetText()) {
	case "true":
		result = &JsonBool{
			// MyToken: MyToken{Token: ctx.GetKw(), TType: parser.ADLParserJsonBool},
			bool: true,
		}
		v.bldr.AddNode(ctx.GetStart(), parser.ADLParserJsonBool, result)
	case "false":
		result = &JsonBool{
			// MyToken: MyToken{Token: ctx.GetKw(), TType: parser.ADLParserJsonBool},
			bool: false,
		}
		v.bldr.AddNode(ctx.GetStart(), parser.ADLParserJsonBool, result)
	case "null":
		result = &JsonNull{
			// MyToken: MyToken{Token: ctx.GetKw(), TType: parser.ADLParserJsonNull},
		}
		v.bldr.AddNode(ctx.GetStart(), parser.ADLParserJsonNull, result)
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
			// MyToken: MyToken{Token: ctx.GetStart(), TType: parser.ADLParserJsonInt},
			int64: i,
		}
		v.bldr.AddNode(ctx.GetStart(), parser.ADLParserJsonInt, result)
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
			// MyToken: MyToken{Token: ctx.GetStart(), TType: parser.ADLParserJsonFloat},
			float64: i,
		}
		v.bldr.AddNode(ctx.GetStart(), parser.ADLParserJsonFloat, result)
	}
	return
}
func (v *jsonVisitor) VisitArrayStatement(ctx parser.IArrayStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	fmt.Printf("ctv %T\n", ctx)
	ja := &JsonArray{
		// MyToken: MyToken{Token: ctx.GetStart(), TType: parser.ADLParserJsonArray}
	}
	v.bldr.AddNode(ctx.GetStart(), parser.ADLParserJsonArray, ja).Down()
	defer v.bldr.Up()
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
	jo := &JsonObj{
		// MyToken:  MyToken{Token: ctx.GetStart(), TType: parser.ADLParserJsonObj},
		JsonObjs: make(map[string]Json),
	}
	if v.bldr == nil {
		v.bldr = ctree.NewBuild("TREE", ctx.GetStart(), parser.ADLParserJsonObj, jo)
	} else {
		v.bldr.AddNode(ctx.GetStart(), parser.ADLParserJsonObj, jo).Down()
		defer v.bldr.Up()
	}
	vs := ctx.GetV()
	for i, child := range ctx.GetK() {
		k := child.GetText()
		jo.JsonObjs[k[1:len(k)-1]] = vs[i].Visit(delegate, args...).(Json)
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
