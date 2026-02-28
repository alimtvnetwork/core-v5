package corestr

import "reflect"

func reflectInterfaceVal(item any) any {
	rVal := reflect.ValueOf(item)

	if rVal.Kind() != reflect.Ptr {
		return rVal.Interface()
	}

	if rVal.Kind() == reflect.Ptr {
		rVal = rVal.Elem()
	}

	return rVal.Interface()
}
