package converters

import (
	"fmt"
	"sort"
	"strconv"
)

type mapConverter struct{}

func (it mapConverter) Keys(
	anyMap interface{},
) (keys []string, err error) {
	switch v := anyMap.(type) {
	case map[string]string:
		for key := range v {
			keys = append(keys, key)
		}

		return keys, nil
	case map[string]interface{}:
		for key := range v {
			keys = append(keys, key)
		}

		return keys, nil

	case map[int]interface{}:
		for key := range v {
			keys = append(keys, strconv.Itoa(key))
		}

		return keys, nil
	case map[int]string:
		for key := range v {
			keys = append(keys, strconv.Itoa(key))
		}

		return keys, nil
	case map[float64]interface{}:
		for key := range v {
			keys = append(keys, AnyToSmartString(key))
		}

		return keys, nil

	case map[interface{}]interface{}:
		for key := range v {
			keys = append(keys, AnyToSmartString(key))
		}

		return keys, nil
	case map[interface{}]string:
		for key := range v {
			keys = append(keys, AnyToSmartString(key))
		}

		return keys, nil
	default:
		return keys, fmt.Errorf(
			"current type %T is not support by the function",
			anyMap,
		)
	}
}

func (it mapConverter) SortedKeys(
	anyMap interface{},
) (sortedKeys []string, err error) {
	keys, err := it.Keys(anyMap)

	if err != nil || len(keys) <= 1 {
		return keys, err
	}

	sort.Strings(keys)

	return keys, err
}

func (it mapConverter) Values(
	anyMap interface{},
) (values []string, err error) {
	switch casted := anyMap.(type) {
	case map[string]string:
		for _, value := range casted {
			values = append(values, value)
		}

		return values, nil
	case map[string]interface{}:
		for _, value := range casted {
			values = append(values, AnyToSmartString(value))
		}

		return values, nil

	case map[int]interface{}:
		for _, value := range casted {
			values = append(values, AnyToSmartString(value))
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
	case map[float64]interface{}:
		for _, value := range casted {
			values = append(values, AnyToSmartString(value))
		}

		return values, nil

	case map[interface{}]interface{}:
		for _, value := range casted {
			values = append(values, AnyToSmartString(value))
		}

		return values, nil
	case map[interface{}]string:
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
