package aukast

import (
	"fmt"
	"reflect"

	"gitlab.com/auk-go/core/coredata/corestr"
)

type AstCollection struct {
	Parent     *AstElem
	childNodes []AstElem
}

func (it *AstCollection) IsEmpty() bool {
	return it == nil ||
		it.Parent == nil ||
		it.Parent.astReader.IsEmpty() ||
		len(it.childNodes) == 0
}

func (it *AstCollection) IsInvalid() bool {
	return it.IsEmpty()
}

func (it *AstCollection) IsValid() bool {
	return !it.IsEmpty()
}

func (it *AstCollection) FullCode() string {
	if it.IsEmpty() {
		return ""
	}

	return it.Parent.FullCode()
}

func (it *AstCollection) RawChildNodes() []AstElem {
	if it.IsEmpty() {
		return []AstElem{}
	}

	return it.childNodes
}

func (it *AstCollection) RawChildNodesStrings() corestr.SimpleSlice {
	if it.IsEmpty() {
		return []string{}
	}

	slice := corestr.New.SimpleSlice.ByLen(it.childNodes)

	for _, elem := range it.RawChildNodes() {
		slice.AddSplit("- "+elem.String(), "\n")
		slice.Add("")
	}

	return slice.NonPtr()
}

// Filter
//
// It will only filter the child nodes (flat - not nested) that returns true in the filter function.
func (it *AstCollection) Filter(isFilterFunc AstWithBreakFilterFunc) *AstCollection {
	if it.IsEmpty() {
		return nil
	}

	var slice []AstElem

	for _, elem := range it.childNodes {
		isTake, isBreak := isFilterFunc(&elem)

		if isTake {
			slice = append(slice, elem)
		}

		if isBreak {
			break
		}
	}

	return &AstCollection{
		Parent:     it.Parent,
		childNodes: slice,
	}
}

func (it *AstCollection) FilterMatchTypesOf(
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

func (it AstCollection) String() string {
	if it.IsEmpty() {
		return ""
	}

	slice := it.RawChildNodesStrings()
	toJoin := slice.TranspileJoin(
		func(s string) string {
			return "  " + s
		}, "\n  ",
	)

	toStr := fmt.Sprintf(
		"\n  Parent:\n"+
			"  %s\n\n"+
			"  Childs:%d\n"+
			"  Childs:\n"+
			"  %s",
		it.Parent.String(),
		slice.Count(),
		toJoin,
	)

	return fmt.Sprintf("AstCollection {\n%s\n}", toStr)
}
