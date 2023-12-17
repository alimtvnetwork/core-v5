package aukast

import "gitlab.com/auk-go/core/coredata/corestr"

type AstStructFuncMap map[string]AstFuncMap

func (it *AstStructFuncMap) IsEmpty() bool {
	return it == nil || len(*it) == 0
}

func (it *AstStructFuncMap) HasAnyItem() bool {
	return it != nil && len(*it) > 0
}

func (it *AstStructFuncMap) Length() int {
	if it.IsEmpty() {
		return 0
	}

	return len(*it)
}

func (it *AstStructFuncMap) AddsValues(astFunctions ...AstFunction) *AstStructFuncMap {
	if it == nil {
		it = new(AstStructFuncMap)
	}

	if len(astFunctions) == 0 {
		return it
	}

	for _, function := range astFunctions {
		existingMap, isFound := (*it)[function.StructName]

		if isFound {
			existingMap[function.Name] = function
		} else {
			funcMap := map[string]AstFunction{}
			funcMap[function.Name] = function
			(*it)[function.StructName] = funcMap
		}
	}

	return it
}

func (it *AstStructFuncMap) Adds(astFunctions ...*AstFunction) *AstStructFuncMap {
	if it.IsEmpty() {
		it = new(AstStructFuncMap)
	}

	if len(astFunctions) == 0 {
		return it
	}

	for _, function := range astFunctions {
		if function == nil {
			continue
		}

		existingMap, isFound := (*it)[function.StructName]

		if isFound {
			existingMap[function.Name] = *function
		} else {
			funcMap := map[string]AstFunction{}
			funcMap[function.Name] = *function
			(*it)[function.StructName] = funcMap
		}
	}

	return it
}

func (it *AstStructFuncMap) AddsAstStructFuncMaps(
	astStructFuncMap ...*AstStructFuncMap,
) *AstStructFuncMap {
	if it.IsEmpty() {
		it = new(AstStructFuncMap)
	}

	if len(astStructFuncMap) == 0 {
		return it
	}

	for _, astStructMap := range astStructFuncMap {
		if astStructMap == nil || astStructMap.IsEmpty() {
			continue
		}

		allFunctions := astStructMap.AllFunctions()

		it.AddsValues(allFunctions...)
	}

	return it
}

func (it *AstStructFuncMap) AddsFuncMap(astFuncMaps ...*AstFuncMap) *AstStructFuncMap {
	if it.IsEmpty() {
		it = new(AstStructFuncMap)
	}

	if len(astFuncMaps) == 0 {
		return it
	}

	for _, astFuncMap := range astFuncMaps {
		if astFuncMap == nil || astFuncMap.IsEmpty() {
			continue
		}

		for _, function := range *astFuncMap {
			existingMap, isFound := (*it)[function.StructName]

			if isFound {
				existingMap[function.Name] = function
			} else {
				funcMap := map[string]AstFunction{}
				funcMap[function.Name] = function
				(*it)[function.StructName] = funcMap
			}
		}
	}

	return it
}

func (it *AstStructFuncMap) AddFuncMap(astFuncMap *AstFuncMap) *AstStructFuncMap {
	if it.IsEmpty() {
		it = new(AstStructFuncMap)
	}

	if astFuncMap.IsEmpty() {
		return it
	}

	for _, function := range *astFuncMap {
		existingMap, isFound := (*it)[function.StructName]

		if isFound {
			existingMap[function.Name] = function
		} else {
			funcMap := map[string]AstFunction{}
			funcMap[function.Name] = function
			(*it)[function.StructName] = funcMap
		}
	}

	return it
}

func (it AstStructFuncMap) Get(name string) *AstFuncMap {
	if it.IsEmpty() {
		return nil
	}

	toFunc, has := it[name]

	if has {
		return &toFunc
	}

	return nil
}

func (it *AstStructFuncMap) Contains(name string) bool {
	if it.IsEmpty() {
		return false
	}

	_, has := (*it)[name]

	return has
}

func (it AstStructFuncMap) IsMissing(name string) bool {
	if it.IsEmpty() {
		return true
	}

	_, has := it[name]

	return !has
}

func (it *AstStructFuncMap) FuncNames() *corestr.SimpleSlice {
	if it.IsEmpty() {
		return corestr.Empty.SimpleSlice()
	}

	slice := corestr.New.SimpleSlice.Cap(20)

	for _, funcMap := range *it {
		slice.Adds(funcMap.SortedFuncNames().List()...)
	}

	return slice
}

func (it *AstStructFuncMap) StructNames() *corestr.SimpleSlice {
	if it.IsEmpty() {
		return corestr.Empty.SimpleSlice()
	}

	slice := corestr.New.SimpleSlice.ByLen(*it)

	for key := range *it {
		slice.Add(key)
	}

	return slice
}

func (it *AstStructFuncMap) AllFunctions() []AstFunction {
	var slice []AstFunction

	if it.IsEmpty() {
		return slice
	}

	for _, funcMap := range *it {
		getAll := funcMap.AllFunctions()

		if len(getAll) == 0 {
			continue
		}

		slice = append(slice, getAll...)
	}

	return slice
}

func (it *AstStructFuncMap) SortedFuncNames() *corestr.SimpleSlice {
	slice := it.FuncNames()

	slice.Sort()

	return slice
}

func (it AstStructFuncMap) StructFunctions() AstStructFuncMap {
	if it.IsEmpty() {
		return map[string]AstFuncMap{}
	}

	return it
}

func (it AstStructFuncMap) StructFunc(structName string) *AstFuncMap {
	if it.IsEmpty() {
		return new(AstFuncMap)
	}

	return it.Get(structName)
}

func (it AstStructFuncMap) String() string {
	if it.IsEmpty() {
		return ""
	}

	slice := corestr.New.SimpleSlice.ByLen(it)

	for s, funcMap := range it {
		slice.AppendFmt(
			"\"%s\":%s",
			s,
			funcMap.String(),
		)
	}

	return slice.JoinCsvLine()
}
