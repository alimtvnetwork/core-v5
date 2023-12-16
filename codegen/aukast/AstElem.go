package aukast

import "go/ast"

type AstElem struct {
	RealNode       interface{}
	Path           *ast.BasicLit
	Name           string
	TypeName       string
	NameIdentifier *ast.Ident
	GenericNode    ast.Node
}
