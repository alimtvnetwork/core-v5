package aukast

import (
	"go/ast"

	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/isany"
)

type AstFilter struct {
	AstReader  *AstReader
	ParentNode ast.Node
	Node       ast.Node
	fullCode   string
}

func (it *AstFilter) Filter(
	node ast.Node,
	filter AstWithBreakFilterFunc,
) *AstCollection {
	if isany.Null(node) {
		return nil
	}

	creatorFunc := New.AstElem.Create
	fullCode := it.FullCode()
	var slice []AstElem
	var rawErr errcore.RawErrCollection

	ast.Inspect(
		node, func(n ast.Node) bool {
			if isany.Null(n) {
				return true
			}

			elem, err := creatorFunc(it.AstReader, fullCode, n)
			rawErr.Add(err)
			isTake, isBreak := filter(elem)

			if err == nil && isTake {
				slice = append(slice, *elem)
			}

			if isBreak {
				return false
			}

			return true
		},
	)

	parent, _ := creatorFunc(it.AstReader, fullCode, node)

	collection := &AstCollection{
		Parent:     parent,
		childNodes: slice,
	}

	return collection
}

func (it AstFilter) CompositeLitKeyValExprs() *AstCollection {
	return it.Filter(
		it.Node,
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

func (it AstFilter) KeyValExprs() *AstCollection {
	return it.Filter(
		it.Node,
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
	return it.Filter(
		it.Node,
		func(elem *AstElem) (isTake, isBreak bool) {
			_, isOkay := elem.Node.(*ast.ImportSpec)

			if !isOkay {
				return false, false
			}

			return true, false
		},
	)
}

func (it AstFilter) FuncTypes() *AstCollection {
	return it.Filter(
		it.Node,
		func(elem *AstElem) (isTake, isBreak bool) {
			_, isOkay := elem.Node.(*ast.FuncType)

			if !isOkay {
				return false, false
			}

			return true, false
		},
	)
}

func (it AstFilter) FuncDecls() *AstCollection {
	return it.Filter(
		it.Node,
		func(elem *AstElem) (isTake, isBreak bool) {
			_, isOkay := elem.Node.(*ast.FuncDecl)

			if !isOkay {
				return false, false
			}

			return true, false
		},
	)
}

func (it AstFilter) CallExprs() *AstCollection {
	return it.Filter(
		it.Node,
		func(elem *AstElem) (isTake, isBreak bool) {
			_, isOkay := elem.Node.(*ast.FuncDecl)

			if !isOkay {
				return false, false
			}

			return true, false
		},
	)
}

func (it AstFilter) RangeStmts() *AstCollection {
	return it.Filter(
		it.Node,
		func(elem *AstElem) (isTake, isBreak bool) {
			_, isOkay := elem.Node.(*ast.RangeStmt)

			if !isOkay {
				return false, false
			}

			return true, false
		},
	)
}

func (it AstFilter) Idents() *AstCollection {
	return it.Filter(
		it.Node,
		func(elem *AstElem) (isTake, isBreak bool) {
			_, isOkay := elem.Node.(*ast.Ident)

			if !isOkay {
				return false, false
			}

			return true, false
		},
	)
}

func (it AstFilter) AssignStmts() *AstCollection {
	return it.Filter(
		it.Node,
		func(elem *AstElem) (isTake, isBreak bool) {
			_, isOkay := elem.Node.(*ast.AssignStmt)

			if !isOkay {
				return false, false
			}

			return true, false
		},
	)
}

func (it AstFilter) SelectorExprs() *AstCollection {
	return it.Filter(
		it.Node,
		func(elem *AstElem) (isTake, isBreak bool) {
			_, isOkay := elem.Node.(*ast.SelectorExpr)

			if !isOkay {
				return false, false
			}

			return true, false
		},
	)
}

func (it AstFilter) FieldLists() *AstCollection {
	return it.Filter(
		it.Node,
		func(elem *AstElem) (isTake, isBreak bool) {
			_, isOkay := elem.Node.(*ast.FieldList)

			if !isOkay {
				return false, false
			}

			return true, false
		},
	)
}

func (it AstFilter) ChildNodes() *AstCollection {
	return it.Filter(
		it.Node,
		func(elem *AstElem) (isTake, isBreak bool) {
			return true, false
		},
	)
}

func (it *AstFilter) FullCode() string {
	if it.IsEmpty() {
		return ""
	}

	it.fullCode = it.AstReader.fullCode

	return it.fullCode
}

func (it *AstFilter) IsEmpty() bool {
	return it == nil || it.AstReader == nil || it.Node == nil
}
