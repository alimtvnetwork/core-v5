package aukast

import "go/ast"

type AstFunction struct {
	Name           string
	StructVarName  string
	StructName     string
	IsAttached     bool
	IsPublic       bool
	IsPrivate      bool
	FieldsCount    int
	Parent         *AstElem
	ReceiverStruct *AstElem
	Comments       *AstElem
	Type           *ast.FuncType
	FuncArg        *RootFuncArgs
	Code           string
}

func (it *AstFunction) IsValid() bool {
	return !it.IsInvalid()
}

func (it *AstFunction) IsInvalid() bool {
	return it == nil || it.Name == "" || it.Type == nil
}
