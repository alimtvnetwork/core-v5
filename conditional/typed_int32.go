package conditional

// IfInt32 is a typed convenience wrapper for If[int32].
func IfInt32(
	isTrue bool,
	trueValue, falseValue int32,
) int32 {
	return If[int32](isTrue, trueValue, falseValue)
}

// IfFuncInt32 is a typed convenience wrapper for IfFunc[int32].
func IfFuncInt32(
	isTrue bool,
	trueValueFunc, falseValueFunc func() int32,
) int32 {
	return IfFunc[int32](isTrue, trueValueFunc, falseValueFunc)
}

// IfTrueFuncInt32 is a typed convenience wrapper for IfTrueFunc[int32].
func IfTrueFuncInt32(
	isTrue bool,
	trueValueFunc func() int32,
) int32 {
	return IfTrueFunc[int32](isTrue, trueValueFunc)
}

// IfSliceInt32 is a typed convenience wrapper for IfSlice[int32].
func IfSliceInt32(
	isTrue bool,
	trueValue, falseValue []int32,
) []int32 {
	return IfSlice[int32](isTrue, trueValue, falseValue)
}

// Deprecated: Use IfSliceInt32 instead.
func IfSlicePtrInt32(
	isTrue bool,
	trueValue, falseValue []int32,
) []int32 {
	return IfSlice[int32](isTrue, trueValue, falseValue)
}

// Deprecated: Use IfSlice[int32] with func wrappers instead.
func IfSlicePtrFuncInt32(
	isTrue bool,
	trueValueFunc, falseValueFunc func() []int32,
) []int32 {
	if isTrue {
		return trueValueFunc()
	}
	return falseValueFunc()
}

// IfPtrInt32 is a typed convenience wrapper for IfPtr[int32].
func IfPtrInt32(
	isTrue bool,
	trueValue, falseValue *int32,
) *int32 {
	return IfPtr[int32](isTrue, trueValue, falseValue)
}

// NilDefInt32 is a typed convenience wrapper for NilDef[int32].
func NilDefInt32(
	valuePointer *int32,
	defVal int32,
) int32 {
	return NilDef[int32](valuePointer, defVal)
}

// NilDefPtrInt32 is a typed convenience wrapper for NilDefPtr[int32].
func NilDefPtrInt32(
	valuePointer *int32,
	defVal int32,
) *int32 {
	return NilDefPtr[int32](valuePointer, defVal)
}

// NilDerefInt32 is a typed convenience wrapper for NilDeref[int32].
func NilDerefInt32(valuePointer *int32) int32 {
	return NilDeref[int32](valuePointer)
}

// NilDerefPtrInt32 is a typed convenience wrapper for NilDerefPtr[int32].
func NilDerefPtrInt32(valuePointer *int32) *int32 {
	return NilDerefPtr[int32](valuePointer)
}

// NilValInt32 is a typed convenience wrapper for NilVal[int32].
func NilValInt32(valuePointer *int32, onNil, onNonNil int32) int32 {
	return NilVal[int32](valuePointer, onNil, onNonNil)
}

// NilValPtrInt32 is a typed convenience wrapper for NilValPtr[int32].
func NilValPtrInt32(valuePointer *int32, onNil, onNonNil int32) *int32 {
	return NilValPtr[int32](valuePointer, onNil, onNonNil)
}
