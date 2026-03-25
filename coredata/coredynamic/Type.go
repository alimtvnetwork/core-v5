package coredynamic

import "reflect"

func Type(any any) reflect.Type {
	return reflect.TypeOf(any)
}
