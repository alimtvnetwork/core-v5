package aukast

import (
	"gitlab.com/auk-go/core/coredata/corestr"
)

type AstFuncMap map[string]AstFunction

func (it *AstFuncMap) IsEmpty() bool {
	return it == nil || len(*it) == 0
}

func (it *AstFuncMap) HasAnyItem() bool {
	return it != nil && len(*it) > 0
}

func (it *AstFuncMap) Length() int {
	if it.IsEmpty() {
		return 0
	}

	return len(*it)
}

func (it *AstFuncMap) AddsValues(astFunctions ...AstFunction) *AstFuncMap {
	if it == nil {
		it = new(AstFuncMap)
	}

	if len(astFunctions) == 0 {
		return it
	}

	for _, function := range astFunctions {
		(*it)[function.Name] = function
	}

	return it
}

func (it *AstFuncMap) Adds(astFunctions ...*AstFunction) *AstFuncMap {
	if it.IsEmpty() {
		it = new(AstFuncMap)
	}

	if len(astFunctions) == 0 {
		return it
	}

	for _, function := range astFunctions {
		(*it)[function.Name] = *function
	}

	return it
}

func (it AstFuncMap) Get(name string) *AstFunction {
	if it.IsEmpty() {
		return nil
	}

	toFunc, has := it[name]

	if has {
		return &toFunc
	}

	return nil
}

func (it *AstFuncMap) IsContains(name string) bool {
	if it.IsEmpty() {
		return false
	}

	_, has := (*it)[name]

	return has
}

func (it AstFuncMap) IsMissing(name string) bool {
	if it.IsEmpty() {
		return true
	}

	_, has := it[name]

	return !has
}

func (it *AstFuncMap) AllFunctions() []AstFunction {
	var slice []AstFunction

	if it.IsEmpty() {
		return slice
	}

	for _, f := range *it {
		slice = append(slice, f)
	}

	return slice
}

func (it *AstFuncMap) FuncNames() *corestr.SimpleSlice {
	if it.IsEmpty() {
		return corestr.Empty.SimpleSlice()
	}

	slice := corestr.New.SimpleSlice.ByLen(*it)

	for key := range *it {
		slice.Add(key)
	}

	return slice
}

func (it *AstFuncMap) SortedFuncNames() *corestr.SimpleSlice {
	slice := it.FuncNames()

	slice.Sort()

	return slice
}

func (it AstFuncMap) StructFunctions() map[string]AstFuncMap {
	if it.IsEmpty() {
		return map[string]AstFuncMap{}
	}

	newMap := map[string]AstFuncMap{}

	for funcName, function := range it {
		if function.IsInvalid() || function.IsNotAttached() {
			continue
		}

		existingMap, isFound := newMap[function.StructName]

		if isFound {
			existingMap[funcName] = function
		} else {
			astMap := map[string]AstFunction{}
			astMap[funcName] = function

			newMap[function.StructName] = astMap
		}
	}

	return newMap
}

func (it AstFuncMap) StructFunc(structName string) AstFuncMap {
	if it.IsEmpty() {
		return map[string]AstFunction{}
	}

	newMap := make(map[string]AstFunction, it.Length())

	for funcName, function := range it {
		if !function.IsAttachToStructOf(structName) {
			continue
		}

		newMap[funcName] = function
	}

	return newMap
}

func (it AstFuncMap) String() string {
	if it.IsEmpty() {
		return ""
	}

	slice := corestr.New.SimpleSlice.ByLen(it)

	for s, function := range it {
		slice.AppendFmt(
			"\"%s\":%s",
			s,
			function.DefCode,
		)
	}

	return slice.JoinCsvLine()
}
