package adlproc

import (
	"fmt"
	"strconv"
	"strings"

	antlr "github.com/wxio/goantlr"
	"github.com/wxio/wxb/internal/ctree"
	parser "github.com/wxio/wxb/internal/grammars/adl"
	walker "github.com/wxio/wxb/internal/grammars/adlwalker"
)

func UnmarshalADL(by []byte) (ctree.Tree, *ADL, error) {
	el := &antlrEL{}
	lexer := parser.NewadlLexer(antlr.NewInputStream(string(by)))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(el)
	//
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewADLParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.AddErrorListener(el)
	// p.BuildParseTrees = true
	ctx := p.Adl()
	if el.err != nil {
		return nil, nil, el.err
	}
	vi := &adlastVisitor{}
	ctx.Visit(vi)
	tr := vi.bldr.Build()
	return tr, vi.adl, vi.err
}

func WalkADL(tr ctree.Tree, list antlr.ParseTreeListener) error {
	var tttype *TTType
	var tts antlr.TokenStream = ctree.NewTreeTokenSource(tr, tttype)
	p := walker.NewADLWalker(tts)
	// debugTreeToken(tts, p)
	p.SetTokenStream(tts)
	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	el := &antlrEL{}
	p.AddErrorListener(el)
	p.BuildParseTrees = true
	jv := p.Json()
	antlr.ParseTreeWalkerDefault.Walk(list, jv)
	return el.err
}

func VisitADL(tr ctree.Tree, vi antlr.ParseTreeVisitor) error {
	var tttype *TTType
	var tts antlr.TokenStream = ctree.NewTreeTokenSource(tr, tttype)
	p := walker.NewADLWalker(tts)
	debugTreeToken(tts, p)
	p.SetTokenStream(tts)
	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	el := &antlrEL{}
	p.AddErrorListener(el)
	ctx := p.Adl()
	ctx.Visit(vi)
	return el.err
}

type ADL struct {
	Modules []*Module
}

func (a ADL) String() string {
	return fmt.Sprintf("%v", a.Modules)
}

type adlastVisitor struct {
	*antlr.BaseParseTreeVisitor
	bldr    ctree.WalkableBuilder
	adl     *ADL
	errToks []ctree.TreeToken
	err     error
}

func (v *adlastVisitor) VisitAdl(ctx parser.IAdlContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	//fmt.Printf("aslast %T\n", ctx)
	v.bldr = ctree.NewBuild("ADL", ctx.GetStart(), parser.ADLParserADL, nil)
	v.adl = &ADL{}
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *adlastVisitor) VisitModuleStatement(ctx parser.IModuleStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	//fmt.Printf("aslast %T\n", ctx)
	var obj interface{}
	if ctx.GetKw().GetText() != "module" {
		obj = Error{
			Expected: []string{"module"},
			Received: ctx.GetKw().GetText(),
		}
		v.bldr.AddNode(ctx.GetKw(), parser.ADLParserERROR, obj)
	} else {
		mod := &Module{
			Name: strings.Join(tokens2strings(ctx.GetName()), "."),
		}
		v.bldr.AddNode(ctx.GetKw(), parser.ADLParserModule, obj)
		v.adl.Modules = append(v.adl.Modules, mod)
		obj = mod
	}
	result = obj
	v.bldr.Down()
	defer v.bldr.Up()
	fmt.Printf("     %v %[1]T\n", obj)

	v.VisitChildren(ctx, delegate, obj)
	return
}
func (v *adlastVisitor) VisitImportScopedName(ctx parser.IImportScopedNameContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	fmt.Printf("aslast %T %v\n", ctx, args[0])
	if ctx.GetKw().GetText() != "import" {
		obj := Error{
			Expected: []string{"import"},
			Received: ctx.GetKw().GetText(),
		}
		v.bldr.AddNode(ctx.GetKw(), parser.ADLParserERROR, obj)
		return
	}
	im := Import{}
	toks := ctx.GetA()
	im.ScopedName = &ScopedName{
		ModuleName: strings.Join(tokens2strings(toks[:len(toks)-1]), "."),
		Name:       toks[len(toks)-1].GetText(),
	}
	v.bldr.AddNode(ctx.GetKw(), parser.ADLParserImportScopedName, im)
	if mo, ok := args[0].(ImportableAble); ok {
		fmt.Println("1yes")
		mo.AddImport(im)
	} else {
		fmt.Println("no")
	}
	return
}
func (v *adlastVisitor) VisitImportModuleName(ctx parser.IImportModuleNameContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	fmt.Printf("aslast %T %v\n", ctx, args[0])
	if ctx.GetKw().GetText() != "import" {
		obj := Error{
			Expected: []string{"import"},
			Received: ctx.GetKw().GetText(),
		}
		v.bldr.AddNode(ctx.GetKw(), parser.ADLParserERROR, obj)
		return
	}
	im := Import{}
	mn := strings.Join(tokens2strings(ctx.GetA()), ".")
	im.ModuleName = &mn
	v.bldr.AddNode(ctx.GetKw(), parser.ADLParserImportModule, im)
	if mo, ok := args[0].(ImportableAble); ok {
		fmt.Println("2yes")
		mo.AddImport(im)
	} else {
		fmt.Println("no")
	}
	return
}

func (v *adlastVisitor) VisitLocalAnno(ctx parser.ILocalAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	// //fmt.Printf("aslast %T\n", ctx)
	v.bldr.AddNode(ctx.GetA(), parser.ADLParserAnnotationNotScoped, ctx.GetA().GetText()).Down()
	defer v.bldr.Up()
	if mo, ok := args[0].(AnnotateAble); ok {
		jo := v.VisitChildren(ctx, delegate, nil)
		an := Annotation{
			Key: ScopedName{
				Name: ctx.GetA().GetText(),
			},
			Val: jo,
		}
		mo.AddAnnotation(an)
	} else {
		v.VisitChildren(ctx, delegate, nil)
	}
	return
}
func (v *adlastVisitor) VisitDocAnno(ctx parser.IDocAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	// //fmt.Printf("aslast %T\n", ctx)
	// n := &AnnoNode{
	// 	MyToken: MyToken{Token: ctx.GetA(), TType: parser.ADLParserAnnotation},
	// 	Name:    ctx.GetA().GetText(),
	// 	Doc:     true,
	// }
	// v.bldr.AddNode(ctx.GetA(),parser.ADLParserAnnotation, )

	// result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *adlastVisitor) VisitStructOrUnion(ctx parser.IStructOrUnionContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	//fmt.Printf("aslast %T\n", ctx)
	switch ctx.GetKw().GetText() {
	case "struct":
		v.bldr.AddNode(ctx.GetKw(), parser.ADLParserStruct, ctx.GetA().GetText()).Down()
		defer v.bldr.Up()
	case "union":
		v.bldr.AddNode(ctx.GetKw(), parser.ADLParserUnion, ctx.GetA().GetText()).Down()
		defer v.bldr.Up()
	default:
		n := &ErrorNode{
			MyToken:  MyToken{Token: ctx.GetKw(), TType: parser.ADLParserERROR},
			Expected: "struct|union", Received: ctx.GetKw().GetText()}
		tr.Builder.Add(n)
		tr.Builder.Down()
	}

	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *adlastVisitor) VisitTypeOrNewtype(ctx parser.ITypeOrNewtypeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	//fmt.Printf("aslast %T\n", ctx)
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *adlastVisitor) VisitModuleAnnotation(ctx parser.IModuleAnnotationContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	//fmt.Printf("aslast %T\n", ctx)
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *adlastVisitor) VisitDeclAnnotation(ctx parser.IDeclAnnotationContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	//fmt.Printf("aslast %T\n", ctx)
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *adlastVisitor) VisitFieldAnnotation(ctx parser.IFieldAnnotationContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	//fmt.Printf("aslast %T\n", ctx)
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *adlastVisitor) VisitTypeParameter(ctx parser.ITypeParameterContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	//fmt.Printf("aslast %T\n", ctx)
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *adlastVisitor) VisitErrorTypeParam(ctx parser.IErrorTypeParamContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	//fmt.Printf("aslast %T\n", ctx)
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *adlastVisitor) VisitTypeParamError(ctx parser.ITypeParamErrorContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	//fmt.Printf("aslast %T\n", ctx)
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *adlastVisitor) VisitTypeExpression(ctx parser.ITypeExpressionContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	//fmt.Printf("aslast %T\n", ctx)
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *adlastVisitor) VisitTypeExpressionElem(ctx parser.ITypeExpressionElemContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	//fmt.Printf("aslast %T\n", ctx)
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *adlastVisitor) VisitFieldStatement(ctx parser.IFieldStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	//fmt.Printf("aslast %T\n", ctx)
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *adlastVisitor) VisitStringStatement(ctx parser.IStringStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	s := ctx.GetS().GetText()
	result = s[1 : len(s)-1]
	v.bldr.AddNode(ctx.GetStart(), parser.ADLParserJsonStr, result)
	return
}
func (v *adlastVisitor) VisitTrueFalseNull(ctx parser.ITrueFalseNullContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch strings.ToLower(ctx.GetKw().GetText()) {
	case "true":
		result = true
		v.bldr.AddNode(ctx.GetStart(), parser.ADLParserJsonBool, result)
	case "false":
		result = false
		v.bldr.AddNode(ctx.GetStart(), parser.ADLParserJsonBool, result)
	case "null":
		result = JsonNull{}
		v.bldr.AddNode(ctx.GetStart(), parser.ADLParserJsonNull, result)
	default:
		v.err = fmt.Errorf("Expected: \"true|false|null\", Received: %s", ctx.GetKw().GetText())
	}
	return
}
func (v *adlastVisitor) VisitNumberStatement(ctx parser.INumberStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	i, err := strconv.ParseInt(ctx.GetN().GetText(), 10, 64)
	if err != nil {
		v.err = fmt.Errorf("Expected: <number>, Received: %s", ctx.GetN().GetText())
	} else {
		result = i
		v.bldr.AddNode(ctx.GetStart(), parser.ADLParserJsonInt, result)
	}
	return
}
func (v *adlastVisitor) VisitFloatStatement(ctx parser.IFloatStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	i, err := strconv.ParseFloat(ctx.GetF().GetText(), 64)
	if err != nil {
		v.err = fmt.Errorf("Expected: <float>, Received: %s", ctx.GetF().GetText())
	} else {
		result = i
		v.bldr.AddNode(ctx.GetStart(), parser.ADLParserJsonFloat, result)
	}
	return
}
func (v *adlastVisitor) VisitArrayStatement(ctx parser.IArrayStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	ja := JsonArray{}
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
func (v *adlastVisitor) VisitObjStatement(ctx parser.IObjStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	jo := JsonObj{
		JsonObjs: make(map[string]Json),
	}
	if v.bldr == nil {
		v.bldr = ctree.NewBuild("JSON", ctx.GetStart(), parser.ADLParserJsonObj, jo)
	} else {
		v.bldr.AddNode(ctx.GetStart(), parser.ADLParserJsonObj, jo).Down()
		defer v.bldr.Up()
	}
	vs := ctx.GetV()
	for i, child := range ctx.GetK() {
		k := child.GetText()
		val := vs[i].Visit(delegate, nil).(Json)
		if val != nil {
			jo.JsonObjs[k[1:len(k)-1]] = val
		}
	}
	result = jo
	return
}
