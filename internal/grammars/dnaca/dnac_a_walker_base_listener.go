// Code generated from DNAC_A_Walker.g4 by ANTLR 4.7.1. DO NOT EDIT.

package dnaca // DNAC_A_Walker
import "github.com/wxio/goantlr"

// BaseDNAC_A_WalkerListener is a complete listener for a parse tree produced by DNAC_A_Walker.
type BaseDNAC_A_WalkerListener struct{}

var _ DNAC_A_WalkerListener = &BaseDNAC_A_WalkerListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDNAC_A_WalkerListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDNAC_A_WalkerListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDNAC_A_WalkerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDNAC_A_WalkerListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterAdl is called when production adl is entered.
func (s *BaseDNAC_A_WalkerListener) EnterAdl(ctx *AdlContext) {}

// ExitAdl is called when production adl is exited.
func (s *BaseDNAC_A_WalkerListener) ExitAdl(ctx *AdlContext) {}

// EnterNameNode is called when production NameNode is entered.
func (s *BaseDNAC_A_WalkerListener) EnterNameNode(ctx *NameNodeContext) {}

// ExitNameNode is called when production NameNode is exited.
func (s *BaseDNAC_A_WalkerListener) ExitNameNode(ctx *NameNodeContext) {}

// EnterTypeNode is called when production TypeNode is entered.
func (s *BaseDNAC_A_WalkerListener) EnterTypeNode(ctx *TypeNodeContext) {}

// ExitTypeNode is called when production TypeNode is exited.
func (s *BaseDNAC_A_WalkerListener) ExitTypeNode(ctx *TypeNodeContext) {}

// EnterExnotationNode is called when production ExnotationNode is entered.
func (s *BaseDNAC_A_WalkerListener) EnterExnotationNode(ctx *ExnotationNodeContext) {}

// ExitExnotationNode is called when production ExnotationNode is exited.
func (s *BaseDNAC_A_WalkerListener) ExitExnotationNode(ctx *ExnotationNodeContext) {}

// EnterNameBodyNode is called when production NameBodyNode is entered.
func (s *BaseDNAC_A_WalkerListener) EnterNameBodyNode(ctx *NameBodyNodeContext) {}

// ExitNameBodyNode is called when production NameBodyNode is exited.
func (s *BaseDNAC_A_WalkerListener) ExitNameBodyNode(ctx *NameBodyNodeContext) {}

// EnterAnnotation is called when production annotation is entered.
func (s *BaseDNAC_A_WalkerListener) EnterAnnotation(ctx *AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *BaseDNAC_A_WalkerListener) ExitAnnotation(ctx *AnnotationContext) {}

// EnterTypeExpr_ is called when production typeExpr_ is entered.
func (s *BaseDNAC_A_WalkerListener) EnterTypeExpr_(ctx *TypeExpr_Context) {}

// ExitTypeExpr_ is called when production typeExpr_ is exited.
func (s *BaseDNAC_A_WalkerListener) ExitTypeExpr_(ctx *TypeExpr_Context) {}

// EnterTypeParams is called when production TypeParams is entered.
func (s *BaseDNAC_A_WalkerListener) EnterTypeParams(ctx *TypeParamsContext) {}

// ExitTypeParams is called when production TypeParams is exited.
func (s *BaseDNAC_A_WalkerListener) ExitTypeParams(ctx *TypeParamsContext) {}

// EnterJsonStr is called when production JsonStr is entered.
func (s *BaseDNAC_A_WalkerListener) EnterJsonStr(ctx *JsonStrContext) {}

// ExitJsonStr is called when production JsonStr is exited.
func (s *BaseDNAC_A_WalkerListener) ExitJsonStr(ctx *JsonStrContext) {}

// EnterJsonBool is called when production JsonBool is entered.
func (s *BaseDNAC_A_WalkerListener) EnterJsonBool(ctx *JsonBoolContext) {}

// ExitJsonBool is called when production JsonBool is exited.
func (s *BaseDNAC_A_WalkerListener) ExitJsonBool(ctx *JsonBoolContext) {}

// EnterJsonNull is called when production JsonNull is entered.
func (s *BaseDNAC_A_WalkerListener) EnterJsonNull(ctx *JsonNullContext) {}

// ExitJsonNull is called when production JsonNull is exited.
func (s *BaseDNAC_A_WalkerListener) ExitJsonNull(ctx *JsonNullContext) {}

// EnterJsonInt is called when production JsonInt is entered.
func (s *BaseDNAC_A_WalkerListener) EnterJsonInt(ctx *JsonIntContext) {}

// ExitJsonInt is called when production JsonInt is exited.
func (s *BaseDNAC_A_WalkerListener) ExitJsonInt(ctx *JsonIntContext) {}

// EnterJsonFloat is called when production JsonFloat is entered.
func (s *BaseDNAC_A_WalkerListener) EnterJsonFloat(ctx *JsonFloatContext) {}

// ExitJsonFloat is called when production JsonFloat is exited.
func (s *BaseDNAC_A_WalkerListener) ExitJsonFloat(ctx *JsonFloatContext) {}

// EnterJsonArray is called when production JsonArray is entered.
func (s *BaseDNAC_A_WalkerListener) EnterJsonArray(ctx *JsonArrayContext) {}

// ExitJsonArray is called when production JsonArray is exited.
func (s *BaseDNAC_A_WalkerListener) ExitJsonArray(ctx *JsonArrayContext) {}

// EnterJsonObj is called when production JsonObj is entered.
func (s *BaseDNAC_A_WalkerListener) EnterJsonObj(ctx *JsonObjContext) {}

// ExitJsonObj is called when production JsonObj is exited.
func (s *BaseDNAC_A_WalkerListener) ExitJsonObj(ctx *JsonObjContext) {}
