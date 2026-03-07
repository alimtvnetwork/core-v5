package conditional

// IfUint32 is a typed convenience wrapper for If[uint32].
func IfUint32(
	isTrue bool,
	trueValue, falseValue uint32,
) uint32 {
	return If[uint32](isTrue, trueValue, falseValue)
}

// IfFuncUint32 is a typed convenience wrapper for IfFunc[uint32].
func IfFuncUint32(
	isTrue bool,
	trueValueFunc, falseValueFunc func() uint32,
) uint32 {
	return IfFunc[uint32](isTrue, trueValueFunc, falseValueFunc)
}

// IfTrueFuncUint32 is a typed convenience wrapper for IfTrueFunc[uint32].
func IfTrueFuncUint32(
	isTrue bool,
	trueValueFunc func() uint32,
) uint32 {
	return IfTrueFunc[uint32](isTrue, trueValueFunc)
}

// IfSliceUint32 is a typed convenience wrapper for IfSlice[uint32].
func IfSliceUint32(
	isTrue bool,
	trueValue, falseValue []uint32,
) []uint32 {
	return IfSlice[uint32](isTrue, trueValue, falseValue)
}

// IfPtrUint32 is a typed convenience wrapper for IfPtr[uint32].
func IfPtrUint32(
	isTrue bool,
	trueValue, falseValue *uint32,
) *uint32 {
	return IfPtr[uint32](isTrue, trueValue, falseValue)
}

// NilDefUint32 is a typed convenience wrapper for NilDef[uint32].
func NilDefUint32(
	valuePointer *uint32,
	defVal uint32,
) uint32 {
	return NilDef[uint32](valuePointer, defVal)
}

// NilDefPtrUint32 is a typed convenience wrapper for NilDefPtr[uint32].
func NilDefPtrUint32(
	valuePointer *uint32,
	defVal uint32,
) *uint32 {
	return NilDefPtr[uint32](valuePointer, defVal)
}

// ValueOrZeroUint32 is a typed convenience wrapper for ValueOrZero[uint32].
func ValueOrZeroUint32(valuePointer *uint32) uint32 {
	return ValueOrZero[uint32](valuePointer)
}

// PtrOrZeroUint32 is a typed convenience wrapper for PtrOrZero[uint32].
func PtrOrZeroUint32(valuePointer *uint32) *uint32 {
	return PtrOrZero[uint32](valuePointer)
}

// NilValUint32 is a typed convenience wrapper for NilVal[uint32].
func NilValUint32(valuePointer *uint32, onNil, onNonNil uint32) uint32 {
	return NilVal[uint32](valuePointer, onNil, onNonNil)
}

// NilValPtrUint32 is a typed convenience wrapper for NilValPtr[uint32].
func NilValPtrUint32(valuePointer *uint32, onNil, onNonNil uint32) *uint32 {
	return NilValPtr[uint32](valuePointer, onNil, onNonNil)
}
