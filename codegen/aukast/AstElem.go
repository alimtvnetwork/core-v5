package aukast

import (
	"go/ast"
	"reflect"
)

type AstElem struct {
	Parent                *AstElem
	RealNode              interface{}
	Path                  *ast.BasicLit
	Name                  string
	TypeName              string
	Code                  string
	NameIdentifier        *ast.Ident
	ParentTypeName        string // https://prnt.sc/ZffPQKPrAh7m
	InnerTypeName         string // https://prnt.sc/OKbh4Q5JahXr
	GenericNode           ast.Node
	parentType, innerType reflect.Type
	IsFunction            bool
	IsIdent               bool
	IsStatement           bool
	IsExpression          bool
	IsBlock               bool
	HasParent             bool
}

func (it AstElem) name() {

}
