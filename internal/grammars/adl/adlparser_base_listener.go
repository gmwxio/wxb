// Generated from ADLParser.g4 by ANTLR 4.7.

package parser // ADLParser

//import "github.com/wxio/goantlr"
//import "generated code if in another package"

//// Commented out basic implementation for your convenience.

//func Example(s string) {
//  // Setup
//  input := antlr.NewInputStream(s)
//  lexer := NewADLParserLexer(input)
//  stream := antlr.NewCommonTokenStream(lexer, 0)
//  p := NewADLParserParser(stream)

//  // Antlr error listener - turns reports (ambiguity etc) into syntax errors
//  p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

//  // Custom error listener, register before the parse
//  el := &ADLParserErrorListener{}
//  p.AddErrorListener(el)

//  // Parse - start rule
//  tree := p.Start()

//  // Antlr provided parse tree representation
//  sexpr := antlr.TreesStringTree(tree, nil, p)
//  fmt.Printf("%s\n", sexpr)

//  // Custom listener
//  l := &ADLParserListener{}
//  antlr.ParseTreeWalkerDefault.Walk(l, tree)

//  // Custom visitor
//  v := &ADLParserVisitor{}
//  tree.Accept(v)
// }

//// implemented all listeners methods
//var _ ADLParserListener = &ADLParserListener{}
//// implemented specific
//var _ AdlEntryListener = &ADLParserListener{}
//var _ AdlExitListener = &ADLParserListener{}

//var _ ModuleStatementEntryListener = &ADLParserListener{}
//var _ ModuleStatementExitListener = &ADLParserListener{}

//var _ ImportStatementEntryListener = &ADLParserListener{}
//var _ ImportStatementExitListener = &ADLParserListener{}

//var _ LocalAnnoEntryListener = &ADLParserListener{}
//var _ LocalAnnoExitListener = &ADLParserListener{}

//var _ DocAnnoEntryListener = &ADLParserListener{}
//var _ DocAnnoExitListener = &ADLParserListener{}

//var _ StructOrUnionEntryListener = &ADLParserListener{}
//var _ StructOrUnionExitListener = &ADLParserListener{}

//var _ TypeOrNewtypeEntryListener = &ADLParserListener{}
//var _ TypeOrNewtypeExitListener = &ADLParserListener{}

//var _ ModuleAnnotationEntryListener = &ADLParserListener{}
//var _ ModuleAnnotationExitListener = &ADLParserListener{}

//var _ DeclAnnotationEntryListener = &ADLParserListener{}
//var _ DeclAnnotationExitListener = &ADLParserListener{}

//var _ FieldAnnotationEntryListener = &ADLParserListener{}
//var _ FieldAnnotationExitListener = &ADLParserListener{}

//var _ TypeParameterEntryListener = &ADLParserListener{}
//var _ TypeParameterExitListener = &ADLParserListener{}

//var _ ErrorTypeParamEntryListener = &ADLParserListener{}
//var _ ErrorTypeParamExitListener = &ADLParserListener{}

//var _ TypeParamErrorEntryListener = &ADLParserListener{}
//var _ TypeParamErrorExitListener = &ADLParserListener{}

//var _ TypeExpressionEntryListener = &ADLParserListener{}
//var _ TypeExpressionExitListener = &ADLParserListener{}

//var _ TypeExpressionElemEntryListener = &ADLParserListener{}
//var _ TypeExpressionElemExitListener = &ADLParserListener{}

//var _ FieldStatementEntryListener = &ADLParserListener{}
//var _ FieldStatementExitListener = &ADLParserListener{}

//var _ StringStatementEntryListener = &ADLParserListener{}
//var _ StringStatementExitListener = &ADLParserListener{}

//var _ TrueFalseNullEntryListener = &ADLParserListener{}
//var _ TrueFalseNullExitListener = &ADLParserListener{}

//var _ NumberStatementEntryListener = &ADLParserListener{}
//var _ NumberStatementExitListener = &ADLParserListener{}

//var _ FloatStatementEntryListener = &ADLParserListener{}
//var _ FloatStatementExitListener = &ADLParserListener{}

//var _ ArrayStatementEntryListener = &ADLParserListener{}
//var _ ArrayStatementExitListener = &ADLParserListener{}

//var _ ObjStatementEntryListener = &ADLParserListener{}
//var _ ObjStatementExitListener = &ADLParserListener{}

//type ADLParserListener struct {
//}

//type ADLParserErrorListener struct {
//    Warning string
//    Err     error
//    Debug   bool
//}

// func (cb *ADLParserErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
//  if cb.Debug {
//      fmt.Printf("SyntaxError %d:%d <%s>\n", line, column, msg)
//  }
//  if strings.HasPrefix(msg, "report") { // TODO remove NewDiagnosticErrorListener and move warning to ReportAmbiguity etc. when getDecisionDescription is make public
//      cb.Warning = fmt.Sprintf("At %d:%d <%s>", line, column, msg)
//  } else {
//      cb.Err = fmt.Errorf("SyntaxError %d:%d <%s>", line, column, msg)
//  }
// }
// func (cb *ADLParserErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int,
//  exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportAmbiguity rec:%v dfs:%v start:%d stop:%d, exact:%v, ambigAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
//  }
// }
// func (cb *ADLParserErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportAttemptingFullContext rec:%v dfs:%v start:%d stop:%d, conflictingAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
//  }
// }
// func (cb *ADLParserErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportContextSensitivity rec:%v dfs:%v start:%d stop:%d, prediction:%v configs:%v\n", recognizer, dfa, startIndex, stopIndex, prediction, configs)
//  }
// }

//// antlr.ParseTreeListener implementation.
//// All required.

//func (s *ADLParserListener ) VisitTerminal(node  antlr.TerminalNode) {   }
//func (s *ADLParserListener ) VisitErrorNode(node antlr.ErrorNode)    {   }
//func (s *ADLParserListener ) EnterEveryRule(ctx antlr.ParserRuleContext) {  }
//func (s *ADLParserListener ) ExitEveryRule(ctx antlr.ParserRuleContext) {  }

//// Only implemented as needed.

//func (s *ADLParserListener) EnterAdl(ctx *AdlContext) {}
//func (s *ADLParserListener) ExitAdl(ctx *AdlContext) {}

//func (s *ADLParserListener) EnterModuleStatement(ctx *ModuleStatementContext) {}
//func (s *ADLParserListener) ExitModuleStatement(ctx *ModuleStatementContext) {}

//func (s *ADLParserListener) EnterImportStatement(ctx *ImportStatementContext) {}
//func (s *ADLParserListener) ExitImportStatement(ctx *ImportStatementContext) {}

//func (s *ADLParserListener) EnterLocalAnno(ctx *LocalAnnoContext) {}
//func (s *ADLParserListener) ExitLocalAnno(ctx *LocalAnnoContext) {}

//func (s *ADLParserListener) EnterDocAnno(ctx *DocAnnoContext) {}
//func (s *ADLParserListener) ExitDocAnno(ctx *DocAnnoContext) {}

//func (s *ADLParserListener) EnterStructOrUnion(ctx *StructOrUnionContext) {}
//func (s *ADLParserListener) ExitStructOrUnion(ctx *StructOrUnionContext) {}

//func (s *ADLParserListener) EnterTypeOrNewtype(ctx *TypeOrNewtypeContext) {}
//func (s *ADLParserListener) ExitTypeOrNewtype(ctx *TypeOrNewtypeContext) {}

//func (s *ADLParserListener) EnterModuleAnnotation(ctx *ModuleAnnotationContext) {}
//func (s *ADLParserListener) ExitModuleAnnotation(ctx *ModuleAnnotationContext) {}

//func (s *ADLParserListener) EnterDeclAnnotation(ctx *DeclAnnotationContext) {}
//func (s *ADLParserListener) ExitDeclAnnotation(ctx *DeclAnnotationContext) {}

//func (s *ADLParserListener) EnterFieldAnnotation(ctx *FieldAnnotationContext) {}
//func (s *ADLParserListener) ExitFieldAnnotation(ctx *FieldAnnotationContext) {}

//func (s *ADLParserListener) EnterTypeParameter(ctx *TypeParameterContext) {}
//func (s *ADLParserListener) ExitTypeParameter(ctx *TypeParameterContext) {}

//func (s *ADLParserListener) EnterErrorTypeParam(ctx *ErrorTypeParamContext) {}
//func (s *ADLParserListener) ExitErrorTypeParam(ctx *ErrorTypeParamContext) {}

//func (s *ADLParserListener) EnterTypeParamError(ctx *TypeParamErrorContext) {}
//func (s *ADLParserListener) ExitTypeParamError(ctx *TypeParamErrorContext) {}

//func (s *ADLParserListener) EnterTypeExpression(ctx *TypeExpressionContext) {}
//func (s *ADLParserListener) ExitTypeExpression(ctx *TypeExpressionContext) {}

//func (s *ADLParserListener) EnterTypeExpressionElem(ctx *TypeExpressionElemContext) {}
//func (s *ADLParserListener) ExitTypeExpressionElem(ctx *TypeExpressionElemContext) {}

//func (s *ADLParserListener) EnterFieldStatement(ctx *FieldStatementContext) {}
//func (s *ADLParserListener) ExitFieldStatement(ctx *FieldStatementContext) {}

//func (s *ADLParserListener) EnterStringStatement(ctx *StringStatementContext) {}
//func (s *ADLParserListener) ExitStringStatement(ctx *StringStatementContext) {}

//func (s *ADLParserListener) EnterTrueFalseNull(ctx *TrueFalseNullContext) {}
//func (s *ADLParserListener) ExitTrueFalseNull(ctx *TrueFalseNullContext) {}

//func (s *ADLParserListener) EnterNumberStatement(ctx *NumberStatementContext) {}
//func (s *ADLParserListener) ExitNumberStatement(ctx *NumberStatementContext) {}

//func (s *ADLParserListener) EnterFloatStatement(ctx *FloatStatementContext) {}
//func (s *ADLParserListener) ExitFloatStatement(ctx *FloatStatementContext) {}

//func (s *ADLParserListener) EnterArrayStatement(ctx *ArrayStatementContext) {}
//func (s *ADLParserListener) ExitArrayStatement(ctx *ArrayStatementContext) {}

//func (s *ADLParserListener) EnterObjStatement(ctx *ObjStatementContext) {}
//func (s *ADLParserListener) ExitObjStatement(ctx *ObjStatementContext) {}
