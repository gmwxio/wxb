// Generated from DNAC_B_Walker.g4 by ANTLR 4.7.

package dnacb // DNAC_B_Walker
// See base listener file for example implementations

//import "github.com/wxio/goantlr"

// import "generate package"

//type DNAC_B_WalkerVisitor struct {
//    *antlr.BaseParseTreeVisitor
//    indent string
//}

//var _ antlr.ParseTreeVisitor = &DNAC_B_WalkerVisitor{}
//var _ antlr.AggregateResultVisitor = &DNAC_B_WalkerVisitor{}
//var _ antlr.VisitNextCheck = &DNAC_B_WalkerVisitor{}
//var _ antlr.VisitRestCheck = &DNAC_B_WalkerVisitor{}
//var _ antlr.EnterEveryRuleVisitor = &DNAC_B_WalkerVisitor{}
//var _ antlr.ExitEveryRuleVisitor = &DNAC_B_WalkerVisitor{}

//var _ dnacb.DnacContextVisitor = &DNAC_B_WalkerVisitor{}
//var _ dnacb.NameNodeContextVisitor = &DNAC_B_WalkerVisitor{}
//var _ dnacb.NameRuleContextVisitor = &DNAC_B_WalkerVisitor{}
//var _ dnacb.TypeNodeContextVisitor = &DNAC_B_WalkerVisitor{}
//var _ dnacb.NameBodyNodeContextVisitor = &DNAC_B_WalkerVisitor{}
//var _ dnacb.ExnotationNodeContextVisitor = &DNAC_B_WalkerVisitor{}

//func (v *DNAC_B_WalkerVisitor) VisitNext(node antlr.Tree, current interface{}) bool {
//    return true
//}
//func (v *DNAC_B_WalkerVisitor) VisitRest(node antlr.RuleNode, current interface{}) bool {
//    return true
//}
//func (v *DNAC_B_WalkerVisitor) AggregateResult(aggregate, nextResult interface{}) (result interface{}) {
//    return nextResult
//}
//func (v *DNAC_B_WalkerVisitor) VisitTerminal(node antlr.TerminalNode) {
//}
//func (v *DNAC_B_WalkerVisitor) VisitErrorNode(node antlr.ErrorNode) {
//}
//func (v *DNAC_B_WalkerVisitor) EnterEveryRule(ctx antlr.RuleNode) {
//    v.indent += " "
//    fmt.Printf("%s %T\n", v.indent, ctx)
//}
//func (v *DNAC_B_WalkerVisitor) ExitEveryRule(ctx antlr.RuleNode) {
//    v.indent = v.indent[:len(v.indent)-1]
//}

//indent := ""
//parsetree.VisitFunc(&dnacb.DNAC_B_WalkerHandlers{
//EnterEveryRule: func(ctx antlr.RuleNode){indent += "\t"},
//ExitEveryRule:  func(ctx antlr.RuleNode){indent = indent[1:]},
//Dnac: func(ctx *dnacb.DnacContext, this *dnacb.DNAC_B_WalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//NameNode: func(ctx *dnacb.NameNodeContext, this *dnacb.DNAC_B_WalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//NameRule: func(ctx *dnacb.NameRuleContext, this *dnacb.DNAC_B_WalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TypeNode: func(ctx *dnacb.TypeNodeContext, this *dnacb.DNAC_B_WalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//NameBodyNode: func(ctx *dnacb.NameBodyNodeContext, this *dnacb.DNAC_B_WalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//ExnotationNode: func(ctx *dnacb.ExnotationNodeContext, this *dnacb.DNAC_B_WalkerHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},

//}

//func (v *DNAC_B_WalkerVisitor) VisitDnac(ctx *dnacb.DnacContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *DNAC_B_WalkerVisitor) VisitNameNode(ctx *dnacb.NameNodeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *DNAC_B_WalkerVisitor) VisitNameRule(ctx *dnacb.NameRuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *DNAC_B_WalkerVisitor) VisitTypeNode(ctx *dnacb.TypeNodeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *DNAC_B_WalkerVisitor) VisitNameBodyNode(ctx *dnacb.NameBodyNodeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *DNAC_B_WalkerVisitor) VisitExnotationNode(ctx *dnacb.ExnotationNodeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}

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
