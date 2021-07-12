package coredynamic

import "reflect"

func SliceItemsAsStringsAnyMust(any interface{}) []string {
	reflectVal := reflect.ValueOf(any)

	items, err := SliceItemsAsStrings(reflectVal)

	if err != nil {
		panic(err)
	}

	return items
}
