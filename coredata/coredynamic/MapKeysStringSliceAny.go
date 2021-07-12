package coredynamic

import (
	"reflect"
)

func MapKeysStringSliceAny(any interface{}) ([]string, error) {
	reflectVal := reflect.ValueOf(any)

	return MapKeysStringSlice(reflectVal)
}
