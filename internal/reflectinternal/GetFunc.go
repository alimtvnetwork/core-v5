package reflectinternal

import (
	"reflect"
	"runtime"
)

func GetFunc(i interface{}) *runtime.Func {
	if IsNull(i) {
		return nil
	}

	rv := reflect.ValueOf(i)

	if rv.Kind() != reflect.Func {
		return nil
	}

	return runtime.FuncForPC(rv.Pointer())
}
