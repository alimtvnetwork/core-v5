package conditional

// IfInt16 is a typed convenience wrapper for If[int16].
func IfInt16(
	isTrue bool,
	trueValue, falseValue int16,
) int16 {
	return If[int16](isTrue, trueValue, falseValue)
}

// IfFuncInt16 is a typed convenience wrapper for IfFunc[int16].
func IfFuncInt16(
	isTrue bool,
	trueValueFunc, falseValueFunc func() int16,
) int16 {
	return IfFunc[int16](isTrue, trueValueFunc, falseValueFunc)
}

// IfTrueFuncInt16 is a typed convenience wrapper for IfTrueFunc[int16].
func IfTrueFuncInt16(
	isTrue bool,
	trueValueFunc func() int16,
) int16 {
	return IfTrueFunc[int16](isTrue, trueValueFunc)
}

// IfSliceInt16 is a typed convenience wrapper for IfSlice[int16].
func IfSliceInt16(
	isTrue bool,
	trueValue, falseValue []int16,
) []int16 {
	return IfSlice[int16](isTrue, trueValue, falseValue)
}

// IfSlicePtrInt16 is a typed convenience wrapper for IfSlicePtr[int16].
func IfSlicePtrInt16(
	isTrue bool,
	trueValue, falseValue *[]int16,
) *[]int16 {
	return IfSlicePtr[int16](isTrue, trueValue, falseValue)
}

// IfSlicePtrFuncInt16 is a typed convenience wrapper for IfSlicePtrFunc[int16].
func IfSlicePtrFuncInt16(
	isTrue bool,
	trueValueFunc, falseValueFunc func() *[]int16,
) *[]int16 {
	return IfSlicePtrFunc[int16](isTrue, trueValueFunc, falseValueFunc)
}

// IfPtrInt16 is a typed convenience wrapper for IfPtr[int16].
func IfPtrInt16(
	isTrue bool,
	trueValue, falseValue *int16,
) *int16 {
	return IfPtr[int16](isTrue, trueValue, falseValue)
}

// NilDefInt16 is a typed convenience wrapper for NilDef[int16].
func NilDefInt16(
	valuePointer *int16,
	defVal int16,
) int16 {
	return NilDef[int16](valuePointer, defVal)
}

// NilDefPtrInt16 is a typed convenience wrapper for NilDefPtr[int16].
func NilDefPtrInt16(
	valuePointer *int16,
	defVal int16,
) *int16 {
	return NilDefPtr[int16](valuePointer, defVal)
}
