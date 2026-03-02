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
	isFalse := !isTrue

	if isFalse {
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

// Deprecated: Use IfSlice instead.
func IfSlicePtr[T any](
	isTrue bool,
	trueValue, falseValue []T,
) []T {
	if isTrue {
		return trueValue
	}

	return falseValue
}

// Deprecated: Use IfSlice with func wrappers instead.
func IfSlicePtrFunc[T any](
	isTrue bool,
	trueValueFunc, falseValueFunc func() []T,
) []T {
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

// NilVal returns onNil if the pointer is nil, otherwise onNonNil.
// Unlike NilDef, it does NOT dereference the pointer — it chooses between two provided values.
//
// It replaces NilStr, NilCheck, and similar per-type nil-conditional functions.
//
// Usage:
//
//	label := conditional.NilVal[string](namePtr, "(unknown)", "has name")
func NilVal[T any](
	valuePointer *T,
	onNil, onNonNil T,
) T {
	if valuePointer == nil {
		return onNil
	}

	return onNonNil
}

// NilValPtr is like NilVal but returns a pointer to the chosen value.
//
// Usage:
//
//	labelPtr := conditional.NilValPtr[string](namePtr, "(unknown)", "has name")
func NilValPtr[T any](
	valuePointer *T,
	onNil, onNonNil T,
) *T {
	if valuePointer == nil {
		return &onNil
	}

	return &onNonNil
}

// ValueOrZero dereferences a pointer, returning the zero value of T if nil.
// It replaces NilDefBool, NilDefByte, NilDefInt, NilDefStr (zero-default variants).
//
// Usage:
//
//	active := conditional.ValueOrZero[bool](flagPtr)   // false if nil
//	name := conditional.ValueOrZero[string](namePtr)   // "" if nil
func ValueOrZero[T any](
	valuePointer *T,
) T {
	if valuePointer == nil {
		var zero T
		return zero
	}

	return *valuePointer
}

// Deprecated: Use ValueOrZero instead.
func NilDeref[T any](valuePointer *T) T {
	return ValueOrZero[T](valuePointer)
}

// PtrOrZero returns the pointer itself if non-nil, otherwise a pointer to the zero value.
// It replaces NilDefBoolPtr, NilDefBytePtr, NilDefIntPtr, NilDefStrPtr.
//
// Usage:
//
//	ptr := conditional.PtrOrZero[int](intPtr) // guaranteed non-nil
func PtrOrZero[T any](
	valuePointer *T,
) *T {
	if valuePointer == nil {
		var zero T
		return &zero
	}

	return valuePointer
}

// Deprecated: Use PtrOrZero instead.
func NilDerefPtr[T any](valuePointer *T) *T {
	return PtrOrZero[T](valuePointer)
}
