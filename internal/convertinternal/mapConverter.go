package convertinternal

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"

	"github.com/alimtvnetwork/core/constants"
)

type mapConverter struct{}

// keysFromIntKeyMap extracts string keys from int-keyed maps.
func keysFromIntKeyMap[V any](m map[int]V) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, strconv.Itoa(key))
	}
	return keys
}

// keysFromStringKeyMap extracts keys from string-keyed maps.
func keysFromStringKeyMap[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// keysFromAnyKeyMap extracts string keys from any-keyed maps using SmartString.
func keysFromAnyKeyMap[K comparable, V any](m map[K]V) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, AnyTo.SmartString(key))
	}
	return keys
}

func (it mapConverter) Keys(
	anyMap any,
) (keys []string, err error) {
	switch v := anyMap.(type) {
	case map[string]string:
		return keysFromStringKeyMap(v), nil
	case map[string]any:
		return keysFromStringKeyMap(v), nil
	case map[int]any:
		return keysFromIntKeyMap(v), nil
	case map[int]string:
		return keysFromIntKeyMap(v), nil
	case map[float64]any:
		return keysFromAnyKeyMap(v), nil
	case map[any]any:
		return keysFromAnyKeyMap(v), nil
	case map[any]string:
		return keysFromAnyKeyMap(v), nil
	case map[reflect.Type]string:
		return keysFromAnyKeyMap(v), nil
	default:
		return keys, fmt.Errorf(
			"current type %T is not support by the function",
			anyMap,
		)
	}
}

func (it mapConverter) KeysValues(
	anyMap any,
) (keys, values []string, err error) {
	switch v := anyMap.(type) {
	case map[string]string:
		for key, value := range v {
			keys = append(keys, key)
			values = append(values, value)
		}
		return keys, values, nil
	case map[string]any:
		for key, value := range v {
			keys = append(keys, key)
			values = append(values, AnyTo.SmartString(value))
		}
		return keys, values, nil
	case map[int]any:
		for key, value := range v {
			keys = append(keys, strconv.Itoa(key))
			values = append(values, AnyTo.SmartString(value))
		}
		return keys, values, nil
	case map[int]string:
		for key, value := range v {
			keys = append(keys, strconv.Itoa(key))
			values = append(values, value)
		}
		return keys, values, nil
	case map[float64]any:
		for key, value := range v {
			keys = append(keys, AnyTo.SmartString(key))
			values = append(values, AnyTo.SmartString(value))
		}
		return keys, values, nil
	case map[any]any:
		for key, value := range v {
			keys = append(keys, AnyTo.SmartString(key))
			values = append(values, AnyTo.SmartString(value))
		}
		return keys, values, nil
	case map[any]string:
		for key, value := range v {
			keys = append(keys, AnyTo.SmartString(key))
			values = append(values, value)
		}
		return keys, values, nil
	default:
		return keys, values, fmt.Errorf(
			"current type %T is not support by the function",
			anyMap,
		)
	}
}

func (it mapConverter) SortedKeys(
	anyMap any,
) (sortedKeys []string, err error) {
	keys, err := it.Keys(anyMap)
	if err != nil || len(keys) <= 1 {
		return keys, err
	}
	sort.Strings(keys)
	return keys, err
}

func (it mapConverter) SortedKeysValues(
	anyMap any,
) (keys, values []string, err error) {
	keys, values, err = it.KeysValues(anyMap)
	if err != nil {
		return keys, values, err
	}
	toMap := make(map[string]string, len(keys))
	for i, key := range keys {
		toMap[key] = values[i]
	}
	sort.Strings(keys)
	for i, key := range keys {
		values[i] = toMap[key]
	}
	return keys, values, err
}

// valuesFromStringValMap extracts string values directly.
func valuesFromStringValMap[K comparable](m map[K]string) []string {
	values := make([]string, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

// valuesFromAnyValMap extracts values using SmartString conversion.
func valuesFromAnyValMap[K comparable](m map[K]any) []string {
	values := make([]string, 0, len(m))
	for _, value := range m {
		values = append(values, AnyTo.SmartString(value))
	}
	return values
}

func (it mapConverter) Values(
	anyMap any,
) (values []string, err error) {
	switch casted := anyMap.(type) {
	case map[string]string:
		return valuesFromStringValMap(casted), nil
	case map[string]any:
		return valuesFromAnyValMap(casted), nil
	case map[int]any:
		return valuesFromAnyValMap(casted), nil
	case map[string]int:
		vals := make([]string, 0, len(casted))
		for _, value := range casted {
			vals = append(vals, strconv.Itoa(value))
		}
		return vals, nil
	case map[int]string:
		return valuesFromStringValMap(casted), nil
	case map[float64]any:
		return valuesFromAnyValMap(casted), nil
	case map[any]any:
		return valuesFromAnyValMap(casted), nil
	case map[any]string:
		return valuesFromStringValMap(casted), nil
	default:
		return values, fmt.Errorf(
			"current type %T is not support by the function",
			anyMap,
		)
	}
}

func (it mapConverter) StringAnyToStringString(
	isSkipEmpty bool,
	additionalMapItems map[string]any,
) map[string]string {
	if len(additionalMapItems) == 0 {
		return map[string]string{}
	}

	newMap := make(map[string]string, len(additionalMapItems))

	for key, valInf := range additionalMapItems {
		val := fmt.Sprintf(
			constants.SprintValueFormat,
			valInf,
		)

		if isSkipEmpty && val == "" {
			continue
		}

		newMap[key] = val
	}

	return newMap
}

func (it mapConverter) CombineMapStringAny(
	isSkipEmpty bool,
	mainMap map[string]any,
	additionalMapItems map[string]any,
) map[string]string {
	if len(mainMap) == 0 && len(additionalMapItems) == 0 {
		return map[string]string{}
	}

	newMap := make(
		map[string]string,
		len(mainMap)+
			len(additionalMapItems)+
			constants.Capacity3,
	)

	for key, valInf := range mainMap {
		val := fmt.Sprintf(
			constants.SprintValueFormat,
			valInf,
		)

		if isSkipEmpty && val == "" {
			continue
		}

		newMap[key] = val
	}

	for key, valInf := range additionalMapItems {
		val := fmt.Sprintf(
			constants.SprintValueFormat,
			valInf,
		)

		if isSkipEmpty && val == "" {
			continue
		}

		newMap[key] = val
	}

	return newMap
}

func (it mapConverter) FromIntegersToMap(inputArray ...int) map[int]bool {
	if len(inputArray) == 0 {
		return map[int]bool{}
	}

	length := len(inputArray)
	hashset := make(map[int]bool, length)

	for _, s := range inputArray {
		hashset[s] = true
	}

	return hashset
}
