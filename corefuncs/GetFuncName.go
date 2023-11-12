package corefuncs

import (
	"reflect"
	"runtime"

	"gitlab.com/auk-go/core/internal/reflectinternal"
)

// GetFuncName
//
// Get the function name, passing non function may result panic
func GetFuncName(i interface{}) string {
	if reflectinternal.IsNull(i) {
		return ""
	}

	rv := reflect.ValueOf(i)

	if rv.Kind() != reflect.Func {
		return ""
	}

	return runtime.FuncForPC(rv.Pointer()).Name()
}
