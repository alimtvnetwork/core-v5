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

func (it *AstFuncCollection) GetFunc(name string) *AstFunction {
	if it.IsEmpty() {
		return nil
	}

	return it.Map.Get(name)
}

func (it *AstFuncCollection) IsContains(name string) bool {
	if it.IsEmpty() {
		return false
	}

	return it.Map.IsContains(name)
}

func (it *AstFuncCollection) String() string {
	if it.IsEmpty() {
		return ""
	}

	slice := corestr.New.SimpleSlice.Cap(len(it.Map))

	for i, name := range *it.Names {
		slice.AppendFmt(
			"%s:%s\n"+
				"  - %s",
			"Func Name",
			name,
		)
	}
}
