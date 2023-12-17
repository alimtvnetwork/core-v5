package aukast

import "gitlab.com/auk-go/core/coredata/corestr"

type AstFuncCollection struct {
	Names  *corestr.SimpleSlice
	Map    AstFuncMap
	Parent *AstElem
}

func (it *AstFuncCollection) IsEmpty() bool {
	return it == nil || len(it.Map) == 0
}

func (it *AstFuncCollection) IsValid() bool {
	return it != nil && len(it.Map) > 0
}

func (it *AstFuncCollection) IsInvalid() bool {
	return it.IsEmpty()
}

func (it *AstFuncCollection) Count() int {
	if it.IsEmpty() {
		return 0
	}

	return len(it.Map)
}
