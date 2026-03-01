package conditional

// IfUint64 is a typed convenience wrapper for If[uint64].
func IfUint64(
	isTrue bool,
	trueValue, falseValue uint64,
) uint64 {
	return If[uint64](isTrue, trueValue, falseValue)
}

// IfFuncUint64 is a typed convenience wrapper for IfFunc[uint64].
func IfFuncUint64(
	isTrue bool,
	trueValueFunc, falseValueFunc func() uint64,
) uint64 {
	return IfFunc[uint64](isTrue, trueValueFunc, falseValueFunc)
}

// IfTrueFuncUint64 is a typed convenience wrapper for IfTrueFunc[uint64].
func IfTrueFuncUint64(
	isTrue bool,
	trueValueFunc func() uint64,
) uint64 {
	return IfTrueFunc[uint64](isTrue, trueValueFunc)
}

// IfSliceUint64 is a typed convenience wrapper for IfSlice[uint64].
func IfSliceUint64(
	isTrue bool,
	trueValue, falseValue []uint64,
) []uint64 {
	return IfSlice[uint64](isTrue, trueValue, falseValue)
}

// IfPtrUint64 is a typed convenience wrapper for IfPtr[uint64].
func IfPtrUint64(
	isTrue bool,
	trueValue, falseValue *uint64,
) *uint64 {
	return IfPtr[uint64](isTrue, trueValue, falseValue)
}

// NilDefUint64 is a typed convenience wrapper for NilDef[uint64].
func NilDefUint64(
	valuePointer *uint64,
	defVal uint64,
) uint64 {
	return NilDef[uint64](valuePointer, defVal)
}

// NilDefPtrUint64 is a typed convenience wrapper for NilDefPtr[uint64].
func NilDefPtrUint64(
	valuePointer *uint64,
	defVal uint64,
) *uint64 {
	return NilDefPtr[uint64](valuePointer, defVal)
}
