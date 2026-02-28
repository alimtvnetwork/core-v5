package isany

import "reflect"

func FloatingPointType(anyItem any) bool {
	return FloatingPointTypeRv(reflect.ValueOf(anyItem))
}
