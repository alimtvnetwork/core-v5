package aukast

import (
	"go/ast"

	"gitlab.com/auk-go/core/isany"
)

type newAstFunctionCreator struct{}

func (it newAstFunctionCreator) Create(
	astReader *AstReader,
	fullCode string,
	n ast.Node,
) (*AstFunction, error) {
	if isany.Null(n) {
		return nil, nil
	}

	toFunc, isOkay := n.(*ast.FuncDecl)

	if !isOkay || toFunc == nil {
		return nil, nil
	}

	// https://prnt.sc/eQZm-iCDdj-H
	parentElem, err := New.AstElem.Create(astReader, fullCode, n)

	if err != nil {
		return nil, err
	}

	name := astUtil.Name(fullCode, toFunc)
	StructName := astUtil.Name(fullCode, toFunc.Recv)
	StructTypeName := astUtil.NodeTypeName(fullCode, toFunc.Recv)
	receiver, _ := New.AstElem.Create(astReader, fullCode, toFunc.Recv)
	comments, _ := New.AstElem.Create(astReader, fullCode, toFunc.Doc)
	funcArgs := New.ArgsParams.Root(parentElem, fullCode, toFunc.Type)

	return &AstFunction{
		Name:           name,
		StructVarName:  StructName,
		StructName:     StructTypeName,
		IsAttached:     false,
		IsPublic:       true,
		IsPrivate:      false,
		FieldsCount:    toFunc.Recv.NumFields(),
		Parent:         parentElem,
		ReceiverStruct: receiver,
		Comments:       comments,
		Type:           toFunc.Type,
		FuncArg:        funcArgs,
		Code:           astUtil.NodeToStringSafe(fullCode, n),
	}, nil
}
