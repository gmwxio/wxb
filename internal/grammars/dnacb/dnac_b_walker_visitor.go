// Code generated from DNAC_B_Walker.g4 by ANTLR 4.7.1. DO NOT EDIT.

package dnacb // DNAC_B_Walker
import "github.com/wxio/goantlr"

// A complete Visitor for a parse tree produced by DNAC_B_Walker.
type DNAC_B_WalkerVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DNAC_B_Walker#dnac.
	VisitDnac(ctx *DnacContext) interface{}

	// Visit a parse tree produced by DNAC_B_Walker#NameNode.
	VisitNameNode(ctx *NameNodeContext) interface{}

	// Visit a parse tree produced by DNAC_B_Walker#TypeNode.
	VisitTypeNode(ctx *TypeNodeContext) interface{}

	// Visit a parse tree produced by DNAC_B_Walker#NameBodyNode.
	VisitNameBodyNode(ctx *NameBodyNodeContext) interface{}

	// Visit a parse tree produced by DNAC_B_Walker#ExnotationNode.
	VisitExnotationNode(ctx *ExnotationNodeContext) interface{}
}
