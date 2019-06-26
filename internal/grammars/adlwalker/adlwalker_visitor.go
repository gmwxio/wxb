// Generated from ADLWalker.g4 by ANTLR 4.7.

package walker // ADLWalker
import "github.com/wxio/goantlr"

// Struct of Handlers
type ADLWalkerHandlers struct {
	EnterEveryRule func(ctx antlr.RuleNode)
	ExitEveryRule  func(ctx antlr.RuleNode)

	Adl            func(ctx IAdlContext, this *ADLWalkerHandlers, args ...interface{}) (result interface{})
	Json           func(ctx IJsonContext, this *ADLWalkerHandlers, args ...interface{}) (result interface{})
	Module         func(ctx IModuleContext, this *ADLWalkerHandlers, args ...interface{}) (result interface{})
	Struct         func(ctx IStructContext, this *ADLWalkerHandlers, args ...interface{}) (result interface{})
	Union          func(ctx IUnionContext, this *ADLWalkerHandlers, args ...interface{}) (result interface{})
	Type           func(ctx ITypeContext, this *ADLWalkerHandlers, args ...interface{}) (result interface{})
	Newtype        func(ctx INewtypeContext, this *ADLWalkerHandlers, args ...interface{}) (result interface{})
	ModAnno        func(ctx IModAnnoContext, this *ADLWalkerHandlers, args ...interface{}) (result interface{})
	DeclAnno       func(ctx IDeclAnnoContext, this *ADLWalkerHandlers, args ...interface{}) (result interface{})
	FieldAnno      func(ctx IFieldAnnoContext, this *ADLWalkerHandlers, args ...interface{}) (result interface{})
	TypeParamError func(ctx ITypeParamErrorContext, this *ADLWalkerHandlers, args ...interface{}) (result interface{})
	Field          func(ctx IFieldContext, this *ADLWalkerHandlers, args ...interface{}) (result interface{})
	Annotation     func(ctx IAnnotationContext, this *ADLWalkerHandlers, args ...interface{}) (result interface{})
	TypeExpr_      func(ctx ITypeExpr_Context, this *ADLWalkerHandlers, args ...interface{}) (result interface{})
	TypeParams     func(ctx ITypeParamsContext, this *ADLWalkerHandlers, args ...interface{}) (result interface{})
	JsonStr        func(ctx IJsonStrContext, this *ADLWalkerHandlers, args ...interface{}) (result interface{})
	JsonBool       func(ctx IJsonBoolContext, this *ADLWalkerHandlers, args ...interface{}) (result interface{})
	JsonNull       func(ctx IJsonNullContext, this *ADLWalkerHandlers, args ...interface{}) (result interface{})
	JsonInt        func(ctx IJsonIntContext, this *ADLWalkerHandlers, args ...interface{}) (result interface{})
	JsonFloat      func(ctx IJsonFloatContext, this *ADLWalkerHandlers, args ...interface{}) (result interface{})
	JsonArray      func(ctx IJsonArrayContext, this *ADLWalkerHandlers, args ...interface{}) (result interface{})
	JsonObj        func(ctx IJsonObjContext, this *ADLWalkerHandlers, args ...interface{}) (result interface{})
}

// A complete Visitor for a parse tree produced by ADLWalker.
type ADLWalkerVisitor interface {
	antlr.ParseTreeVisitor
	AdlContextVisitor
	JsonContextVisitor
	ModuleContextVisitor
	StructContextVisitor
	UnionContextVisitor
	TypeContextVisitor
	NewtypeContextVisitor
	ModAnnoContextVisitor
	DeclAnnoContextVisitor
	FieldAnnoContextVisitor
	TypeParamErrorContextVisitor
	FieldContextVisitor
	AnnotationContextVisitor
	TypeExpr_ContextVisitor
	TypeParamsContextVisitor
	JsonStrContextVisitor
	JsonBoolContextVisitor
	JsonNullContextVisitor
	JsonIntContextVisitor
	JsonFloatContextVisitor
	JsonArrayContextVisitor
	JsonObjContextVisitor
}

type AdlContextVisitor interface {
	VisitAdl(ctx IAdlContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type JsonContextVisitor interface {
	VisitJson(ctx IJsonContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type ModuleContextVisitor interface {
	VisitModule(ctx IModuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type StructContextVisitor interface {
	VisitStruct(ctx IStructContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type UnionContextVisitor interface {
	VisitUnion(ctx IUnionContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type TypeContextVisitor interface {
	VisitType(ctx ITypeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type NewtypeContextVisitor interface {
	VisitNewtype(ctx INewtypeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type ModAnnoContextVisitor interface {
	VisitModAnno(ctx IModAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type DeclAnnoContextVisitor interface {
	VisitDeclAnno(ctx IDeclAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type FieldAnnoContextVisitor interface {
	VisitFieldAnno(ctx IFieldAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type TypeParamErrorContextVisitor interface {
	VisitTypeParamError(ctx ITypeParamErrorContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type FieldContextVisitor interface {
	VisitField(ctx IFieldContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type AnnotationContextVisitor interface {
	VisitAnnotation(ctx IAnnotationContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type TypeExpr_ContextVisitor interface {
	VisitTypeExpr_(ctx ITypeExpr_Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type TypeParamsContextVisitor interface {
	VisitTypeParams(ctx ITypeParamsContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type JsonStrContextVisitor interface {
	VisitJsonStr(ctx IJsonStrContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type JsonBoolContextVisitor interface {
	VisitJsonBool(ctx IJsonBoolContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type JsonNullContextVisitor interface {
	VisitJsonNull(ctx IJsonNullContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type JsonIntContextVisitor interface {
	VisitJsonInt(ctx IJsonIntContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type JsonFloatContextVisitor interface {
	VisitJsonFloat(ctx IJsonFloatContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type JsonArrayContextVisitor interface {
	VisitJsonArray(ctx IJsonArrayContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type JsonObjContextVisitor interface {
	VisitJsonObj(ctx IJsonObjContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
