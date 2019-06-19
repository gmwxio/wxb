package adlproc

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"strings"

	"github.com/golang/glog"
	"github.com/jpillora/opts"
	antlr "github.com/wxio/goantlr"

	"github.com/wxio/wxb/internal/ctree"
	parser "github.com/wxio/wxb/internal/grammars/adl"
	"github.com/wxio/wxb/internal/schemas/build"
)

type rootTron struct {
}

type exerciseTron struct {
	File string `type:"arg" help:"proto file" predict:"files"`
}

func Register(parent opts.Opts) opts.Opts {
	rt := rootTron{}
	et := exerciseTron{}
	wt := walkADL{}
	rto := parent.
		AddCommand(opts.New(&rt).Name("adl").
			AddCommand(opts.New(&et).Name("exec")).
			AddCommand(opts.New(&wt).Name("walk")),
		)
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

	var tr ctree.Tree
	{
		errListener := &errorListener{}
		is := antlr.NewInputStream(string(by))
		lexer := parser.NewadlLexer(is)
		lexer.RemoveErrorListeners()
		lexer.AddErrorListener(errListener)

		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		p := parser.NewADLParser(stream)
		p.RemoveErrorListeners()
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

		r := &ADLNode{
			MyToken: MyToken{
				Token: antlr.NewCommonToken(nil, parser.ADLParserADL, antlr.TokenDefaultChannel, 0, 0),
				TType: parser.ADLParserADL},
		}
		tbl := &ADLBuildListener{
			debug:   true,
			Builder: ctree.NewWalkableBuild("TREE", r),
		}

		p.AddErrorListener(tbl)
		p.BuildParseTrees = true
		ctx := p.Adl()
		fmt.Println("--------")
		antlr.ParseTreeWalkerDefault.Walk(tbl, ctx)
		tr = tbl.Builder.Build()
	}
	// tr, _, _, err := BuildAdlAST(string(by))
	// if err != nil {
	// 	glog.Warningf("BuildTronAST err:%v", err)
	// 	return nil
	// }
	fmt.Printf("%v\n", tr)
	// return nil
	// is := antlr.NewInputStream(string(by))
	// lexer := parser.NewadlLexer(is)
	// stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// p := parser.NewADLParser(stream)
	// p.AddErrorListener(antlr.NewConsoleErrorListener())
	// ctx := p.Adl()
	// fmt.Println("--------")
	// antlr.ParseTreeWalkerDefault.Walk(&adl_listener{baseR: lexer.BaseRecognizer}, ctx)
	return nil
}

type adl_listener struct {
	*parser.BaseADLParserListener
	//
	baseR  *antlr.BaseRecognizer
	indent string
}

func (tr *adl_listener) VisitTerminal(node antlr.TerminalNode) {
	// fmt.Printf(" %v", node)
}
func (tr *adl_listener) VisitErrorNode(node antlr.ErrorNode) {
	tok := node.GetSymbol()
	tid := tok.GetTokenType()
	fmt.Printf("ERROR %d %v\n", tid, tok)
	if tid >= len(tr.baseR.LiteralNames) {
		fmt.Printf("ERROR %d\n", tid)
	} else {
		fmt.Printf("ERROR %d %#+v\n", tid, tr.baseR.LiteralNames[tid])
	}
}
func (tr *adl_listener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("%s>>%T\n", tr.indent, ctx)
	tr.indent += "  "
}
func (tr *adl_listener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	tr.indent = tr.indent[2:]
	fmt.Printf("%s<<%T\n", tr.indent, ctx)
}

func BuildAdlAST(str string) (ctree.Tree, *antlr.BaseLexer, antlr.TokenStream, error) {
	errListener := &errorListener{}
	tbl := &ADLBuildListener{
		debug: true,
	}
	is := antlr.NewInputStream(str)
	lexer := parser.NewadlLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errListener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewADLParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.AddErrorListener(tbl)
	p.BuildParseTrees = true
	ctx := p.Adl()
	fmt.Println("--------")
	antlr.ParseTreeWalkerDefault.Walk(tbl, ctx)
	// if tbl.err != "" {
	// 	return nil, lexer.BaseLexer, stream, fmt.Errorf("ERROR:%v", tbl.err)
	// }
	// return nil, nil, nil, nil
	return tbl.Builder.Build(), lexer.BaseLexer, stream, nil
}

type ADLBuildListener struct {
	*parser.BaseADLParserListener
	Builder ctree.WalkableBuilder
	baseR   *antlr.BaseRecognizer
	indent  string

	warning string
	err     string
	debug   bool
}

func tokens2strings(arr []antlr.Token) []string {
	name := make([]string, len(arr))
	for i, tks := range arr {
		name[i] = tks.GetText()
	}
	return name
}

// EnterEveryRule is called when any rule is entered.
func (tr *ADLBuildListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	switch ctx := ctx.(type) {
	case *parser.AdlContext:
		// r := &ADLNode{MyToken: MyToken{Token: ctx.GetStart(), TType: parser.ADLParserADL}}
		// tr.Builder = ctree.NewWalkableBuild("TREE", r)
	case *parser.ModuleStatementContext:
		if ctx.GetKw().GetText() != "module" {
			n := &ErrorNode{
				MyToken:  MyToken{Token: ctx.GetKw(), TType: parser.ADLParserERROR},
				Expected: "module", Received: ctx.GetKw().GetText()}
			tr.Builder.Add(n)
		} else {
			n := &ModuleNode{
				MyToken: MyToken{Token: ctx.GetKw(), TType: parser.ADLParserModule},
				Name:    tokens2strings(ctx.GetName()),
			}
			tr.Builder.Add(n)
		}
		tr.Builder.Down()
	case *parser.ImportStatementContext:
		if ctx.GetKw().GetText() != "import" {
			n := &ErrorNode{
				MyToken:  MyToken{Token: ctx.GetKw(), TType: parser.ADLParserERROR},
				Expected: "import", Received: ctx.GetKw().GetText()}
			tr.Builder.Add(n)
		} else {
			n := &ImportNode{
				MyToken: MyToken{Token: ctx.GetKw(), TType: parser.ADLParserImport},
				Path:    tokens2strings(ctx.GetA()),
				Star:    ctx.GetS() != nil,
			}
			tr.Builder.Add(n)
		}
	case *parser.LocalAnnoContext:
	case *parser.DocAnnoContext:
	case *parser.StructOrUnionContext:
		switch ctx.GetKw().GetText() {
		case "struct":
			n := &StructNode{
				MyToken: MyToken{Token: ctx.GetKw(), TType: parser.ADLParserStruct},
				Name:    ctx.GetA().GetText(),
			}
			tr.Builder.Add(n)
			tr.Builder.Down()
		case "union":
			n := &UnionNode{
				MyToken: MyToken{Token: ctx.GetKw(), TType: parser.ADLParserUnion},
				Name:    ctx.GetA().GetText(),
			}
			tr.Builder.Add(n)
			tr.Builder.Down()
		default:
			n := &ErrorNode{
				MyToken:  MyToken{Token: ctx.GetKw(), TType: parser.ADLParserERROR},
				Expected: "struct|union", Received: ctx.GetKw().GetText()}
			tr.Builder.Add(n)
			tr.Builder.Down()
		}
	case *parser.TypeOrNewtypeContext:
		switch ctx.GetKw().GetText() {
		case "type":
			n := &TypeNode{
				MyToken: MyToken{Token: ctx.GetKw(), TType: parser.ADLParserType},
				Name:    ctx.GetA().GetText(),
				TypeRef: ctx.GetB().GetText(),
			}
			tr.Builder.Add(n)
			tr.Builder.Down()
		case "newtype":
			n := &NewTypeNode{
				MyToken: MyToken{Token: ctx.GetKw(), TType: parser.ADLParserNewtype},
				Name:    ctx.GetA().GetText(),
				TypeRef: ctx.GetB().GetText(),
			}
			tr.Builder.Add(n)
			tr.Builder.Down()
		default:
			n := &ErrorNode{
				MyToken:  MyToken{Token: ctx.GetKw(), TType: parser.ADLParserERROR},
				Expected: "type|newtype", Received: ctx.GetKw().GetText()}
			tr.Builder.Add(n)
			tr.Builder.Down()
		}
	case *parser.ModuleAnnotationContext:
	case *parser.DeclAnnotationContext:
	case *parser.FieldAnnotationContext:
	case *parser.TypeParameterContext:
		n := &TypeParamNode{
			MyToken: MyToken{Token: ctx.GetStart(), TType: parser.ADLParserTypeParam},
			Params:  tokens2strings(ctx.GetTypep()),
		}
		tr.Builder.Add(n)
	case *parser.TypeExpressionContext:
	case *parser.TypeExpressionElemContext:
	case *parser.FieldStatementContext:
	case *parser.StringStatementContext:
	case *parser.TrueFalseNullContext:
	case *parser.NumberStatementContext:
	case *parser.FloatStatementContext:
	case *parser.ArrayStatementContext:
	case *parser.ObjStatementContext:
		// rules name occur on errors
	case *parser.Top_level_statementContext:
		n := &ErrorNode{
			MyToken:  MyToken{Token: ctx.GetStart(), TType: parser.ADLParserERROR},
			Expected: "@|struct|union|annotation", Received: ctx.GetStart().GetText()}
		tr.Builder.Add(n)

	// case *parser.ProtoContext:
	// 	r := &ENode{MyToken: MyToken{Token: ctx.GetStart(), TType: parser.ADLParserROOT}}
	// 	tr.Builder = ctree.NewWalkableBuild("TREE", r)
	// case *parser.ImportStatementContext:
	// 	if ctx.GetImp().GetText() != "import" {
	// 		n := &ErrorNode{MyToken: MyToken{Token: ctx.GetImp(), TType: parser.ADLParserERROR}, Expected: "import", Received: ctx.GetImp().GetText()}
	// 		tr.Builder.Add(n)
	// 	} else {
	// 		n := &ImportNode{
	// 			MyToken: MyToken{Token: ctx.GetImp(), TType: parser.ADLParserImport},
	// 			Path:    strings.Trim(ctx.GetB().GetText(), `"`),
	// 		}
	// 		tr.Builder.Add(n)
	// 	}
	// case *parser.PackageStatementContext:
	// 	switch ctx.GetPkg().GetText() {
	// 	case "package":
	// 		n := &MessageNode{
	// 			MyToken: MyToken{Token: ctx.GetPkg(), TType: parser.ADLParserPackage},
	// 		}
	// 		name := make([]string, len(ctx.GetA()))
	// 		for i, tks := range ctx.GetA() {
	// 			name[i] = tks.GetText()
	// 		}
	// 		n.Name = name
	// 		tr.Builder.Add(n)
	// 	default:
	// 		n := &ErrorNode{
	// 			MyToken:  MyToken{Token: ctx.GetPkg(), TType: parser.ADLParserERROR},
	// 			Expected: "package",
	// 			Received: ctx.GetPkg().GetText(),
	// 		}
	// 		tr.Builder.Add(n)
	// 	}
	// case *parser.OptionFileDefContext:
	// case *parser.Ext_Msg_Enum_SvcContext:
	// 	switch ctx.GetMese().GetText() {
	// 	case "message":
	// 		n := &MessageNode{
	// 			MyToken: MyToken{Token: ctx.GetMese(), TType: parser.ADLParserMessage},
	// 			//Name: ctx.GetA(),
	// 		}
	// 		tr.Builder.Add(n)
	// 		tr.Builder.Down()
	// 	case "enum":
	// 		n := &EnumNode{
	// 			MyToken: MyToken{Token: ctx.GetMese(), TType: parser.ADLParserEnum},
	// 		} //, Name: ctx.GetA()}
	// 		tr.Builder.Add(n)
	// 		tr.Builder.Down()
	// 	case "service":
	// 		n := &ServiceNode{
	// 			MyToken: MyToken{Token: ctx.GetMese(), TType: parser.ADLParserService},
	// 		} //,, Name: ctx.GetA()}
	// 		tr.Builder.Add(n)
	// 		tr.Builder.Down()
	// 	case "extend":
	// 		n := &ExtendNode{
	// 			MyToken: MyToken{Token: ctx.GetMese(), TType: parser.ADLParserExtend},
	// 		} //,, Name: ctx.GetA()}
	// 		tr.Builder.Add(n)
	// 		tr.Builder.Down()
	// 	// case "oneof":
	// 	// 	n := &OneofNode{
	// 	// 		MyToken: MyToken{Token: ctx.GetMese(), TType: parser.ADLParserOneof},
	// 	// 	} //,, Name: ctx.GetA()}
	// 	// 	tr.Builder.Add(n)
	// 	default:
	// 		n := &ErrorNode{
	// 			MyToken:  MyToken{Token: ctx.GetMese(), TType: parser.ADLParserERROR},
	// 			Expected: "msg|enum|svc|extend",
	// 			Received: ctx.GetMese().GetText(),
	// 		}
	// 		tr.Builder.Add(n)
	// 		tr.Builder.Down()
	// 	}
	// case *parser.EmptyStmContext:
	// case *parser.AssocContext:
	// case *parser.EmptyTopLvlContext:
	// 	switch ctx.GetMsg().GetText() {
	// 	case "message":
	// 		n := &MessageNode{
	// 			MyToken: MyToken{Token: ctx.GetMsg(), TType: parser.ADLParserMessage},
	// 		} //, Name: ctx.GetA()}
	// 		tr.Builder.Add(n)
	// 	case "enum":
	// 		n := &EnumNode{
	// 			MyToken: MyToken{Token: ctx.GetMsg(), TType: parser.ADLParserEnum},
	// 		} //, Name: ctx.GetA()}
	// 		tr.Builder.Add(n)
	// 	case "service":
	// 		//TODO this should be an error - as per the spec
	// 		n := &ServiceNode{
	// 			MyToken: MyToken{Token: ctx.GetMsg(), TType: parser.ADLParserService},
	// 		} //,, Name: ctx.GetA()}
	// 		tr.Builder.Add(n)
	// 	case "extend":
	// 		n := &ExtendNode{
	// 			MyToken: MyToken{Token: ctx.GetMsg(), TType: parser.ADLParserExtend},
	// 		} //,, Name: ctx.GetA()}
	// 		tr.Builder.Add(n)
	// 	case "oneof":
	// 		n := &OneofNode{
	// 			MyToken: MyToken{Token: ctx.GetMsg(), TType: parser.ADLParserOneof},
	// 		} //,, Name: ctx.GetA()}
	// 		tr.Builder.Add(n)
	// 	default:
	// 		n := &ErrorNode{
	// 			MyToken:  MyToken{Token: ctx.GetMsg(), TType: parser.ADLParserERROR},
	// 			Expected: "msg|enum|svc|extend",
	// 			Received: ctx.GetMsg().GetText(),
	// 		}
	// 		tr.Builder.Add(n)
	// 	}
	// case *parser.MsgEnumSvcExtContext:
	// 	switch ctx.GetMese().GetText() {
	// 	case "message":
	// 		n := &MessageNode{
	// 			MyToken: MyToken{Token: ctx.GetMese(), TType: parser.ADLParserMessage},
	// 		} //, Name: ctx.GetA()}
	// 		tr.Builder.Add(n)
	// 		tr.Builder.Down()
	// 	case "enum":
	// 		n := &EnumNode{
	// 			MyToken: MyToken{Token: ctx.GetMese(), TType: parser.ADLParserEnum},
	// 		} //, Name: ctx.GetA()}
	// 		tr.Builder.Add(n)
	// 		tr.Builder.Down()
	// 	case "service":
	// 		n := &ServiceNode{
	// 			MyToken: MyToken{Token: ctx.GetMese(), TType: parser.ADLParserService},
	// 		} //,, Name: ctx.GetA()}
	// 		tr.Builder.Add(n)
	// 		tr.Builder.Down()
	// 	case "extend":
	// 		n := &ExtendNode{
	// 			MyToken: MyToken{Token: ctx.GetMese(), TType: parser.ADLParserExtend},
	// 		} //,, Name: ctx.GetA()}
	// 		tr.Builder.Add(n)
	// 		tr.Builder.Down()
	// 	case "oneof":
	// 		n := &OneofNode{
	// 			MyToken: MyToken{Token: ctx.GetMese(), TType: parser.ADLParserOneof},
	// 		} //,, Name: ctx.GetA()}
	// 		tr.Builder.Add(n)
	// 		tr.Builder.Down()
	// 	default:
	// 		n := &ErrorNode{
	// 			MyToken:  MyToken{Token: ctx.GetMese(), TType: parser.ADLParserERROR},
	// 			Expected: "msg|enum|svc|extend",
	// 			Received: ctx.GetMese().GetText(),
	// 		}
	// 		tr.Builder.Add(n)
	// 		tr.Builder.Down()
	// 	}
	// case *parser.EmptyTopLvlStmContext:
	// case *parser.RangeContext:
	// case *parser.TLIOptionContext:
	// case *parser.RPCSigContext:
	// 	switch ctx.GetRpcID().GetText() {
	// 	case "rpc":
	// 		n := &MessageNode{
	// 			MyToken: MyToken{Token: ctx.GetRpcID(), TType: parser.ADLParserRpc},
	// 		} //, Name: ctx.GetA()}
	// 		tr.Builder.Add(n)
	// 		tr.Builder.Down()
	// 	default:
	// 		n := &ErrorNode{
	// 			MyToken:  MyToken{Token: ctx.GetRpcID(), TType: parser.ADLParserERROR},
	// 			Expected: "rpc",
	// 			Received: ctx.GetRpcID().GetText(),
	// 		}
	// 		tr.Builder.Add(n)
	// 		tr.Builder.Down()
	// 	}
	// case *parser.EmptyStmStmContext:
	// case *parser.EnumLeftContext:
	// 	v, err := strconv.Atoi(ctx.GetV().GetText())
	// 	if err != nil {
	// 		panic("")
	// 	}
	// 	n := &EnumValueNode{
	// 		MyToken: MyToken{Token: ctx.GetA(), TType: parser.ADLParserEnumValue},
	// 		Name:    ctx.GetA().GetText(),
	// 		Index:   int32(v),
	// 	}
	// 	tr.Builder.Add(n)
	// case *parser.Opt_SingleContext:
	// 	if ctx.GetA().GetText() == "option" {
	// 		n := &OptionNode{
	// 			MyToken: MyToken{Token: ctx.GetA(), TType: parser.ADLParserOption},
	// 		} //,, Name: ctx.GetA()}
	// 		tr.Builder.Add(n)
	// 	} else {
	// 		n := &FieldNode{
	// 			MyToken: MyToken{Token: ctx.GetA(), TType: parser.ADLParserField},
	// 		} //,, Name: ctx.GetA()}
	// 		tr.Builder.Add(n)
	// 	}
	// case *parser.OptContext:
	// 	if ctx.GetA().GetText() != "option" {
	// 		n := &ErrorNode{
	// 			MyToken:  MyToken{Token: ctx.GetA(), TType: parser.ADLParserERROR},
	// 			Expected: "option",
	// 			Received: ctx.GetA().GetText(),
	// 		}
	// 		tr.Builder.Add(n)
	// 	} else {
	// 		n := &OptionNode{
	// 			MyToken: MyToken{Token: ctx.GetA(), TType: parser.ADLParserOption},
	// 		} //, Path: ctx.GetB().GetText()}
	// 		tr.Builder.Add(n)
	// 	}
	// case *parser.SingleFull_RepLocalContext:
	// 	n := &FieldNode{
	// 		MyToken: MyToken{Token: ctx.GetB()[0], TType: parser.ADLParserField},
	// 	} //,, Name: ctx.GetA()}
	// 	tname := make([]string, len(ctx.GetB()))
	// 	for i, tkb := range ctx.GetB() {
	// 		tname[i] = tkb.GetText()
	// 	}
	// 	n.TypeName = tname
	// 	n.Name = ctx.GetC().GetText()
	// 	tr.Builder.Add(n)
	// case *parser.SingleLocalContext:
	// 	// fmt.Printf("\n%s>>%T %v", tr.indent, ctx, "field local")
	// case *parser.RepeatedContext:
	// 	tname := make([]string, len(ctx.GetB()))
	// 	for i, tks := range ctx.GetB() {
	// 		tname[i] = tks.GetText()
	// 	}
	// 	n := &FieldNode{
	// 		MyToken:  MyToken{Token: ctx.GetB()[0], TType: parser.ADLParserField},
	// 		TypeName: tname,
	// 		Name:     ctx.GetC().GetText(),
	// 	} //,, Name: ctx.GetA()}
	// 	tr.Builder.Add(n)
	// 	ds := &DatastructNode{
	// 		MyToken: MyToken{Token: ctx.GetA(), TType: parser.ADLParserDatastructure},
	// 		Type:    ctx.GetA().GetText(),
	// 	}
	// 	tr.Builder.Add(ds)
	// 	// fmt.Printf("\n%s>>%T %v", tr.indent, ctx, "repeated field")
	// case *parser.MapLeftContext:
	// 	n := &MapFieldNode{
	// 		MyToken: MyToken{Token: ctx.GetMpt(), TType: parser.ADLParserMap},
	// 	} //,, Name: ctx.GetA()}
	// 	tr.Builder.Add(n)
	// // fmt.Printf("\n%s>>%T %v", tr.indent, ctx, "map")
	// case *parser.MessageTypeContext:
	// case *parser.RpcDelimContext:
	// case *parser.SyntaxContext:
	// case *parser.ConstantContext:
	// case *parser.ConstantObjContext:
	// case *parser.FieldOptionContext:
	// case *parser.FieldOptionsContext:
	// 	n := &OptionNode{
	// 		MyToken: MyToken{Token: ctx.GetStart(), TType: parser.ADLParserOption},
	// 	} //,, Name: ctx.GetA()}
	// 	tr.Builder.Add(n)
	// 	tr.Builder.Down()
	// case *parser.MsgSvcExtContext:
	// case *parser.OptionNameContext:
	// case *parser.PronSTRContext:
	// case *parser.RangeeContext:
	// case *parser.RangesContext:
	// case *parser.Right_assocContext:
	// case *parser.ADLObjContext:
	// case *parser.PronARRAYContext:
	// // case *parser.ConstantObjElemContext:
	// case *parser.ADLObjsContext:
	default:
		glog.Warningf("Unhandled in >EnterEveryRule case %T:\n", ctx)
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
func (tr *ADLBuildListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	switch ctx.(type) {
	case *parser.AdlContext:
	case *parser.ModuleStatementContext:
		tr.Builder.Up()
	case *parser.ImportStatementContext:
	case *parser.LocalAnnoContext:
	case *parser.DocAnnoContext:
	case *parser.StructOrUnionContext:
		tr.Builder.Up()
	case *parser.TypeOrNewtypeContext:
		tr.Builder.Up()
	case *parser.ModuleAnnotationContext:
	case *parser.DeclAnnotationContext:
	case *parser.FieldAnnotationContext:
	case *parser.TypeParameterContext:
	case *parser.TypeExpressionContext:
	case *parser.TypeExpressionElemContext:
	case *parser.FieldStatementContext:
	case *parser.StringStatementContext:
	case *parser.TrueFalseNullContext:
	case *parser.NumberStatementContext:
	case *parser.FloatStatementContext:
	case *parser.ArrayStatementContext:
	case *parser.ObjStatementContext:

	// case *parser.ProtoContext:
	// case *parser.ImportStatementContext:
	// case *parser.PackageStatementContext:
	// case *parser.OptionFileDefContext:
	// case *parser.Ext_Msg_Enum_SvcContext:
	// 	tr.Builder.Up()
	// case *parser.EmptyStmContext:
	// case *parser.AssocContext:
	// case *parser.EmptyTopLvlContext:
	// case *parser.MsgEnumSvcExtContext:
	// 	tr.Builder.Up()
	// case *parser.EmptyTopLvlStmContext:
	// case *parser.RangeContext:
	// case *parser.TLIOptionContext:
	// case *parser.RPCSigContext:
	// 	tr.Builder.Up()
	// case *parser.EmptyStmStmContext:
	// case *parser.EnumLeftContext:
	// case *parser.Opt_SingleContext:
	// case *parser.OptContext:
	// case *parser.SingleFull_RepLocalContext:
	// case *parser.SingleLocalContext:
	// case *parser.RepeatedContext:
	// case *parser.MapLeftContext:
	// case *parser.MessageTypeContext:
	// case *parser.RpcDelimContext:
	// case *parser.SyntaxContext:
	// case *parser.ConstantContext:
	// case *parser.ConstantObjContext:
	// case *parser.FieldOptionContext:
	// case *parser.FieldOptionsContext:
	// 	tr.Builder.Up()
	// case *parser.MsgSvcExtContext:
	// case *parser.OptionNameContext:
	// case *parser.PronSTRContext:
	// case *parser.RangeeContext:
	// case *parser.RangesContext:
	// case *parser.Right_assocContext:
	// case *parser.ADLObjContext:
	// case *parser.PronARRAYContext:
	// case *parser.ADLObjsContext:
	// case *parser.ConstantObjElemContext:
	// 	fmt.Printf("%#+v %v\n", ctx.GetText(), ctx.GetStart())
	default:
		glog.Warningf("Unhandled in <ExitEveryRule case %T:\n", ctx)
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

func (tbl *ADLBuildListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	t, ok := offendingSymbol.(antlr.Token)
	if !ok && e != nil {
		t = e.GetOffendingToken()
	}
	n := &ErrorNode{
		MyToken: MyToken{
			Token: t,
			TType: parser.ADLParserERROR,
		},
		Expected: msg,
		Received: fmt.Sprintf("%v", offendingSymbol),
	}
	tbl.Builder.Add(n)

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

func (tbl *ADLBuildListener) ReportAmbiguity(
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

func (tbl *ADLBuildListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	// panic("ReportAttemptingFullContext")
	if tbl.debug {
		fmt.Printf("ReportAttemptingFullContext rec:%v dfs:%v start:%d stop:%d, conflictingAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
	}
}

func (tbl *ADLBuildListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	//	panic("ReportContextSensitivity")
}
