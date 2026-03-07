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

// IfPtrInt is a typed convenience wrapper for IfPtr[int].
func IfPtrInt(
	isTrue bool,
	trueValue, falseValue *int,
) *int {
	return IfPtr[int](isTrue, trueValue, falseValue)
}

// NilDefInt is a typed convenience wrapper for NilDef[int].
func NilDefInt(
	valuePointer *int,
	defVal int,
) int {
	return NilDef[int](valuePointer, defVal)
}

// NilDefPtrInt is a typed convenience wrapper for NilDefPtr[int].
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

// NilValInt is a typed convenience wrapper for NilVal[int].
func NilValInt(valuePointer *int, onNil, onNonNil int) int {
	return NilVal[int](valuePointer, onNil, onNonNil)
}

// NilValPtrInt is a typed convenience wrapper for NilValPtr[int].
func NilValPtrInt(valuePointer *int, onNil, onNonNil int) *int {
	return NilValPtr[int](valuePointer, onNil, onNonNil)
}
