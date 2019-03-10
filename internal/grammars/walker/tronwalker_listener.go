// Code generated from TronWalker.g4 by ANTLR 4.7.1. DO NOT EDIT.

package walker // TronWalker
import "github.com/wxio/goantlr"

// TronWalkerListener is a complete listener for a parse tree produced by TronWalker.
type TronWalkerListener interface {
	antlr.ParseTreeListener

	// EnterProto is called when entering the proto production.
	EnterProto(c *ProtoContext)

	// EnterSyntax is called when entering the syntax production.
	EnterSyntax(c *SyntaxContext)

	// EnterTld is called when entering the tld production.
	EnterTld(c *TldContext)

	// EnterMsgBody is called when entering the msgBody production.
	EnterMsgBody(c *MsgBodyContext)

	// EnterEnumBody is called when entering the enumBody production.
	EnterEnumBody(c *EnumBodyContext)

	// EnterSvcBody is called when entering the svcBody production.
	EnterSvcBody(c *SvcBodyContext)

	// EnterExtendBody is called when entering the extendBody production.
	EnterExtendBody(c *ExtendBodyContext)

	// ExitProto is called when exiting the proto production.
	ExitProto(c *ProtoContext)

	// ExitSyntax is called when exiting the syntax production.
	ExitSyntax(c *SyntaxContext)

	// ExitTld is called when exiting the tld production.
	ExitTld(c *TldContext)

	// ExitMsgBody is called when exiting the msgBody production.
	ExitMsgBody(c *MsgBodyContext)

	// ExitEnumBody is called when exiting the enumBody production.
	ExitEnumBody(c *EnumBodyContext)

	// ExitSvcBody is called when exiting the svcBody production.
	ExitSvcBody(c *SvcBodyContext)

	// ExitExtendBody is called when exiting the extendBody production.
	ExitExtendBody(c *ExtendBodyContext)
}
