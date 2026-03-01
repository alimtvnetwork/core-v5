package conditional

// IfUint is a typed convenience wrapper for If[uint].
func IfUint(
	isTrue bool,
	trueValue, falseValue uint,
) uint {
	return If[uint](isTrue, trueValue, falseValue)
}

// IfFuncUint is a typed convenience wrapper for IfFunc[uint].
func IfFuncUint(
	isTrue bool,
	trueValueFunc, falseValueFunc func() uint,
) uint {
	return IfFunc[uint](isTrue, trueValueFunc, falseValueFunc)
}

// IfTrueFuncUint is a typed convenience wrapper for IfTrueFunc[uint].
func IfTrueFuncUint(
	isTrue bool,
	trueValueFunc func() uint,
) uint {
	return IfTrueFunc[uint](isTrue, trueValueFunc)
}

// IfSliceUint is a typed convenience wrapper for IfSlice[uint].
func IfSliceUint(
	isTrue bool,
	trueValue, falseValue []uint,
) []uint {
	return IfSlice[uint](isTrue, trueValue, falseValue)
}

// IfPtrUint is a typed convenience wrapper for IfPtr[uint].
func IfPtrUint(
	isTrue bool,
	trueValue, falseValue *uint,
) *uint {
	return IfPtr[uint](isTrue, trueValue, falseValue)
}

// NilDefUint is a typed convenience wrapper for NilDef[uint].
func NilDefUint(
	valuePointer *uint,
	defVal uint,
) uint {
	return NilDef[uint](valuePointer, defVal)
}

// NilDefPtrUint is a typed convenience wrapper for NilDefPtr[uint].
func NilDefPtrUint(
	valuePointer *uint,
	defVal uint,
) *uint {
	return NilDefPtr[uint](valuePointer, defVal)
}
