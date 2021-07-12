package coredynamic

import "reflect"

func MapKeysStringSliceAnyMust(any interface{}) []string {
	reflectVal := reflect.ValueOf(any)

	mapKeys, err := MapKeysStringSlice(reflectVal)

	if err != nil {
		panic(err)
	}

	return mapKeys
}
