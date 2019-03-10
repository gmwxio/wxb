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

// EnterSyntax is called when production syntax is entered.
func (s *BaseTronWalkerListener) EnterSyntax(ctx *SyntaxContext) {}

// ExitSyntax is called when production syntax is exited.
func (s *BaseTronWalkerListener) ExitSyntax(ctx *SyntaxContext) {}

// EnterTld is called when production tld is entered.
func (s *BaseTronWalkerListener) EnterTld(ctx *TldContext) {}

// ExitTld is called when production tld is exited.
func (s *BaseTronWalkerListener) ExitTld(ctx *TldContext) {}

// EnterMsgBody is called when production msgBody is entered.
func (s *BaseTronWalkerListener) EnterMsgBody(ctx *MsgBodyContext) {}

// ExitMsgBody is called when production msgBody is exited.
func (s *BaseTronWalkerListener) ExitMsgBody(ctx *MsgBodyContext) {}

// EnterEnumBody is called when production enumBody is entered.
func (s *BaseTronWalkerListener) EnterEnumBody(ctx *EnumBodyContext) {}

// ExitEnumBody is called when production enumBody is exited.
func (s *BaseTronWalkerListener) ExitEnumBody(ctx *EnumBodyContext) {}

// EnterSvcBody is called when production svcBody is entered.
func (s *BaseTronWalkerListener) EnterSvcBody(ctx *SvcBodyContext) {}

// ExitSvcBody is called when production svcBody is exited.
func (s *BaseTronWalkerListener) ExitSvcBody(ctx *SvcBodyContext) {}

// EnterExtendBody is called when production extendBody is entered.
func (s *BaseTronWalkerListener) EnterExtendBody(ctx *ExtendBodyContext) {}

// ExitExtendBody is called when production extendBody is exited.
func (s *BaseTronWalkerListener) ExitExtendBody(ctx *ExtendBodyContext) {}
