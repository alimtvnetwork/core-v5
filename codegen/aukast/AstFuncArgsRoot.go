package aukast

import (
	"go/ast"
)

type AstFuncArgsRoot struct {
	Parent          *AstElem
	FuncType        *ast.FuncType
	Params, Results *AstElem
	InArgs, OutArgs []AstParam
	Code            string
}

func (it *AstFuncArgsRoot) IsValid() bool {
	return !it.IsInvalid()
}

func (it *AstFuncArgsRoot) IsInvalid() bool {
	return it == nil || it.FuncType == nil
}

func (it *AstFuncArgsRoot) HasInArgs() bool {
	return it.IsValid() && len(it.InArgs) > 0
}

func (it *AstFuncArgsRoot) HasOutArgs() bool {
	return it.IsValid() && len(it.OutArgs) > 0
}

func (it *AstFuncArgsRoot) IsEmptyInArgs() bool {
	return !it.HasInArgs()
}

func (it *AstFuncArgsRoot) IsEmptyOutArgs() bool {
	return !it.HasOutArgs()
}

func (it *AstFuncArgsRoot) String() any {

}
