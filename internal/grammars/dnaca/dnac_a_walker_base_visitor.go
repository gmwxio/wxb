// Code generated from DNAC_A_Walker.g4 by ANTLR 4.7.1. DO NOT EDIT.

package dnaca // DNAC_A_Walker
import "github.com/wxio/goantlr"

type BaseDNAC_A_WalkerVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDNAC_A_WalkerVisitor) VisitAdl(ctx *AdlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDNAC_A_WalkerVisitor) VisitNameNode(ctx *NameNodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDNAC_A_WalkerVisitor) VisitTypeNode(ctx *TypeNodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDNAC_A_WalkerVisitor) VisitExnotationNode(ctx *ExnotationNodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDNAC_A_WalkerVisitor) VisitNameBodyNode(ctx *NameBodyNodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDNAC_A_WalkerVisitor) VisitAnnotation(ctx *AnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDNAC_A_WalkerVisitor) VisitTypeExpr_(ctx *TypeExpr_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDNAC_A_WalkerVisitor) VisitTypeParams(ctx *TypeParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDNAC_A_WalkerVisitor) VisitJsonStr(ctx *JsonStrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDNAC_A_WalkerVisitor) VisitJsonBool(ctx *JsonBoolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDNAC_A_WalkerVisitor) VisitJsonNull(ctx *JsonNullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDNAC_A_WalkerVisitor) VisitJsonInt(ctx *JsonIntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDNAC_A_WalkerVisitor) VisitJsonFloat(ctx *JsonFloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDNAC_A_WalkerVisitor) VisitJsonArray(ctx *JsonArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDNAC_A_WalkerVisitor) VisitJsonObj(ctx *JsonObjContext) interface{} {
	return v.VisitChildren(ctx)
}
