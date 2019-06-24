// Generated from DNAC_B_Walker.g4 by ANTLR 4.7.

package dnacb // DNAC_B_Walker
//import "github.com/wxio/goantlr"
//import "generated code if in another package"

//// Commented out basic implementation for your convenience.

//func Example(s string) {
//  // Setup
//  input := antlr.NewInputStream(s)
//  lexer := dnacb.NewDNAC_B_WalkerLexer(input)
//  stream := antlr.NewCommonTokenStream(lexer, 0)
//  p := dnacb.NewDNAC_B_WalkerParser(stream)

//  // Antlr error listener - turns reports (ambiguity etc) into syntax errors
//  p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

//  // Custom error listener, register before the parse
//  el := &DNAC_B_WalkerErrorListener{}
//  p.AddErrorListener(el)

//  // Parse - start rule
//  tree := p.Start()

//  // Antlr provided parse tree representation
//  sexpr := antlr.TreesStringTree(tree, nil, p)
//  fmt.Printf("%s\n", sexpr)

//  // Custom listener
//  l := &DNAC_B_WalkerListener{}
//  antlr.ParseTreeWalkerDefault.Walk(l, tree)

//  // Custom visitor
//  v := &DNAC_B_WalkerVisitor{}
//  tree.Accept(v)
// }

//// implemented all listeners methods
//var _ dnacb.DNAC_B_WalkerListener = &DNAC_B_WalkerListener{}
//// implemented specific
//var _ dnacb.DnacEntryListener = &DNAC_B_WalkerListener{}
//var _ dnacb.DnacExitListener = &DNAC_B_WalkerListener{}

//var _ dnacb.NameNodeEntryListener = &DNAC_B_WalkerListener{}
//var _ dnacb.NameNodeExitListener = &DNAC_B_WalkerListener{}

//var _ dnacb.NameRuleEntryListener = &DNAC_B_WalkerListener{}
//var _ dnacb.NameRuleExitListener = &DNAC_B_WalkerListener{}

//var _ dnacb.TypeNodeEntryListener = &DNAC_B_WalkerListener{}
//var _ dnacb.TypeNodeExitListener = &DNAC_B_WalkerListener{}

//var _ dnacb.NameBodyNodeEntryListener = &DNAC_B_WalkerListener{}
//var _ dnacb.NameBodyNodeExitListener = &DNAC_B_WalkerListener{}

//var _ dnacb.ExnotationNodeEntryListener = &DNAC_B_WalkerListener{}
//var _ dnacb.ExnotationNodeExitListener = &DNAC_B_WalkerListener{}

//type DNAC_B_WalkerListener struct {
//}

//type DNAC_B_WalkerErrorListener struct {
//    Warning string
//    Err     error
//    Debug   bool
//}

// func (cb *DNAC_B_WalkerErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
//  if cb.Debug {
//      fmt.Printf("SyntaxError %d:%d <%s>\n", line, column, msg)
//  }
//  if strings.HasPrefix(msg, "report") { // TODO remove NewDiagnosticErrorListener and move warning to ReportAmbiguity etc. when getDecisionDescription is make public
//      cb.Warning = fmt.Sprintf("At %d:%d <%s>", line, column, msg)
//  } else {
//      cb.Err = fmt.Errorf("SyntaxError %d:%d <%s>", line, column, msg)
//  }
// }
// func (cb *DNAC_B_WalkerErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int,
//  exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportAmbiguity rec:%v dfs:%v start:%d stop:%d, exact:%v, ambigAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
//  }
// }
// func (cb *DNAC_B_WalkerErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportAttemptingFullContext rec:%v dfs:%v start:%d stop:%d, conflictingAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
//  }
// }
// func (cb *DNAC_B_WalkerErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportContextSensitivity rec:%v dfs:%v start:%d stop:%d, prediction:%v configs:%v\n", recognizer, dfa, startIndex, stopIndex, prediction, configs)
//  }
// }

//// antlr.ParseTreeListener implementation.
//// All required.

//func (s *DNAC_B_WalkerListener ) VisitTerminal(node  antlr.TerminalNode) {   }
//func (s *DNAC_B_WalkerListener ) VisitErrorNode(node antlr.ErrorNode)    {   }
//func (s *DNAC_B_WalkerListener ) EnterEveryRule(ctx antlr.ParserRuleContext) {  }
//func (s *DNAC_B_WalkerListener ) ExitEveryRule(ctx antlr.ParserRuleContext) {  }

//// Only implemented as needed.

//func (s *DNAC_B_WalkerListener) EnterDnac(ctx dnacb.*dnacb.DnacContext) {}
//func (s *DNAC_B_WalkerListener) ExitDnac(ctx dnacb.*dnacb.DnacContext) {}

//func (s *DNAC_B_WalkerListener) EnterNameNode(ctx dnacb.*dnacb.NameNodeContext) {}
//func (s *DNAC_B_WalkerListener) ExitNameNode(ctx dnacb.*dnacb.NameNodeContext) {}

//func (s *DNAC_B_WalkerListener) EnterNameRule(ctx dnacb.*dnacb.NameRuleContext) {}
//func (s *DNAC_B_WalkerListener) ExitNameRule(ctx dnacb.*dnacb.NameRuleContext) {}

//func (s *DNAC_B_WalkerListener) EnterTypeNode(ctx dnacb.*dnacb.TypeNodeContext) {}
//func (s *DNAC_B_WalkerListener) ExitTypeNode(ctx dnacb.*dnacb.TypeNodeContext) {}

//func (s *DNAC_B_WalkerListener) EnterNameBodyNode(ctx dnacb.*dnacb.NameBodyNodeContext) {}
//func (s *DNAC_B_WalkerListener) ExitNameBodyNode(ctx dnacb.*dnacb.NameBodyNodeContext) {}

//func (s *DNAC_B_WalkerListener) EnterExnotationNode(ctx dnacb.*dnacb.ExnotationNodeContext) {}
//func (s *DNAC_B_WalkerListener) ExitExnotationNode(ctx dnacb.*dnacb.ExnotationNodeContext) {}
