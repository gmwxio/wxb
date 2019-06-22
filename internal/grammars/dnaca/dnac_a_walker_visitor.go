// Code generated from DNAC_A_Walker.g4 by ANTLR 4.7.1. DO NOT EDIT.

package dnaca // DNAC_A_Walker
import "github.com/wxio/goantlr"

// A complete Visitor for a parse tree produced by DNAC_A_Walker.
type DNAC_A_WalkerVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DNAC_A_Walker#adl.
	VisitAdl(ctx *AdlContext) interface{}

	// Visit a parse tree produced by DNAC_A_Walker#NameNode.
	VisitNameNode(ctx *NameNodeContext) interface{}

	// Visit a parse tree produced by DNAC_A_Walker#TypeNode.
	VisitTypeNode(ctx *TypeNodeContext) interface{}

	// Visit a parse tree produced by DNAC_A_Walker#ExnotationNode.
	VisitExnotationNode(ctx *ExnotationNodeContext) interface{}

	// Visit a parse tree produced by DNAC_A_Walker#NameBodyNode.
	VisitNameBodyNode(ctx *NameBodyNodeContext) interface{}

	// Visit a parse tree produced by DNAC_A_Walker#annotation.
	VisitAnnotation(ctx *AnnotationContext) interface{}

	// Visit a parse tree produced by DNAC_A_Walker#typeExpr_.
	VisitTypeExpr_(ctx *TypeExpr_Context) interface{}

	// Visit a parse tree produced by DNAC_A_Walker#TypeParams.
	VisitTypeParams(ctx *TypeParamsContext) interface{}

	// Visit a parse tree produced by DNAC_A_Walker#JsonStr.
	VisitJsonStr(ctx *JsonStrContext) interface{}

	// Visit a parse tree produced by DNAC_A_Walker#JsonBool.
	VisitJsonBool(ctx *JsonBoolContext) interface{}

	// Visit a parse tree produced by DNAC_A_Walker#JsonNull.
	VisitJsonNull(ctx *JsonNullContext) interface{}

	// Visit a parse tree produced by DNAC_A_Walker#JsonInt.
	VisitJsonInt(ctx *JsonIntContext) interface{}

	// Visit a parse tree produced by DNAC_A_Walker#JsonFloat.
	VisitJsonFloat(ctx *JsonFloatContext) interface{}

	// Visit a parse tree produced by DNAC_A_Walker#JsonArray.
	VisitJsonArray(ctx *JsonArrayContext) interface{}

	// Visit a parse tree produced by DNAC_A_Walker#JsonObj.
	VisitJsonObj(ctx *JsonObjContext) interface{}
}
