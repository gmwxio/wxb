// Generated from ADLWalker.g4 by ANTLR 4.7.

package walker // ADLWalker
//import "github.com/wxio/goantlr"
//import "generated code if in another package"

//// Commented out basic implementation for your convenience.

//func Example(s string) {
//  // Setup
//  input := antlr.NewInputStream(s)
//  lexer := walker.NewADLWalkerLexer(input)
//  stream := antlr.NewCommonTokenStream(lexer, 0)
//  p := walker.NewADLWalkerParser(stream)

//  // Antlr error listener - turns reports (ambiguity etc) into syntax errors
//  p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

//  // Custom error listener, register before the parse
//  el := &ADLWalkerErrorListener{}
//  p.AddErrorListener(el)

//  // Parse - start rule
//  tree := p.Start()

//  // Antlr provided parse tree representation
//  sexpr := antlr.TreesStringTree(tree, nil, p)
//  fmt.Printf("%s\n", sexpr)

//  // Custom listener
//  l := &ADLWalkerListener{}
//  antlr.ParseTreeWalkerDefault.Walk(l, tree)

//  // Custom visitor
//  v := &ADLWalkerVisitor{}
//  tree.Accept(v)
// }

//// implemented all listeners methods
//var _ walker.ADLWalkerListener = &ADLWalkerListener{}
//// implemented specific
//var _ walker.AdlEntryListener = &ADLWalkerListener{}
//var _ walker.AdlExitListener = &ADLWalkerListener{}

//var _ walker.JsonEntryListener = &ADLWalkerListener{}
//var _ walker.JsonExitListener = &ADLWalkerListener{}

//var _ walker.ModuleEntryListener = &ADLWalkerListener{}
//var _ walker.ModuleExitListener = &ADLWalkerListener{}

//var _ walker.ImportModuleEntryListener = &ADLWalkerListener{}
//var _ walker.ImportModuleExitListener = &ADLWalkerListener{}

//var _ walker.ImportScopedModuleEntryListener = &ADLWalkerListener{}
//var _ walker.ImportScopedModuleExitListener = &ADLWalkerListener{}

//var _ walker.StructEntryListener = &ADLWalkerListener{}
//var _ walker.StructExitListener = &ADLWalkerListener{}

//var _ walker.UnionEntryListener = &ADLWalkerListener{}
//var _ walker.UnionExitListener = &ADLWalkerListener{}

//var _ walker.TypeEntryListener = &ADLWalkerListener{}
//var _ walker.TypeExitListener = &ADLWalkerListener{}

//var _ walker.NewtypeEntryListener = &ADLWalkerListener{}
//var _ walker.NewtypeExitListener = &ADLWalkerListener{}

//var _ walker.ModAnnoEntryListener = &ADLWalkerListener{}
//var _ walker.ModAnnoExitListener = &ADLWalkerListener{}

//var _ walker.DeclAnnoEntryListener = &ADLWalkerListener{}
//var _ walker.DeclAnnoExitListener = &ADLWalkerListener{}

//var _ walker.FieldAnnoEntryListener = &ADLWalkerListener{}
//var _ walker.FieldAnnoExitListener = &ADLWalkerListener{}

//var _ walker.TypeParamErrorEntryListener = &ADLWalkerListener{}
//var _ walker.TypeParamErrorExitListener = &ADLWalkerListener{}

//var _ walker.FieldEntryListener = &ADLWalkerListener{}
//var _ walker.FieldExitListener = &ADLWalkerListener{}

//var _ walker.AnnotationEntryListener = &ADLWalkerListener{}
//var _ walker.AnnotationExitListener = &ADLWalkerListener{}

//var _ walker.TypeExpr_EntryListener = &ADLWalkerListener{}
//var _ walker.TypeExpr_ExitListener = &ADLWalkerListener{}

//var _ walker.TypeParamsEntryListener = &ADLWalkerListener{}
//var _ walker.TypeParamsExitListener = &ADLWalkerListener{}

//var _ walker.JsonStrEntryListener = &ADLWalkerListener{}
//var _ walker.JsonStrExitListener = &ADLWalkerListener{}

//var _ walker.JsonBoolEntryListener = &ADLWalkerListener{}
//var _ walker.JsonBoolExitListener = &ADLWalkerListener{}

//var _ walker.JsonNullEntryListener = &ADLWalkerListener{}
//var _ walker.JsonNullExitListener = &ADLWalkerListener{}

//var _ walker.JsonIntEntryListener = &ADLWalkerListener{}
//var _ walker.JsonIntExitListener = &ADLWalkerListener{}

//var _ walker.JsonFloatEntryListener = &ADLWalkerListener{}
//var _ walker.JsonFloatExitListener = &ADLWalkerListener{}

//var _ walker.JsonArrayEntryListener = &ADLWalkerListener{}
//var _ walker.JsonArrayExitListener = &ADLWalkerListener{}

//var _ walker.JsonObjEntryListener = &ADLWalkerListener{}
//var _ walker.JsonObjExitListener = &ADLWalkerListener{}

//type ADLWalkerListener struct {
//}

//type ADLWalkerErrorListener struct {
//    Warning string
//    Err     error
//    Debug   bool
//}

// func (cb *ADLWalkerErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
//  if cb.Debug {
//      fmt.Printf("SyntaxError %d:%d <%s>\n", line, column, msg)
//  }
//  if strings.HasPrefix(msg, "report") { // TODO remove NewDiagnosticErrorListener and move warning to ReportAmbiguity etc. when getDecisionDescription is make public
//      cb.Warning = fmt.Sprintf("At %d:%d <%s>", line, column, msg)
//  } else {
//      cb.Err = fmt.Errorf("SyntaxError %d:%d <%s>", line, column, msg)
//  }
// }
// func (cb *ADLWalkerErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int,
//  exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportAmbiguity rec:%v dfs:%v start:%d stop:%d, exact:%v, ambigAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
//  }
// }
// func (cb *ADLWalkerErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportAttemptingFullContext rec:%v dfs:%v start:%d stop:%d, conflictingAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
//  }
// }
// func (cb *ADLWalkerErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportContextSensitivity rec:%v dfs:%v start:%d stop:%d, prediction:%v configs:%v\n", recognizer, dfa, startIndex, stopIndex, prediction, configs)
//  }
// }

//// antlr.ParseTreeListener implementation.
//// All required.

//func (s *ADLWalkerListener ) VisitTerminal(node  antlr.TerminalNode) {   }
//func (s *ADLWalkerListener ) VisitErrorNode(node antlr.ErrorNode)    {   }
//func (s *ADLWalkerListener ) EnterEveryRule(ctx antlr.ParserRuleContext) {  }
//func (s *ADLWalkerListener ) ExitEveryRule(ctx antlr.ParserRuleContext) {  }

//// Only implemented as needed.

//func (s *ADLWalkerListener) EnterAdl(ctx walker.*walker.AdlContext) {}
//func (s *ADLWalkerListener) ExitAdl(ctx walker.*walker.AdlContext) {}

//func (s *ADLWalkerListener) EnterJson(ctx walker.*walker.JsonContext) {}
//func (s *ADLWalkerListener) ExitJson(ctx walker.*walker.JsonContext) {}

//func (s *ADLWalkerListener) EnterModule(ctx walker.*walker.ModuleContext) {}
//func (s *ADLWalkerListener) ExitModule(ctx walker.*walker.ModuleContext) {}

//func (s *ADLWalkerListener) EnterImportModule(ctx walker.*walker.ImportModuleContext) {}
//func (s *ADLWalkerListener) ExitImportModule(ctx walker.*walker.ImportModuleContext) {}

//func (s *ADLWalkerListener) EnterImportScopedModule(ctx walker.*walker.ImportScopedModuleContext) {}
//func (s *ADLWalkerListener) ExitImportScopedModule(ctx walker.*walker.ImportScopedModuleContext) {}

//func (s *ADLWalkerListener) EnterStruct(ctx walker.*walker.StructContext) {}
//func (s *ADLWalkerListener) ExitStruct(ctx walker.*walker.StructContext) {}

//func (s *ADLWalkerListener) EnterUnion(ctx walker.*walker.UnionContext) {}
//func (s *ADLWalkerListener) ExitUnion(ctx walker.*walker.UnionContext) {}

//func (s *ADLWalkerListener) EnterType(ctx walker.*walker.TypeContext) {}
//func (s *ADLWalkerListener) ExitType(ctx walker.*walker.TypeContext) {}

//func (s *ADLWalkerListener) EnterNewtype(ctx walker.*walker.NewtypeContext) {}
//func (s *ADLWalkerListener) ExitNewtype(ctx walker.*walker.NewtypeContext) {}

//func (s *ADLWalkerListener) EnterModAnno(ctx walker.*walker.ModAnnoContext) {}
//func (s *ADLWalkerListener) ExitModAnno(ctx walker.*walker.ModAnnoContext) {}

//func (s *ADLWalkerListener) EnterDeclAnno(ctx walker.*walker.DeclAnnoContext) {}
//func (s *ADLWalkerListener) ExitDeclAnno(ctx walker.*walker.DeclAnnoContext) {}

//func (s *ADLWalkerListener) EnterFieldAnno(ctx walker.*walker.FieldAnnoContext) {}
//func (s *ADLWalkerListener) ExitFieldAnno(ctx walker.*walker.FieldAnnoContext) {}

//func (s *ADLWalkerListener) EnterTypeParamError(ctx walker.*walker.TypeParamErrorContext) {}
//func (s *ADLWalkerListener) ExitTypeParamError(ctx walker.*walker.TypeParamErrorContext) {}

//func (s *ADLWalkerListener) EnterField(ctx walker.*walker.FieldContext) {}
//func (s *ADLWalkerListener) ExitField(ctx walker.*walker.FieldContext) {}

//func (s *ADLWalkerListener) EnterAnnotation(ctx walker.*walker.AnnotationContext) {}
//func (s *ADLWalkerListener) ExitAnnotation(ctx walker.*walker.AnnotationContext) {}

//func (s *ADLWalkerListener) EnterTypeExpr_(ctx walker.*walker.TypeExpr_Context) {}
//func (s *ADLWalkerListener) ExitTypeExpr_(ctx walker.*walker.TypeExpr_Context) {}

//func (s *ADLWalkerListener) EnterTypeParams(ctx walker.*walker.TypeParamsContext) {}
//func (s *ADLWalkerListener) ExitTypeParams(ctx walker.*walker.TypeParamsContext) {}

//func (s *ADLWalkerListener) EnterJsonStr(ctx walker.*walker.JsonStrContext) {}
//func (s *ADLWalkerListener) ExitJsonStr(ctx walker.*walker.JsonStrContext) {}

//func (s *ADLWalkerListener) EnterJsonBool(ctx walker.*walker.JsonBoolContext) {}
//func (s *ADLWalkerListener) ExitJsonBool(ctx walker.*walker.JsonBoolContext) {}

//func (s *ADLWalkerListener) EnterJsonNull(ctx walker.*walker.JsonNullContext) {}
//func (s *ADLWalkerListener) ExitJsonNull(ctx walker.*walker.JsonNullContext) {}

//func (s *ADLWalkerListener) EnterJsonInt(ctx walker.*walker.JsonIntContext) {}
//func (s *ADLWalkerListener) ExitJsonInt(ctx walker.*walker.JsonIntContext) {}

//func (s *ADLWalkerListener) EnterJsonFloat(ctx walker.*walker.JsonFloatContext) {}
//func (s *ADLWalkerListener) ExitJsonFloat(ctx walker.*walker.JsonFloatContext) {}

//func (s *ADLWalkerListener) EnterJsonArray(ctx walker.*walker.JsonArrayContext) {}
//func (s *ADLWalkerListener) ExitJsonArray(ctx walker.*walker.JsonArrayContext) {}

//func (s *ADLWalkerListener) EnterJsonObj(ctx walker.*walker.JsonObjContext) {}
//func (s *ADLWalkerListener) ExitJsonObj(ctx walker.*walker.JsonObjContext) {}
