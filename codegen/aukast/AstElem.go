package aukast

import (
	"fmt"
	"go/ast"

	"gitlab.com/auk-go/core/errcore"
)

type AstElem struct {
	astReader         *AstReader
	Parent            *AstElem
	Name              string
	TypeName          string
	Code              string
	NameIdentifier    *ast.Ident
	Node              ast.Node
	childNodes        *AstCollection
	astFuncCollection *AstFuncCollection
}

func (it *AstElem) IsFieldList() bool {
	if it.IsEmpty() {
		return false
	}

	_, isOkay := it.Node.(*ast.FieldList)

	return isOkay
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

func (it *AstElem) Functions() (*AstFuncCollection, error) {
	if it.IsEmpty() {
		return nil, nil
	}

	if it.astFuncCollection != nil {
		return it.astFuncCollection, nil
	}

	astFuncCollection, err := New.AstFuncCollection.Create(
		it.AstReader(),
		it.Node,
	)

	it.astFuncCollection = astFuncCollection

	return astFuncCollection, err
}

func (it *AstElem) FieldsList() []*ast.Field {
	if it.IsEmpty() || !it.IsFieldList() {
		return []*ast.Field{}
	}

	fieldsList, isOkay := it.Node.(*ast.FieldList)

	if isOkay {
		return fieldsList.List
	}

	return nil
}

func (it *AstElem) String() string {
	if it == nil || it.IsInvalid() {
		return ""
	}

	return fmt.Sprintf(
		" - Name: %s, Type: %s, Code (20) : %s",
		it.Name, it.TypeName, it.CodeTakeMax(20),
	)
}

func (it *AstElem) CodeTakeMax(charsCount int) string {
	if it.IsEmpty() {
		return ""
	}

	if len(it.Code) > charsCount {
		return it.Code[:charsCount]
	}

	return it.Code
}
