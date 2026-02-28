package coredynamic

import "reflect"

func PointerOrNonPointer(
	isPointerOutput bool,
	input any,
) (output any, finalReflectVal reflect.Value) {
	return PointerOrNonPointerUsingReflectValue(
		isPointerOutput,
		reflect.ValueOf(input))
}
