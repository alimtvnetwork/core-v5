package aukast

import (
	"fmt"
	"go/ast"
	"reflect"

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

func (it *AstElem) Filter(filter AstWithBreakFilterFunc) *AstCollection {
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

	if it.Name == "" {
		return fmt.Sprintf(
			"Type: %s,\n"+
				"  Code (20) : `%s`",
			it.TypeName,
			it.CodeTakeMax(20),
		)
	}

	return fmt.Sprintf(
		"Name: %s, Type: %s,\n"+
			"  Code (20) : `%s`",
		it.Name,
		it.TypeName,
		it.CodeTakeMax(40),
	)
}

func (it *AstElem) CodeTakeMax(charsCount int) string {
	if it.IsEmpty() {
		return ""
	}

	return astUtil.MaxSubstringTrimSpaces(it.Code, charsCount)
}

// IsAnyNodeTypeMatches
//
//   - given nothing given returns true if and only if it is valid node (Not IsEmpty)
//   - if IsEmpty then always returns false
//   - if any type matches with current node type then returns true
func (it *AstElem) IsAnyNodeTypeMatches(matches ...reflect.Type) bool {
	if it.IsEmpty() {
		return false
	}

	if len(matches) == 0 {
		return true
	}

	currMatchType := reflect.TypeOf(it.Node)

	for _, match := range matches {
		if currMatchType == match {
			return true
		}
	}

	return false
}

func (it *AstElem) ChildOf(
	typeMatches ...reflect.Type,
) *AstCollection {
	return it.Filter(
		func(elem *AstElem) (isTake, isBreak bool) {
			if elem.IsAnyNodeTypeMatches(typeMatches...) {
				return true, false
			}

			return false, false
		},
	)
}
