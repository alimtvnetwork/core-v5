package conditional

// IfFloat32 is a typed convenience wrapper for If[float32].
func IfFloat32(
	isTrue bool,
	trueValue, falseValue float32,
) float32 {
	return If[float32](isTrue, trueValue, falseValue)
}

// IfFuncFloat32 is a typed convenience wrapper for IfFunc[float32].
func IfFuncFloat32(
	isTrue bool,
	trueValueFunc, falseValueFunc func() float32,
) float32 {
	return IfFunc[float32](isTrue, trueValueFunc, falseValueFunc)
}

// IfTrueFuncFloat32 is a typed convenience wrapper for IfTrueFunc[float32].
func IfTrueFuncFloat32(
	isTrue bool,
	trueValueFunc func() float32,
) float32 {
	return IfTrueFunc[float32](isTrue, trueValueFunc)
}

// IfSliceFloat32 is a typed convenience wrapper for IfSlice[float32].
func IfSliceFloat32(
	isTrue bool,
	trueValue, falseValue []float32,
) []float32 {
	return IfSlice[float32](isTrue, trueValue, falseValue)
}

// IfSlicePtrFloat32 is a typed convenience wrapper for IfSlicePtr[float32].
func IfSlicePtrFloat32(
	isTrue bool,
	trueValue, falseValue *[]float32,
) *[]float32 {
	return IfSlicePtr[float32](isTrue, trueValue, falseValue)
}

// IfSlicePtrFuncFloat32 is a typed convenience wrapper for IfSlicePtrFunc[float32].
func IfSlicePtrFuncFloat32(
	isTrue bool,
	trueValueFunc, falseValueFunc func() *[]float32,
) *[]float32 {
	return IfSlicePtrFunc[float32](isTrue, trueValueFunc, falseValueFunc)
}

// IfPtrFloat32 is a typed convenience wrapper for IfPtr[float32].
func IfPtrFloat32(
	isTrue bool,
	trueValue, falseValue *float32,
) *float32 {
	return IfPtr[float32](isTrue, trueValue, falseValue)
}

// NilDefFloat32 is a typed convenience wrapper for NilDef[float32].
func NilDefFloat32(
	valuePointer *float32,
	defVal float32,
) float32 {
	return NilDef[float32](valuePointer, defVal)
}

// NilDefPtrFloat32 is a typed convenience wrapper for NilDefPtr[float32].
func NilDefPtrFloat32(
	valuePointer *float32,
	defVal float32,
) *float32 {
	return NilDefPtr[float32](valuePointer, defVal)
}
