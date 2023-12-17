package aukast

import (
	"go/ast"

	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/isany"
)

type newAstElemCreator struct{}

func (it newAstElemCreator) Create(
	astReader *AstReader,
	fullCode string,
	node ast.Node,
) (*AstElem, error) {
	if isany.Null(node) {
		return nil, errcore.FailedToParseType.ErrorNoRefs("node is nil")
	}

	code, _ := astUtil.NodeToString(fullCode, node)
	typeName := astUtil.TypeName(node)
	name := astUtil.Name(fullCode, node)
	identifier := astUtil.ToIdent(node)

	return &AstElem{
		astReader:      astReader,
		Name:           name,
		TypeName:       typeName,
		Code:           code,
		NameIdentifier: identifier,
		ParentTypeName: "",
		InnerTypeName:  "",
		Node:           node,
		parentType:     nil,
		innerType:      nil,
		properties:     nil,
		childNodes:     nil,
	}, nil
}

func (it newAstElemCreator) CreateByParent(
	parent *AstElem,
	fullCode string,
	node ast.Node,
) (*AstElem, error) {
	if isany.Null(node) {
		return nil, errcore.FailedToParseType.ErrorNoRefs("node is nil")
	}

	code, _ := astUtil.NodeToString(fullCode, node)
	typeName := astUtil.TypeName(node)
	name := astUtil.Name(fullCode, node)
	identifier := astUtil.ToIdent(node)

	return &AstElem{
		astReader:      parent.AstReader(),
		Parent:         parent,
		Name:           name,
		TypeName:       typeName,
		Code:           code,
		NameIdentifier: identifier,
		ParentTypeName: "",
		InnerTypeName:  "",
		Node:           node,
	}, nil
}

func (it newAstElemCreator) CreateByAstReader(astReader *AstReader, node ast.Node) (*AstElem, error) {
	if isany.Null(node) {
		return nil, errcore.FailedToParseType.ErrorNoRefs("node is nil")
	}

	if astReader == nil {
		return nil, errcore.FailedToParseType.ErrorNoRefs("AstReader is nil")
	}

	fullCode, err := astReader.FullCode()

	if err != nil {
		return nil, err
	}

	return it.Create(
		astReader,
		fullCode,
		node,
	)
}
