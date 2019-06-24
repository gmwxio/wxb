// Generated from DNAC_B_Walker.g4 by ANTLR 4.7.

package dnacb // DNAC_B_Walker
import "github.com/wxio/goantlr"

// DNAC_B_WalkerListener is a complete listener for a parse tree produced by DNAC_B_Walker.
type DNAC_B_WalkerListener interface {
	antlr.ParseTreeListener

	DnacEntryListener
	DnacExitListener

	NameNodeEntryListener
	NameNodeExitListener

	NameRuleEntryListener
	NameRuleExitListener

	TypeNodeEntryListener
	TypeNodeExitListener

	NameBodyNodeEntryListener
	NameBodyNodeExitListener

	ExnotationNodeEntryListener
	ExnotationNodeExitListener
}

//
// Rules with unnamed alternatives
//
type DnacEntryListener interface {
	EnterDnac(c *DnacContext)
}
type DnacExitListener interface {
	ExitDnac(c *DnacContext)
}

//
// Named alternatives
//
//
// From Rule 'name'
type NameNodeEntryListener interface {
	EnterNameNode(c *NameNodeContext)
}
type NameNodeExitListener interface {
	ExitNameNode(c *NameNodeContext)
}

// From Rule 'tld'
type NameRuleEntryListener interface {
	EnterNameRule(c *NameRuleContext)
}
type NameRuleExitListener interface {
	ExitNameRule(c *NameRuleContext)
}

// From Rule 'tld'
type TypeNodeEntryListener interface {
	EnterTypeNode(c *TypeNodeContext)
}
type TypeNodeExitListener interface {
	ExitTypeNode(c *TypeNodeContext)
}

// From Rule 'tld'
type NameBodyNodeEntryListener interface {
	EnterNameBodyNode(c *NameBodyNodeContext)
}
type NameBodyNodeExitListener interface {
	ExitNameBodyNode(c *NameBodyNodeContext)
}

// From Rule 'tld'
type ExnotationNodeEntryListener interface {
	EnterExnotationNode(c *ExnotationNodeContext)
}
type ExnotationNodeExitListener interface {
	ExitExnotationNode(c *ExnotationNodeContext)
}
