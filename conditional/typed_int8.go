package conditional

// IfInt8 is a typed convenience wrapper for If[int8].
func IfInt8(
	isTrue bool,
	trueValue, falseValue int8,
) int8 {
	return If[int8](isTrue, trueValue, falseValue)
}

// IfFuncInt8 is a typed convenience wrapper for IfFunc[int8].
func IfFuncInt8(
	isTrue bool,
	trueValueFunc, falseValueFunc func() int8,
) int8 {
	return IfFunc[int8](isTrue, trueValueFunc, falseValueFunc)
}

// IfTrueFuncInt8 is a typed convenience wrapper for IfTrueFunc[int8].
func IfTrueFuncInt8(
	isTrue bool,
	trueValueFunc func() int8,
) int8 {
	return IfTrueFunc[int8](isTrue, trueValueFunc)
}

// IfSliceInt8 is a typed convenience wrapper for IfSlice[int8].
func IfSliceInt8(
	isTrue bool,
	trueValue, falseValue []int8,
) []int8 {
	return IfSlice[int8](isTrue, trueValue, falseValue)
}

// IfPtrInt8 is a typed convenience wrapper for IfPtr[int8].
func IfPtrInt8(
	isTrue bool,
	trueValue, falseValue *int8,
) *int8 {
	return IfPtr[int8](isTrue, trueValue, falseValue)
}

// NilDefInt8 is a typed convenience wrapper for NilDef[int8].
func NilDefInt8(
	valuePointer *int8,
	defVal int8,
) int8 {
	return NilDef[int8](valuePointer, defVal)
}

// NilDefPtrInt8 is a typed convenience wrapper for NilDefPtr[int8].
func NilDefPtrInt8(
	valuePointer *int8,
	defVal int8,
) *int8 {
	return NilDefPtr[int8](valuePointer, defVal)
}

// ValueOrZeroInt8 is a typed convenience wrapper for ValueOrZero[int8].
func ValueOrZeroInt8(valuePointer *int8) int8 {
	return ValueOrZero[int8](valuePointer)
}

// PtrOrZeroInt8 is a typed convenience wrapper for PtrOrZero[int8].
func PtrOrZeroInt8(valuePointer *int8) *int8 {
	return PtrOrZero[int8](valuePointer)
}

// NilValInt8 is a typed convenience wrapper for NilVal[int8].
func NilValInt8(valuePointer *int8, onNil, onNonNil int8) int8 {
	return NilVal[int8](valuePointer, onNil, onNonNil)
}

// NilValPtrInt8 is a typed convenience wrapper for NilValPtr[int8].
func NilValPtrInt8(valuePointer *int8, onNil, onNonNil int8) *int8 {
	return NilValPtr[int8](valuePointer, onNil, onNonNil)
}
