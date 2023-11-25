package reflectinternal

import "reflect"

func IsFunc(item interface{}) bool {
	if item == nil {
		return true
	}

	typeOf := reflect.TypeOf(item)

	return IsFuncTypeOf(typeOf)
}

func IsFuncTypeOf(typeOf reflect.Type) bool {
	kind := typeOf.Kind()

	switch kind {
	case reflect.Func:
		return true
	}

	return false
}
