// Generated from DNAC_B_Walker.g4 by ANTLR 4.7.

package dnacb // DNAC_B_Walker
import "github.com/wxio/goantlr"

// Struct of Handlers
type DNAC_B_WalkerHandlers struct {
	EnterEveryRule func(ctx antlr.RuleNode)
	ExitEveryRule  func(ctx antlr.RuleNode)

	Dnac           func(ctx IDnacContext, this *DNAC_B_WalkerHandlers, args ...interface{}) (result interface{})
	NameNode       func(ctx INameNodeContext, this *DNAC_B_WalkerHandlers, args ...interface{}) (result interface{})
	NameRule       func(ctx INameRuleContext, this *DNAC_B_WalkerHandlers, args ...interface{}) (result interface{})
	TypeNode       func(ctx ITypeNodeContext, this *DNAC_B_WalkerHandlers, args ...interface{}) (result interface{})
	NameBodyNode   func(ctx INameBodyNodeContext, this *DNAC_B_WalkerHandlers, args ...interface{}) (result interface{})
	ExnotationNode func(ctx IExnotationNodeContext, this *DNAC_B_WalkerHandlers, args ...interface{}) (result interface{})
}

// A complete Visitor for a parse tree produced by DNAC_B_Walker.
type DNAC_B_WalkerVisitor interface {
	antlr.ParseTreeVisitor
	DnacContextVisitor
	NameNodeContextVisitor
	NameRuleContextVisitor
	TypeNodeContextVisitor
	NameBodyNodeContextVisitor
	ExnotationNodeContextVisitor
}

type DnacContextVisitor interface {
	VisitDnac(ctx IDnacContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type NameNodeContextVisitor interface {
	VisitNameNode(ctx INameNodeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type NameRuleContextVisitor interface {
	VisitNameRule(ctx INameRuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type TypeNodeContextVisitor interface {
	VisitTypeNode(ctx ITypeNodeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type NameBodyNodeContextVisitor interface {
	VisitNameBodyNode(ctx INameBodyNodeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type ExnotationNodeContextVisitor interface {
	VisitExnotationNode(ctx IExnotationNodeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
