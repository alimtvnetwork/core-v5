package coredynamic

import (
	"reflect"
	"sort"
)

func MapKeysStringSliceAnySorted(any interface{}) ([]string, error) {
	reflectVal := reflect.ValueOf(any)

	keys, err := MapKeysStringSlice(reflectVal)

	if err != nil {
		return keys, err
	}

	sort.Strings(keys)

	return keys, nil
}
