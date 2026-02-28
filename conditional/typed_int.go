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

// IfSlicePtrInt is a typed convenience wrapper for IfSlicePtr[int].
func IfSlicePtrInt(
	isTrue bool,
	trueValue, falseValue *[]int,
) *[]int {
	return IfSlicePtr[int](isTrue, trueValue, falseValue)
}

// IfSlicePtrFuncInt is a typed convenience wrapper for IfSlicePtrFunc[int].
func IfSlicePtrFuncInt(
	isTrue bool,
	trueValueFunc, falseValueFunc func() *[]int,
) *[]int {
	return IfSlicePtrFunc[int](isTrue, trueValueFunc, falseValueFunc)
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
