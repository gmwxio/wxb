// Code generated from ADLParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // ADLParser

import "github.com/wxio/goantlr"

// ADLParserListener is a complete listener for a parse tree produced by ADLParser.
type ADLParserListener interface {
	antlr.ParseTreeListener

	// EnterAdl is called when entering the adl production.
	EnterAdl(c *AdlContext)

	// EnterModuleStatement is called when entering the ModuleStatement production.
	EnterModuleStatement(c *ModuleStatementContext)

	// EnterImportStatement is called when entering the ImportStatement production.
	EnterImportStatement(c *ImportStatementContext)

	// EnterStructOrUnion is called when entering the StructOrUnion production.
	EnterStructOrUnion(c *StructOrUnionContext)

	// EnterTypeOrNewtype is called when entering the TypeOrNewtype production.
	EnterTypeOrNewtype(c *TypeOrNewtypeContext)

	// EnterModuleAnnotation is called when entering the ModuleAnnotation production.
	EnterModuleAnnotation(c *ModuleAnnotationContext)

	// EnterDeclAnnotation is called when entering the DeclAnnotation production.
	EnterDeclAnnotation(c *DeclAnnotationContext)

	// EnterFieldAnnotation is called when entering the FieldAnnotation production.
	EnterFieldAnnotation(c *FieldAnnotationContext)

	// EnterTypeParam is called when entering the typeParam production.
	EnterTypeParam(c *TypeParamContext)

	// EnterSoruBody is called when entering the soruBody production.
	EnterSoruBody(c *SoruBodyContext)

	// EnterStringStatement is called when entering the StringStatement production.
	EnterStringStatement(c *StringStatementContext)

	// EnterTrueFalseNull is called when entering the TrueFalseNull production.
	EnterTrueFalseNull(c *TrueFalseNullContext)

	// EnterNumberStatement is called when entering the NumberStatement production.
	EnterNumberStatement(c *NumberStatementContext)

	// EnterFloatStatement is called when entering the FloatStatement production.
	EnterFloatStatement(c *FloatStatementContext)

	// EnterArrayStatement is called when entering the ArrayStatement production.
	EnterArrayStatement(c *ArrayStatementContext)

	// EnterObjStatement is called when entering the ObjStatement production.
	EnterObjStatement(c *ObjStatementContext)

	// ExitAdl is called when exiting the adl production.
	ExitAdl(c *AdlContext)

	// ExitModuleStatement is called when exiting the ModuleStatement production.
	ExitModuleStatement(c *ModuleStatementContext)

	// ExitImportStatement is called when exiting the ImportStatement production.
	ExitImportStatement(c *ImportStatementContext)

	// ExitStructOrUnion is called when exiting the StructOrUnion production.
	ExitStructOrUnion(c *StructOrUnionContext)

	// ExitTypeOrNewtype is called when exiting the TypeOrNewtype production.
	ExitTypeOrNewtype(c *TypeOrNewtypeContext)

	// ExitModuleAnnotation is called when exiting the ModuleAnnotation production.
	ExitModuleAnnotation(c *ModuleAnnotationContext)

	// ExitDeclAnnotation is called when exiting the DeclAnnotation production.
	ExitDeclAnnotation(c *DeclAnnotationContext)

	// ExitFieldAnnotation is called when exiting the FieldAnnotation production.
	ExitFieldAnnotation(c *FieldAnnotationContext)

	// ExitTypeParam is called when exiting the typeParam production.
	ExitTypeParam(c *TypeParamContext)

	// ExitSoruBody is called when exiting the soruBody production.
	ExitSoruBody(c *SoruBodyContext)

	// ExitStringStatement is called when exiting the StringStatement production.
	ExitStringStatement(c *StringStatementContext)

	// ExitTrueFalseNull is called when exiting the TrueFalseNull production.
	ExitTrueFalseNull(c *TrueFalseNullContext)

	// ExitNumberStatement is called when exiting the NumberStatement production.
	ExitNumberStatement(c *NumberStatementContext)

	// ExitFloatStatement is called when exiting the FloatStatement production.
	ExitFloatStatement(c *FloatStatementContext)

	// ExitArrayStatement is called when exiting the ArrayStatement production.
	ExitArrayStatement(c *ArrayStatementContext)

	// ExitObjStatement is called when exiting the ObjStatement production.
	ExitObjStatement(c *ObjStatementContext)
}
