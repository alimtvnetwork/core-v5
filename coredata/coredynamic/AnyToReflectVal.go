package coredynamic

import "reflect"

func AnyToReflectVal(any interface{}) reflect.Value {
	return reflect.ValueOf(any)
}
