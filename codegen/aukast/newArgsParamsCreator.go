package aukast

import (
	"go/ast"
	"strings"
)

type newArgsParamsCreator struct{}

func (it newArgsParamsCreator) Root(
	parent *AstElem,
	fullCode string,
	f *ast.FuncType,
) *AstFuncArgsRoot {
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

	var inArgs, outArgs []AstParam

	for _, field := range inParams.FieldsList() {
		toArgs := it.Params(fullCode, field)
		inArgs = append(inArgs, toArgs...)
	}

	for _, field := range outParams.FieldsList() {
		toArgs := it.Params(fullCode, field)
		outArgs = append(outArgs, toArgs...)
	}

	code := astUtil.NodeToStringSafe(fullCode, f)

	return &AstFuncArgsRoot{
		Parent:   parent,
		FuncType: f,
		Params:   inParams,
		Results:  outParams,
		InArgs:   inArgs,
		OutArgs:  outArgs,
		Code:     code,
	}
}

func (it newArgsParamsCreator) ParamsUsingFieldsList(
	fullCode string,
	fieldsList []*ast.Field,
) []AstParam {
	if len(fieldsList) == 0 || fullCode == "" {
		return []AstParam{}
	}

	var toParams []AstParam

	for _, field := range fieldsList {
		toArgs := it.Params(fullCode, field)
		toParams = append(toParams, toArgs...)
	}

	return toParams
}

func (it newArgsParamsCreator) Params(
	fullCode string,
	f *ast.Field,
) []AstParam {
	if f == nil {
		return []AstParam{}
	}

	var args []AstParam
	subCode := astUtil.NodeToStringSafe(fullCode, f)

	for _, ident := range f.Names {
		typeIdent := astUtil.ExprToIdent(f.Type)
		typeName, _ := astUtil.NodeToString(fullCode, f.Type)
		isArray := astUtil.HasAnyPrefix(typeName, "[]", "*[]")
		isArrayPointerElement := astUtil.HasAnyPrefix(typeName, "*[]*", "[]*")

		a := AstParam{
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
