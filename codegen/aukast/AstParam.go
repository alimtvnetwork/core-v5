package aukast

import "go/ast"

type AstParam struct {
	Name, TypeName        string
	NameIdent             *ast.Ident
	TypeExpr              ast.Expr
	TypeIdent             *ast.Ident
	Comment               *ast.CommentGroup
	IsPointerType         bool
	IsArray               bool
	IsArrayPointerElement bool
	Code                  string
}
