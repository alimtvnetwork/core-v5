package stringslice

import (
	"reflect"
	"sync"
)

func AnyLinesProcessAsyncUsingProcessor(
	lines any,
	lineProcessor func(index int, lineIn any) (lineOut string),
) []string {
	if lines == nil {
		return []string{}
	}

	reflectType := reflect.TypeOf(lines)
	kind := reflectType.Kind()
	isArrayOrSlice := kind == reflect.Slice ||
		kind == reflect.Array

	isNotSliceOrArray := !isArrayOrSlice

	if isNotSliceOrArray {
		return []string{}
	}

	reflectValue := reflect.ValueOf(lines)
	length := reflectValue.Len()

	if length == 0 {
		return []string{}
	}

	slice := MakeLen(length)
	wg := sync.WaitGroup{}

	wg.Add(length)

	asyncProcessor := func(index int, lineIn any) {
		slice[index] = lineProcessor(index, lineIn)

		wg.Done()
	}

	for i := 0; i < length; i++ {
		itemAt := reflectValue.Index(i)
		go asyncProcessor(i, itemAt.Interface())
	}

	wg.Wait()

	return slice
}
