package tron

import (
	"fmt"
	"testing"

	antlr "github.com/wxio/goantlr"

	parser "github.com/wxio/wxb/internal/grammars/tron"
)

var example = `
syntax = "proto3"; 
import public "other.proto";
option java_package = "com.example.foo";
enum EnumAllowingAlias {
  option allow_alias = true;
  UNKNOWN = 0;
  STARTED = 1;
  RUNNING = 2 [(custom_option) = "hello world"];
}
message outer {
  option (my_option).a = true;
  message inner {   // Level 2
    int64 ival = 1;
  }
  repeated inner inner_message = 2;
  EnumAllowingAlias enum_field =3;
  map<int32, string> my_map = 4;
}
`

type tron_test_listener struct {
	*parser.BasetronParserListener
	indent string
}

func (tr *tron_test_listener) VisitTerminal(node antlr.TerminalNode) {
	fmt.Printf("%s%v\n", tr.indent, node)
}
func (tr *tron_test_listener) VisitErrorNode(node antlr.ErrorNode) {
	fmt.Printf("ERROR %#+v\n", node.GetSymbol().GetTokenType())
}
func (tr *tron_test_listener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("%s>>%T\n", tr.indent, ctx)
	tr.indent += "  "
}
func (tr *tron_test_listener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	tr.indent = tr.indent[2:]
}

func TestExample(t *testing.T) {
	// Setup the input
	is := antlr.NewInputStream(example)
	lexer := parser.NewtronLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewtronParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(&tron_test_listener{}, p.Proto())
}
