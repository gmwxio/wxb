// Generated from ADLParser.g4 by ANTLR 4.7.

package parser // ADLParser

import "github.com/wxio/goantlr"

// Struct of Handlers
type ADLParserHandlers struct {
	EnterEveryRule func(ctx antlr.RuleNode)
	ExitEveryRule  func(ctx antlr.RuleNode)

	Adl                func(ctx IAdlContext, this *ADLParserHandlers, args ...interface{}) (result interface{})
	ModuleStatement    func(ctx IModuleStatementContext, this *ADLParserHandlers, args ...interface{}) (result interface{})
	ImportStatement    func(ctx IImportStatementContext, this *ADLParserHandlers, args ...interface{}) (result interface{})
	LocalAnno          func(ctx ILocalAnnoContext, this *ADLParserHandlers, args ...interface{}) (result interface{})
	DocAnno            func(ctx IDocAnnoContext, this *ADLParserHandlers, args ...interface{}) (result interface{})
	StructOrUnion      func(ctx IStructOrUnionContext, this *ADLParserHandlers, args ...interface{}) (result interface{})
	TypeOrNewtype      func(ctx ITypeOrNewtypeContext, this *ADLParserHandlers, args ...interface{}) (result interface{})
	ModuleAnnotation   func(ctx IModuleAnnotationContext, this *ADLParserHandlers, args ...interface{}) (result interface{})
	DeclAnnotation     func(ctx IDeclAnnotationContext, this *ADLParserHandlers, args ...interface{}) (result interface{})
	FieldAnnotation    func(ctx IFieldAnnotationContext, this *ADLParserHandlers, args ...interface{}) (result interface{})
	TypeParameter      func(ctx ITypeParameterContext, this *ADLParserHandlers, args ...interface{}) (result interface{})
	ErrorTypeParam     func(ctx IErrorTypeParamContext, this *ADLParserHandlers, args ...interface{}) (result interface{})
	TypeParamError     func(ctx ITypeParamErrorContext, this *ADLParserHandlers, args ...interface{}) (result interface{})
	TypeExpression     func(ctx ITypeExpressionContext, this *ADLParserHandlers, args ...interface{}) (result interface{})
	TypeExpressionElem func(ctx ITypeExpressionElemContext, this *ADLParserHandlers, args ...interface{}) (result interface{})
	FieldStatement     func(ctx IFieldStatementContext, this *ADLParserHandlers, args ...interface{}) (result interface{})
	StringStatement    func(ctx IStringStatementContext, this *ADLParserHandlers, args ...interface{}) (result interface{})
	TrueFalseNull      func(ctx ITrueFalseNullContext, this *ADLParserHandlers, args ...interface{}) (result interface{})
	NumberStatement    func(ctx INumberStatementContext, this *ADLParserHandlers, args ...interface{}) (result interface{})
	FloatStatement     func(ctx IFloatStatementContext, this *ADLParserHandlers, args ...interface{}) (result interface{})
	ArrayStatement     func(ctx IArrayStatementContext, this *ADLParserHandlers, args ...interface{}) (result interface{})
	ObjStatement       func(ctx IObjStatementContext, this *ADLParserHandlers, args ...interface{}) (result interface{})
}

// A complete Visitor for a parse tree produced by ADLParser.
type ADLParserVisitor interface {
	antlr.ParseTreeVisitor
	AdlContextVisitor
	ModuleStatementContextVisitor
	ImportStatementContextVisitor
	LocalAnnoContextVisitor
	DocAnnoContextVisitor
	StructOrUnionContextVisitor
	TypeOrNewtypeContextVisitor
	ModuleAnnotationContextVisitor
	DeclAnnotationContextVisitor
	FieldAnnotationContextVisitor
	TypeParameterContextVisitor
	ErrorTypeParamContextVisitor
	TypeParamErrorContextVisitor
	TypeExpressionContextVisitor
	TypeExpressionElemContextVisitor
	FieldStatementContextVisitor
	StringStatementContextVisitor
	TrueFalseNullContextVisitor
	NumberStatementContextVisitor
	FloatStatementContextVisitor
	ArrayStatementContextVisitor
	ObjStatementContextVisitor
}

type AdlContextVisitor interface {
	VisitAdl(ctx IAdlContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type ModuleStatementContextVisitor interface {
	VisitModuleStatement(ctx IModuleStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type ImportStatementContextVisitor interface {
	VisitImportStatement(ctx IImportStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type LocalAnnoContextVisitor interface {
	VisitLocalAnno(ctx ILocalAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type DocAnnoContextVisitor interface {
	VisitDocAnno(ctx IDocAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type StructOrUnionContextVisitor interface {
	VisitStructOrUnion(ctx IStructOrUnionContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type TypeOrNewtypeContextVisitor interface {
	VisitTypeOrNewtype(ctx ITypeOrNewtypeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type ModuleAnnotationContextVisitor interface {
	VisitModuleAnnotation(ctx IModuleAnnotationContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type DeclAnnotationContextVisitor interface {
	VisitDeclAnnotation(ctx IDeclAnnotationContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type FieldAnnotationContextVisitor interface {
	VisitFieldAnnotation(ctx IFieldAnnotationContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type TypeParameterContextVisitor interface {
	VisitTypeParameter(ctx ITypeParameterContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type ErrorTypeParamContextVisitor interface {
	VisitErrorTypeParam(ctx IErrorTypeParamContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type TypeParamErrorContextVisitor interface {
	VisitTypeParamError(ctx ITypeParamErrorContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type TypeExpressionContextVisitor interface {
	VisitTypeExpression(ctx ITypeExpressionContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type TypeExpressionElemContextVisitor interface {
	VisitTypeExpressionElem(ctx ITypeExpressionElemContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type FieldStatementContextVisitor interface {
	VisitFieldStatement(ctx IFieldStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type StringStatementContextVisitor interface {
	VisitStringStatement(ctx IStringStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type TrueFalseNullContextVisitor interface {
	VisitTrueFalseNull(ctx ITrueFalseNullContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type NumberStatementContextVisitor interface {
	VisitNumberStatement(ctx INumberStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type FloatStatementContextVisitor interface {
	VisitFloatStatement(ctx IFloatStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type ArrayStatementContextVisitor interface {
	VisitArrayStatement(ctx IArrayStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type ObjStatementContextVisitor interface {
	VisitObjStatement(ctx IObjStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
