// Code generated from ADLWalker.g4 by ANTLR 4.7.1. DO NOT EDIT.

package walker // ADLWalker
import "github.com/wxio/goantlr"

// BaseADLWalkerListener is a complete listener for a parse tree produced by ADLWalker.
type BaseADLWalkerListener struct{}

var _ ADLWalkerListener = &BaseADLWalkerListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseADLWalkerListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseADLWalkerListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseADLWalkerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseADLWalkerListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterAdl is called when production adl is entered.
func (s *BaseADLWalkerListener) EnterAdl(ctx *AdlContext) {}

// ExitAdl is called when production adl is exited.
func (s *BaseADLWalkerListener) ExitAdl(ctx *AdlContext) {}

// EnterModule is called when production module is entered.
func (s *BaseADLWalkerListener) EnterModule(ctx *ModuleContext) {}

// ExitModule is called when production module is exited.
func (s *BaseADLWalkerListener) ExitModule(ctx *ModuleContext) {}

// EnterPkgNode is called when production PkgNode is entered.
func (s *BaseADLWalkerListener) EnterPkgNode(ctx *PkgNodeContext) {}

// ExitPkgNode is called when production PkgNode is exited.
func (s *BaseADLWalkerListener) ExitPkgNode(ctx *PkgNodeContext) {}

// EnterImportNode is called when production ImportNode is entered.
func (s *BaseADLWalkerListener) EnterImportNode(ctx *ImportNodeContext) {}

// ExitImportNode is called when production ImportNode is exited.
func (s *BaseADLWalkerListener) ExitImportNode(ctx *ImportNodeContext) {}

// EnterMsgNode is called when production MsgNode is entered.
func (s *BaseADLWalkerListener) EnterMsgNode(ctx *MsgNodeContext) {}

// ExitMsgNode is called when production MsgNode is exited.
func (s *BaseADLWalkerListener) ExitMsgNode(ctx *MsgNodeContext) {}

// EnterExtNode is called when production ExtNode is entered.
func (s *BaseADLWalkerListener) EnterExtNode(ctx *ExtNodeContext) {}

// ExitExtNode is called when production ExtNode is exited.
func (s *BaseADLWalkerListener) ExitExtNode(ctx *ExtNodeContext) {}

// EnterDeclAnnoNode is called when production DeclAnnoNode is entered.
func (s *BaseADLWalkerListener) EnterDeclAnnoNode(ctx *DeclAnnoNodeContext) {}

// ExitDeclAnnoNode is called when production DeclAnnoNode is exited.
func (s *BaseADLWalkerListener) ExitDeclAnnoNode(ctx *DeclAnnoNodeContext) {}

// EnterFieldAnnoNode is called when production FieldAnnoNode is entered.
func (s *BaseADLWalkerListener) EnterFieldAnnoNode(ctx *FieldAnnoNodeContext) {}

// ExitFieldAnnoNode is called when production FieldAnnoNode is exited.
func (s *BaseADLWalkerListener) ExitFieldAnnoNode(ctx *FieldAnnoNodeContext) {}

// EnterNameBody is called when production nameBody is entered.
func (s *BaseADLWalkerListener) EnterNameBody(ctx *NameBodyContext) {}

// ExitNameBody is called when production nameBody is exited.
func (s *BaseADLWalkerListener) ExitNameBody(ctx *NameBodyContext) {}

// EnterJsonVal is called when production jsonVal is entered.
func (s *BaseADLWalkerListener) EnterJsonVal(ctx *JsonValContext) {}

// ExitJsonVal is called when production jsonVal is exited.
func (s *BaseADLWalkerListener) ExitJsonVal(ctx *JsonValContext) {}
