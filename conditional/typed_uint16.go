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

// NilDerefUint16 is a typed convenience wrapper for NilDeref[uint16].
func NilDerefUint16(valuePointer *uint16) uint16 {
	return NilDeref[uint16](valuePointer)
}

// NilDerefPtrUint16 is a typed convenience wrapper for NilDerefPtr[uint16].
func NilDerefPtrUint16(valuePointer *uint16) *uint16 {
	return NilDerefPtr[uint16](valuePointer)
}

// NilValUint16 is a typed convenience wrapper for NilVal[uint16].
func NilValUint16(valuePointer *uint16, onNil, onNonNil uint16) uint16 {
	return NilVal[uint16](valuePointer, onNil, onNonNil)
}

// NilValPtrUint16 is a typed convenience wrapper for NilValPtr[uint16].
func NilValPtrUint16(valuePointer *uint16, onNil, onNonNil uint16) *uint16 {
	return NilValPtr[uint16](valuePointer, onNil, onNonNil)
}
