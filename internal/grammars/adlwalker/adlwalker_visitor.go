// Code generated from ADLWalker.g4 by ANTLR 4.7.1. DO NOT EDIT.

package walker // ADLWalker
import "github.com/wxio/goantlr"

// A complete Visitor for a parse tree produced by ADLWalker.
type ADLWalkerVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ADLWalker#adl.
	VisitAdl(ctx *AdlContext) interface{}

	// Visit a parse tree produced by ADLWalker#module.
	VisitModule(ctx *ModuleContext) interface{}

	// Visit a parse tree produced by ADLWalker#Struct.
	VisitStruct(ctx *StructContext) interface{}

	// Visit a parse tree produced by ADLWalker#Union.
	VisitUnion(ctx *UnionContext) interface{}

	// Visit a parse tree produced by ADLWalker#Type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by ADLWalker#Newtype.
	VisitNewtype(ctx *NewtypeContext) interface{}

	// Visit a parse tree produced by ADLWalker#ModAnno.
	VisitModAnno(ctx *ModAnnoContext) interface{}

	// Visit a parse tree produced by ADLWalker#DeclAnno.
	VisitDeclAnno(ctx *DeclAnnoContext) interface{}

	// Visit a parse tree produced by ADLWalker#FieldAnno.
	VisitFieldAnno(ctx *FieldAnnoContext) interface{}

	// Visit a parse tree produced by ADLWalker#TypeParamError.
	VisitTypeParamError(ctx *TypeParamErrorContext) interface{}

	// Visit a parse tree produced by ADLWalker#Field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by ADLWalker#annotation.
	VisitAnnotation(ctx *AnnotationContext) interface{}

	// Visit a parse tree produced by ADLWalker#typeExpr_.
	VisitTypeExpr_(ctx *TypeExpr_Context) interface{}

	// Visit a parse tree produced by ADLWalker#TypeParams.
	VisitTypeParams(ctx *TypeParamsContext) interface{}

	// Visit a parse tree produced by ADLWalker#JsonStr.
	VisitJsonStr(ctx *JsonStrContext) interface{}

	// Visit a parse tree produced by ADLWalker#JsonBool.
	VisitJsonBool(ctx *JsonBoolContext) interface{}

	// Visit a parse tree produced by ADLWalker#JsonNull.
	VisitJsonNull(ctx *JsonNullContext) interface{}

	// Visit a parse tree produced by ADLWalker#JsonInt.
	VisitJsonInt(ctx *JsonIntContext) interface{}

	// Visit a parse tree produced by ADLWalker#JsonFloat.
	VisitJsonFloat(ctx *JsonFloatContext) interface{}

	// Visit a parse tree produced by ADLWalker#JsonArray.
	VisitJsonArray(ctx *JsonArrayContext) interface{}

	// Visit a parse tree produced by ADLWalker#JsonObj.
	VisitJsonObj(ctx *JsonObjContext) interface{}
}
