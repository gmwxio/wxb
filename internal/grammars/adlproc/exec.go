package adlproc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	antlr "github.com/wxio/goantlr"
	walker "github.com/wxio/wxb/internal/grammars/adlwalker"
)

type readast struct {
	File string `type:"arg" help:"adl ast file" predict:"files"`
}

func (et *readast) Run() error {
	by, err := ioutil.ReadFile(et.File)
	if err != nil {
		return err
	}
	var m Module
	err = json.Unmarshal(by, &m)
	if err != nil {
		return err
	}
	// fmt.Printf("%v\n", m)
	b2, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", string(b2))
	return nil
}

type jsonObj struct {
}

func (et *jsonObj) Run() error {
	x := JsonObj{
		JsonObjs: map[string]Json{"a": "b"},
	}
	// x := map[string]interface{}{
	// 	"a": "b",
	// }
	b2, err := json.MarshalIndent(x, "", "    ")
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", string(b2))

	jstr := `{"a" : [
		
		"b",
		 1, 
		 
		 2, "3", true], "c" : 2 }`
	tr, js, err := UnmarshalJSON([]byte(jstr))
	fmt.Printf("js:%v\ntr:%v\n", js, tr)
	// err = json.Unmarshal(by, &m)
	if err != nil {
		return err
	}
	b2, err = json.MarshalIndent(js, "", "    ")
	if err != nil {
		return err
	}
	fmt.Printf("js - %v\n", string(b2))
	// err = WalkJSON(tr, &walkListener{})
	// if err != nil {
	// 	return err
	// }
	obj, err := VisitJSON(tr, &treeVisitor{})
	if err != nil {
		return err
	}
	fmt.Printf("obj - %v\n", obj)
	return nil

}

type walkListener struct {
	// walker.ADLWalkerListener
	*antlr.BaseParseTreeListener
	ruleCount int
	indent    string
}

func (tr *walkListener) VisitTerminal(node antlr.TerminalNode) {
	tr.ruleCount++
	fmt.Printf("%s  '%T'\n", tr.indent, node.GetPayload())
}
func (tr *walkListener) VisitErrorNode(node antlr.ErrorNode) {
	tr.ruleCount++
	tid := node.GetSymbol().GetTokenType()
	sym := node.GetSymbol()
	fmt.Printf("2.ERROR #%d %d:%d  len:%d start_stop:%d:%d  tok_type:%d %v %+v\n", tr.ruleCount, sym.GetLine(), sym.GetColumn(), len(sym.GetText()), sym.GetStart(), sym.GetStop(), tid, sym, node)
}
func (tr *walkListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	// tr.ruleCount++
	fmt.Printf("%s>>%T\n", tr.indent, ctx)
	tr.indent += "  "
}

type treeVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *treeVisitor) VisitJsonStr(ctx walker.IJsonStrContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	fmt.Printf("i'm a string '%s'\n", ctx.GetTok())
	return
}

type adlVisitor struct {
	File string `type:"arg" help:"proto file" predict:"files"`

	*antlr.BaseParseTreeVisitor `opts:"-"`
	indent                      string
}

func (et *adlVisitor) Run() error {
	by, err := ioutil.ReadFile(et.File)
	if err != nil {
		return err
	}
	tr, adl, err := UnmarshalADL(by)
	if err != nil {
		return err
	}
	fmt.Printf("%v\n%v\n", tr, adl)
	fmt.Printf("---\n")
	VisitADL(tr, et)
	if err != nil {
		return err
	}
	b2, err := json.MarshalIndent(adl, "", "    ")
	if err != nil {
		return err
	}
	fmt.Printf("js - %v\n", string(b2))

	return nil
}

func (v *adlVisitor) VisitTerminal(node antlr.TerminalNode) {
	fmt.Printf("%s term:%v\n", v.indent, node.GetPayload())
}
func (v *adlVisitor) VisitErrorNode(node antlr.ErrorNode) {
	fmt.Printf("ERROR%s %T\n", v.indent, node)
}
func (v *adlVisitor) EnterEveryRule(ctx antlr.RuleNode) {
	v.indent += " "
	fmt.Printf("%s %T\n", v.indent, ctx)
}
func (v *adlVisitor) ExitEveryRule(ctx antlr.RuleNode) {
	v.indent = v.indent[:len(v.indent)-1]
}
