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

// Deprecated: Use IfSliceBool instead.
func IfSlicePtrBool(
	isTrue bool,
	trueValue, falseValue []bool,
) []bool {
	return IfSlice[bool](isTrue, trueValue, falseValue)
}

// Deprecated: Use IfSlice[bool] with func wrappers instead.
func IfSlicePtrFuncBool(
	isTrue bool,
	trueValueFunc, falseValueFunc func() []bool,
) []bool {
	if isTrue {
		return trueValueFunc()
	}
	return falseValueFunc()
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

// ValueOrZeroBool is a typed convenience wrapper for ValueOrZero[bool].
func ValueOrZeroBool(valuePointer *bool) bool {
	return ValueOrZero[bool](valuePointer)
}

// PtrOrZeroBool is a typed convenience wrapper for PtrOrZero[bool].
func PtrOrZeroBool(valuePointer *bool) *bool {
	return PtrOrZero[bool](valuePointer)
}

// Deprecated: Use ValueOrZeroBool instead.
func NilDerefBool(valuePointer *bool) bool {
	return ValueOrZero[bool](valuePointer)
}

// Deprecated: Use PtrOrZeroBool instead.
func NilDerefPtrBool(valuePointer *bool) *bool {
	return PtrOrZero[bool](valuePointer)
}

// NilValBool is a typed convenience wrapper for NilVal[bool].
func NilValBool(valuePointer *bool, onNil, onNonNil bool) bool {
	return NilVal[bool](valuePointer, onNil, onNonNil)
}

// NilValPtrBool is a typed convenience wrapper for NilValPtr[bool].
func NilValPtrBool(valuePointer *bool, onNil, onNonNil bool) *bool {
	return NilValPtr[bool](valuePointer, onNil, onNonNil)
}
