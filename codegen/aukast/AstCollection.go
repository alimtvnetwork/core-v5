package aukast

import (
	"fmt"

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

func (it *AstCollection) RawChildNodesString() corestr.SimpleSlice {
	if it.IsEmpty() {
		return []string{}
	}

	slice := corestr.New.SimpleSlice.ByLen(it.childNodes)

	for i, elem := range it.RawChildNodes() {
		slice.Add(elem.String())
	}
}

func (it AstCollection) String() string {
	if it.IsEmpty() {
		return ""
	}

	return fmt.Sprintf(
		"Parent:%s" +
			"Childs:%d" +
			"Parent:%s" +
			"Parent:%s" +
			"Parent:%s" +
			"Parent:%s" +
			"Parent:%s" +
			"Parent:%s" +,
	)
}
