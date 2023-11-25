package coretests

import (
	"errors"
	"reflect"

	"gitlab.com/auk-go/core/internal/reflectinternal"
)

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

func (it *FuncWrap) IsInvalid() bool {
	return it == nil || !it.HasValidFunc()
}

func (it *FuncWrap) Invoke(items ...interface{}) (results []interface{}, err error) {
	if it.IsInvalid() {
		return nil, errors.New("not a valid func")
	}

	rv := reflect.ValueOf(it.Func)

	rv.Call()
}
