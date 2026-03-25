package isany

import "reflect"

func NumberType(anyItem any) bool {
	return NumberTypeRv(reflect.ValueOf(anyItem))
}
