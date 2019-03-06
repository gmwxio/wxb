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
	is := antlr.NewInputStream(string(by))
	lexer := parser.NewtronLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewtronParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(&tron_listener{baseR: lexer.BaseRecognizer}, p.Proto())
	return nil
}

type tron_listener struct {
	*parser.BasetronParserListener
	//
	baseR  *antlr.BaseRecognizer
	indent string
}

func (tr *tron_listener) VisitTerminal(node antlr.TerminalNode) {
	fmt.Printf("%s%v\n", tr.indent, node)
}
func (tr *tron_listener) VisitErrorNode(node antlr.ErrorNode) {
	tid := node.GetSymbol().GetTokenType()
	fmt.Printf("ERROR %d %#+v\n", tid, tr.baseR.LiteralNames[tid])
}
func (tr *tron_listener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("%s>>%T\n", tr.indent, ctx)
	tr.indent += "  "
}
func (tr *tron_listener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("%s<<%T\n", tr.indent, ctx)
	tr.indent = tr.indent[2:]
}
