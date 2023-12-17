package aukast

type AstFuncMap map[string]AstFunction

func (it AstFuncMap) IsEmpty() bool {
	return it == nil || len(it) == 0
}

func (it AstFuncMap) HasAnyItem() bool {
	return it != nil && len(it) > 0
}

func (it AstFuncMap) Length() int {
	return len(it)
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
	if it == nil {
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

func (it AstFuncMap) Contains(name string) bool {
	if it.IsEmpty() {
		return false
	}

	_, has := it[name]

	return has
}

func (it AstFuncMap) IsMissing(name string) bool {
	if it.IsEmpty() {
		return true
	}

	_, has := it[name]

	return !has
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
