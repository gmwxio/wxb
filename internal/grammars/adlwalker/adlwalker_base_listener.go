// Code generated from ADLWalker.g4 by ANTLR 4.7.1. DO NOT EDIT.

package walker // ADLWalker
import "github.com/wxio/goantlr"

// BaseADLWalkerListener is a complete listener for a parse tree produced by ADLWalker.
type BaseADLWalkerListener struct{}

var _ ADLWalkerListener = &BaseADLWalkerListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseADLWalkerListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseADLWalkerListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseADLWalkerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseADLWalkerListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterAdl is called when production adl is entered.
func (s *BaseADLWalkerListener) EnterAdl(ctx *AdlContext) {}

// ExitAdl is called when production adl is exited.
func (s *BaseADLWalkerListener) ExitAdl(ctx *AdlContext) {}

// EnterModule is called when production module is entered.
func (s *BaseADLWalkerListener) EnterModule(ctx *ModuleContext) {}

// ExitModule is called when production module is exited.
func (s *BaseADLWalkerListener) ExitModule(ctx *ModuleContext) {}

// EnterStruct is called when production Struct is entered.
func (s *BaseADLWalkerListener) EnterStruct(ctx *StructContext) {}

// ExitStruct is called when production Struct is exited.
func (s *BaseADLWalkerListener) ExitStruct(ctx *StructContext) {}

// EnterUnion is called when production Union is entered.
func (s *BaseADLWalkerListener) EnterUnion(ctx *UnionContext) {}

// ExitUnion is called when production Union is exited.
func (s *BaseADLWalkerListener) ExitUnion(ctx *UnionContext) {}

// EnterType is called when production Type is entered.
func (s *BaseADLWalkerListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production Type is exited.
func (s *BaseADLWalkerListener) ExitType(ctx *TypeContext) {}

// EnterNewtype is called when production Newtype is entered.
func (s *BaseADLWalkerListener) EnterNewtype(ctx *NewtypeContext) {}

// ExitNewtype is called when production Newtype is exited.
func (s *BaseADLWalkerListener) ExitNewtype(ctx *NewtypeContext) {}

// EnterModAnno is called when production ModAnno is entered.
func (s *BaseADLWalkerListener) EnterModAnno(ctx *ModAnnoContext) {}

// ExitModAnno is called when production ModAnno is exited.
func (s *BaseADLWalkerListener) ExitModAnno(ctx *ModAnnoContext) {}

// EnterDeclAnno is called when production DeclAnno is entered.
func (s *BaseADLWalkerListener) EnterDeclAnno(ctx *DeclAnnoContext) {}

// ExitDeclAnno is called when production DeclAnno is exited.
func (s *BaseADLWalkerListener) ExitDeclAnno(ctx *DeclAnnoContext) {}

// EnterFieldAnno is called when production FieldAnno is entered.
func (s *BaseADLWalkerListener) EnterFieldAnno(ctx *FieldAnnoContext) {}

// ExitFieldAnno is called when production FieldAnno is exited.
func (s *BaseADLWalkerListener) ExitFieldAnno(ctx *FieldAnnoContext) {}

// EnterTypeParamError is called when production TypeParamError is entered.
func (s *BaseADLWalkerListener) EnterTypeParamError(ctx *TypeParamErrorContext) {}

// ExitTypeParamError is called when production TypeParamError is exited.
func (s *BaseADLWalkerListener) ExitTypeParamError(ctx *TypeParamErrorContext) {}

// EnterField is called when production Field is entered.
func (s *BaseADLWalkerListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production Field is exited.
func (s *BaseADLWalkerListener) ExitField(ctx *FieldContext) {}

// EnterAnnotation is called when production annotation is entered.
func (s *BaseADLWalkerListener) EnterAnnotation(ctx *AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *BaseADLWalkerListener) ExitAnnotation(ctx *AnnotationContext) {}

// EnterTypeExpr_ is called when production typeExpr_ is entered.
func (s *BaseADLWalkerListener) EnterTypeExpr_(ctx *TypeExpr_Context) {}

// ExitTypeExpr_ is called when production typeExpr_ is exited.
func (s *BaseADLWalkerListener) ExitTypeExpr_(ctx *TypeExpr_Context) {}

// EnterTypeParams is called when production TypeParams is entered.
func (s *BaseADLWalkerListener) EnterTypeParams(ctx *TypeParamsContext) {}

// ExitTypeParams is called when production TypeParams is exited.
func (s *BaseADLWalkerListener) ExitTypeParams(ctx *TypeParamsContext) {}

// EnterJsonStr is called when production JsonStr is entered.
func (s *BaseADLWalkerListener) EnterJsonStr(ctx *JsonStrContext) {}

// ExitJsonStr is called when production JsonStr is exited.
func (s *BaseADLWalkerListener) ExitJsonStr(ctx *JsonStrContext) {}

// EnterJsonBool is called when production JsonBool is entered.
func (s *BaseADLWalkerListener) EnterJsonBool(ctx *JsonBoolContext) {}

// ExitJsonBool is called when production JsonBool is exited.
func (s *BaseADLWalkerListener) ExitJsonBool(ctx *JsonBoolContext) {}

// EnterJsonNull is called when production JsonNull is entered.
func (s *BaseADLWalkerListener) EnterJsonNull(ctx *JsonNullContext) {}

// ExitJsonNull is called when production JsonNull is exited.
func (s *BaseADLWalkerListener) ExitJsonNull(ctx *JsonNullContext) {}

// EnterJsonInt is called when production JsonInt is entered.
func (s *BaseADLWalkerListener) EnterJsonInt(ctx *JsonIntContext) {}

// ExitJsonInt is called when production JsonInt is exited.
func (s *BaseADLWalkerListener) ExitJsonInt(ctx *JsonIntContext) {}

// EnterJsonFloat is called when production JsonFloat is entered.
func (s *BaseADLWalkerListener) EnterJsonFloat(ctx *JsonFloatContext) {}

// ExitJsonFloat is called when production JsonFloat is exited.
func (s *BaseADLWalkerListener) ExitJsonFloat(ctx *JsonFloatContext) {}

// EnterJsonArray is called when production JsonArray is entered.
func (s *BaseADLWalkerListener) EnterJsonArray(ctx *JsonArrayContext) {}

// ExitJsonArray is called when production JsonArray is exited.
func (s *BaseADLWalkerListener) ExitJsonArray(ctx *JsonArrayContext) {}

// EnterJsonObj is called when production JsonObj is entered.
func (s *BaseADLWalkerListener) EnterJsonObj(ctx *JsonObjContext) {}

// ExitJsonObj is called when production JsonObj is exited.
func (s *BaseADLWalkerListener) ExitJsonObj(ctx *JsonObjContext) {}
