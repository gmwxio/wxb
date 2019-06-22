// Code generated from DNAC_B_Walker.g4 by ANTLR 4.7.1. DO NOT EDIT.

package dnacb // DNAC_B_Walker
import "github.com/wxio/goantlr"

type BaseDNAC_B_WalkerVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDNAC_B_WalkerVisitor) VisitDnac(ctx *DnacContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDNAC_B_WalkerVisitor) VisitNameNode(ctx *NameNodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDNAC_B_WalkerVisitor) VisitTypeNode(ctx *TypeNodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDNAC_B_WalkerVisitor) VisitNameBodyNode(ctx *NameBodyNodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDNAC_B_WalkerVisitor) VisitExnotationNode(ctx *ExnotationNodeContext) interface{} {
	return v.VisitChildren(ctx)
}
