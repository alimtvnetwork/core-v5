package coredynamic

import "reflect"

func AnyToReflectVal(item any) reflect.Value {
	return reflect.ValueOf(item)
}
