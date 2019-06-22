// Code generated from DNAC_B_Walker.g4 by ANTLR 4.7.1. DO NOT EDIT.

package dnacb // DNAC_B_Walker
import "github.com/wxio/goantlr"

// DNAC_B_WalkerListener is a complete listener for a parse tree produced by DNAC_B_Walker.
type DNAC_B_WalkerListener interface {
	antlr.ParseTreeListener

	// EnterDnac is called when entering the dnac production.
	EnterDnac(c *DnacContext)

	// EnterNameNode is called when entering the NameNode production.
	EnterNameNode(c *NameNodeContext)

	// EnterTypeNode is called when entering the TypeNode production.
	EnterTypeNode(c *TypeNodeContext)

	// EnterNameBodyNode is called when entering the NameBodyNode production.
	EnterNameBodyNode(c *NameBodyNodeContext)

	// EnterExnotationNode is called when entering the ExnotationNode production.
	EnterExnotationNode(c *ExnotationNodeContext)

	// ExitDnac is called when exiting the dnac production.
	ExitDnac(c *DnacContext)

	// ExitNameNode is called when exiting the NameNode production.
	ExitNameNode(c *NameNodeContext)

	// ExitTypeNode is called when exiting the TypeNode production.
	ExitTypeNode(c *TypeNodeContext)

	// ExitNameBodyNode is called when exiting the NameBodyNode production.
	ExitNameBodyNode(c *NameBodyNodeContext)

	// ExitExnotationNode is called when exiting the ExnotationNode production.
	ExitExnotationNode(c *ExnotationNodeContext)
}
