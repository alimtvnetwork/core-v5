package conditional

// If is a generic ternary helper.
// It replaces the per-type Bool, Int, String, Byte, Interface functions.
//
// Usage:
//
//	result := conditional.If[int](true, 2, 7)   // returns 2
//	name := conditional.If[string](len(s) > 0, s, "default")
func If[T any](
	isTrue bool,
	trueValue, falseValue T,
) T {
	if isTrue {
		return trueValue
	}

	return falseValue
}

// IfFunc evaluates the appropriate function based on condition and returns the result.
// It replaces the per-type BoolFunc, StringFunc, InterfaceFunc functions.
func IfFunc[T any](
	isTrue bool,
	trueValueFunc, falseValueFunc func() T,
) T {
	if isTrue {
		return trueValueFunc()
	}

	return falseValueFunc()
}

// IfTrueFunc evaluates trueValueFunc only when isTrue, otherwise returns the zero value.
// It replaces the per-type BooleanTrueFunc, StringTrueFunc, BytesTrueFunc functions.
func IfTrueFunc[T any](
	isTrue bool,
	trueValueFunc func() T,
) T {
	if !isTrue {
		var zero T
		return zero
	}

	return trueValueFunc()
}

// IfSlice is a generic ternary for slice types.
// It replaces the per-type Integers, Booleans, Strings, Interfaces, Bytes functions.
func IfSlice[T any](
	isTrue bool,
	trueValue, falseValue []T,
) []T {
	if isTrue {
		return trueValue
	}

	return falseValue
}

// IfSlicePtr is a generic ternary for pointer-to-slice types.
// It replaces the per-type IntegersPtr, BooleansPtr, StringsPtr, InterfacesPtr, BytesPtr functions.
func IfSlicePtr[T any](
	isTrue bool,
	trueValue, falseValue *[]T,
) *[]T {
	if isTrue {
		return trueValue
	}

	return falseValue
}

// IfSlicePtrFunc evaluates the appropriate function and returns a pointer to a slice.
// It replaces IntegersPtrFunc and similar per-type functions.
func IfSlicePtrFunc[T any](
	isTrue bool,
	trueValueFunc, falseValueFunc func() *[]T,
) *[]T {
	if isTrue {
		return trueValueFunc()
	}

	return falseValueFunc()
}

// NilDef dereferences a pointer, returning defVal if the pointer is nil.
// It replaces NilDefValInt, NilBoolVal, NilByteVal and similar functions.
func NilDef[T any](
	valuePointer *T,
	defVal T,
) T {
	if valuePointer == nil {
		return defVal
	}

	return *valuePointer
}

// NilDefPtr returns the pointer itself if non-nil, otherwise a pointer to defVal.
// It replaces NilDefIntPtr, NilDefBoolPtr, NilDefBytePtr, NilDefStrPtr and similar.
func NilDefPtr[T any](
	valuePointer *T,
	defVal T,
) *T {
	if valuePointer == nil {
		return &defVal
	}

	return valuePointer
}

// IfPtr is a generic ternary for pointer types.
// It replaces StringPtr and similar per-type pointer conditionals.
func IfPtr[T any](
	isTrue bool,
	trueValue, falseValue *T,
) *T {
	if isTrue {
		return trueValue
	}

	return falseValue
}
