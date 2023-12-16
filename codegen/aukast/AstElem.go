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
	properties            map[string]bool
	childNodes            []AstElem
}

func (it *AstElem) IsEmpty() bool {
	return it == nil ||
		it.GenericNode == nil ||
		it.RealNode == nil
}

func (it *AstElem) IsInvalid() bool {
	return it.IsEmpty()
}

func (it *AstElem) IsValid() bool {
	return !it.IsEmpty()
}

func (it *AstElem) ChildNodes() []AstElem {
	ast.Inspect(
		it.GenericNode, func(n ast.Node) bool {
			if n == nil {
				return true
			}

			return true
		},
	)

}
