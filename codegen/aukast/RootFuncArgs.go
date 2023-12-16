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
