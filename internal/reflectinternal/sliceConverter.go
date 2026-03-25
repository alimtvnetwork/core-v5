package reflectinternal

import (
	"fmt"
	"reflect"

	"github.com/alimtvnetwork/core/constants"
)

type sliceConverter struct{}

func (it sliceConverter) Length(i any) int {
	if Is.Null(i) {
		return 0
	}

	reflectVal := reflect.ValueOf(i)

	k := reflectVal.Kind()
	isSliceOrArray := k == reflect.Slice ||
		k == reflect.Array ||
		k == reflect.Map

	if isSliceOrArray {
		return reflectVal.Len()
	}

	return 0
}

func (it sliceConverter) ToStringsRv(reflectVal reflect.Value) ([]string, error) {
	if reflectVal.Kind() == reflect.Ptr {
		return it.ToStringsRv(
			reflect.Indirect(reflectVal),
		)
	}

	k := reflectVal.Kind()
	isSliceOrArray := k == reflect.Slice ||
		k == reflect.Array

	isNotSliceOrArray := !isSliceOrArray

	if isNotSliceOrArray {
		return []string{},
			fmt.Errorf("reflection is not a slice nor array but %s", reflectVal.String())
	}

	length := reflectVal.Len()
	slice := make([]string, length)

	if length == 0 {
		return slice, nil
	}

	for i := 0; i < length; i++ {
		value := reflectVal.Index(i)
		toString := fmt.Sprintf(
			constants.SprintValueFormat,
			value.Interface(),
		)
		slice[i] = toString
	}

	return slice, nil
}

func (it sliceConverter) ToStrings(anyItem any) ([]string, error) {
	reflectVal := reflect.ValueOf(anyItem)

	return it.ToStringsRv(reflectVal)
}

func (it sliceConverter) ToStringsMust(anyItem any) []string {
	reflectVal := reflect.ValueOf(anyItem)

	items, err := it.ToStringsRv(reflectVal)

	if err != nil {
		panic(err)
	}

	return items
}

func (it sliceConverter) ToStringsRvUsingProcessor(
	reflectVal reflect.Value,
	processor func(index int, item any) (result string, isTake, isBreak bool),
) ([]string, error) {
	if reflectVal.Kind() == reflect.Ptr {
		return it.ToStringsRvUsingProcessor(
			reflect.Indirect(reflectVal),
			processor,
		)
	}

	k := reflectVal.Kind()
	isSliceOrArray := k == reflect.Slice ||
		k == reflect.Array

	isNotSliceOrArray := !isSliceOrArray

	if isNotSliceOrArray {
		return []string{},
			fmt.Errorf("reflection is not a slice nor array but %s", reflectVal.String())
	}

	length := reflectVal.Len()
	slice := make([]string, 0, length)

	for i := 0; i < length; i++ {
		value := reflectVal.Index(i)
		result, isTake, isBreak := processor(
			i, value,
		)

		if isTake {
			slice = append(
				slice,
				result,
			)
		}

		if isBreak {
			return slice, nil
		}
	}

	return slice, nil
}

func (it sliceConverter) ToStringsRvUsingSimpleProcessor(
	reflectVal reflect.Value,
	isSkipEmpty bool,
	processor func(index int, item any) (result string),
) ([]string, error) {
	if reflectVal.Kind() == reflect.Ptr {
		return it.ToStringsRvUsingSimpleProcessor(
			reflect.Indirect(reflectVal),
			isSkipEmpty,
			processor,
		)
	}

	k := reflectVal.Kind()
	isSliceOrArray := k == reflect.Slice ||
		k == reflect.Array

	isNotSliceOrArray := !isSliceOrArray

	if isNotSliceOrArray {
		return []string{},
			fmt.Errorf("reflection is not a slice nor array but %s", reflectVal.String())
	}

	length := reflectVal.Len()
	slice := make([]string, 0, length)

	for i := 0; i < length; i++ {
		value := reflectVal.Index(i)
		result := processor(
			i, value,
		)

		if isSkipEmpty && result == "" {
			continue
		}

		slice = append(
			slice,
			result,
		)
	}

	return slice, nil
}

func (it sliceConverter) ToAnyItemsAsync(
	slice any,
) []any {
	if slice == nil {
		return []any{}
	}

	return Converter.ReflectValToInterfacesAsync(
		reflect.ValueOf(slice),
	)
}
