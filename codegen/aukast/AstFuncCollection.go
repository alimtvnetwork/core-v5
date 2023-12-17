package aukast

import "gitlab.com/auk-go/core/coredata/corestr"

type AstFuncCollection struct {
	Names  *corestr.SimpleSlice
	Map    map[string]AstFunction
	Parent *AstElem
}

func (it *AstFuncCollection) IsEmpty() bool {
	return it == nil || len(it.Map) == 0
}

func (it *AstFuncCollection) Count() int {
	if it.IsEmpty() {
		return 0
	}

	return len(it.Map)
}
