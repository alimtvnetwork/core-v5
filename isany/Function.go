package isany

import (
	"reflect"
	"runtime"
)

// Function
//
// Returns false for nil, struct, anything else
func Function(item any) (isFunc bool, name string) {
	if item == nil {
		return false, ""
	}

	rv := reflect.ValueOf(item)

	if rv.Kind() != reflect.Func {
		return false, ""
	}

	if rv.IsNil() {
		return true, ""
	}

	name = runtime.FuncForPC(rv.Pointer()).Name()

	return name != "", name
}
