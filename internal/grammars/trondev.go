package tron

import (
	"fmt"
	"io/ioutil"

	"github.com/jpillora/opts"
	antlr "github.com/wxio/goantlr"

	parser "github.com/wxio/wxb/internal/grammars/tron"
	"github.com/wxio/wxb/internal/schemas/build"
)

type rootTron struct {
}

type exerciseTron struct {
	File string `type:"arg" help:"proto file" predict:"files"`
}

func Register(opts opts.Builder) opts.Builder {
	rt := rootTron{}
	rto := opts.SubCmd("tron", &rt)
	et := exerciseTron{}
	eto := rto.SubCmd("exec", &et)
	_ = eto
	wt := walkTron{}
	wto := rto.SubCmd("walk", &wt)
	_ = wto
	return rto
}

func (et *exerciseTron) Run() error {
	wxb := build.WxbBuild{}
	_ = wxb
	//
	by, err := ioutil.ReadFile(et.File)
	if err != nil {
		return err
	}
	tr, _, _, err := BuildTronAST(string(by))
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", tr)
	// return nil
	is := antlr.NewInputStream(string(by))
	lexer := parser.NewtronLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewTronParser(stream)
	p.AddErrorListener(antlr.NewConsoleErrorListener())
	ctx := p.Proto()
	fmt.Println("--------")
	antlr.ParseTreeWalkerDefault.Walk(&tron_listener{baseR: lexer.BaseRecognizer}, ctx)
	return nil
}

type tron_listener struct {
	*parser.BaseTronParserListener
	//
	baseR  *antlr.BaseRecognizer
	indent string
}

func (tr *tron_listener) VisitTerminal(node antlr.TerminalNode) {
	fmt.Printf(" %v", node)
}
func (tr *tron_listener) VisitErrorNode(node antlr.ErrorNode) {
	tid := node.GetSymbol().GetTokenType()
	if tid >= len(tr.baseR.LiteralNames) {
		fmt.Printf("ERROR %d\n", tid)
	} else {
		fmt.Printf("ERROR %d %#+v\n", tid, tr.baseR.LiteralNames[tid])
	}
}
func (tr *tron_listener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("\n%s>>%T ", tr.indent, ctx)
	tr.indent += "  "
	// switch ctx := ctx.(type) {
	// case *parser.EnumLeftContext:
	// 	fmt.Printf("\n%s>>%T ", tr.indent, ctx)
	// case *parser.Opt_SingleContext:
	// 	if ctx.GetA().GetText() == "option" {
	// 		fmt.Printf("\n%s>>%T %v", tr.indent, ctx, "option")
	// 	} else {
	// 		fmt.Printf("\n%s>>%T %v", tr.indent, ctx, "field")
	// 	}
	// case *parser.OptContext:
	// 	fmt.Printf("\n%s>>%T %v", tr.indent, ctx, "option")
	// case *parser.SingleFull_RepLocalContext:
	// 	if ctx.GetA().GetText() == "repeated" {
	// 		fmt.Printf("\n%s>>%T %v", tr.indent, ctx, "repeated local")
	// 	} else {
	// 		fmt.Printf("\n%s>>%T %v", tr.indent, ctx, "field")
	// 	}
	// case *parser.SingleLocalContext:
	// 	fmt.Printf("\n%s>>%T %v", tr.indent, ctx, "field local")
	// case *parser.RepeatedContext:
	// 	fmt.Printf("\n%s>>%T %v", tr.indent, ctx, "repeated field")
	// case *parser.MapLeftContext:
	// 	fmt.Printf("\n%s>>%T %v", tr.indent, ctx, "map")
	// }
}
func (tr *tron_listener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	tr.indent = tr.indent[2:]
	// fmt.Printf("\n%s<<%T", tr.indent, ctx)
}
