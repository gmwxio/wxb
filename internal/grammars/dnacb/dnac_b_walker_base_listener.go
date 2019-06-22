// Code generated from DNAC_B_Walker.g4 by ANTLR 4.7.1. DO NOT EDIT.

package dnacb // DNAC_B_Walker
import "github.com/wxio/goantlr"

// BaseDNAC_B_WalkerListener is a complete listener for a parse tree produced by DNAC_B_Walker.
type BaseDNAC_B_WalkerListener struct{}

var _ DNAC_B_WalkerListener = &BaseDNAC_B_WalkerListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDNAC_B_WalkerListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDNAC_B_WalkerListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDNAC_B_WalkerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDNAC_B_WalkerListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDnac is called when production dnac is entered.
func (s *BaseDNAC_B_WalkerListener) EnterDnac(ctx *DnacContext) {}

// ExitDnac is called when production dnac is exited.
func (s *BaseDNAC_B_WalkerListener) ExitDnac(ctx *DnacContext) {}

// EnterNameNode is called when production NameNode is entered.
func (s *BaseDNAC_B_WalkerListener) EnterNameNode(ctx *NameNodeContext) {}

// ExitNameNode is called when production NameNode is exited.
func (s *BaseDNAC_B_WalkerListener) ExitNameNode(ctx *NameNodeContext) {}

// EnterTypeNode is called when production TypeNode is entered.
func (s *BaseDNAC_B_WalkerListener) EnterTypeNode(ctx *TypeNodeContext) {}

// ExitTypeNode is called when production TypeNode is exited.
func (s *BaseDNAC_B_WalkerListener) ExitTypeNode(ctx *TypeNodeContext) {}

// EnterNameBodyNode is called when production NameBodyNode is entered.
func (s *BaseDNAC_B_WalkerListener) EnterNameBodyNode(ctx *NameBodyNodeContext) {}

// ExitNameBodyNode is called when production NameBodyNode is exited.
func (s *BaseDNAC_B_WalkerListener) ExitNameBodyNode(ctx *NameBodyNodeContext) {}

// EnterExnotationNode is called when production ExnotationNode is entered.
func (s *BaseDNAC_B_WalkerListener) EnterExnotationNode(ctx *ExnotationNodeContext) {}

// ExitExnotationNode is called when production ExnotationNode is exited.
func (s *BaseDNAC_B_WalkerListener) ExitExnotationNode(ctx *ExnotationNodeContext) {}
