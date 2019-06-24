// Generated from DNAC_A_Walker.g4 by ANTLR 4.7.

package dnaca // DNAC_A_Walker
import "github.com/wxio/goantlr"

// DNAC_A_WalkerListener is a complete listener for a parse tree produced by DNAC_A_Walker.
type DNAC_A_WalkerListener interface {
	antlr.ParseTreeListener

	AdlEntryListener
	AdlExitListener

	AnnotationEntryListener
	AnnotationExitListener

	TypeExpr_EntryListener
	TypeExpr_ExitListener

	NameNodeEntryListener
	NameNodeExitListener

	NameRuleEntryListener
	NameRuleExitListener

	TypeNodeEntryListener
	TypeNodeExitListener

	ExnotationNodeEntryListener
	ExnotationNodeExitListener

	NameBodyNodeEntryListener
	NameBodyNodeExitListener

	TypeParamsEntryListener
	TypeParamsExitListener

	JsonStrEntryListener
	JsonStrExitListener

	JsonBoolEntryListener
	JsonBoolExitListener

	JsonNullEntryListener
	JsonNullExitListener

	JsonIntEntryListener
	JsonIntExitListener

	JsonFloatEntryListener
	JsonFloatExitListener

	JsonArrayEntryListener
	JsonArrayExitListener

	JsonObjEntryListener
	JsonObjExitListener
}

//
// Rules with unnamed alternatives
//
type AdlEntryListener interface {
	EnterAdl(c *AdlContext)
}
type AdlExitListener interface {
	ExitAdl(c *AdlContext)
}

type AnnotationEntryListener interface {
	EnterAnnotation(c *AnnotationContext)
}
type AnnotationExitListener interface {
	ExitAnnotation(c *AnnotationContext)
}

type TypeExpr_EntryListener interface {
	EnterTypeExpr_(c *TypeExpr_Context)
}
type TypeExpr_ExitListener interface {
	ExitTypeExpr_(c *TypeExpr_Context)
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
type ExnotationNodeEntryListener interface {
	EnterExnotationNode(c *ExnotationNodeContext)
}
type ExnotationNodeExitListener interface {
	ExitExnotationNode(c *ExnotationNodeContext)
}

// From Rule 'tld'
type NameBodyNodeEntryListener interface {
	EnterNameBodyNode(c *NameBodyNodeContext)
}
type NameBodyNodeExitListener interface {
	ExitNameBodyNode(c *NameBodyNodeContext)
}

// From Rule 'typeExprElem_'
type TypeParamsEntryListener interface {
	EnterTypeParams(c *TypeParamsContext)
}
type TypeParamsExitListener interface {
	ExitTypeParams(c *TypeParamsContext)
}

// From Rule 'jsonVal'
type JsonStrEntryListener interface {
	EnterJsonStr(c *JsonStrContext)
}
type JsonStrExitListener interface {
	ExitJsonStr(c *JsonStrContext)
}

// From Rule 'jsonVal'
type JsonBoolEntryListener interface {
	EnterJsonBool(c *JsonBoolContext)
}
type JsonBoolExitListener interface {
	ExitJsonBool(c *JsonBoolContext)
}

// From Rule 'jsonVal'
type JsonNullEntryListener interface {
	EnterJsonNull(c *JsonNullContext)
}
type JsonNullExitListener interface {
	ExitJsonNull(c *JsonNullContext)
}

// From Rule 'jsonVal'
type JsonIntEntryListener interface {
	EnterJsonInt(c *JsonIntContext)
}
type JsonIntExitListener interface {
	ExitJsonInt(c *JsonIntContext)
}

// From Rule 'jsonVal'
type JsonFloatEntryListener interface {
	EnterJsonFloat(c *JsonFloatContext)
}
type JsonFloatExitListener interface {
	ExitJsonFloat(c *JsonFloatContext)
}

// From Rule 'jsonVal'
type JsonArrayEntryListener interface {
	EnterJsonArray(c *JsonArrayContext)
}
type JsonArrayExitListener interface {
	ExitJsonArray(c *JsonArrayContext)
}

// From Rule 'jsonVal'
type JsonObjEntryListener interface {
	EnterJsonObj(c *JsonObjContext)
}
type JsonObjExitListener interface {
	ExitJsonObj(c *JsonObjContext)
}
