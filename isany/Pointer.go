package isany

import "reflect"

func Pointer(anyItem any) (isPtr bool) {
	rv := reflect.ValueOf(anyItem)

	return rv.Kind() == reflect.Ptr
}
