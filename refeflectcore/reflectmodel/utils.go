package reflectmodel

import (
	"reflect"
)

func argsReflectValues(args []interface{}) []reflect.Value {
	if len(args) == 0 {
		return []reflect.Value{}
	}

	list := make(
		[]reflect.Value,
		len(args),
	)

	for i, arg := range args {
		list[i] = reflect.ValueOf(arg)
	}

	return list
}

func reflectValuesToInterfaces(
	reflectValues []reflect.Value,
) []interface{} {
	if len(reflectValues) == 0 {
		return []interface{}{}
	}

	list := make(
		[]interface{},
		len(reflectValues),
	)

	for i, rv := range reflectValues {
		list[i] = reflectValueToAnyValue(rv)
	}

	return list
}

func reflectValueToAnyValue(rv reflect.Value) interface{} {
	if isNull(rv) {
		return nil
	}

	k := rv.Kind()

	switch k {
	case reflect.Ptr, reflect.Interface:
		return rv.Elem().Interface()
	default:
		return rv.Interface()
	}
}
