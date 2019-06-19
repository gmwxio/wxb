// Code generated from ADLWalker.g4 by ANTLR 4.7.1. DO NOT EDIT.

package walker // ADLWalker
import "github.com/wxio/goantlr"

// ADLWalkerListener is a complete listener for a parse tree produced by ADLWalker.
type ADLWalkerListener interface {
	antlr.ParseTreeListener

	// EnterAdl is called when entering the adl production.
	EnterAdl(c *AdlContext)

	// EnterModule is called when entering the module production.
	EnterModule(c *ModuleContext)

	// EnterPkgNode is called when entering the PkgNode production.
	EnterPkgNode(c *PkgNodeContext)

	// EnterImportNode is called when entering the ImportNode production.
	EnterImportNode(c *ImportNodeContext)

	// EnterMsgNode is called when entering the MsgNode production.
	EnterMsgNode(c *MsgNodeContext)

	// EnterExtNode is called when entering the ExtNode production.
	EnterExtNode(c *ExtNodeContext)

	// EnterDeclAnnoNode is called when entering the DeclAnnoNode production.
	EnterDeclAnnoNode(c *DeclAnnoNodeContext)

	// EnterFieldAnnoNode is called when entering the FieldAnnoNode production.
	EnterFieldAnnoNode(c *FieldAnnoNodeContext)

	// EnterNameBody is called when entering the nameBody production.
	EnterNameBody(c *NameBodyContext)

	// EnterJsonVal is called when entering the jsonVal production.
	EnterJsonVal(c *JsonValContext)

	// ExitAdl is called when exiting the adl production.
	ExitAdl(c *AdlContext)

	// ExitModule is called when exiting the module production.
	ExitModule(c *ModuleContext)

	// ExitPkgNode is called when exiting the PkgNode production.
	ExitPkgNode(c *PkgNodeContext)

	// ExitImportNode is called when exiting the ImportNode production.
	ExitImportNode(c *ImportNodeContext)

	// ExitMsgNode is called when exiting the MsgNode production.
	ExitMsgNode(c *MsgNodeContext)

	// ExitExtNode is called when exiting the ExtNode production.
	ExitExtNode(c *ExtNodeContext)

	// ExitDeclAnnoNode is called when exiting the DeclAnnoNode production.
	ExitDeclAnnoNode(c *DeclAnnoNodeContext)

	// ExitFieldAnnoNode is called when exiting the FieldAnnoNode production.
	ExitFieldAnnoNode(c *FieldAnnoNodeContext)

	// ExitNameBody is called when exiting the nameBody production.
	ExitNameBody(c *NameBodyContext)

	// ExitJsonVal is called when exiting the jsonVal production.
	ExitJsonVal(c *JsonValContext)
}
