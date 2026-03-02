package conditional

// IfInt is a typed convenience wrapper for If[int].
func IfInt(
	isTrue bool,
	trueValue, falseValue int,
) int {
	return If[int](isTrue, trueValue, falseValue)
}

// IfFuncInt is a typed convenience wrapper for IfFunc[int].
func IfFuncInt(
	isTrue bool,
	trueValueFunc, falseValueFunc func() int,
) int {
	return IfFunc[int](isTrue, trueValueFunc, falseValueFunc)
}

// IfTrueFuncInt is a typed convenience wrapper for IfTrueFunc[int].
func IfTrueFuncInt(
	isTrue bool,
	trueValueFunc func() int,
) int {
	return IfTrueFunc[int](isTrue, trueValueFunc)
}

// IfSliceInt is a typed convenience wrapper for IfSlice[int].
func IfSliceInt(
	isTrue bool,
	trueValue, falseValue []int,
) []int {
	return IfSlice[int](isTrue, trueValue, falseValue)
}

// Deprecated: Use IfSliceInt instead.
func IfSlicePtrInt(
	isTrue bool,
	trueValue, falseValue []int,
) []int {
	return IfSlice[int](isTrue, trueValue, falseValue)
}

// Deprecated: Use IfSlice[int] with func wrappers instead.
func IfSlicePtrFuncInt(
	isTrue bool,
	trueValueFunc, falseValueFunc func() []int,
) []int {
	if isTrue {
		return trueValueFunc()
	}
	return falseValueFunc()
}

// IfPtrInt is a typed convenience wrapper for IfPtr[int].
func IfPtrInt(
	isTrue bool,
	trueValue, falseValue *int,
) *int {
	return IfPtr[int](isTrue, trueValue, falseValue)
}

// NilDefPtrInt is a typed convenience wrapper for NilDefPtr[int].
// Note: NilDef[int] typed wrapper is omitted due to naming conflict
// with the deprecated NilDefInt(ptr) in NilDefInt.go.
// Use NilDef[int](ptr, defVal) directly.
func NilDefPtrInt(
	valuePointer *int,
	defVal int,
) *int {
	return NilDefPtr[int](valuePointer, defVal)
}

// ValueOrZeroInt is a typed convenience wrapper for ValueOrZero[int].
func ValueOrZeroInt(valuePointer *int) int {
	return ValueOrZero[int](valuePointer)
}

// PtrOrZeroInt is a typed convenience wrapper for PtrOrZero[int].
func PtrOrZeroInt(valuePointer *int) *int {
	return PtrOrZero[int](valuePointer)
}

// Deprecated: Use ValueOrZeroInt instead.
func NilDerefInt(valuePointer *int) int {
	return ValueOrZero[int](valuePointer)
}

// Deprecated: Use PtrOrZeroInt instead.
func NilDerefPtrInt(valuePointer *int) *int {
	return PtrOrZero[int](valuePointer)
}

// NilValInt is a typed convenience wrapper for NilVal[int].
func NilValInt(valuePointer *int, onNil, onNonNil int) int {
	return NilVal[int](valuePointer, onNil, onNonNil)
}

// NilValPtrInt is a typed convenience wrapper for NilValPtr[int].
func NilValPtrInt(valuePointer *int, onNil, onNonNil int) *int {
	return NilValPtr[int](valuePointer, onNil, onNonNil)
}
