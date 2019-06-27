package adlproc

import (
	"fmt"
)

type Annotations []Annotation

type Annotation struct {
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
	Struct  *Name    `json:"struct_,omitempty"`
	Union   *Name    `json:"union_,omitempty"`
	Type    *TypeDef `json:"type_,omitempty"`
	Newtype *NewType `json:"newtype_,omitempty"`
}

type Decl struct {
	Name        string      `json:"name"`
	Version     *string     `json:"version,omitempty"`
	Type        DeclType    `json:"type_"`
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
	Name        string          `json:"name"`
	Imports     []Import        `json:"imports"`
	Decls       map[string]Decl `json:"decls"`
	Annotations `json:"annotations"`
}

type AnnotateAble interface {
	AddAnnotation(an Annotation)
}
type ImportableAble interface {
	AddImport(Import)
}
type Setable interface {
	Set(val interface{})
}

func (ans *Annotations) AddAnnotation(an Annotation) {
	*ans = append(*ans, an)
}
func (mo *Module) AddImport(im Import) {
	mo.Imports = append(mo.Imports, im)
}
func (an *Annotation) Set(val interface{}) {
	an.Val = val
}

func (m Module) String() string { return fmt.Sprintf("name: %s imports: %v", m.Name, m.Imports) }
func (i Import) String() string {
	if i.ModuleName != nil {
		return *i.ModuleName
	}
	return i.ScopedName.ModuleName + ":" + i.ScopedName.Name
}
