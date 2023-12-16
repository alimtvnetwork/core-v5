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

	inParams, _ := New.AstElem.CreateByParent(
		parent,
		fullCode,
		f.Params,
	)

	outParams, _ := New.AstElem.CreateByParent(
		parent,
		fullCode,
		f.Results,
	)

	var inArgs, outArgs []Arg

	for _, field := range inParams.FieldsList() {
		toArgs := NewAstArgs(field)
		inArgs = append(inArgs, toArgs...)
	}

	for _, field := range outParams.FieldsList() {
		toArgs := NewAstArgs(field)
		outArgs = append(inArgs, toArgs...)
	}

	return &AstFuncArg{
		Parent:   parent,
		FuncType: f,
		Params:   inParams,
		Results:  outParams,
		InArgs:   inArgs,
		OutArgs:  outArgs,
	}
}

func NewAstArgs(
	f *ast.Field,
) []Arg {
	if f == nil {
		return []Arg{}
	}

	var args []Arg

	for _, ident := range f.Names {
		typeIdent := f.Type.(*ast.Ident)

		a := Arg{
			Name:      ident.Name,
			TypeName:  typeIdent.Name,
			NameIdent: ident,
			TypeExpr:  typeIdent,
			Comment:   f.Comment,
		}

		args = append(args, a)
	}

	return args
}
