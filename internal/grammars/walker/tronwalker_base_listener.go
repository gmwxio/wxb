// Code generated from TronWalker.g4 by ANTLR 4.7.1. DO NOT EDIT.

package walker // TronWalker
import "github.com/wxio/goantlr"

// BaseTronWalkerListener is a complete listener for a parse tree produced by TronWalker.
type BaseTronWalkerListener struct{}

var _ TronWalkerListener = &BaseTronWalkerListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTronWalkerListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTronWalkerListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTronWalkerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTronWalkerListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProto is called when production proto is entered.
func (s *BaseTronWalkerListener) EnterProto(ctx *ProtoContext) {}

// ExitProto is called when production proto is exited.
func (s *BaseTronWalkerListener) ExitProto(ctx *ProtoContext) {}

// EnterSyntaxNode is called when production SyntaxNode is entered.
func (s *BaseTronWalkerListener) EnterSyntaxNode(ctx *SyntaxNodeContext) {}

// ExitSyntaxNode is called when production SyntaxNode is exited.
func (s *BaseTronWalkerListener) ExitSyntaxNode(ctx *SyntaxNodeContext) {}

// EnterPkgNode is called when production PkgNode is entered.
func (s *BaseTronWalkerListener) EnterPkgNode(ctx *PkgNodeContext) {}

// ExitPkgNode is called when production PkgNode is exited.
func (s *BaseTronWalkerListener) ExitPkgNode(ctx *PkgNodeContext) {}

// EnterImportNode is called when production ImportNode is entered.
func (s *BaseTronWalkerListener) EnterImportNode(ctx *ImportNodeContext) {}

// ExitImportNode is called when production ImportNode is exited.
func (s *BaseTronWalkerListener) ExitImportNode(ctx *ImportNodeContext) {}

// EnterMsgNode is called when production MsgNode is entered.
func (s *BaseTronWalkerListener) EnterMsgNode(ctx *MsgNodeContext) {}

// ExitMsgNode is called when production MsgNode is exited.
func (s *BaseTronWalkerListener) ExitMsgNode(ctx *MsgNodeContext) {}

// EnterEnumNode is called when production EnumNode is entered.
func (s *BaseTronWalkerListener) EnterEnumNode(ctx *EnumNodeContext) {}

// ExitEnumNode is called when production EnumNode is exited.
func (s *BaseTronWalkerListener) ExitEnumNode(ctx *EnumNodeContext) {}

// EnterSvcNode is called when production SvcNode is entered.
func (s *BaseTronWalkerListener) EnterSvcNode(ctx *SvcNodeContext) {}

// ExitSvcNode is called when production SvcNode is exited.
func (s *BaseTronWalkerListener) ExitSvcNode(ctx *SvcNodeContext) {}

// EnterExtNode is called when production ExtNode is entered.
func (s *BaseTronWalkerListener) EnterExtNode(ctx *ExtNodeContext) {}

// ExitExtNode is called when production ExtNode is exited.
func (s *BaseTronWalkerListener) ExitExtNode(ctx *ExtNodeContext) {}

// EnterMsgOptNode is called when production MsgOptNode is entered.
func (s *BaseTronWalkerListener) EnterMsgOptNode(ctx *MsgOptNodeContext) {}

// ExitMsgOptNode is called when production MsgOptNode is exited.
func (s *BaseTronWalkerListener) ExitMsgOptNode(ctx *MsgOptNodeContext) {}

// EnterMsgFldNode is called when production MsgFldNode is entered.
func (s *BaseTronWalkerListener) EnterMsgFldNode(ctx *MsgFldNodeContext) {}

// ExitMsgFldNode is called when production MsgFldNode is exited.
func (s *BaseTronWalkerListener) ExitMsgFldNode(ctx *MsgFldNodeContext) {}

// EnterMsgMsgNode is called when production MsgMsgNode is entered.
func (s *BaseTronWalkerListener) EnterMsgMsgNode(ctx *MsgMsgNodeContext) {}

// ExitMsgMsgNode is called when production MsgMsgNode is exited.
func (s *BaseTronWalkerListener) ExitMsgMsgNode(ctx *MsgMsgNodeContext) {}

// EnterMsgEnumNode is called when production MsgEnumNode is entered.
func (s *BaseTronWalkerListener) EnterMsgEnumNode(ctx *MsgEnumNodeContext) {}

// ExitMsgEnumNode is called when production MsgEnumNode is exited.
func (s *BaseTronWalkerListener) ExitMsgEnumNode(ctx *MsgEnumNodeContext) {}

// EnterMsgOneofNode is called when production MsgOneofNode is entered.
func (s *BaseTronWalkerListener) EnterMsgOneofNode(ctx *MsgOneofNodeContext) {}

// ExitMsgOneofNode is called when production MsgOneofNode is exited.
func (s *BaseTronWalkerListener) ExitMsgOneofNode(ctx *MsgOneofNodeContext) {}

// EnterOneofOptNode is called when production OneofOptNode is entered.
func (s *BaseTronWalkerListener) EnterOneofOptNode(ctx *OneofOptNodeContext) {}

// ExitOneofOptNode is called when production OneofOptNode is exited.
func (s *BaseTronWalkerListener) ExitOneofOptNode(ctx *OneofOptNodeContext) {}

// EnterOneofFldNode is called when production OneofFldNode is entered.
func (s *BaseTronWalkerListener) EnterOneofFldNode(ctx *OneofFldNodeContext) {}

// ExitOneofFldNode is called when production OneofFldNode is exited.
func (s *BaseTronWalkerListener) ExitOneofFldNode(ctx *OneofFldNodeContext) {}

// EnterEnumOptNode is called when production EnumOptNode is entered.
func (s *BaseTronWalkerListener) EnterEnumOptNode(ctx *EnumOptNodeContext) {}

// ExitEnumOptNode is called when production EnumOptNode is exited.
func (s *BaseTronWalkerListener) ExitEnumOptNode(ctx *EnumOptNodeContext) {}

// EnterEnumValNode is called when production EnumValNode is entered.
func (s *BaseTronWalkerListener) EnterEnumValNode(ctx *EnumValNodeContext) {}

// ExitEnumValNode is called when production EnumValNode is exited.
func (s *BaseTronWalkerListener) ExitEnumValNode(ctx *EnumValNodeContext) {}

// EnterSvcOptNode is called when production SvcOptNode is entered.
func (s *BaseTronWalkerListener) EnterSvcOptNode(ctx *SvcOptNodeContext) {}

// ExitSvcOptNode is called when production SvcOptNode is exited.
func (s *BaseTronWalkerListener) ExitSvcOptNode(ctx *SvcOptNodeContext) {}

// EnterSvcRpcNode is called when production SvcRpcNode is entered.
func (s *BaseTronWalkerListener) EnterSvcRpcNode(ctx *SvcRpcNodeContext) {}

// ExitSvcRpcNode is called when production SvcRpcNode is exited.
func (s *BaseTronWalkerListener) ExitSvcRpcNode(ctx *SvcRpcNodeContext) {}

// EnterExtOptNode is called when production ExtOptNode is entered.
func (s *BaseTronWalkerListener) EnterExtOptNode(ctx *ExtOptNodeContext) {}

// ExitExtOptNode is called when production ExtOptNode is exited.
func (s *BaseTronWalkerListener) ExitExtOptNode(ctx *ExtOptNodeContext) {}

// EnterExtFldNode is called when production ExtFldNode is entered.
func (s *BaseTronWalkerListener) EnterExtFldNode(ctx *ExtFldNodeContext) {}

// ExitExtFldNode is called when production ExtFldNode is exited.
func (s *BaseTronWalkerListener) ExitExtFldNode(ctx *ExtFldNodeContext) {}
