package conditional

// IfInt64 is a typed convenience wrapper for If[int64].
func IfInt64(
	isTrue bool,
	trueValue, falseValue int64,
) int64 {
	return If[int64](isTrue, trueValue, falseValue)
}

// IfFuncInt64 is a typed convenience wrapper for IfFunc[int64].
func IfFuncInt64(
	isTrue bool,
	trueValueFunc, falseValueFunc func() int64,
) int64 {
	return IfFunc[int64](isTrue, trueValueFunc, falseValueFunc)
}

// IfTrueFuncInt64 is a typed convenience wrapper for IfTrueFunc[int64].
func IfTrueFuncInt64(
	isTrue bool,
	trueValueFunc func() int64,
) int64 {
	return IfTrueFunc[int64](isTrue, trueValueFunc)
}

// IfSliceInt64 is a typed convenience wrapper for IfSlice[int64].
func IfSliceInt64(
	isTrue bool,
	trueValue, falseValue []int64,
) []int64 {
	return IfSlice[int64](isTrue, trueValue, falseValue)
}

// Deprecated: Use IfSliceInt64 instead.
func IfSlicePtrInt64(
	isTrue bool,
	trueValue, falseValue []int64,
) []int64 {
	return IfSlice[int64](isTrue, trueValue, falseValue)
}

// Deprecated: Use IfSlice[int64] with func wrappers instead.
func IfSlicePtrFuncInt64(
	isTrue bool,
	trueValueFunc, falseValueFunc func() []int64,
) []int64 {
	if isTrue {
		return trueValueFunc()
	}
	return falseValueFunc()
}

// IfPtrInt64 is a typed convenience wrapper for IfPtr[int64].
func IfPtrInt64(
	isTrue bool,
	trueValue, falseValue *int64,
) *int64 {
	return IfPtr[int64](isTrue, trueValue, falseValue)
}

// NilDefInt64 is a typed convenience wrapper for NilDef[int64].
func NilDefInt64(
	valuePointer *int64,
	defVal int64,
) int64 {
	return NilDef[int64](valuePointer, defVal)
}

// NilDefPtrInt64 is a typed convenience wrapper for NilDefPtr[int64].
func NilDefPtrInt64(
	valuePointer *int64,
	defVal int64,
) *int64 {
	return NilDefPtr[int64](valuePointer, defVal)
}
