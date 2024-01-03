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

func (it *AstParam) CodeTakeMax(charsCount int) string {
	if it.IsEmpty() {
		return ""
	}

	return astUtil.MaxSubstringTrimSpaces(it.Code, charsCount)
}

func (it *AstParam) IsEmpty() bool {
	return it == nil || it.TypeIdent == nil || it.NameIdent == nil
}
