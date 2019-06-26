// Generated from DNAC_A_Walker.g4 by ANTLR 4.7.

package dnaca // DNAC_A_Walker
// See base listener file for example implementations

//import "github.com/wxio/goantlr"

// import "generate package"

//type DNAC_A_WalkerVisitor struct {
//    *antlr.BaseParseTreeVisitor
//    indent string
//}

//var _ antlr.ParseTreeVisitor = &DNAC_A_WalkerVisitor{}
//var _ antlr.AggregateResultVisitor = &DNAC_A_WalkerVisitor{}
//var _ antlr.VisitNextCheck = &DNAC_A_WalkerVisitor{}
//var _ antlr.VisitRestCheck = &DNAC_A_WalkerVisitor{}
//var _ antlr.EnterEveryRuleVisitor = &DNAC_A_WalkerVisitor{}
//var _ antlr.ExitEveryRuleVisitor = &DNAC_A_WalkerVisitor{}

//var _ dnaca.AdlContextVisitor = &DNAC_A_WalkerVisitor{}
//var _ dnaca.NameNodeContextVisitor = &DNAC_A_WalkerVisitor{}
//var _ dnaca.NameRuleContextVisitor = &DNAC_A_WalkerVisitor{}
//var _ dnaca.TypeNodeContextVisitor = &DNAC_A_WalkerVisitor{}
//var _ dnaca.ExnotationNodeContextVisitor = &DNAC_A_WalkerVisitor{}
//var _ dnaca.NameBodyNodeContextVisitor = &DNAC_A_WalkerVisitor{}
//var _ dnaca.AnnotationContextVisitor = &DNAC_A_WalkerVisitor{}
//var _ dnaca.TypeExpr_ContextVisitor = &DNAC_A_WalkerVisitor{}
//var _ dnaca.TypeParamsContextVisitor = &DNAC_A_WalkerVisitor{}
//var _ dnaca.JsonStrContextVisitor = &DNAC_A_WalkerVisitor{}
//var _ dnaca.JsonBoolContextVisitor = &DNAC_A_WalkerVisitor{}
//var _ dnaca.JsonNullContextVisitor = &DNAC_A_WalkerVisitor{}
//var _ dnaca.JsonIntContextVisitor = &DNAC_A_WalkerVisitor{}
//var _ dnaca.JsonFloatContextVisitor = &DNAC_A_WalkerVisitor{}
//var _ dnaca.JsonArrayContextVisitor = &DNAC_A_WalkerVisitor{}
//var _ dnaca.JsonObjContextVisitor = &DNAC_A_WalkerVisitor{}

//func (v *DNAC_A_WalkerVisitor) VisitNext(node antlr.Tree, current interface{}) bool {
//    return true
//}
//func (v *DNAC_A_WalkerVisitor) VisitRest(node antlr.RuleNode, current interface{}) bool {
//    return true
//}
//func (v *DNAC_A_WalkerVisitor) AggregateResult(aggregate, nextResult interface{}) (result interface{}) {
//    return nextResult
//}
//func (v *DNAC_A_WalkerVisitor) VisitTerminal(node antlr.TerminalNode) {
//}
//func (v *DNAC_A_WalkerVisitor) VisitErrorNode(node antlr.ErrorNode) {
//}
//func (v *DNAC_A_WalkerVisitor) EnterEveryRule(ctx antlr.RuleNode) {
//    v.indent += " "
//    fmt.Printf("%s %T\n", v.indent, ctx)
//}
//func (v *DNAC_A_WalkerVisitor) ExitEveryRule(ctx antlr.RuleNode) {
//    v.indent = v.indent[:len(v.indent)-1]
//}

//indent := ""
//parsetree.VisitFunc(&dnaca.DNAC_A_WalkerHandlers{
//EnterEveryRule: func(ctx antlr.RuleNode){indent += "\t"},
//ExitEveryRule:  func(ctx antlr.RuleNode){indent = indent[1:]},
//Adl: func(ctx *dnaca.AdlContext, this *dnaca.DNAC_A_WalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//NameNode: func(ctx *dnaca.NameNodeContext, this *dnaca.DNAC_A_WalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//NameRule: func(ctx *dnaca.NameRuleContext, this *dnaca.DNAC_A_WalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TypeNode: func(ctx *dnaca.TypeNodeContext, this *dnaca.DNAC_A_WalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//ExnotationNode: func(ctx *dnaca.ExnotationNodeContext, this *dnaca.DNAC_A_WalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//NameBodyNode: func(ctx *dnaca.NameBodyNodeContext, this *dnaca.DNAC_A_WalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Annotation: func(ctx *dnaca.AnnotationContext, this *dnaca.DNAC_A_WalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TypeExpr_: func(ctx *dnaca.TypeExpr_Context, this *dnaca.DNAC_A_WalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TypeParams: func(ctx *dnaca.TypeParamsContext, this *dnaca.DNAC_A_WalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonStr: func(ctx *dnaca.JsonStrContext, this *dnaca.DNAC_A_WalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonBool: func(ctx *dnaca.JsonBoolContext, this *dnaca.DNAC_A_WalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonNull: func(ctx *dnaca.JsonNullContext, this *dnaca.DNAC_A_WalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonInt: func(ctx *dnaca.JsonIntContext, this *dnaca.DNAC_A_WalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonFloat: func(ctx *dnaca.JsonFloatContext, this *dnaca.DNAC_A_WalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonArray: func(ctx *dnaca.JsonArrayContext, this *dnaca.DNAC_A_WalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonObj: func(ctx *dnaca.JsonObjContext, this *dnaca.DNAC_A_WalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},

//}

//func (v *DNAC_A_WalkerVisitor) VisitAdl(ctx *dnaca.AdlContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *DNAC_A_WalkerVisitor) VisitNameNode(ctx *dnaca.NameNodeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *DNAC_A_WalkerVisitor) VisitNameRule(ctx *dnaca.NameRuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *DNAC_A_WalkerVisitor) VisitTypeNode(ctx *dnaca.TypeNodeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *DNAC_A_WalkerVisitor) VisitExnotationNode(ctx *dnaca.ExnotationNodeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *DNAC_A_WalkerVisitor) VisitNameBodyNode(ctx *dnaca.NameBodyNodeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *DNAC_A_WalkerVisitor) VisitAnnotation(ctx *dnaca.AnnotationContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *DNAC_A_WalkerVisitor) VisitTypeExpr_(ctx *dnaca.TypeExpr_Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *DNAC_A_WalkerVisitor) VisitTypeParams(ctx *dnaca.TypeParamsContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *DNAC_A_WalkerVisitor) VisitJsonStr(ctx *dnaca.JsonStrContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *DNAC_A_WalkerVisitor) VisitJsonBool(ctx *dnaca.JsonBoolContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *DNAC_A_WalkerVisitor) VisitJsonNull(ctx *dnaca.JsonNullContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *DNAC_A_WalkerVisitor) VisitJsonInt(ctx *dnaca.JsonIntContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *DNAC_A_WalkerVisitor) VisitJsonFloat(ctx *dnaca.JsonFloatContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *DNAC_A_WalkerVisitor) VisitJsonArray(ctx *dnaca.JsonArrayContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *DNAC_A_WalkerVisitor) VisitJsonObj(ctx *dnaca.JsonObjContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}

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
