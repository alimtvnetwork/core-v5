package aukast

import (
	"go/ast"

	"gitlab.com/auk-go/core/errcore"
)

type newAstElemCreator struct{}

func (it newAstElemCreator) Create(fullCode string, node ast.Node) (*AstElem, error) {
	if node == nil {
		return nil, errcore.FailedToParseType.ErrorNoRefs("node is nil")
	}

	code, _ := astUtil.NodeToString(fullCode, node)
	typeName := astUtil.TypeName(node)
	name := astUtil.Name(fullCode, node)
	identifier := astUtil.ToIdent(node)

	return &AstElem{
		Node:           node,
		Name:           name,
		TypeName:       typeName,
		Code:           code,
		NameIdentifier: identifier,
		ParentTypeName: "",
		InnerTypeName:  "",
	}, nil
}

func (it newAstElemCreator) CreateByAstReader(astReader *AstReader, node ast.Node) (*AstElem, error) {
	if node == nil {
		return nil, errcore.FailedToParseType.ErrorNoRefs("node is nil")
	}

	if astReader == nil {
		return nil, errcore.FailedToParseType.ErrorNoRefs("astReader is nil")
	}

	if astReader.HasError() {
		return nil, errcore.FailedToParseType.ErrorNoRefs("astReader has error : " + astReader.parseErr.Error())
	}

	return it.Create(astReader.fullCode, node)
}
