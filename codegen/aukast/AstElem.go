package aukast

import (
	"go/ast"
	"reflect"

	"gitlab.com/auk-go/core/errcore"
)

type AstElem struct {
	astReader             *AstReader
	Parent                *AstElem
	Name                  string
	TypeName              string
	Code                  string
	NameIdentifier        *ast.Ident
	ParentTypeName        string // https://prnt.sc/ZffPQKPrAh7m
	InnerTypeName         string // https://prnt.sc/OKbh4Q5JahXr
	Node                  ast.Node
	parentType, innerType reflect.Type
	properties            map[string]bool
	childNodes            *AstCollection
}

func (it *AstElem) IsEmpty() bool {
	return it == nil ||
		it.Node == nil ||
		it.astReader == nil
}

func (it *AstElem) IsInvalid() bool {
	return it.IsEmpty()
}

func (it *AstElem) IsValid() bool {
	return !it.IsEmpty()
}

func (it *AstElem) AstReader() *AstReader {
	if it.IsEmpty() {
		return nil
	}

	return it.astReader
}

func (it *AstElem) FullCode() string {
	if it.IsEmpty() {
		return ""
	}

	return it.AstReader().fullCode
}

func (it *AstElem) ChildNodes() *AstCollection {
	if it.IsEmpty() {
		return nil
	}

	if it.childNodes != nil {
		return it.childNodes
	}

	creatorFunc := New.AstElem.CreateByParent
	fullCode := it.FullCode()
	var slice []AstElem
	var rawErr errcore.RawErrCollection

	ast.Inspect(
		it.Node, func(n ast.Node) bool {
			if n == nil {
				return true
			}

			elem, err := creatorFunc(it, fullCode, n)
			rawErr.Add(err)

			if err == nil {
				slice = append(slice, *elem)
			}

			return true
		},
	)

	collection := &AstCollection{
		Parent:     it,
		childNodes: slice,
	}

	it.childNodes = collection

	return it.childNodes
}

func (it *AstElem) Filter(filter func(elem *AstElem) (isTake, isBreak bool)) *AstCollection {
	if it.IsEmpty() {
		return nil
	}

	if it.childNodes != nil {
		return it.childNodes
	}

	creatorFunc := New.AstElem.CreateByParent
	fullCode := it.FullCode()
	var slice []AstElem
	var rawErr errcore.RawErrCollection

	ast.Inspect(
		it.Node, func(n ast.Node) bool {
			if n == nil {
				return true
			}

			elem, err := creatorFunc(it, fullCode, n)
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

	collection := &AstCollection{
		Parent:     it,
		childNodes: slice,
	}

	it.childNodes = collection

	return it.childNodes
}

func (it *AstElem) Functions() *AstCollection {
	if it.IsEmpty() {
		return nil
	}

	if it.childNodes != nil {
		return it.childNodes
	}

	creatorFunc := New.AstElem.CreateByParent
	fullCode := it.FullCode()
	var slice []AstElem
	var rawErr errcore.RawErrCollection

	ast.Inspect(
		it.Node, func(n ast.Node) bool {
			if n == nil {
				return true
			}

			elem, err := creatorFunc(it, fullCode, n)
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

	collection := &AstCollection{
		Parent:     it,
		childNodes: slice,
	}

	it.childNodes = collection

	return it.childNodes
}
