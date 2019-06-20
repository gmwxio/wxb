package adlproc

import (
	"strings"

	antlr "github.com/wxio/goantlr"
)

type MyToken struct {
	antlr.Token
	TType int
}

func (t *MyToken) GetTokenType() int { return t.TType }

type ADLNode struct {
	MyToken
}

func (ADLNode) String() string { return "" }

type ModuleNode struct {
	MyToken
	Name []string
}

func (mn ModuleNode) String() string { return strings.Join(mn.Name, ".") }

type ErrorNode struct {
	MyToken
	Expected string
	Received string
}

func (node ErrorNode) String() string {
	return "error - expected:" + node.Expected + " received:" + node.Received
}

type ImportNode struct {
	MyToken
	Path []string
	Star bool
}

func (node ImportNode) String() string {
	if node.Star {
		return strings.Join(node.Path, ".") + ".*"
	}
	return strings.Join(node.Path, ".")
}

type AnnoNode struct {
	MyToken
	Name string
	Doc  bool
}
type StructNode struct {
	MyToken
	Name string
}

func (node StructNode) String() string {
	return "struct " + node.Name
}

type UnionNode struct {
	MyToken
	Name string
}

func (node UnionNode) String() string {
	return "union " + node.Name
}

type TypeNode struct {
	MyToken
	Name    string
	TypeRef string
}

func (node TypeNode) String() string {
	return "type " + node.Name + " = " + node.TypeRef
}

type NewTypeNode struct {
	MyToken
	Name    string
	TypeRef string
}

func (node NewTypeNode) String() string {
	return "type " + node.Name + " = " + node.TypeRef
}

type ModuleAnnoNode struct {
	MyToken
	TypeRef string
}
type DeclAnnoNode struct {
	MyToken
	DeclRef string
	TypeRef string
}
type FieldAnnoNode struct {
	MyToken
	DeclRef  string
	FieldRef string
	TypeRef  string
}
type TypeParamNode struct {
	MyToken
	Params []string
}

func (node TypeParamNode) String() string {
	return "<" + strings.Join(node.Params, ",") + ">"
}

type TypeExprNode struct {
	MyToken
}
type FieldNode struct {
	MyToken
	TypeRef string
	Name    string
}
type JsonNode struct {
	MyToken
}
type JsonStrNode struct {
	MyToken
	Value string
}
type JsonBoolNode struct {
	MyToken
	Value bool
}
type JsonNullNode struct {
	MyToken
}
type JsonIntNode struct {
	MyToken
	Value int64
}
type JsonFloatNode struct {
	MyToken
	Value float64
}
type JsonArrayNode struct {
	MyToken
}
type JsonObjNode struct {
	MyToken
}
