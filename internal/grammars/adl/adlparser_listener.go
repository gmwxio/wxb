// Generated from ADLParser.g4 by ANTLR 4.7.

package parser // ADLParser

import "github.com/wxio/goantlr"

// ADLParserListener is a complete listener for a parse tree produced by ADLParser.
type ADLParserListener interface {
	antlr.ParseTreeListener

	AdlEntryListener
	AdlExitListener

	TypeParamErrorEntryListener
	TypeParamErrorExitListener

	ModuleStatementEntryListener
	ModuleStatementExitListener

	ImportStatementEntryListener
	ImportStatementExitListener

	LocalAnnoEntryListener
	LocalAnnoExitListener

	DocAnnoEntryListener
	DocAnnoExitListener

	StructOrUnionEntryListener
	StructOrUnionExitListener

	TypeOrNewtypeEntryListener
	TypeOrNewtypeExitListener

	ModuleAnnotationEntryListener
	ModuleAnnotationExitListener

	DeclAnnotationEntryListener
	DeclAnnotationExitListener

	FieldAnnotationEntryListener
	FieldAnnotationExitListener

	TypeParameterEntryListener
	TypeParameterExitListener

	ErrorTypeParamEntryListener
	ErrorTypeParamExitListener

	TypeExpressionEntryListener
	TypeExpressionExitListener

	TypeExpressionElemEntryListener
	TypeExpressionElemExitListener

	FieldStatementEntryListener
	FieldStatementExitListener

	StringStatementEntryListener
	StringStatementExitListener

	TrueFalseNullEntryListener
	TrueFalseNullExitListener

	NumberStatementEntryListener
	NumberStatementExitListener

	FloatStatementEntryListener
	FloatStatementExitListener

	ArrayStatementEntryListener
	ArrayStatementExitListener

	ObjStatementEntryListener
	ObjStatementExitListener
}

//
// Rules with unnamed alternatives
//
type AdlEntryListener interface {
	EnterAdl(c *AdlContext)
}
type AdlExitListener interface {
	ExitAdl(c *AdlContext)
}

type TypeParamErrorEntryListener interface {
	EnterTypeParamError(c *TypeParamErrorContext)
}
type TypeParamErrorExitListener interface {
	ExitTypeParamError(c *TypeParamErrorContext)
}

//
// Named alternatives
//
//
// From Rule 'module'
type ModuleStatementEntryListener interface {
	EnterModuleStatement(c *ModuleStatementContext)
}
type ModuleStatementExitListener interface {
	ExitModuleStatement(c *ModuleStatementContext)
}

// From Rule 'imports'
type ImportStatementEntryListener interface {
	EnterImportStatement(c *ImportStatementContext)
}
type ImportStatementExitListener interface {
	ExitImportStatement(c *ImportStatementContext)
}

// From Rule 'annon'
type LocalAnnoEntryListener interface {
	EnterLocalAnno(c *LocalAnnoContext)
}
type LocalAnnoExitListener interface {
	ExitLocalAnno(c *LocalAnnoContext)
}

// From Rule 'annon'
type DocAnnoEntryListener interface {
	EnterDocAnno(c *DocAnnoContext)
}
type DocAnnoExitListener interface {
	ExitDocAnno(c *DocAnnoContext)
}

// From Rule 'top_level_statement'
type StructOrUnionEntryListener interface {
	EnterStructOrUnion(c *StructOrUnionContext)
}
type StructOrUnionExitListener interface {
	ExitStructOrUnion(c *StructOrUnionContext)
}

// From Rule 'top_level_statement'
type TypeOrNewtypeEntryListener interface {
	EnterTypeOrNewtype(c *TypeOrNewtypeContext)
}
type TypeOrNewtypeExitListener interface {
	ExitTypeOrNewtype(c *TypeOrNewtypeContext)
}

// From Rule 'top_level_statement'
type ModuleAnnotationEntryListener interface {
	EnterModuleAnnotation(c *ModuleAnnotationContext)
}
type ModuleAnnotationExitListener interface {
	ExitModuleAnnotation(c *ModuleAnnotationContext)
}

// From Rule 'top_level_statement'
type DeclAnnotationEntryListener interface {
	EnterDeclAnnotation(c *DeclAnnotationContext)
}
type DeclAnnotationExitListener interface {
	ExitDeclAnnotation(c *DeclAnnotationContext)
}

// From Rule 'top_level_statement'
type FieldAnnotationEntryListener interface {
	EnterFieldAnnotation(c *FieldAnnotationContext)
}
type FieldAnnotationExitListener interface {
	ExitFieldAnnotation(c *FieldAnnotationContext)
}

// From Rule 'typeParam'
type TypeParameterEntryListener interface {
	EnterTypeParameter(c *TypeParameterContext)
}
type TypeParameterExitListener interface {
	ExitTypeParameter(c *TypeParameterContext)
}

// From Rule 'typeParam'
type ErrorTypeParamEntryListener interface {
	EnterErrorTypeParam(c *ErrorTypeParamContext)
}
type ErrorTypeParamExitListener interface {
	ExitErrorTypeParam(c *ErrorTypeParamContext)
}

// From Rule 'typeExpr'
type TypeExpressionEntryListener interface {
	EnterTypeExpression(c *TypeExpressionContext)
}
type TypeExpressionExitListener interface {
	ExitTypeExpression(c *TypeExpressionContext)
}

// From Rule 'typeExprElem'
type TypeExpressionElemEntryListener interface {
	EnterTypeExpressionElem(c *TypeExpressionElemContext)
}
type TypeExpressionElemExitListener interface {
	ExitTypeExpressionElem(c *TypeExpressionElemContext)
}

// From Rule 'soruBody'
type FieldStatementEntryListener interface {
	EnterFieldStatement(c *FieldStatementContext)
}
type FieldStatementExitListener interface {
	ExitFieldStatement(c *FieldStatementContext)
}

// From Rule 'jsonValue'
type StringStatementEntryListener interface {
	EnterStringStatement(c *StringStatementContext)
}
type StringStatementExitListener interface {
	ExitStringStatement(c *StringStatementContext)
}

// From Rule 'jsonValue'
type TrueFalseNullEntryListener interface {
	EnterTrueFalseNull(c *TrueFalseNullContext)
}
type TrueFalseNullExitListener interface {
	ExitTrueFalseNull(c *TrueFalseNullContext)
}

// From Rule 'jsonValue'
type NumberStatementEntryListener interface {
	EnterNumberStatement(c *NumberStatementContext)
}
type NumberStatementExitListener interface {
	ExitNumberStatement(c *NumberStatementContext)
}

// From Rule 'jsonValue'
type FloatStatementEntryListener interface {
	EnterFloatStatement(c *FloatStatementContext)
}
type FloatStatementExitListener interface {
	ExitFloatStatement(c *FloatStatementContext)
}

// From Rule 'jsonValue'
type ArrayStatementEntryListener interface {
	EnterArrayStatement(c *ArrayStatementContext)
}
type ArrayStatementExitListener interface {
	ExitArrayStatement(c *ArrayStatementContext)
}

// From Rule 'jsonValue'
type ObjStatementEntryListener interface {
	EnterObjStatement(c *ObjStatementContext)
}
type ObjStatementExitListener interface {
	ExitObjStatement(c *ObjStatementContext)
}
