// Generated from ADLParser.g4 by ANTLR 4.7.

package parser // ADLParser

// See base listener file for example implementations

//import "github.com/wxio/goantlr"

// import "generate package"

//type ADLParserVisitor struct {
//    *antlr.BaseParseTreeVisitor
//    indent string
//}

//var _ antlr.ParseTreeVisitor = &ADLParserVisitor{}
//var _ antlr.AggregateResultVisitor = &ADLParserVisitor{}
//var _ antlr.VisitNextCheck = &ADLParserVisitor{}
//var _ antlr.VisitRestCheck = &ADLParserVisitor{}
//var _ antlr.EnterEveryRuleVisitor = &ADLParserVisitor{}
//var _ antlr.ExitEveryRuleVisitor = &ADLParserVisitor{}

//var _ AdlContextVisitor = &ADLParserVisitor{}
//var _ ModuleStatementContextVisitor = &ADLParserVisitor{}
//var _ ImportStatementContextVisitor = &ADLParserVisitor{}
//var _ LocalAnnoContextVisitor = &ADLParserVisitor{}
//var _ DocAnnoContextVisitor = &ADLParserVisitor{}
//var _ StructOrUnionContextVisitor = &ADLParserVisitor{}
//var _ TypeOrNewtypeContextVisitor = &ADLParserVisitor{}
//var _ ModuleAnnotationContextVisitor = &ADLParserVisitor{}
//var _ DeclAnnotationContextVisitor = &ADLParserVisitor{}
//var _ FieldAnnotationContextVisitor = &ADLParserVisitor{}
//var _ TypeParameterContextVisitor = &ADLParserVisitor{}
//var _ ErrorTypeParamContextVisitor = &ADLParserVisitor{}
//var _ TypeParamErrorContextVisitor = &ADLParserVisitor{}
//var _ TypeExpressionContextVisitor = &ADLParserVisitor{}
//var _ TypeExpressionElemContextVisitor = &ADLParserVisitor{}
//var _ FieldStatementContextVisitor = &ADLParserVisitor{}
//var _ StringStatementContextVisitor = &ADLParserVisitor{}
//var _ TrueFalseNullContextVisitor = &ADLParserVisitor{}
//var _ NumberStatementContextVisitor = &ADLParserVisitor{}
//var _ FloatStatementContextVisitor = &ADLParserVisitor{}
//var _ ArrayStatementContextVisitor = &ADLParserVisitor{}
//var _ ObjStatementContextVisitor = &ADLParserVisitor{}

//func (v *ADLParserVisitor) VisitNext(node antlr.Tree, current interface{}) bool {
//    return true
//}
//func (v *ADLParserVisitor) VisitRest(node antlr.RuleNode, current interface{}) bool {
//    return true
//}
//func (v *ADLParserVisitor) AggregateResult(aggregate, nextResult interface{}) (result interface{}) {
//    return nextResult
//}
//func (v *ADLParserVisitor) VisitTerminal(node antlr.TerminalNode) {
//}
//func (v *ADLParserVisitor) VisitErrorNode(node antlr.ErrorNode) {
//}
//func (v *ADLParserVisitor) EnterEveryRule(ctx antlr.RuleNode) {
//    v.indent += " "
//    fmt.Printf("%s %T\n", v.indent, ctx)
//}
//func (v *ADLParserVisitor) ExitEveryRule(ctx antlr.RuleNode) {
//    v.indent = v.indent[:len(v.indent)-1]
//}

//indent := ""
//parsetree.VisitFunc(&parser.ADLParserHandlers{
//EnterEveryRule: func(ctx antlr.RuleNode){indent += "\t"},
//ExitEveryRule:  func(ctx antlr.RuleNode){indent = indent[1:]},
//Adl: func(ctx *parser.AdlContext, this *parser.ADLParserHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//ModuleStatement: func(ctx *parser.ModuleStatementContext, this *parser.ADLParserHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//ImportStatement: func(ctx *parser.ImportStatementContext, this *parser.ADLParserHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//LocalAnno: func(ctx *parser.LocalAnnoContext, this *parser.ADLParserHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//DocAnno: func(ctx *parser.DocAnnoContext, this *parser.ADLParserHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//StructOrUnion: func(ctx *parser.StructOrUnionContext, this *parser.ADLParserHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TypeOrNewtype: func(ctx *parser.TypeOrNewtypeContext, this *parser.ADLParserHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//ModuleAnnotation: func(ctx *parser.ModuleAnnotationContext, this *parser.ADLParserHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//DeclAnnotation: func(ctx *parser.DeclAnnotationContext, this *parser.ADLParserHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//FieldAnnotation: func(ctx *parser.FieldAnnotationContext, this *parser.ADLParserHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TypeParameter: func(ctx *parser.TypeParameterContext, this *parser.ADLParserHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//ErrorTypeParam: func(ctx *parser.ErrorTypeParamContext, this *parser.ADLParserHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TypeParamError: func(ctx *parser.TypeParamErrorContext, this *parser.ADLParserHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TypeExpression: func(ctx *parser.TypeExpressionContext, this *parser.ADLParserHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TypeExpressionElem: func(ctx *parser.TypeExpressionElemContext, this *parser.ADLParserHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//FieldStatement: func(ctx *parser.FieldStatementContext, this *parser.ADLParserHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//StringStatement: func(ctx *parser.StringStatementContext, this *parser.ADLParserHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TrueFalseNull: func(ctx *parser.TrueFalseNullContext, this *parser.ADLParserHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//NumberStatement: func(ctx *parser.NumberStatementContext, this *parser.ADLParserHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//FloatStatement: func(ctx *parser.FloatStatementContext, this *parser.ADLParserHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//ArrayStatement: func(ctx *parser.ArrayStatementContext, this *parser.ADLParserHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//ObjStatement: func(ctx *parser.ObjStatementContext, this *parser.ADLParserHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},

//}

//func (v *ADLParserVisitor) VisitAdl(ctx parser.IAdlContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLParserVisitor) VisitModuleStatement(ctx parser.IModuleStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLParserVisitor) VisitImportStatement(ctx parser.IImportStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLParserVisitor) VisitLocalAnno(ctx parser.ILocalAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLParserVisitor) VisitDocAnno(ctx parser.IDocAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLParserVisitor) VisitStructOrUnion(ctx parser.IStructOrUnionContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLParserVisitor) VisitTypeOrNewtype(ctx parser.ITypeOrNewtypeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLParserVisitor) VisitModuleAnnotation(ctx parser.IModuleAnnotationContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLParserVisitor) VisitDeclAnnotation(ctx parser.IDeclAnnotationContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLParserVisitor) VisitFieldAnnotation(ctx parser.IFieldAnnotationContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLParserVisitor) VisitTypeParameter(ctx parser.ITypeParameterContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLParserVisitor) VisitErrorTypeParam(ctx parser.IErrorTypeParamContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLParserVisitor) VisitTypeParamError(ctx parser.ITypeParamErrorContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLParserVisitor) VisitTypeExpression(ctx parser.ITypeExpressionContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLParserVisitor) VisitTypeExpressionElem(ctx parser.ITypeExpressionElemContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLParserVisitor) VisitFieldStatement(ctx parser.IFieldStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLParserVisitor) VisitStringStatement(ctx parser.IStringStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLParserVisitor) VisitTrueFalseNull(ctx parser.ITrueFalseNullContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLParserVisitor) VisitNumberStatement(ctx parser.INumberStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLParserVisitor) VisitFloatStatement(ctx parser.IFloatStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLParserVisitor) VisitArrayStatement(ctx parser.IArrayStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *ADLParserVisitor) VisitObjStatement(ctx parser.IObjStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}

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
