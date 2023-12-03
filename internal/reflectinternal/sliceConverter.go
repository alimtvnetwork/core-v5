package reflectinternal

import (
	"fmt"
	"reflect"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
)

type sliceConverter struct{}

func (it sliceConverter) ToStrings(reflectVal reflect.Value) ([]string, error) {
	if reflectVal.Kind() == reflect.Ptr {
		return SliceItemsAsStrings(
			reflect.Indirect(reflect.ValueOf(reflectVal)),
		)
	}

	k := reflectVal.Kind()
	isSliceOrArray := k == reflect.Slice ||
		k == reflect.Array

	if !isSliceOrArray {
		return []string{},
			errcore.TypeMismatchType.Error("Reflection is not Slice or Array", reflectVal)
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

func SliceItemsAsStringsAny(any interface{}) ([]string, error) {
	reflectVal := reflect.ValueOf(any)

	return SliceItemsAsStrings(reflectVal)
}

func SliceItemsAsStringsAnyMust(any interface{}) []string {
	reflectVal := reflect.ValueOf(any)

	items, err := SliceItemsAsStrings(reflectVal)

	if err != nil {
		panic(err)
	}

	return items
}
func SliceItemsProcessorAsStrings(
	reflectVal reflect.Value,
	processor func(index int, item interface{}) (result string, isTake, isBreak bool),
) ([]string, error) {
	if reflectVal.Kind() == reflect.Ptr {
		return SliceItemsAsStrings(
			reflect.Indirect(reflect.ValueOf(reflectVal)),
		)
	}

	k := reflectVal.Kind()
	isSliceOrArray := k == reflect.Slice ||
		k == reflect.Array

	if !isSliceOrArray {
		return []string{},
			errcore.TypeMismatchType.Error("Reflection is not Slice or Array", reflectVal)
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

func SliceItemsSimpleProcessorAsStrings(
	reflectVal reflect.Value,
	isSkipEmpty bool,
	processor func(index int, item interface{}) (result string),
) ([]string, error) {
	if reflectVal.Kind() == reflect.Ptr {
		return SliceItemsAsStrings(
			reflect.Indirect(reflect.ValueOf(reflectVal)),
		)
	}

	k := reflectVal.Kind()
	isSliceOrArray := k == reflect.Slice ||
		k == reflect.Array

	if !isSliceOrArray {
		return []string{},
			errcore.TypeMismatchType.Error(
				"Reflection is not Slice or Array", reflectVal,
			)
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
