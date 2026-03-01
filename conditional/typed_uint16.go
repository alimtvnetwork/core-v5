package conditional

// IfUint16 is a typed convenience wrapper for If[uint16].
func IfUint16(
	isTrue bool,
	trueValue, falseValue uint16,
) uint16 {
	return If[uint16](isTrue, trueValue, falseValue)
}

// IfFuncUint16 is a typed convenience wrapper for IfFunc[uint16].
func IfFuncUint16(
	isTrue bool,
	trueValueFunc, falseValueFunc func() uint16,
) uint16 {
	return IfFunc[uint16](isTrue, trueValueFunc, falseValueFunc)
}

// IfTrueFuncUint16 is a typed convenience wrapper for IfTrueFunc[uint16].
func IfTrueFuncUint16(
	isTrue bool,
	trueValueFunc func() uint16,
) uint16 {
	return IfTrueFunc[uint16](isTrue, trueValueFunc)
}

// IfSliceUint16 is a typed convenience wrapper for IfSlice[uint16].
func IfSliceUint16(
	isTrue bool,
	trueValue, falseValue []uint16,
) []uint16 {
	return IfSlice[uint16](isTrue, trueValue, falseValue)
}

// IfPtrUint16 is a typed convenience wrapper for IfPtr[uint16].
func IfPtrUint16(
	isTrue bool,
	trueValue, falseValue *uint16,
) *uint16 {
	return IfPtr[uint16](isTrue, trueValue, falseValue)
}

// NilDefUint16 is a typed convenience wrapper for NilDef[uint16].
func NilDefUint16(
	valuePointer *uint16,
	defVal uint16,
) uint16 {
	return NilDef[uint16](valuePointer, defVal)
}

// NilDefPtrUint16 is a typed convenience wrapper for NilDefPtr[uint16].
func NilDefPtrUint16(
	valuePointer *uint16,
	defVal uint16,
) *uint16 {
	return NilDefPtr[uint16](valuePointer, defVal)
}
