package results

import (
	"reflect"
	"runtime"
	"strings"
)

// MethodName extracts the method name from a direct function reference.
//
// Example:
//
//	MethodName((*MyStruct).IsValid) → "IsValid"
//	MethodName((*MyStruct).String)  → "String"
func MethodName(funcRef any) string {
	if funcRef == nil {
		return ""
	}

	rv := reflect.ValueOf(funcRef)

	if rv.Kind() != reflect.Func {
		return ""
	}

	fullName := runtime.FuncForPC(rv.Pointer()).Name()

	// fullName looks like: "github.com/alimtvnetwork/core/pkg.(*Type).Method-fm"
	// We want just "Method"

	lastDot := strings.LastIndex(fullName, ".")

	if lastDot < 0 {
		return fullName
	}

	name := fullName[lastDot+1:]

	// Strip Go's "-fm" suffix for method expressions
	name = strings.TrimSuffix(name, "-fm")

	return name
}
