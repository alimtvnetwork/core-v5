package isany

import "reflect"

// ReflectNull
//
// # Returns true for any nil given
//
// Reference : https://stackoverflow.com/a/43896204
func ReflectNull(item any) bool {
	if item == nil {
		return true
	}

	if rv, ok := item.(reflect.Value); ok {
		return ReflectValueNull(rv)
	}

	rv := reflect.ValueOf(item)

	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Slice:
		return rv.IsNil()
	default:
		return false
	}
}
