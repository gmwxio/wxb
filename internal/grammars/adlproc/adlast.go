package adlproc

type Annotations []struct {
	Key ScopedName  `json:"v1"`
	Val interface{} `json:"v2"`
}

type ScopedName struct {
	ModuleName string `json:"moduleName"`
	Name       string `json:"name"`
}

type TypeRef struct {
	Primitive *string     `json:"primitive,omitempty"`
	TypeParam *string     `json:"typeParam,omitempty"`
	Reference *ScopedName `json:"reference,omitempty"`
}

type TypeExpr struct {
	TypeRef    TypeRef    `json:"typeRef"`
	Parameters []TypeExpr `json:"parameters"`
}

type Field struct {
	Name           string      `json:"name"`
	SerializedName string      `json:"serializedName"`
	TypeExpr       TypeExpr    `json:"typeExpr"`
	Default        interface{} `json:"default,omitempty"`
	Annotations    Annotations `json:"annotations"`
}

// Struct & Union
type Name struct {
	TypeParams []string `json:"typeParams"`
	Field      []Field  `json:"fields"`
}

type TypeDef struct {
	TypeParams []string `json:"typeParams"`
	TypeExpr   TypeExpr `json:"typeExpr"`
}

type NewType struct {
	TypeParams []string    `json:"typeParams"`
	TypeExpr   TypeExpr    `json:"typeExpr"`
	Default    interface{} `json:"default,omitempty"`
}

type DeclType struct {
	Struct_  *Name    `json:"struct_,omitempty"`
	Union_   *Name    `json:"union_,omitempty"`
	Type_    *TypeDef `json:"type_,omitempty"`
	Newtype_ *NewType `json:"newtype_,omitempty"`
}

type Decl struct {
	MyToken     `json:"-"`
	Name        string      `json:"name"`
	Version     *string     `json:"version,omitempty"`
	Type_       DeclType    `json:"type_"`
	Annotations Annotations `json:"annotations"`
}

// struct ScopedDecl
// {
//     ModuleName moduleName;
//     Decl decl;
// };

// type DeclVersions = Vector<Decl>;

type Import struct {
	ModuleName *string     `json:"moduleName,omitempty"`
	ScopedName *ScopedName `json:"scopedName,omitempty"`
}

type Module struct {
	MyToken     `json:"-"`
	Name        string          `json:"name"`
	Imports     *[]Import       `json:"imports"`
	Decls       map[string]Decl `json:"decls"`
	Annotations Annotations     `json:"annotations"`
}
