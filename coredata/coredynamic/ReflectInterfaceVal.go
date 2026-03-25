package coredynamic

import "reflect"

// ReflectInterfaceVal
//
// Reduce pointer to value (one step only)
func ReflectInterfaceVal(item any) any {
	rVal := reflect.ValueOf(item)
	k := rVal.Kind()

	if k != reflect.Ptr && k != reflect.Interface {
		return rVal.Interface()
	}

	if k == reflect.Ptr || k == reflect.Interface {
		return rVal.Elem().Interface()
	}

	return rVal.Interface()
}
