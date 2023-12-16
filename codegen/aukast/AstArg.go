package aukast

import "go/ast"

type AstFuncArg struct {
	Parent          *AstElem
	FuncType        *ast.FuncType
	Params, Results *AstElem
	InArgs, OutArgs []Arg
}

type Arg struct {
	Name, TypeName string
	NameIdent      *ast.Ident
	TypeExpr       *ast.Ident
	Comment        *ast.CommentGroup
}

func NewAstFuncArg(
	parent *AstElem,
	fullCode string,
	f *ast.FuncType,
) *AstFuncArg {
	if f == nil {
		return nil
	}

	p, _ := New.AstElem.CreateByParent(
		parent,
		fullCode,
		f.Params,
	)

	r, _ := New.AstElem.CreateByParent(
		parent,
		fullCode,
		f.Results,
	)

	var inArgs, outArgs []Arg

	for i, i := range p {

	}

	return &AstFuncArg{
		Parent:   nil,
		FuncType: nil,
		Params:   p,
		Results:  r,
		InArgs:   nil,
		OutArgs:  nil,
	}
}
