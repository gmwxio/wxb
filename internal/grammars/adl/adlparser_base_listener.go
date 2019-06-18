// Code generated from ADLParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // ADLParser

import "github.com/wxio/goantlr"

// BaseADLParserListener is a complete listener for a parse tree produced by ADLParser.
type BaseADLParserListener struct{}

var _ ADLParserListener = &BaseADLParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseADLParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseADLParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseADLParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseADLParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterAdl is called when production adl is entered.
func (s *BaseADLParserListener) EnterAdl(ctx *AdlContext) {}

// ExitAdl is called when production adl is exited.
func (s *BaseADLParserListener) ExitAdl(ctx *AdlContext) {}

// EnterModuleStatement is called when production ModuleStatement is entered.
func (s *BaseADLParserListener) EnterModuleStatement(ctx *ModuleStatementContext) {}

// ExitModuleStatement is called when production ModuleStatement is exited.
func (s *BaseADLParserListener) ExitModuleStatement(ctx *ModuleStatementContext) {}

// EnterImportStatement is called when production ImportStatement is entered.
func (s *BaseADLParserListener) EnterImportStatement(ctx *ImportStatementContext) {}

// ExitImportStatement is called when production ImportStatement is exited.
func (s *BaseADLParserListener) ExitImportStatement(ctx *ImportStatementContext) {}

// EnterStructOrUnion is called when production StructOrUnion is entered.
func (s *BaseADLParserListener) EnterStructOrUnion(ctx *StructOrUnionContext) {}

// ExitStructOrUnion is called when production StructOrUnion is exited.
func (s *BaseADLParserListener) ExitStructOrUnion(ctx *StructOrUnionContext) {}

// EnterTypeOrNewtype is called when production TypeOrNewtype is entered.
func (s *BaseADLParserListener) EnterTypeOrNewtype(ctx *TypeOrNewtypeContext) {}

// ExitTypeOrNewtype is called when production TypeOrNewtype is exited.
func (s *BaseADLParserListener) ExitTypeOrNewtype(ctx *TypeOrNewtypeContext) {}

// EnterModuleAnnotation is called when production ModuleAnnotation is entered.
func (s *BaseADLParserListener) EnterModuleAnnotation(ctx *ModuleAnnotationContext) {}

// ExitModuleAnnotation is called when production ModuleAnnotation is exited.
func (s *BaseADLParserListener) ExitModuleAnnotation(ctx *ModuleAnnotationContext) {}

// EnterDeclAnnotation is called when production DeclAnnotation is entered.
func (s *BaseADLParserListener) EnterDeclAnnotation(ctx *DeclAnnotationContext) {}

// ExitDeclAnnotation is called when production DeclAnnotation is exited.
func (s *BaseADLParserListener) ExitDeclAnnotation(ctx *DeclAnnotationContext) {}

// EnterFieldAnnotation is called when production FieldAnnotation is entered.
func (s *BaseADLParserListener) EnterFieldAnnotation(ctx *FieldAnnotationContext) {}

// ExitFieldAnnotation is called when production FieldAnnotation is exited.
func (s *BaseADLParserListener) ExitFieldAnnotation(ctx *FieldAnnotationContext) {}

// EnterTypeParam is called when production typeParam is entered.
func (s *BaseADLParserListener) EnterTypeParam(ctx *TypeParamContext) {}

// ExitTypeParam is called when production typeParam is exited.
func (s *BaseADLParserListener) ExitTypeParam(ctx *TypeParamContext) {}

// EnterTypeExpr is called when production typeExpr is entered.
func (s *BaseADLParserListener) EnterTypeExpr(ctx *TypeExprContext) {}

// ExitTypeExpr is called when production typeExpr is exited.
func (s *BaseADLParserListener) ExitTypeExpr(ctx *TypeExprContext) {}

// EnterSoruBody is called when production soruBody is entered.
func (s *BaseADLParserListener) EnterSoruBody(ctx *SoruBodyContext) {}

// ExitSoruBody is called when production soruBody is exited.
func (s *BaseADLParserListener) ExitSoruBody(ctx *SoruBodyContext) {}

// EnterStringStatement is called when production StringStatement is entered.
func (s *BaseADLParserListener) EnterStringStatement(ctx *StringStatementContext) {}

// ExitStringStatement is called when production StringStatement is exited.
func (s *BaseADLParserListener) ExitStringStatement(ctx *StringStatementContext) {}

// EnterTrueFalseNull is called when production TrueFalseNull is entered.
func (s *BaseADLParserListener) EnterTrueFalseNull(ctx *TrueFalseNullContext) {}

// ExitTrueFalseNull is called when production TrueFalseNull is exited.
func (s *BaseADLParserListener) ExitTrueFalseNull(ctx *TrueFalseNullContext) {}

// EnterNumberStatement is called when production NumberStatement is entered.
func (s *BaseADLParserListener) EnterNumberStatement(ctx *NumberStatementContext) {}

// ExitNumberStatement is called when production NumberStatement is exited.
func (s *BaseADLParserListener) ExitNumberStatement(ctx *NumberStatementContext) {}

// EnterFloatStatement is called when production FloatStatement is entered.
func (s *BaseADLParserListener) EnterFloatStatement(ctx *FloatStatementContext) {}

// ExitFloatStatement is called when production FloatStatement is exited.
func (s *BaseADLParserListener) ExitFloatStatement(ctx *FloatStatementContext) {}

// EnterArrayStatement is called when production ArrayStatement is entered.
func (s *BaseADLParserListener) EnterArrayStatement(ctx *ArrayStatementContext) {}

// ExitArrayStatement is called when production ArrayStatement is exited.
func (s *BaseADLParserListener) ExitArrayStatement(ctx *ArrayStatementContext) {}

// EnterObjStatement is called when production ObjStatement is entered.
func (s *BaseADLParserListener) EnterObjStatement(ctx *ObjStatementContext) {}

// ExitObjStatement is called when production ObjStatement is exited.
func (s *BaseADLParserListener) ExitObjStatement(ctx *ObjStatementContext) {}
