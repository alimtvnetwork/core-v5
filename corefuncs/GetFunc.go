package corefuncs

import (
	"reflect"
	"runtime"

	"gitlab.com/auk-go/core/internal/reflectinternal"
)

func GetFunc(i interface{}) *runtime.Func {
	if reflectinternal.IsNull(i) {
		return nil
	}

	rv := reflect.ValueOf(i)

	if rv.Kind() != reflect.Func {
		return nil
	}

	return runtime.FuncForPC(rv.Pointer())
}
