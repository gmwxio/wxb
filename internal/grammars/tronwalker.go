package tron

import (
	"fmt"
	"io/ioutil"
	"strings"

	antlr "github.com/wxio/goantlr"
	"github.com/wxio/wxb/internal/ctree"
	"github.com/wxio/wxb/internal/grammars/walker"
)

type walkTron struct {
	File string `type:"arg" help:"proto file" predict:"files"`
}

type TTType struct{}

func (*TTType) Eof() int  { return walker.TronWalkerEOF }
func (*TTType) Down() int { return walker.TronWalkerDOWN }
func (*TTType) Up() int   { return walker.TronWalkerUP }

func (et *walkTron) Run() error {
	by, err := ioutil.ReadFile(et.File)
	if err != nil {
		return err
	}
	tr, _, _, err := BuildTronAST(string(by))
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", tr)

	// iterator := ctree.NewPreOrderTreeIterator(tr, tr.Root())
	// for iterator.HasNext() {
	// 	n := iterator.Next()
	// 	fmt.Printf("%#+v\n", n)
	// }

	var tttype *TTType
	var tts antlr.TokenStream = ctree.NewTreeTokenSource(tr, tttype)
	p := walker.NewTronWalker(tts)
	i := 1
	for {
		to := tts.Get(i)
		if -1 == to.GetTokenType() {
			break
		}
		fmt.Printf("%d ", to.GetTokenType())
		fmt.Printf("%d : %v\n", i, p.BaseRecognizer.SymbolicNames[to.GetTokenType()])
		i++
	}
	// return nil

	p.SetTokenStream(tts)
	// p.RemoveErrorListeners()
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tbl := &WalkListener{}
	tbl.debug = true
	// p.AddErrorListener(tbl)
	// p.BuildParseTrees = false
	ret := p.Proto()
	// fmt.Printf("ret %#+v\n", tbl)
	antlr.ParseTreeWalkerDefault.Walk(tbl, ret)
	return tbl.err
}

type WalkListener struct {
	*walker.BaseTronWalkerListener
	//
	indent  string
	debug   bool
	warning string
	err     error
}

// EnterTld is called when production tld is entered.
func (tr *WalkListener) EnterTld(ctx *walker.TldContext) {
	// fmt.Printf("tld %v\n", ctx)
}

func (tr *WalkListener) VisitTerminal(node antlr.TerminalNode) {
	fmt.Printf("%v", node)
}
func (tr *WalkListener) VisitErrorNode(node antlr.ErrorNode) {
	tid := node.GetSymbol().GetTokenType()
	fmt.Printf("\n2.ERROR %d", tid)
}
func (tr *WalkListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("\n%s>>%T kids:%d ", tr.indent, ctx, ctx.GetChildCount())
	tr.indent += "  "
}
func (tr *WalkListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	tr.indent = tr.indent[2:]
	// fmt.Printf("\n%s<<%T", tr.indent, ctx)
}

func (tbl *WalkListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	if tbl.debug {
		fmt.Printf("SyntaxError %d:%d <%s>\n", line, column, msg)
	}
	if strings.HasPrefix(msg, "report") { // TODO remove NewDiagnosticErrorListener and move warning to ReportAmbiguity etc. when getDecisionDescription is make public
		tbl.warning = fmt.Sprintf("At %d:%d <%s>", line, column, msg)
	} else {
		tbl.err = fmt.Errorf("SyntaxError %d:%d <%s>", line, column, msg)
	}
	// panic("line " + strconv.Itoa(line) + ":" + strconv.Itoa(column) + " " + msg)
}

func (tbl *WalkListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	// panic("ReportAmbiguity")
	if tbl.debug {
		fmt.Printf("ReportAmbiguity rec:%v dfs:%v start:%d stop:%d, exact:%v, ambigAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
	}
}

func (tbl *WalkListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	// panic("ReportAttemptingFullContext")
	if tbl.debug {
		fmt.Printf("ReportAttemptingFullContext rec:%v dfs:%v start:%d stop:%d, conflictingAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
	}
}

func (tbl *WalkListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	panic("ReportContextSensitivity")
}
