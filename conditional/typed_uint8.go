package conditional

// IfUint8 is a typed convenience wrapper for If[uint8].
func IfUint8(
	isTrue bool,
	trueValue, falseValue uint8,
) uint8 {
	return If[uint8](isTrue, trueValue, falseValue)
}

// IfFuncUint8 is a typed convenience wrapper for IfFunc[uint8].
func IfFuncUint8(
	isTrue bool,
	trueValueFunc, falseValueFunc func() uint8,
) uint8 {
	return IfFunc[uint8](isTrue, trueValueFunc, falseValueFunc)
}

// IfTrueFuncUint8 is a typed convenience wrapper for IfTrueFunc[uint8].
func IfTrueFuncUint8(
	isTrue bool,
	trueValueFunc func() uint8,
) uint8 {
	return IfTrueFunc[uint8](isTrue, trueValueFunc)
}

// IfSliceUint8 is a typed convenience wrapper for IfSlice[uint8].
func IfSliceUint8(
	isTrue bool,
	trueValue, falseValue []uint8,
) []uint8 {
	return IfSlice[uint8](isTrue, trueValue, falseValue)
}

// IfPtrUint8 is a typed convenience wrapper for IfPtr[uint8].
func IfPtrUint8(
	isTrue bool,
	trueValue, falseValue *uint8,
) *uint8 {
	return IfPtr[uint8](isTrue, trueValue, falseValue)
}

// NilDefUint8 is a typed convenience wrapper for NilDef[uint8].
func NilDefUint8(
	valuePointer *uint8,
	defVal uint8,
) uint8 {
	return NilDef[uint8](valuePointer, defVal)
}

// NilDefPtrUint8 is a typed convenience wrapper for NilDefPtr[uint8].
func NilDefPtrUint8(
	valuePointer *uint8,
	defVal uint8,
) *uint8 {
	return NilDefPtr[uint8](valuePointer, defVal)
}

// ValueOrZeroUint8 is a typed convenience wrapper for ValueOrZero[uint8].
func ValueOrZeroUint8(valuePointer *uint8) uint8 {
	return ValueOrZero[uint8](valuePointer)
}

// PtrOrZeroUint8 is a typed convenience wrapper for PtrOrZero[uint8].
func PtrOrZeroUint8(valuePointer *uint8) *uint8 {
	return PtrOrZero[uint8](valuePointer)
}

// NilValUint8 is a typed convenience wrapper for NilVal[uint8].
func NilValUint8(valuePointer *uint8, onNil, onNonNil uint8) uint8 {
	return NilVal[uint8](valuePointer, onNil, onNonNil)
}

// NilValPtrUint8 is a typed convenience wrapper for NilValPtr[uint8].
func NilValPtrUint8(valuePointer *uint8, onNil, onNonNil uint8) *uint8 {
	return NilValPtr[uint8](valuePointer, onNil, onNonNil)
}
