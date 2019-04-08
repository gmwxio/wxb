// Code generated from TronWalker.g4 by ANTLR 4.7.1. DO NOT EDIT.

package walker // TronWalker
import "github.com/wxio/goantlr"

// TronWalkerListener is a complete listener for a parse tree produced by TronWalker.
type TronWalkerListener interface {
	antlr.ParseTreeListener

	// EnterProto is called when entering the proto production.
	EnterProto(c *ProtoContext)

	// EnterSyntaxNode is called when entering the SyntaxNode production.
	EnterSyntaxNode(c *SyntaxNodeContext)

	// EnterPkgNode is called when entering the PkgNode production.
	EnterPkgNode(c *PkgNodeContext)

	// EnterImportNode is called when entering the ImportNode production.
	EnterImportNode(c *ImportNodeContext)

	// EnterMsgNode is called when entering the MsgNode production.
	EnterMsgNode(c *MsgNodeContext)

	// EnterEnumNode is called when entering the EnumNode production.
	EnterEnumNode(c *EnumNodeContext)

	// EnterSvcNode is called when entering the SvcNode production.
	EnterSvcNode(c *SvcNodeContext)

	// EnterExtNode is called when entering the ExtNode production.
	EnterExtNode(c *ExtNodeContext)

	// EnterMsgOptNode is called when entering the MsgOptNode production.
	EnterMsgOptNode(c *MsgOptNodeContext)

	// EnterMsgFldNode is called when entering the MsgFldNode production.
	EnterMsgFldNode(c *MsgFldNodeContext)

	// EnterMsgMsgNode is called when entering the MsgMsgNode production.
	EnterMsgMsgNode(c *MsgMsgNodeContext)

	// EnterMsgEnumNode is called when entering the MsgEnumNode production.
	EnterMsgEnumNode(c *MsgEnumNodeContext)

	// EnterMsgOneofNode is called when entering the MsgOneofNode production.
	EnterMsgOneofNode(c *MsgOneofNodeContext)

	// EnterOneofOptNode is called when entering the OneofOptNode production.
	EnterOneofOptNode(c *OneofOptNodeContext)

	// EnterOneofFldNode is called when entering the OneofFldNode production.
	EnterOneofFldNode(c *OneofFldNodeContext)

	// EnterEnumOptNode is called when entering the EnumOptNode production.
	EnterEnumOptNode(c *EnumOptNodeContext)

	// EnterEnumValNode is called when entering the EnumValNode production.
	EnterEnumValNode(c *EnumValNodeContext)

	// EnterSvcOptNode is called when entering the SvcOptNode production.
	EnterSvcOptNode(c *SvcOptNodeContext)

	// EnterSvcRpcNode is called when entering the SvcRpcNode production.
	EnterSvcRpcNode(c *SvcRpcNodeContext)

	// EnterExtOptNode is called when entering the ExtOptNode production.
	EnterExtOptNode(c *ExtOptNodeContext)

	// EnterExtFldNode is called when entering the ExtFldNode production.
	EnterExtFldNode(c *ExtFldNodeContext)

	// ExitProto is called when exiting the proto production.
	ExitProto(c *ProtoContext)

	// ExitSyntaxNode is called when exiting the SyntaxNode production.
	ExitSyntaxNode(c *SyntaxNodeContext)

	// ExitPkgNode is called when exiting the PkgNode production.
	ExitPkgNode(c *PkgNodeContext)

	// ExitImportNode is called when exiting the ImportNode production.
	ExitImportNode(c *ImportNodeContext)

	// ExitMsgNode is called when exiting the MsgNode production.
	ExitMsgNode(c *MsgNodeContext)

	// ExitEnumNode is called when exiting the EnumNode production.
	ExitEnumNode(c *EnumNodeContext)

	// ExitSvcNode is called when exiting the SvcNode production.
	ExitSvcNode(c *SvcNodeContext)

	// ExitExtNode is called when exiting the ExtNode production.
	ExitExtNode(c *ExtNodeContext)

	// ExitMsgOptNode is called when exiting the MsgOptNode production.
	ExitMsgOptNode(c *MsgOptNodeContext)

	// ExitMsgFldNode is called when exiting the MsgFldNode production.
	ExitMsgFldNode(c *MsgFldNodeContext)

	// ExitMsgMsgNode is called when exiting the MsgMsgNode production.
	ExitMsgMsgNode(c *MsgMsgNodeContext)

	// ExitMsgEnumNode is called when exiting the MsgEnumNode production.
	ExitMsgEnumNode(c *MsgEnumNodeContext)

	// ExitMsgOneofNode is called when exiting the MsgOneofNode production.
	ExitMsgOneofNode(c *MsgOneofNodeContext)

	// ExitOneofOptNode is called when exiting the OneofOptNode production.
	ExitOneofOptNode(c *OneofOptNodeContext)

	// ExitOneofFldNode is called when exiting the OneofFldNode production.
	ExitOneofFldNode(c *OneofFldNodeContext)

	// ExitEnumOptNode is called when exiting the EnumOptNode production.
	ExitEnumOptNode(c *EnumOptNodeContext)

	// ExitEnumValNode is called when exiting the EnumValNode production.
	ExitEnumValNode(c *EnumValNodeContext)

	// ExitSvcOptNode is called when exiting the SvcOptNode production.
	ExitSvcOptNode(c *SvcOptNodeContext)

	// ExitSvcRpcNode is called when exiting the SvcRpcNode production.
	ExitSvcRpcNode(c *SvcRpcNodeContext)

	// ExitExtOptNode is called when exiting the ExtOptNode production.
	ExitExtOptNode(c *ExtOptNodeContext)

	// ExitExtFldNode is called when exiting the ExtFldNode production.
	ExitExtFldNode(c *ExtFldNodeContext)
}
