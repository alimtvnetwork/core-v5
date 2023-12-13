package args

import (
	"reflect"

	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type newFuncWrapCreator struct{}

func (it newFuncWrapCreator) Default(anyFunc interface{}) *FuncWrap {
	if reflectinternal.Is.Null(anyFunc) {
		return &FuncWrap{
			Func:      anyFunc,
			isInvalid: true,
		}
	}

	switch v := anyFunc.(type) {
	case *FuncWrap:
		return v
	case FuncWrapGetter:
		return v.FuncWrap()
	}

	typeOf := reflect.TypeOf(anyFunc)
	kind := typeOf.Kind()

	if kind != reflect.Func {
		// invalid

		return &FuncWrap{
			Func:      anyFunc,
			isInvalid: true,
			rvType:    typeOf,
		}
	}

	// valid
	fullName, nameOnly := reflectinternal.
		GetFunc.
		FullNameWithName(anyFunc)

	return &FuncWrap{
		Name:      nameOnly,
		FullName:  fullName,
		Func:      anyFunc,
		isInvalid: false,
		rvType:    typeOf,
		rv:        reflect.ValueOf(anyFunc),
	}
}

func (it newFuncWrapCreator) Single(
	anyFunc interface{},
) *FuncWrap {
	return it.Default(anyFunc)
}

func (it newFuncWrapCreator) Map(
	anyFunctions ...interface{},
) map[string]*FuncWrap {
	if len(anyFunctions) == 0 {
		return map[string]*FuncWrap{}
	}

	newMap := make(
		map[string]*FuncWrap,
		len(anyFunctions),
	)

	for _, function := range anyFunctions {
		v := it.Default(function)

		if v.IsValid() {
			newMap[v.GetFuncName()] = v
		}
	}

	return newMap
}

func (it newFuncWrapCreator) Many(
	anyFunctions ...interface{},
) []*FuncWrap {
	if len(anyFunctions) == 0 {
		return []*FuncWrap{}
	}

	slice := make(
		[]*FuncWrap,
		len(anyFunctions),
	)

	for i, function := range anyFunctions {
		v := it.Default(function)

		slice[i] = v
	}

	return slice
}

func (it newFuncWrapCreator) StructToMap(
	i interface{},
) map[string]*FuncWrap {
	if len(anyFunctions) == 0 {
		return []*FuncWrap{}
	}

	slice := make(
		[]*FuncWrap,
		len(anyFunctions),
	)

	for i, function := range anyFunctions {
		v := it.Default(function)

		slice[i] = v
	}

	return slice
}
