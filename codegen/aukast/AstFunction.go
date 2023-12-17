package aukast

import "go/ast"

type AstFunction struct {
	AstReader      *AstReader
	Name           string
	StructVarName  string
	StructName     string
	IsAttached     bool
	IsPublic       bool
	IsPrivate      bool
	ReceiverCount  int
	Parent         *AstElem
	ReceiverStruct *AstElem
	Comments       *AstElem
	Type           *ast.FuncType
	FuncArg        *AstFuncArgsRoot
	Code           string
}

func (it *AstFunction) IsNotAttached() bool {
	return it.IsValid() && !it.IsAttached
}

func (it *AstFunction) IsValid() bool {
	return !it.IsInvalid()
}

func (it *AstFunction) IsInvalid() bool {
	return it == nil || it.Name == "" || it.Type == nil
}

func (it *AstFunction) HasInArgs() bool {
	return it.IsValid() && len(it.FuncArg.InArgs) > 0
}

func (it *AstFunction) HasOutArgs() bool {
	return it.IsValid() && len(it.FuncArg.OutArgs) > 0
}

func (it *AstFunction) IsEmptyInArgs() bool {
	return !it.HasInArgs()
}

func (it *AstFunction) IsEmptyOutArgs() bool {
	return !it.HasOutArgs()
}

func (it *AstFunction) IsAttachToStructOf(structName string) bool {
	if it.IsInvalid() {
		return false
	}

	return it.StructName == structName
}
