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

// ValueOrZeroFloat32 is a typed convenience wrapper for ValueOrZero[float32].
func ValueOrZeroFloat32(valuePointer *float32) float32 {
	return ValueOrZero[float32](valuePointer)
}

// PtrOrZeroFloat32 is a typed convenience wrapper for PtrOrZero[float32].
func PtrOrZeroFloat32(valuePointer *float32) *float32 {
	return PtrOrZero[float32](valuePointer)
}

// NilValFloat32 is a typed convenience wrapper for NilVal[float32].
func NilValFloat32(valuePointer *float32, onNil, onNonNil float32) float32 {
	return NilVal[float32](valuePointer, onNil, onNonNil)
}

// NilValPtrFloat32 is a typed convenience wrapper for NilValPtr[float32].
func NilValPtrFloat32(valuePointer *float32, onNil, onNonNil float32) *float32 {
	return NilValPtr[float32](valuePointer, onNil, onNonNil)
}
