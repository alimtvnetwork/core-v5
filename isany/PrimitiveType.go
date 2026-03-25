package isany

import "reflect"

// PrimitiveType
//
// function returns true if the kind passed to it is one of the
// primitive types (boolean, int, uint, float, string)
func PrimitiveType(anyItem any) bool {
	return PrimitiveTypeRv(reflect.ValueOf(anyItem).Kind())
}
