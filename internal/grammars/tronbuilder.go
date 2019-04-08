package tron

import (
	"fmt"
	"strconv"

	"strings"

	"github.com/golang/glog"
	antlr "github.com/wxio/goantlr"

	"github.com/wxio/wxb/internal/ctree"
	parser "github.com/wxio/wxb/internal/grammars/tron"
)

func BuildTronAST(str string) (ctree.Tree, *antlr.BaseLexer, antlr.TokenStream, error) {
	errListener := &errorListener{}
	tbl := &TronBuildListener{}
	is := antlr.NewInputStream(str)
	lexer := parser.NewtronLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errListener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewTronParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.AddErrorListener(tbl)
	p.BuildParseTrees = true
	ctx := p.Proto()
	fmt.Println("--------")
	antlr.ParseTreeWalkerDefault.Walk(tbl, ctx)
	if tbl.err != "" {
		return nil, lexer.BaseLexer, stream, fmt.Errorf("ERROR:%v", tbl.err)
	}
	return tbl.Builder.Build(), lexer.BaseLexer, stream, nil
}

type MyToken struct {
	antlr.Token
	TType int
}

func (t *MyToken) GetTokenType() int { return t.TType }

type ENode struct{ MyToken }
type ErrorNode struct {
	MyToken
	Expected string
	Received string
}
type ImportNode struct {
	MyToken
	Weak, Public bool
	Path         string
}
type PackageNode struct {
	MyToken
	Pkg []string
}
type OptionNode struct {
	MyToken
}
type MessageNode struct {
	MyToken
	Name []string
}
type EnumNode struct {
	MyToken
	Name []string
}
type ServiceNode struct {
	MyToken
	Name []string
}
type ExtendNode struct {
	MyToken
	Name []string
}
type OneofNode struct {
	MyToken
	Name []string
}
type EmptyNode struct {
	MyToken
}
type FieldNode struct {
	MyToken
	Local    bool
	TypeName []string
	Name     string
}
type DatastructNode struct {
	MyToken
	Type string
}
type MapFieldNode struct {
	MyToken
	Local bool
	Key   string
	Value []string
}
type EnumValueNode struct {
	MyToken
	Name  string
	Index int32
}

func (e *ENode) String() string          { return "" }
func (e *ErrorNode) String() string      { return "error expected:" + e.Expected + " got:" + e.Received }
func (e *ImportNode) String() string     { return "import" }
func (e *PackageNode) String() string    { return "package" }
func (e *OptionNode) String() string     { return "option" }
func (e *EmptyNode) String() string      { return ";" }
func (e *FieldNode) String() string      { return "field" }
func (e *MapFieldNode) String() string   { return "map" }
func (e *MessageNode) String() string    { return "message" }
func (e *EnumNode) String() string       { return "enum" }
func (e *EnumValueNode) String() string  { return "enumValue" }
func (e *ServiceNode) String() string    { return "service" }
func (e *ExtendNode) String() string     { return "extend" }
func (e *OneofNode) String() string      { return "oneof" }
func (e *DatastructNode) String() string { return "datastructure" }

type TronBuildListener struct {
	*parser.BaseTronParserListener
	Builder ctree.WalkableBuilder
	baseR   *antlr.BaseRecognizer
	indent  string

	warning string
	err     string
	debug   bool
}

// EnterEveryRule is called when any rule is entered.
func (tr *TronBuildListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	switch ctx := ctx.(type) {
	case *parser.ProtoContext:
		r := &ENode{MyToken: MyToken{Token: ctx.GetStart(), TType: parser.TronParserROOT}}
		tr.Builder = ctree.NewWalkableBuild("TREE", r)
	case *parser.ImportStatementContext:
		if ctx.GetImp().GetText() != "import" {
			n := &ErrorNode{MyToken: MyToken{Token: ctx.GetImp(), TType: parser.TronParserERROR}, Expected: "import", Received: ctx.GetImp().GetText()}
			tr.Builder.Add(n)
		} else {
			n := &ImportNode{
				MyToken: MyToken{Token: ctx.GetImp(), TType: parser.TronParserImport},
				Path:    strings.Trim(ctx.GetB().GetText(), `"`),
			}
			tr.Builder.Add(n)
		}
	case *parser.PackageStatementContext:
		switch ctx.GetPkg().GetText() {
		case "package":
			n := &MessageNode{
				MyToken: MyToken{Token: ctx.GetPkg(), TType: parser.TronParserPackage},
			}
			name := make([]string, len(ctx.GetA()))
			for i, tks := range ctx.GetA() {
				name[i] = tks.GetText()
			}
			n.Name = name
			tr.Builder.Add(n)
		default:
			n := &ErrorNode{
				MyToken:  MyToken{Token: ctx.GetPkg(), TType: parser.TronParserERROR},
				Expected: "package",
				Received: ctx.GetPkg().GetText(),
			}
			tr.Builder.Add(n)
		}
	case *parser.OptionFileDefContext:
	case *parser.Ext_Msg_Enum_SvcContext:
		switch ctx.GetMese().GetText() {
		case "message":
			n := &MessageNode{
				MyToken: MyToken{Token: ctx.GetMese(), TType: parser.TronParserMessage},
				//Name: ctx.GetA(),
			}
			tr.Builder.Add(n)
			tr.Builder.Down()
		case "enum":
			n := &EnumNode{
				MyToken: MyToken{Token: ctx.GetMese(), TType: parser.TronParserEnum},
			} //, Name: ctx.GetA()}
			tr.Builder.Add(n)
			tr.Builder.Down()
		case "service":
			n := &ServiceNode{
				MyToken: MyToken{Token: ctx.GetMese(), TType: parser.TronParserService},
			} //,, Name: ctx.GetA()}
			tr.Builder.Add(n)
			tr.Builder.Down()
		case "extend":
			n := &ExtendNode{
				MyToken: MyToken{Token: ctx.GetMese(), TType: parser.TronParserExtend},
			} //,, Name: ctx.GetA()}
			tr.Builder.Add(n)
			tr.Builder.Down()
		// case "oneof":
		// 	n := &OneofNode{
		// 		MyToken: MyToken{Token: ctx.GetMese(), TType: parser.TronParserOneof},
		// 	} //,, Name: ctx.GetA()}
		// 	tr.Builder.Add(n)
		default:
			n := &ErrorNode{
				MyToken:  MyToken{Token: ctx.GetMese(), TType: parser.TronParserERROR},
				Expected: "msg|enum|svc|extend",
				Received: ctx.GetMese().GetText(),
			}
			tr.Builder.Add(n)
			tr.Builder.Down()
		}
	case *parser.EmptyStmContext:
	case *parser.AssocContext:
	case *parser.EmptyTopLvlContext:
		switch ctx.GetMsg().GetText() {
		case "message":
			n := &MessageNode{
				MyToken: MyToken{Token: ctx.GetMsg(), TType: parser.TronParserMessage},
			} //, Name: ctx.GetA()}
			tr.Builder.Add(n)
		case "enum":
			n := &EnumNode{
				MyToken: MyToken{Token: ctx.GetMsg(), TType: parser.TronParserEnum},
			} //, Name: ctx.GetA()}
			tr.Builder.Add(n)
		case "service":
			//TODO this should be an error - as per the spec
			n := &ServiceNode{
				MyToken: MyToken{Token: ctx.GetMsg(), TType: parser.TronParserService},
			} //,, Name: ctx.GetA()}
			tr.Builder.Add(n)
		case "extend":
			n := &ExtendNode{
				MyToken: MyToken{Token: ctx.GetMsg(), TType: parser.TronParserExtend},
			} //,, Name: ctx.GetA()}
			tr.Builder.Add(n)
		case "oneof":
			n := &OneofNode{
				MyToken: MyToken{Token: ctx.GetMsg(), TType: parser.TronParserOneof},
			} //,, Name: ctx.GetA()}
			tr.Builder.Add(n)
		default:
			n := &ErrorNode{
				MyToken:  MyToken{Token: ctx.GetMsg(), TType: parser.TronParserERROR},
				Expected: "msg|enum|svc|extend",
				Received: ctx.GetMsg().GetText(),
			}
			tr.Builder.Add(n)
		}
	case *parser.MsgEnumSvcExtContext:
		switch ctx.GetMese().GetText() {
		case "message":
			n := &MessageNode{
				MyToken: MyToken{Token: ctx.GetMese(), TType: parser.TronParserMessage},
			} //, Name: ctx.GetA()}
			tr.Builder.Add(n)
			tr.Builder.Down()
		case "enum":
			n := &EnumNode{
				MyToken: MyToken{Token: ctx.GetMese(), TType: parser.TronParserEnum},
			} //, Name: ctx.GetA()}
			tr.Builder.Add(n)
			tr.Builder.Down()
		case "service":
			n := &ServiceNode{
				MyToken: MyToken{Token: ctx.GetMese(), TType: parser.TronParserService},
			} //,, Name: ctx.GetA()}
			tr.Builder.Add(n)
			tr.Builder.Down()
		case "extend":
			n := &ExtendNode{
				MyToken: MyToken{Token: ctx.GetMese(), TType: parser.TronParserExtend},
			} //,, Name: ctx.GetA()}
			tr.Builder.Add(n)
			tr.Builder.Down()
		case "oneof":
			n := &OneofNode{
				MyToken: MyToken{Token: ctx.GetMese(), TType: parser.TronParserOneof},
			} //,, Name: ctx.GetA()}
			tr.Builder.Add(n)
			tr.Builder.Down()
		default:
			n := &ErrorNode{
				MyToken:  MyToken{Token: ctx.GetMese(), TType: parser.TronParserERROR},
				Expected: "msg|enum|svc|extend",
				Received: ctx.GetMese().GetText(),
			}
			tr.Builder.Add(n)
			tr.Builder.Down()
		}
	case *parser.EmptyTopLvlStmContext:
	case *parser.RangeContext:
	case *parser.TLIOptionContext:
	case *parser.RPCSigContext:
		switch ctx.GetRpcID().GetText() {
		case "rpc":
			n := &MessageNode{
				MyToken: MyToken{Token: ctx.GetRpcID(), TType: parser.TronParserRpc},
			} //, Name: ctx.GetA()}
			tr.Builder.Add(n)
			tr.Builder.Down()
		default:
			n := &ErrorNode{
				MyToken:  MyToken{Token: ctx.GetRpcID(), TType: parser.TronParserERROR},
				Expected: "rpc",
				Received: ctx.GetRpcID().GetText(),
			}
			tr.Builder.Add(n)
			tr.Builder.Down()
		}
	case *parser.EmptyStmStmContext:
	case *parser.EnumLeftContext:
		v, err := strconv.Atoi(ctx.GetV().GetText())
		if err != nil {
			panic("")
		}
		n := &EnumValueNode{
			MyToken: MyToken{Token: ctx.GetA(), TType: parser.TronParserEnumValue},
			Name:    ctx.GetA().GetText(),
			Index:   int32(v),
		}
		tr.Builder.Add(n)
	case *parser.Opt_SingleContext:
		if ctx.GetA().GetText() == "option" {
			n := &OptionNode{
				MyToken: MyToken{Token: ctx.GetA(), TType: parser.TronParserOption},
			} //,, Name: ctx.GetA()}
			tr.Builder.Add(n)
		} else {
			n := &FieldNode{
				MyToken: MyToken{Token: ctx.GetA(), TType: parser.TronParserField},
			} //,, Name: ctx.GetA()}
			tr.Builder.Add(n)
		}
	case *parser.OptContext:
		if ctx.GetA().GetText() != "option" {
			n := &ErrorNode{
				MyToken:  MyToken{Token: ctx.GetA(), TType: parser.TronParserERROR},
				Expected: "option",
				Received: ctx.GetA().GetText(),
			}
			tr.Builder.Add(n)
		} else {
			n := &OptionNode{
				MyToken: MyToken{Token: ctx.GetA(), TType: parser.TronParserOption},
			} //, Path: ctx.GetB().GetText()}
			tr.Builder.Add(n)
		}
	case *parser.SingleFull_RepLocalContext:
		n := &FieldNode{
			MyToken: MyToken{Token: ctx.GetB()[0], TType: parser.TronParserField},
		} //,, Name: ctx.GetA()}
		tname := make([]string, len(ctx.GetB()))
		for i, tkb := range ctx.GetB() {
			tname[i] = tkb.GetText()
		}
		n.TypeName = tname
		n.Name = ctx.GetC().GetText()
		tr.Builder.Add(n)
	case *parser.SingleLocalContext:
		// fmt.Printf("\n%s>>%T %v", tr.indent, ctx, "field local")
	case *parser.RepeatedContext:
		tname := make([]string, len(ctx.GetB()))
		for i, tks := range ctx.GetB() {
			tname[i] = tks.GetText()
		}
		n := &FieldNode{
			MyToken:  MyToken{Token: ctx.GetB()[0], TType: parser.TronParserField},
			TypeName: tname,
			Name:     ctx.GetC().GetText(),
		} //,, Name: ctx.GetA()}
		tr.Builder.Add(n)
		ds := &DatastructNode{
			MyToken: MyToken{Token: ctx.GetA(), TType: parser.TronParserDatastructure},
			Type:    ctx.GetA().GetText(),
		}
		tr.Builder.Add(ds)
		// fmt.Printf("\n%s>>%T %v", tr.indent, ctx, "repeated field")
	case *parser.MapLeftContext:
		n := &MapFieldNode{
			MyToken: MyToken{Token: ctx.GetMpt(), TType: parser.TronParserMap},
		} //,, Name: ctx.GetA()}
		tr.Builder.Add(n)
	// fmt.Printf("\n%s>>%T %v", tr.indent, ctx, "map")
	case *parser.MessageTypeContext:
	case *parser.RpcDelimContext:
	case *parser.SyntaxContext:
	case *parser.ConstantContext:
	case *parser.ConstantObjContext:
	case *parser.FieldOptionContext:
	case *parser.FieldOptionsContext:
		n := &OptionNode{
			MyToken: MyToken{Token: ctx.GetStart(), TType: parser.TronParserOption},
		} //,, Name: ctx.GetA()}
		tr.Builder.Add(n)
		tr.Builder.Down()
	case *parser.MsgSvcExtContext:
	case *parser.OptionNameContext:
	case *parser.PronSTRContext:
	case *parser.RangeeContext:
	case *parser.RangesContext:
	case *parser.Right_assocContext:
	case *parser.TronObjContext:
	case *parser.PronARRAYContext:
	// case *parser.ConstantObjElemContext:
	case *parser.TronObjsContext:
	default:
		glog.Warningf("Unhandled in EnterEveryRule case %T:\n", ctx)
	}
	// fmt.Printf("\n%s>>%T ", tr.indent, ctx)
	tr.indent += "  "

	// switch ctx := ctx.(type) {
	// 	v, err := strconv.Atoi(ctx.GetA().GetText())
	// 	var v float64
	// 	_, err := fmt.Sscanf(ctx.GetA().GetText(), "%f", &v)
	// 	name := ctx.GetFirst().GetText()
	// 	for _, r := range ctx.GetRest() {
	// 		name = name + r.GetText()
	// 	}
	// 		Col: ctx.CHAR().GetSymbol().GetText(),
	// 		Row: ctx.INT().GetSymbol().GetText(),
	// 	if ctx.GetRow1() == nil || ctx.GetRow2() == nil { // this only happen under antlr error recovery
	// 		return
	// 	}
}

// ExitEveryRule
func (tr *TronBuildListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	switch ctx.(type) {
	case *parser.ProtoContext:
	case *parser.ImportStatementContext:
	case *parser.PackageStatementContext:
	case *parser.OptionFileDefContext:
	case *parser.Ext_Msg_Enum_SvcContext:
		tr.Builder.Up()
	case *parser.EmptyStmContext:
	case *parser.AssocContext:
	case *parser.EmptyTopLvlContext:
	case *parser.MsgEnumSvcExtContext:
		tr.Builder.Up()
	case *parser.EmptyTopLvlStmContext:
	case *parser.RangeContext:
	case *parser.TLIOptionContext:
	case *parser.RPCSigContext:
		tr.Builder.Up()
	case *parser.EmptyStmStmContext:
	case *parser.EnumLeftContext:
	case *parser.Opt_SingleContext:
	case *parser.OptContext:
	case *parser.SingleFull_RepLocalContext:
	case *parser.SingleLocalContext:
	case *parser.RepeatedContext:
	case *parser.MapLeftContext:
	case *parser.MessageTypeContext:
	case *parser.RpcDelimContext:
	case *parser.SyntaxContext:
	case *parser.ConstantContext:
	case *parser.ConstantObjContext:
	case *parser.FieldOptionContext:
	case *parser.FieldOptionsContext:
		tr.Builder.Up()
	case *parser.MsgSvcExtContext:
	case *parser.OptionNameContext:
	case *parser.PronSTRContext:
	case *parser.RangeeContext:
	case *parser.RangesContext:
	case *parser.Right_assocContext:
	case *parser.TronObjContext:
	case *parser.PronARRAYContext:
	case *parser.TronObjsContext:
	// case *parser.ConstantObjElemContext:
	// 	fmt.Printf("%#+v %v\n", ctx.GetText(), ctx.GetStart())
	default:
		glog.Warningf("Unhandled in ExitEveryRule case %T:\n", ctx)
	}
}

type errorListener struct {
}

func (d *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	panic("line " + strconv.Itoa(line) + ":" + strconv.Itoa(column) + " " + msg)
}

func (d *errorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	// panic("ReportAmbiguity")
	fmt.Printf("ReportAmbiguity rec:%v dfs:%v start:%d stop:%d, exact:%v, ambigAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
}

func (d *errorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	//	panic("ReportAttemptingFullContext")
	fmt.Printf("ReportAttemptingFullContext rec:%v dfs:%v start:%d stop:%d, conflictingAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
}

func (d *errorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	// panic("ReportContextSensitivity")
}

func (tbl *TronBuildListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	if tbl.debug {
		fmt.Printf("SyntaxError %d:%d <%s>\n", line, column, msg)
	}
	if strings.HasPrefix(msg, "report") { // TODO remove NewDiagnosticErrorListener and move warning to ReportAmbiguity etc. when getDecisionDescription is make public
		tbl.warning = fmt.Sprintf("At %d:%d <%s>", line, column, msg)
	} else {
		tbl.err = fmt.Sprintf("SyntaxError %d:%d <%s>", line, column, msg)
	}
	// panic("line " + strconv.Itoa(line) + ":" + strconv.Itoa(column) + " " + msg)
}

func (tbl *TronBuildListener) ReportAmbiguity(
	recognizer antlr.Parser,
	dfa *antlr.DFA,
	startIndex, stopIndex int,
	exact bool,
	ambigAlts *antlr.BitSet,
	configs antlr.ATNConfigSet) {
	// panic("ReportAmbiguity")
	if tbl.debug {
		fmt.Printf("ReportAmbiguity rec:%v dfs:%v start:%d stop:%d, exact:%v, ambigAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
	}
}

func (tbl *TronBuildListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	// panic("ReportAttemptingFullContext")
	if tbl.debug {
		fmt.Printf("ReportAttemptingFullContext rec:%v dfs:%v start:%d stop:%d, conflictingAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
	}
}

func (tbl *TronBuildListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	//	panic("ReportContextSensitivity")
}
