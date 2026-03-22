package coredynamic

import (
	"reflect"
)

func SafeZeroSet(rv reflect.Value) {
	if !rv.IsValid() {
		return
	}

	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return
	}

	elem := rv.Elem()

	if !elem.IsValid() || !elem.CanSet() {
		return
	}

	elem.Set(reflect.Zero(elem.Type()))
}
