package coredynamic

import (
	"reflect"
)

func LengthOfReflect(reflectVal reflect.Value) int {
	if !reflectVal.IsValid() {
		return 0
	}

	if reflectVal.Kind() == reflect.Ptr {
		return LengthOfReflect(reflect.Indirect(reflectVal))
	}

	switch reflectVal.Kind() {
	case reflect.Slice:
		return reflectVal.Len()
	case reflect.Array:
		return reflectVal.Type().Len()
	case reflect.Map:
		return len(reflectVal.MapKeys())
	default:
		return 0
	}
}
