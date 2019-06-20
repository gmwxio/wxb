// Code generated from ADLWalker.g4 by ANTLR 4.7.1. DO NOT EDIT.

package walker // ADLWalker
import "github.com/wxio/goantlr"

// ADLWalkerListener is a complete listener for a parse tree produced by ADLWalker.
type ADLWalkerListener interface {
	antlr.ParseTreeListener

	// EnterAdl is called when entering the adl production.
	EnterAdl(c *AdlContext)

	// EnterModule is called when entering the module production.
	EnterModule(c *ModuleContext)

	// EnterStruct is called when entering the Struct production.
	EnterStruct(c *StructContext)

	// EnterUnion is called when entering the Union production.
	EnterUnion(c *UnionContext)

	// EnterType is called when entering the Type production.
	EnterType(c *TypeContext)

	// EnterNewtype is called when entering the Newtype production.
	EnterNewtype(c *NewtypeContext)

	// EnterModAnno is called when entering the ModAnno production.
	EnterModAnno(c *ModAnnoContext)

	// EnterDeclAnno is called when entering the DeclAnno production.
	EnterDeclAnno(c *DeclAnnoContext)

	// EnterFieldAnno is called when entering the FieldAnno production.
	EnterFieldAnno(c *FieldAnnoContext)

	// EnterTypeParamError is called when entering the TypeParamError production.
	EnterTypeParamError(c *TypeParamErrorContext)

	// EnterField is called when entering the Field production.
	EnterField(c *FieldContext)

	// EnterTypeExpr_ is called when entering the typeExpr_ production.
	EnterTypeExpr_(c *TypeExpr_Context)

	// EnterTypeParams is called when entering the TypeParams production.
	EnterTypeParams(c *TypeParamsContext)

	// EnterJsonStr is called when entering the JsonStr production.
	EnterJsonStr(c *JsonStrContext)

	// EnterJsonBool is called when entering the JsonBool production.
	EnterJsonBool(c *JsonBoolContext)

	// EnterJsonNull is called when entering the JsonNull production.
	EnterJsonNull(c *JsonNullContext)

	// EnterJsonInt is called when entering the JsonInt production.
	EnterJsonInt(c *JsonIntContext)

	// EnterJsonFloat is called when entering the JsonFloat production.
	EnterJsonFloat(c *JsonFloatContext)

	// EnterJsonArray is called when entering the JsonArray production.
	EnterJsonArray(c *JsonArrayContext)

	// EnterJsonObj is called when entering the JsonObj production.
	EnterJsonObj(c *JsonObjContext)

	// ExitAdl is called when exiting the adl production.
	ExitAdl(c *AdlContext)

	// ExitModule is called when exiting the module production.
	ExitModule(c *ModuleContext)

	// ExitStruct is called when exiting the Struct production.
	ExitStruct(c *StructContext)

	// ExitUnion is called when exiting the Union production.
	ExitUnion(c *UnionContext)

	// ExitType is called when exiting the Type production.
	ExitType(c *TypeContext)

	// ExitNewtype is called when exiting the Newtype production.
	ExitNewtype(c *NewtypeContext)

	// ExitModAnno is called when exiting the ModAnno production.
	ExitModAnno(c *ModAnnoContext)

	// ExitDeclAnno is called when exiting the DeclAnno production.
	ExitDeclAnno(c *DeclAnnoContext)

	// ExitFieldAnno is called when exiting the FieldAnno production.
	ExitFieldAnno(c *FieldAnnoContext)

	// ExitTypeParamError is called when exiting the TypeParamError production.
	ExitTypeParamError(c *TypeParamErrorContext)

	// ExitField is called when exiting the Field production.
	ExitField(c *FieldContext)

	// ExitTypeExpr_ is called when exiting the typeExpr_ production.
	ExitTypeExpr_(c *TypeExpr_Context)

	// ExitTypeParams is called when exiting the TypeParams production.
	ExitTypeParams(c *TypeParamsContext)

	// ExitJsonStr is called when exiting the JsonStr production.
	ExitJsonStr(c *JsonStrContext)

	// ExitJsonBool is called when exiting the JsonBool production.
	ExitJsonBool(c *JsonBoolContext)

	// ExitJsonNull is called when exiting the JsonNull production.
	ExitJsonNull(c *JsonNullContext)

	// ExitJsonInt is called when exiting the JsonInt production.
	ExitJsonInt(c *JsonIntContext)

	// ExitJsonFloat is called when exiting the JsonFloat production.
	ExitJsonFloat(c *JsonFloatContext)

	// ExitJsonArray is called when exiting the JsonArray production.
	ExitJsonArray(c *JsonArrayContext)

	// ExitJsonObj is called when exiting the JsonObj production.
	ExitJsonObj(c *JsonObjContext)
}
