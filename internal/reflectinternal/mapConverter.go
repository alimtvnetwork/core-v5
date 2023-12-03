package reflectinternal

import (
	"reflect"
	"sort"

	"gitlab.com/auk-go/core/errcore"
)

type mapConverter struct{}

// MapKeysStringSlice
//
//	expectation : map[key:string]...value don't care.
func MapKeysStringSlice(reflectVal reflect.Value) ([]string, error) {
	if reflectVal.Kind() == reflect.Ptr {
		return MapKeysStringSlice(
			reflect.Indirect(reflect.ValueOf(reflectVal)),
		)
	}

	if reflectVal.Kind() != reflect.Map {
		return []string{},
			errcore.TypeMismatchType.Error("Reflection is not Map", reflectVal)
	}

	mapKeys := reflectVal.MapKeys()
	length := len(mapKeys)
	keys := make([]string, length)

	for i, key := range reflectVal.MapKeys() {
		keyAny := key.Interface()
		keyAsString, isString := keyAny.(string)

		if !isString {
			return keys, errcore.TypeMismatchType.Error("Not string type", keyAny)
		}

		keys[i] = keyAsString
	}

	return keys, nil
}

// MapKeysStringSliceAny
//
//	expectation : map[key:string]don't care values
func MapKeysStringSliceAny(any interface{}) ([]string, error) {
	reflectVal := reflect.ValueOf(any)

	return MapKeysStringSlice(reflectVal)
}

func MapKeysStringSliceAnyMust(any interface{}) []string {
	reflectVal := reflect.ValueOf(any)

	mapKeys, err := MapKeysStringSlice(reflectVal)

	if err != nil {
		panic(err)
	}

	return mapKeys
}

func MapKeysStringSliceAnySorted(any interface{}) ([]string, error) {
	reflectVal := reflect.ValueOf(any)

	keys, err := MapKeysStringSlice(reflectVal)

	if err != nil {
		return keys, err
	}

	sort.Strings(keys)

	return keys, nil
}

func MapKeysStringSliceAnySortedMust(any interface{}) []string {
	reflectVal := reflect.ValueOf(any)

	keys := MapKeysStringSliceAnyMust(reflectVal)
	sort.Strings(keys)

	return keys
}
