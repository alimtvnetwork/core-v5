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
