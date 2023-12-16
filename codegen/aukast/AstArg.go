package aukast

import (
	"go/ast"
	"strings"
)

type AstFuncArg struct {
	Parent          *AstElem
	FuncType        *ast.FuncType
	Params, Results *AstElem
	InArgs, OutArgs []Arg
}

type Arg struct {
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
		toArgs := NewAstArgs(fullCode, field)
		inArgs = append(inArgs, toArgs...)
	}

	for _, field := range outParams.FieldsList() {
		toArgs := NewAstArgs(fullCode, field)
		outArgs = append(outArgs, toArgs...)
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
	code string,
	f *ast.Field,
) []Arg {
	if f == nil {
		return []Arg{}
	}

	var args []Arg
	subCode := astUtil.NodeToStringSafe(code, f)

	for _, ident := range f.Names {
		typeIdent := astUtil.ExprToIdent(f.Type)
		typeName, _ := astUtil.NodeToString(code, f.Type)
		isArray := astUtil.HasAnyPrefix(typeName, "[]", "*[]")
		isArrayPointerElement := astUtil.HasAnyPrefix(typeName, "*[]*", "[]*")

		a := Arg{
			Name:                  ident.Name,
			TypeName:              typeName,
			NameIdent:             ident,
			TypeExpr:              f.Type,
			TypeIdent:             typeIdent,
			Comment:               f.Comment,
			IsPointerType:         strings.HasPrefix(typeName, "*"),
			IsArray:               isArray,
			IsArrayPointerElement: isArrayPointerElement,
			Code:                  subCode,
		}

		args = append(args, a)
	}

	return args
}
