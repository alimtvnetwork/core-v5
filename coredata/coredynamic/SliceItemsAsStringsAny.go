package coredynamic

import "reflect"

func SliceItemsAsStringsAny(any interface{}) ([]string, error) {
	reflectVal := reflect.ValueOf(any)

	return SliceItemsAsStrings(reflectVal)
}
