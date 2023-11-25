package coretests

import "gitlab.com/auk-go/core/internal/reflectinternal"

type FuncWrap struct {
	Name string
	Func interface{}
}

func (it FuncWrap) FuncName() string {
	return it.Name
}

func (it *FuncWrap) HasValidFunc() bool {
	return it != nil && reflectinternal.IsFunc(it.Func)
}

func (it *FuncWrap) Invoke(items ...interface{}) (results []interface{}) {
	return it != nil && reflectinternal.IsFunc(it.Func)
}
