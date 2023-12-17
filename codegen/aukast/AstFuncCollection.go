package aukast

import "gitlab.com/auk-go/core/coredata/corestr"

type AstFuncCollection struct {
	Names         *corestr.SimpleSlice
	Map           AstFuncMap
	Parent        *AstElem
	structFuncMap *AstStructFuncMap
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

func (it *AstFuncCollection) StructFuncMap() *AstStructFuncMap {
	if it.IsEmpty() {
		return new(AstStructFuncMap)
	}

	if it.structFuncMap != nil {
		return it.structFuncMap
	}

	m := it.Map.StructFunctions()
	it.structFuncMap = &m

	return it.structFuncMap
}

func (it *AstFuncCollection) StructFuncMapOf(structName string) *AstFuncMap {
	if it.IsEmpty() {
		return new(AstFuncMap)
	}

	m := it.Map.StructFunc(structName)

	return &m
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

func (it *AstFuncCollection) FuncNamesWithTypeStrings() corestr.SimpleSlice {
	if it.IsEmpty() {
		return []string{}
	}

	slice := corestr.New.SimpleSlice.Cap(len(it.Map))

	for _, name := range *it.Names {
		slice.AppendFmt(
			" - Func: %s \n  - %s \n",
			name,
			it.GetFunc(name).DefCode,
		)
	}

	return slice.NonPtr()
}

func (it *AstFuncCollection) String() string {
	if it.IsEmpty() {
		return ""
	}

	funcNames := it.FuncNamesWithTypeStrings()

	return "AstFuncCollection {\n" + funcNames.JoinLine() + "\n}"
}
