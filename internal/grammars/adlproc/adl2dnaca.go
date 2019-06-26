// Generated from ADLWalker.g4 by ANTLR 4.7.

package adlproc // ADLWalker
// See base listener file for example implementations

import (
	"fmt"

	antlr "github.com/wxio/goantlr"
	"github.com/wxio/wxb/internal/ctree"
	walker "github.com/wxio/wxb/internal/grammars/adlwalker"
)

type ADLWalkerVisitor struct {
	*antlr.BaseParseTreeVisitor
	Builder ctree.WalkableBuilder
	indent  string
}

var _ antlr.ParseTreeVisitor = &ADLWalkerVisitor{}
var _ antlr.AggregateResultVisitor = &ADLWalkerVisitor{}
var _ antlr.VisitNextCheck = &ADLWalkerVisitor{}
var _ antlr.VisitRestCheck = &ADLWalkerVisitor{}
var _ antlr.EnterEveryRuleVisitor = &ADLWalkerVisitor{}
var _ antlr.ExitEveryRuleVisitor = &ADLWalkerVisitor{}

var _ walker.AdlContextVisitor = &ADLWalkerVisitor{}
var _ walker.ModuleContextVisitor = &ADLWalkerVisitor{}
var _ walker.StructContextVisitor = &ADLWalkerVisitor{}
var _ walker.UnionContextVisitor = &ADLWalkerVisitor{}
var _ walker.TypeContextVisitor = &ADLWalkerVisitor{}
var _ walker.NewtypeContextVisitor = &ADLWalkerVisitor{}
var _ walker.ModAnnoContextVisitor = &ADLWalkerVisitor{}
var _ walker.DeclAnnoContextVisitor = &ADLWalkerVisitor{}
var _ walker.FieldAnnoContextVisitor = &ADLWalkerVisitor{}
var _ walker.TypeParamErrorContextVisitor = &ADLWalkerVisitor{}
var _ walker.FieldContextVisitor = &ADLWalkerVisitor{}
var _ walker.AnnotationContextVisitor = &ADLWalkerVisitor{}
var _ walker.TypeExpr_ContextVisitor = &ADLWalkerVisitor{}
var _ walker.TypeParamsContextVisitor = &ADLWalkerVisitor{}
var _ walker.JsonStrContextVisitor = &ADLWalkerVisitor{}
var _ walker.JsonBoolContextVisitor = &ADLWalkerVisitor{}
var _ walker.JsonNullContextVisitor = &ADLWalkerVisitor{}
var _ walker.JsonIntContextVisitor = &ADLWalkerVisitor{}
var _ walker.JsonFloatContextVisitor = &ADLWalkerVisitor{}
var _ walker.JsonArrayContextVisitor = &ADLWalkerVisitor{}
var _ walker.JsonObjContextVisitor = &ADLWalkerVisitor{}

func (v *ADLWalkerVisitor) VisitNext(node antlr.Tree, current interface{}) bool {
	return true
}
func (v *ADLWalkerVisitor) VisitRest(node antlr.RuleNode, current interface{}) bool {
	return true
}
func (v *ADLWalkerVisitor) AggregateResult(aggregate, nextResult interface{}) (result interface{}) {
	return nextResult
}
func (v *ADLWalkerVisitor) VisitTerminal(node antlr.TerminalNode) {
}
func (v *ADLWalkerVisitor) VisitErrorNode(node antlr.ErrorNode) {
	fmt.Printf("%s ERROR:%T\n", v.indent, node)
}
func (v *ADLWalkerVisitor) EnterEveryRule(ctx antlr.RuleNode) {
	v.indent += " "
	fmt.Printf("%s %T\n", v.indent, ctx)
}
func (v *ADLWalkerVisitor) ExitEveryRule(ctx antlr.RuleNode) {
	v.indent = v.indent[:len(v.indent)-1]
}

func (v *ADLWalkerVisitor) VisitAdl(ctx walker.IAdlContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	fmt.Printf("%T\n", ctx.GetTok())
	r := &ADLNode{MyToken: MyToken{Token: ctx.GetTok(), TType: walker.ADLWalkerDNAC}}
	v.Builder = ctree.NewWalkableBuild("TREE", r)
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *ADLWalkerVisitor) VisitModule(ctx walker.IModuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	// mod := ctx.GetTok().(*ModuleNode)
	// for _, imp := range ctx.AllImport() {
	// 	fmt.Printf("imp %T\n", imp.GetPayload())
	// }
	// n := &ModuleNode{
	// 	MyToken: MyToken{Token: mod, TType: walker.ADLWalkerName},
	// 	Name:    mod.Name,
	// }
	// v.Builder.Add(n)
	// v.Builder.Down()
	// result = v.VisitChildren(ctx, delegate, args...)
	// v.Builder.Up()
	return
}
func (v *ADLWalkerVisitor) VisitStruct(ctx walker.IStructContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	fmt.Printf("%T\n", ctx.GetTok())
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *ADLWalkerVisitor) VisitUnion(ctx walker.IUnionContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	fmt.Printf("%T\n", ctx.GetTok())
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *ADLWalkerVisitor) VisitType(ctx walker.ITypeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	fmt.Printf("%T\n", ctx.GetTok())
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *ADLWalkerVisitor) VisitNewtype(ctx walker.INewtypeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *ADLWalkerVisitor) VisitModAnno(ctx walker.IModAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *ADLWalkerVisitor) VisitDeclAnno(ctx walker.IDeclAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *ADLWalkerVisitor) VisitFieldAnno(ctx walker.IFieldAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *ADLWalkerVisitor) VisitTypeParamError(ctx walker.ITypeParamErrorContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *ADLWalkerVisitor) VisitField(ctx walker.IFieldContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *ADLWalkerVisitor) VisitAnnotation(ctx walker.IAnnotationContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *ADLWalkerVisitor) VisitTypeExpr_(ctx walker.ITypeExpr_Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *ADLWalkerVisitor) VisitTypeParams(ctx walker.ITypeParamsContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *ADLWalkerVisitor) VisitJsonStr(ctx walker.IJsonStrContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *ADLWalkerVisitor) VisitJsonBool(ctx walker.IJsonBoolContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *ADLWalkerVisitor) VisitJsonNull(ctx walker.IJsonNullContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *ADLWalkerVisitor) VisitJsonInt(ctx walker.IJsonIntContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *ADLWalkerVisitor) VisitJsonFloat(ctx walker.IJsonFloatContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *ADLWalkerVisitor) VisitJsonArray(ctx walker.IJsonArrayContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *ADLWalkerVisitor) VisitJsonObj(ctx walker.IJsonObjContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
