package aukast

import "go/ast"

type AstFilter struct {
	astReader *AstReader
	n         ast.Node
}

func (it AstFilter) CompositeLitKeyValExpr() *AstCollection {
	return it.astReader.Filter(
		func(elem *AstElem) (isTake, isBreak bool) {
			v, isOkay := elem.Node.(*ast.CompositeLit)

			if !isOkay {
				return false, false
			}

			_, isOkayVal := v.Elts[0].(*ast.KeyValueExpr)

			if !isOkayVal {
				return false, false
			}

			return true, false
		},
	)
}

func (it AstFilter) KeyValExpr() *AstCollection {
	return it.astReader.Filter(
		func(elem *AstElem) (isTake, isBreak bool) {
			_, isOkay := elem.Node.(*ast.KeyValueExpr)

			if !isOkay {
				return false, false
			}

			return true, false
		},
	)
}

func (it AstFilter) Imports() *AstCollection {
	return it.astReader.Filter(
		func(elem *AstElem) (isTake, isBreak bool) {
			_, isOkay := elem.Node.(*ast.ValueSpec)

			if !isOkay {
				return false, false
			}

			return true, false
		},
	)
}
