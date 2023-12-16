package aukast

import (
	"go/ast"
)

type RootFuncArgs struct {
	Parent          *AstElem
	FuncType        *ast.FuncType
	Params, Results *AstElem
	InArgs, OutArgs []Param
	Code            string
}

func (it *RootFuncArgs) IsValid() bool {
	return !it.IsInvalid()
}

func (it *RootFuncArgs) IsInvalid() bool {
	return it == nil || it.FuncType == nil
}

func (it *RootFuncArgs) HasInArgs() bool {
	return it.IsValid() && len(it.InArgs) > 0
}

func (it *RootFuncArgs) HasOutArgs() bool {
	return it.IsValid() && len(it.OutArgs) > 0
}

func (it *RootFuncArgs) IsEmptyInArgs() bool {
	return !it.HasInArgs()
}

func (it *RootFuncArgs) IsEmptyOutArgs() bool {
	return !it.HasOutArgs()
}
