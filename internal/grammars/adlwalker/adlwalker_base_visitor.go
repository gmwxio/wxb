// Generated from ADLWalker.g4 by ANTLR 4.7.

package walker // ADLWalker
// See base listener file for example implementations

//import "github.com/wxio/goantlr"

// import "generate package"

//type ADLWalkerVisitor struct {
//    *antlr.BaseParseTreeVisitor
//    indent string
//}

//var _ antlr.ParseTreeVisitor = &ADLWalkerVisitor{}
//var _ antlr.AggregateResultVisitor = &ADLWalkerVisitor{}
//var _ antlr.VisitNextCheck = &ADLWalkerVisitor{}
//var _ antlr.VisitRestCheck = &ADLWalkerVisitor{}
//var _ antlr.EnterEveryRuleVisitor = &ADLWalkerVisitor{}
//var _ antlr.ExitEveryRuleVisitor = &ADLWalkerVisitor{}

//var _ walker.AdlContextVisitor = &ADLWalkerVisitor{}
//var _ walker.JsonContextVisitor = &ADLWalkerVisitor{}
//var _ walker.ModuleContextVisitor = &ADLWalkerVisitor{}
//var _ walker.StructContextVisitor = &ADLWalkerVisitor{}
//var _ walker.UnionContextVisitor = &ADLWalkerVisitor{}
//var _ walker.TypeContextVisitor = &ADLWalkerVisitor{}
//var _ walker.NewtypeContextVisitor = &ADLWalkerVisitor{}
//var _ walker.ModAnnoContextVisitor = &ADLWalkerVisitor{}
//var _ walker.DeclAnnoContextVisitor = &ADLWalkerVisitor{}
//var _ walker.FieldAnnoContextVisitor = &ADLWalkerVisitor{}
//var _ walker.TypeParamErrorContextVisitor = &ADLWalkerVisitor{}
//var _ walker.FieldContextVisitor = &ADLWalkerVisitor{}
//var _ walker.AnnotationContextVisitor = &ADLWalkerVisitor{}
//var _ walker.TypeExpr_ContextVisitor = &ADLWalkerVisitor{}
//var _ walker.TypeParamsContextVisitor = &ADLWalkerVisitor{}
//var _ walker.JsonStrContextVisitor = &ADLWalkerVisitor{}
//var _ walker.JsonBoolContextVisitor = &ADLWalkerVisitor{}
//var _ walker.JsonNullContextVisitor = &ADLWalkerVisitor{}
//var _ walker.JsonIntContextVisitor = &ADLWalkerVisitor{}
//var _ walker.JsonFloatContextVisitor = &ADLWalkerVisitor{}
//var _ walker.JsonArrayContextVisitor = &ADLWalkerVisitor{}
//var _ walker.JsonObjContextVisitor = &ADLWalkerVisitor{}

//func (v *ADLWalkerVisitor) VisitNext(node antlr.Tree, current interface{}) bool {
//    return true
//}
//func (v *ADLWalkerVisitor) VisitRest(node antlr.RuleNode, current interface{}) bool {
//    return true
//}
//func (v *ADLWalkerVisitor) AggregateResult(aggregate, nextResult interface{}) (result interface{}) {
//    return nextResult
//}
//func (v *ADLWalkerVisitor) VisitTerminal(node antlr.TerminalNode) {
//}
//func (v *ADLWalkerVisitor) VisitErrorNode(node antlr.ErrorNode) {
//}
//func (v *ADLWalkerVisitor) EnterEveryRule(ctx antlr.RuleNode) {
//    v.indent += " "
//    fmt.Printf("%s %T\n", v.indent, ctx)
//}
//func (v *ADLWalkerVisitor) ExitEveryRule(ctx antlr.RuleNode) {
//    v.indent = v.indent[:len(v.indent)-1]
//}

//indent := ""
//parsetree.VisitFunc(&walker.ADLWalkerHandlers{
//EnterEveryRule: func(ctx antlr.RuleNode){indent += "\t"},
//ExitEveryRule:  func(ctx antlr.RuleNode){indent = indent[1:]},
//Adl: func(ctx *walker.AdlContext, this *walker.ADLWalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Json: func(ctx *walker.JsonContext, this *walker.ADLWalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Module: func(ctx *walker.ModuleContext, this *walker.ADLWalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Struct: func(ctx *walker.StructContext, this *walker.ADLWalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Union: func(ctx *walker.UnionContext, this *walker.ADLWalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Type: func(ctx *walker.TypeContext, this *walker.ADLWalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Newtype: func(ctx *walker.NewtypeContext, this *walker.ADLWalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//ModAnno: func(ctx *walker.ModAnnoContext, this *walker.ADLWalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//DeclAnno: func(ctx *walker.DeclAnnoContext, this *walker.ADLWalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//FieldAnno: func(ctx *walker.FieldAnnoContext, this *walker.ADLWalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TypeParamError: func(ctx *walker.TypeParamErrorContext, this *walker.ADLWalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Field: func(ctx *walker.FieldContext, this *walker.ADLWalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Annotation: func(ctx *walker.AnnotationContext, this *walker.ADLWalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TypeExpr_: func(ctx *walker.TypeExpr_Context, this *walker.ADLWalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TypeParams: func(ctx *walker.TypeParamsContext, this *walker.ADLWalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonStr: func(ctx *walker.JsonStrContext, this *walker.ADLWalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonBool: func(ctx *walker.JsonBoolContext, this *walker.ADLWalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonNull: func(ctx *walker.JsonNullContext, this *walker.ADLWalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonInt: func(ctx *walker.JsonIntContext, this *walker.ADLWalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonFloat: func(ctx *walker.JsonFloatContext, this *walker.ADLWalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonArray: func(ctx *walker.JsonArrayContext, this *walker.ADLWalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonObj: func(ctx *walker.JsonObjContext, this *walker.ADLWalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},

//}

//func (v *ADLWalkerVisitor) VisitAdl(ctx walker.IAdlContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLWalkerVisitor) VisitJson(ctx walker.IJsonContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLWalkerVisitor) VisitModule(ctx walker.IModuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLWalkerVisitor) VisitStruct(ctx walker.IStructContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLWalkerVisitor) VisitUnion(ctx walker.IUnionContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLWalkerVisitor) VisitType(ctx walker.ITypeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLWalkerVisitor) VisitNewtype(ctx walker.INewtypeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLWalkerVisitor) VisitModAnno(ctx walker.IModAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLWalkerVisitor) VisitDeclAnno(ctx walker.IDeclAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLWalkerVisitor) VisitFieldAnno(ctx walker.IFieldAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLWalkerVisitor) VisitTypeParamError(ctx walker.ITypeParamErrorContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLWalkerVisitor) VisitField(ctx walker.IFieldContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLWalkerVisitor) VisitAnnotation(ctx walker.IAnnotationContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLWalkerVisitor) VisitTypeExpr_(ctx walker.ITypeExpr_Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLWalkerVisitor) VisitTypeParams(ctx walker.ITypeParamsContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLWalkerVisitor) VisitJsonStr(ctx walker.IJsonStrContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLWalkerVisitor) VisitJsonBool(ctx walker.IJsonBoolContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLWalkerVisitor) VisitJsonNull(ctx walker.IJsonNullContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLWalkerVisitor) VisitJsonInt(ctx walker.IJsonIntContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLWalkerVisitor) VisitJsonFloat(ctx walker.IJsonFloatContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLWalkerVisitor) VisitJsonArray(ctx walker.IJsonArrayContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLWalkerVisitor) VisitJsonObj(ctx walker.IJsonObjContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}

//  TODO list rules here
//  Visit rules manually
//  eg a : b c* | d;
//  if ctx.GetB() != nil {
//    result1 = ctx.GetB(ctx, delegate, args)
//    for _, c := range ctx.GetC() {
//      resultS = c.GetC(ctx, delegate, args)
//    }
//  } else { ... }
//  OR visit all children rules
//  // before children
//  v.VisitChildren(ctx, delegate)
//  // afer children
//
//  return result
