package isany

import "reflect"

func ReflectValueNull(rv reflect.Value) bool {
	if !rv.IsValid() {
		return true
	}

	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Slice:
		return rv.IsNil()
	default:
		return false
	}
}
