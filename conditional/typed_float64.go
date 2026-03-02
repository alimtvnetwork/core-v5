package conditional

// IfFloat64 is a typed convenience wrapper for If[float64].
func IfFloat64(
	isTrue bool,
	trueValue, falseValue float64,
) float64 {
	return If[float64](isTrue, trueValue, falseValue)
}

// IfFuncFloat64 is a typed convenience wrapper for IfFunc[float64].
func IfFuncFloat64(
	isTrue bool,
	trueValueFunc, falseValueFunc func() float64,
) float64 {
	return IfFunc[float64](isTrue, trueValueFunc, falseValueFunc)
}

// IfTrueFuncFloat64 is a typed convenience wrapper for IfTrueFunc[float64].
func IfTrueFuncFloat64(
	isTrue bool,
	trueValueFunc func() float64,
) float64 {
	return IfTrueFunc[float64](isTrue, trueValueFunc)
}

// IfSliceFloat64 is a typed convenience wrapper for IfSlice[float64].
func IfSliceFloat64(
	isTrue bool,
	trueValue, falseValue []float64,
) []float64 {
	return IfSlice[float64](isTrue, trueValue, falseValue)
}

// Deprecated: Use IfSliceFloat64 instead.
func IfSlicePtrFloat64(
	isTrue bool,
	trueValue, falseValue []float64,
) []float64 {
	return IfSlice[float64](isTrue, trueValue, falseValue)
}

// Deprecated: Use IfSlice[float64] with func wrappers instead.
func IfSlicePtrFuncFloat64(
	isTrue bool,
	trueValueFunc, falseValueFunc func() []float64,
) []float64 {
	if isTrue {
		return trueValueFunc()
	}
	return falseValueFunc()
}

// IfPtrFloat64 is a typed convenience wrapper for IfPtr[float64].
func IfPtrFloat64(
	isTrue bool,
	trueValue, falseValue *float64,
) *float64 {
	return IfPtr[float64](isTrue, trueValue, falseValue)
}

// NilDefFloat64 is a typed convenience wrapper for NilDef[float64].
func NilDefFloat64(
	valuePointer *float64,
	defVal float64,
) float64 {
	return NilDef[float64](valuePointer, defVal)
}

// NilDefPtrFloat64 is a typed convenience wrapper for NilDefPtr[float64].
func NilDefPtrFloat64(
	valuePointer *float64,
	defVal float64,
) *float64 {
	return NilDefPtr[float64](valuePointer, defVal)
}

// ValueOrZeroFloat64 is a typed convenience wrapper for ValueOrZero[float64].
func ValueOrZeroFloat64(valuePointer *float64) float64 {
	return ValueOrZero[float64](valuePointer)
}

// PtrOrZeroFloat64 is a typed convenience wrapper for PtrOrZero[float64].
func PtrOrZeroFloat64(valuePointer *float64) *float64 {
	return PtrOrZero[float64](valuePointer)
}

// Deprecated: Use ValueOrZeroFloat64 instead.
func NilDerefFloat64(valuePointer *float64) float64 {
	return ValueOrZero[float64](valuePointer)
}

// Deprecated: Use PtrOrZeroFloat64 instead.
func NilDerefPtrFloat64(valuePointer *float64) *float64 {
	return PtrOrZero[float64](valuePointer)
}

// NilValFloat64 is a typed convenience wrapper for NilVal[float64].
func NilValFloat64(valuePointer *float64, onNil, onNonNil float64) float64 {
	return NilVal[float64](valuePointer, onNil, onNonNil)
}

// NilValPtrFloat64 is a typed convenience wrapper for NilValPtr[float64].
func NilValPtrFloat64(valuePointer *float64, onNil, onNonNil float64) *float64 {
	return NilValPtr[float64](valuePointer, onNil, onNonNil)
}
