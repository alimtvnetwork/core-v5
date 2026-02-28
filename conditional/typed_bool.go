package conditional

// IfBool is a typed convenience wrapper for If[bool].
func IfBool(
	isTrue bool,
	trueValue, falseValue bool,
) bool {
	return If[bool](isTrue, trueValue, falseValue)
}

// IfFuncBool is a typed convenience wrapper for IfFunc[bool].
func IfFuncBool(
	isTrue bool,
	trueValueFunc, falseValueFunc func() bool,
) bool {
	return IfFunc[bool](isTrue, trueValueFunc, falseValueFunc)
}

// IfTrueFuncBool is a typed convenience wrapper for IfTrueFunc[bool].
func IfTrueFuncBool(
	isTrue bool,
	trueValueFunc func() bool,
) bool {
	return IfTrueFunc[bool](isTrue, trueValueFunc)
}

// IfSliceBool is a typed convenience wrapper for IfSlice[bool].
func IfSliceBool(
	isTrue bool,
	trueValue, falseValue []bool,
) []bool {
	return IfSlice[bool](isTrue, trueValue, falseValue)
}

// IfSlicePtrBool is a typed convenience wrapper for IfSlicePtr[bool].
func IfSlicePtrBool(
	isTrue bool,
	trueValue, falseValue *[]bool,
) *[]bool {
	return IfSlicePtr[bool](isTrue, trueValue, falseValue)
}

// IfSlicePtrFuncBool is a typed convenience wrapper for IfSlicePtrFunc[bool].
func IfSlicePtrFuncBool(
	isTrue bool,
	trueValueFunc, falseValueFunc func() *[]bool,
) *[]bool {
	return IfSlicePtrFunc[bool](isTrue, trueValueFunc, falseValueFunc)
}

// IfPtrBool is a typed convenience wrapper for IfPtr[bool].
func IfPtrBool(
	isTrue bool,
	trueValue, falseValue *bool,
) *bool {
	return IfPtr[bool](isTrue, trueValue, falseValue)
}

// NilDefPtrBool is a typed convenience wrapper for NilDefPtr[bool].
// Note: NilDef[bool] typed wrapper is omitted due to naming conflict
// with the deprecated NilDefBool(ptr) in NilDefBool.go.
// Use NilDef[bool](ptr, defVal) directly.
func NilDefPtrBool(
	valuePointer *bool,
	defVal bool,
) *bool {
	return NilDefPtr[bool](valuePointer, defVal)
}
