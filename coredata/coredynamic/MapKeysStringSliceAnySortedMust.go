package coredynamic

import (
	"reflect"
	"sort"
)

func MapKeysStringSliceAnySortedMust(any interface{}) []string {
	reflectVal := reflect.ValueOf(any)

	keys := MapKeysStringSliceAnyMust(reflectVal)
	sort.Strings(keys)

	return keys
}
