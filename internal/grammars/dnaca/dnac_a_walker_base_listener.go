// Generated from DNAC_A_Walker.g4 by ANTLR 4.7.

package dnaca // DNAC_A_Walker
//import "github.com/wxio/goantlr"
//import "generated code if in another package"

//// Commented out basic implementation for your convenience.

//func Example(s string) {
//  // Setup
//  input := antlr.NewInputStream(s)
//  lexer := dnaca.NewDNAC_A_WalkerLexer(input)
//  stream := antlr.NewCommonTokenStream(lexer, 0)
//  p := dnaca.NewDNAC_A_WalkerParser(stream)

//  // Antlr error listener - turns reports (ambiguity etc) into syntax errors
//  p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

//  // Custom error listener, register before the parse
//  el := &DNAC_A_WalkerErrorListener{}
//  p.AddErrorListener(el)

//  // Parse - start rule
//  tree := p.Start()

//  // Antlr provided parse tree representation
//  sexpr := antlr.TreesStringTree(tree, nil, p)
//  fmt.Printf("%s\n", sexpr)

//  // Custom listener
//  l := &DNAC_A_WalkerListener{}
//  antlr.ParseTreeWalkerDefault.Walk(l, tree)

//  // Custom visitor
//  v := &DNAC_A_WalkerVisitor{}
//  tree.Accept(v)
// }

//// implemented all listeners methods
//var _ dnaca.DNAC_A_WalkerListener = &DNAC_A_WalkerListener{}
//// implemented specific
//var _ dnaca.AdlEntryListener = &DNAC_A_WalkerListener{}
//var _ dnaca.AdlExitListener = &DNAC_A_WalkerListener{}

//var _ dnaca.NameNodeEntryListener = &DNAC_A_WalkerListener{}
//var _ dnaca.NameNodeExitListener = &DNAC_A_WalkerListener{}

//var _ dnaca.NameRuleEntryListener = &DNAC_A_WalkerListener{}
//var _ dnaca.NameRuleExitListener = &DNAC_A_WalkerListener{}

//var _ dnaca.TypeNodeEntryListener = &DNAC_A_WalkerListener{}
//var _ dnaca.TypeNodeExitListener = &DNAC_A_WalkerListener{}

//var _ dnaca.ExnotationNodeEntryListener = &DNAC_A_WalkerListener{}
//var _ dnaca.ExnotationNodeExitListener = &DNAC_A_WalkerListener{}

//var _ dnaca.NameBodyNodeEntryListener = &DNAC_A_WalkerListener{}
//var _ dnaca.NameBodyNodeExitListener = &DNAC_A_WalkerListener{}

//var _ dnaca.AnnotationEntryListener = &DNAC_A_WalkerListener{}
//var _ dnaca.AnnotationExitListener = &DNAC_A_WalkerListener{}

//var _ dnaca.TypeExpr_EntryListener = &DNAC_A_WalkerListener{}
//var _ dnaca.TypeExpr_ExitListener = &DNAC_A_WalkerListener{}

//var _ dnaca.TypeParamsEntryListener = &DNAC_A_WalkerListener{}
//var _ dnaca.TypeParamsExitListener = &DNAC_A_WalkerListener{}

//var _ dnaca.JsonStrEntryListener = &DNAC_A_WalkerListener{}
//var _ dnaca.JsonStrExitListener = &DNAC_A_WalkerListener{}

//var _ dnaca.JsonBoolEntryListener = &DNAC_A_WalkerListener{}
//var _ dnaca.JsonBoolExitListener = &DNAC_A_WalkerListener{}

//var _ dnaca.JsonNullEntryListener = &DNAC_A_WalkerListener{}
//var _ dnaca.JsonNullExitListener = &DNAC_A_WalkerListener{}

//var _ dnaca.JsonIntEntryListener = &DNAC_A_WalkerListener{}
//var _ dnaca.JsonIntExitListener = &DNAC_A_WalkerListener{}

//var _ dnaca.JsonFloatEntryListener = &DNAC_A_WalkerListener{}
//var _ dnaca.JsonFloatExitListener = &DNAC_A_WalkerListener{}

//var _ dnaca.JsonArrayEntryListener = &DNAC_A_WalkerListener{}
//var _ dnaca.JsonArrayExitListener = &DNAC_A_WalkerListener{}

//var _ dnaca.JsonObjEntryListener = &DNAC_A_WalkerListener{}
//var _ dnaca.JsonObjExitListener = &DNAC_A_WalkerListener{}

//type DNAC_A_WalkerListener struct {
//}

//type DNAC_A_WalkerErrorListener struct {
//    Warning string
//    Err     error
//    Debug   bool
//}

// func (cb *DNAC_A_WalkerErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
//  if cb.Debug {
//      fmt.Printf("SyntaxError %d:%d <%s>\n", line, column, msg)
//  }
//  if strings.HasPrefix(msg, "report") { // TODO remove NewDiagnosticErrorListener and move warning to ReportAmbiguity etc. when getDecisionDescription is make public
//      cb.Warning = fmt.Sprintf("At %d:%d <%s>", line, column, msg)
//  } else {
//      cb.Err = fmt.Errorf("SyntaxError %d:%d <%s>", line, column, msg)
//  }
// }
// func (cb *DNAC_A_WalkerErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int,
//  exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportAmbiguity rec:%v dfs:%v start:%d stop:%d, exact:%v, ambigAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
//  }
// }
// func (cb *DNAC_A_WalkerErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportAttemptingFullContext rec:%v dfs:%v start:%d stop:%d, conflictingAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
//  }
// }
// func (cb *DNAC_A_WalkerErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportContextSensitivity rec:%v dfs:%v start:%d stop:%d, prediction:%v configs:%v\n", recognizer, dfa, startIndex, stopIndex, prediction, configs)
//  }
// }

//// antlr.ParseTreeListener implementation.
//// All required.

//func (s *DNAC_A_WalkerListener ) VisitTerminal(node  antlr.TerminalNode) {   }
//func (s *DNAC_A_WalkerListener ) VisitErrorNode(node antlr.ErrorNode)    {   }
//func (s *DNAC_A_WalkerListener ) EnterEveryRule(ctx antlr.ParserRuleContext) {  }
//func (s *DNAC_A_WalkerListener ) ExitEveryRule(ctx antlr.ParserRuleContext) {  }

//// Only implemented as needed.

//func (s *DNAC_A_WalkerListener) EnterAdl(ctx dnaca.*dnaca.AdlContext) {}
//func (s *DNAC_A_WalkerListener) ExitAdl(ctx dnaca.*dnaca.AdlContext) {}

//func (s *DNAC_A_WalkerListener) EnterNameNode(ctx dnaca.*dnaca.NameNodeContext) {}
//func (s *DNAC_A_WalkerListener) ExitNameNode(ctx dnaca.*dnaca.NameNodeContext) {}

//func (s *DNAC_A_WalkerListener) EnterNameRule(ctx dnaca.*dnaca.NameRuleContext) {}
//func (s *DNAC_A_WalkerListener) ExitNameRule(ctx dnaca.*dnaca.NameRuleContext) {}

//func (s *DNAC_A_WalkerListener) EnterTypeNode(ctx dnaca.*dnaca.TypeNodeContext) {}
//func (s *DNAC_A_WalkerListener) ExitTypeNode(ctx dnaca.*dnaca.TypeNodeContext) {}

//func (s *DNAC_A_WalkerListener) EnterExnotationNode(ctx dnaca.*dnaca.ExnotationNodeContext) {}
//func (s *DNAC_A_WalkerListener) ExitExnotationNode(ctx dnaca.*dnaca.ExnotationNodeContext) {}

//func (s *DNAC_A_WalkerListener) EnterNameBodyNode(ctx dnaca.*dnaca.NameBodyNodeContext) {}
//func (s *DNAC_A_WalkerListener) ExitNameBodyNode(ctx dnaca.*dnaca.NameBodyNodeContext) {}

//func (s *DNAC_A_WalkerListener) EnterAnnotation(ctx dnaca.*dnaca.AnnotationContext) {}
//func (s *DNAC_A_WalkerListener) ExitAnnotation(ctx dnaca.*dnaca.AnnotationContext) {}

//func (s *DNAC_A_WalkerListener) EnterTypeExpr_(ctx dnaca.*dnaca.TypeExpr_Context) {}
//func (s *DNAC_A_WalkerListener) ExitTypeExpr_(ctx dnaca.*dnaca.TypeExpr_Context) {}

//func (s *DNAC_A_WalkerListener) EnterTypeParams(ctx dnaca.*dnaca.TypeParamsContext) {}
//func (s *DNAC_A_WalkerListener) ExitTypeParams(ctx dnaca.*dnaca.TypeParamsContext) {}

//func (s *DNAC_A_WalkerListener) EnterJsonStr(ctx dnaca.*dnaca.JsonStrContext) {}
//func (s *DNAC_A_WalkerListener) ExitJsonStr(ctx dnaca.*dnaca.JsonStrContext) {}

//func (s *DNAC_A_WalkerListener) EnterJsonBool(ctx dnaca.*dnaca.JsonBoolContext) {}
//func (s *DNAC_A_WalkerListener) ExitJsonBool(ctx dnaca.*dnaca.JsonBoolContext) {}

//func (s *DNAC_A_WalkerListener) EnterJsonNull(ctx dnaca.*dnaca.JsonNullContext) {}
//func (s *DNAC_A_WalkerListener) ExitJsonNull(ctx dnaca.*dnaca.JsonNullContext) {}

//func (s *DNAC_A_WalkerListener) EnterJsonInt(ctx dnaca.*dnaca.JsonIntContext) {}
//func (s *DNAC_A_WalkerListener) ExitJsonInt(ctx dnaca.*dnaca.JsonIntContext) {}

//func (s *DNAC_A_WalkerListener) EnterJsonFloat(ctx dnaca.*dnaca.JsonFloatContext) {}
//func (s *DNAC_A_WalkerListener) ExitJsonFloat(ctx dnaca.*dnaca.JsonFloatContext) {}

//func (s *DNAC_A_WalkerListener) EnterJsonArray(ctx dnaca.*dnaca.JsonArrayContext) {}
//func (s *DNAC_A_WalkerListener) ExitJsonArray(ctx dnaca.*dnaca.JsonArrayContext) {}

//func (s *DNAC_A_WalkerListener) EnterJsonObj(ctx dnaca.*dnaca.JsonObjContext) {}
//func (s *DNAC_A_WalkerListener) ExitJsonObj(ctx dnaca.*dnaca.JsonObjContext) {}
