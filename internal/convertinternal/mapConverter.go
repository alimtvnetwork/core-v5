package convertinternal

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"

	"gitlab.com/auk-go/core/constants"
)

type mapConverter struct{}

func (it mapConverter) Keys(
	anyMap any,
) (keys []string, err error) {
	switch v := anyMap.(type) {
	case map[string]string:
		for key := range v {
			keys = append(keys, key)
		}
		return keys, nil
	case map[string]any:
		for key := range v {
			keys = append(keys, key)
		}
		return keys, nil
	case map[int]any:
		for key := range v {
			keys = append(keys, strconv.Itoa(key))
		}
		return keys, nil
	case map[int]string:
		for key := range v {
			keys = append(keys, strconv.Itoa(key))
		}
		return keys, nil
	case map[float64]any:
		for key := range v {
			keys = append(keys, AnyTo.SmartString(key))
		}
		return keys, nil
	case map[any]any:
		for key := range v {
			keys = append(keys, AnyTo.SmartString(key))
		}
		return keys, nil
	case map[any]string:
		for key := range v {
			keys = append(keys, AnyTo.SmartString(key))
		}
		return keys, nil
	case map[reflect.Type]string:
		for key := range v {
			keys = append(keys, AnyTo.SmartString(key))
		}
		return keys, nil
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

func (it mapConverter) Values(
	anyMap any,
) (values []string, err error) {
	switch casted := anyMap.(type) {
	case map[string]string:
		for _, value := range casted {
			values = append(values, value)
		}
		return values, nil
	case map[string]any:
		for _, value := range casted {
			values = append(values, AnyTo.SmartString(value))
		}
		return values, nil
	case map[int]any:
		for _, value := range casted {
			values = append(values, AnyTo.SmartString(value))
		}
		return values, nil
	case map[string]int:
		for _, value := range casted {
			values = append(values, strconv.Itoa(value))
		}
		return values, nil
	case map[int]string:
		for _, value := range casted {
			values = append(values, value)
		}
		return values, nil
	case map[float64]any:
		for _, value := range casted {
			values = append(values, AnyTo.SmartString(value))
		}
		return values, nil
	case map[any]any:
		for _, value := range casted {
			values = append(values, AnyTo.SmartString(value))
		}
		return values, nil
	case map[any]string:
		for _, value := range casted {
			values = append(values, value)
		}
		return values, nil
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
