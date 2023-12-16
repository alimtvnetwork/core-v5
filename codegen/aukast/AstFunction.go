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
	FuncArg        *AstFuncArg
}
