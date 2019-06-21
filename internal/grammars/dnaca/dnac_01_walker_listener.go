// Code generated from DNAC_01_Walker.g4 by ANTLR 4.7.1. DO NOT EDIT.

package dnaca // DNAC_01_Walker
import "github.com/wxio/goantlr"

// DNAC_01_WalkerListener is a complete listener for a parse tree produced by DNAC_01_Walker.
type DNAC_01_WalkerListener interface {
	antlr.ParseTreeListener

	// EnterAdl is called when entering the adl production.
	EnterAdl(c *AdlContext)

	// EnterNameNode is called when entering the NameNode production.
	EnterNameNode(c *NameNodeContext)

	// EnterTypeNode is called when entering the TypeNode production.
	EnterTypeNode(c *TypeNodeContext)

	// EnterExnotationNode is called when entering the ExnotationNode production.
	EnterExnotationNode(c *ExnotationNodeContext)

	// EnterNameBodyNode is called when entering the NameBodyNode production.
	EnterNameBodyNode(c *NameBodyNodeContext)

	// EnterAnnotation is called when entering the annotation production.
	EnterAnnotation(c *AnnotationContext)

	// EnterTypeExpr_ is called when entering the typeExpr_ production.
	EnterTypeExpr_(c *TypeExpr_Context)

	// EnterTypeParams is called when entering the TypeParams production.
	EnterTypeParams(c *TypeParamsContext)

	// EnterJsonStr is called when entering the JsonStr production.
	EnterJsonStr(c *JsonStrContext)

	// EnterJsonBool is called when entering the JsonBool production.
	EnterJsonBool(c *JsonBoolContext)

	// EnterJsonNull is called when entering the JsonNull production.
	EnterJsonNull(c *JsonNullContext)

	// EnterJsonInt is called when entering the JsonInt production.
	EnterJsonInt(c *JsonIntContext)

	// EnterJsonFloat is called when entering the JsonFloat production.
	EnterJsonFloat(c *JsonFloatContext)

	// EnterJsonArray is called when entering the JsonArray production.
	EnterJsonArray(c *JsonArrayContext)

	// EnterJsonObj is called when entering the JsonObj production.
	EnterJsonObj(c *JsonObjContext)

	// ExitAdl is called when exiting the adl production.
	ExitAdl(c *AdlContext)

	// ExitNameNode is called when exiting the NameNode production.
	ExitNameNode(c *NameNodeContext)

	// ExitTypeNode is called when exiting the TypeNode production.
	ExitTypeNode(c *TypeNodeContext)

	// ExitExnotationNode is called when exiting the ExnotationNode production.
	ExitExnotationNode(c *ExnotationNodeContext)

	// ExitNameBodyNode is called when exiting the NameBodyNode production.
	ExitNameBodyNode(c *NameBodyNodeContext)

	// ExitAnnotation is called when exiting the annotation production.
	ExitAnnotation(c *AnnotationContext)

	// ExitTypeExpr_ is called when exiting the typeExpr_ production.
	ExitTypeExpr_(c *TypeExpr_Context)

	// ExitTypeParams is called when exiting the TypeParams production.
	ExitTypeParams(c *TypeParamsContext)

	// ExitJsonStr is called when exiting the JsonStr production.
	ExitJsonStr(c *JsonStrContext)

	// ExitJsonBool is called when exiting the JsonBool production.
	ExitJsonBool(c *JsonBoolContext)

	// ExitJsonNull is called when exiting the JsonNull production.
	ExitJsonNull(c *JsonNullContext)

	// ExitJsonInt is called when exiting the JsonInt production.
	ExitJsonInt(c *JsonIntContext)

	// ExitJsonFloat is called when exiting the JsonFloat production.
	ExitJsonFloat(c *JsonFloatContext)

	// ExitJsonArray is called when exiting the JsonArray production.
	ExitJsonArray(c *JsonArrayContext)

	// ExitJsonObj is called when exiting the JsonObj production.
	ExitJsonObj(c *JsonObjContext)
}
