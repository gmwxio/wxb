// Code generated from ADLWalker.g4 by ANTLR 4.7.1. DO NOT EDIT.

package walker // ADLWalker
import "github.com/wxio/goantlr"

type BaseADLWalkerVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseADLWalkerVisitor) VisitAdl(ctx *AdlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADLWalkerVisitor) VisitModule(ctx *ModuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADLWalkerVisitor) VisitStruct(ctx *StructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADLWalkerVisitor) VisitUnion(ctx *UnionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADLWalkerVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADLWalkerVisitor) VisitNewtype(ctx *NewtypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADLWalkerVisitor) VisitModAnno(ctx *ModAnnoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADLWalkerVisitor) VisitDeclAnno(ctx *DeclAnnoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADLWalkerVisitor) VisitFieldAnno(ctx *FieldAnnoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADLWalkerVisitor) VisitTypeParamError(ctx *TypeParamErrorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADLWalkerVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADLWalkerVisitor) VisitAnnotation(ctx *AnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADLWalkerVisitor) VisitTypeExpr_(ctx *TypeExpr_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADLWalkerVisitor) VisitTypeParams(ctx *TypeParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADLWalkerVisitor) VisitJsonStr(ctx *JsonStrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADLWalkerVisitor) VisitJsonBool(ctx *JsonBoolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADLWalkerVisitor) VisitJsonNull(ctx *JsonNullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADLWalkerVisitor) VisitJsonInt(ctx *JsonIntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADLWalkerVisitor) VisitJsonFloat(ctx *JsonFloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADLWalkerVisitor) VisitJsonArray(ctx *JsonArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADLWalkerVisitor) VisitJsonObj(ctx *JsonObjContext) interface{} {
	return v.VisitChildren(ctx)
}
