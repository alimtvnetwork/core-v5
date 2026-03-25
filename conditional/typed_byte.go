package conditional

// IfByte is a typed convenience wrapper for If[byte].
func IfByte(
	isTrue bool,
	trueValue, falseValue byte,
) byte {
	return If[byte](isTrue, trueValue, falseValue)
}

// IfFuncByte is a typed convenience wrapper for IfFunc[byte].
func IfFuncByte(
	isTrue bool,
	trueValueFunc, falseValueFunc func() byte,
) byte {
	return IfFunc[byte](isTrue, trueValueFunc, falseValueFunc)
}

// IfTrueFuncByte is a typed convenience wrapper for IfTrueFunc[byte].
func IfTrueFuncByte(
	isTrue bool,
	trueValueFunc func() byte,
) byte {
	return IfTrueFunc[byte](isTrue, trueValueFunc)
}

// IfSliceByte is a typed convenience wrapper for IfSlice[byte].
func IfSliceByte(
	isTrue bool,
	trueValue, falseValue []byte,
) []byte {
	return IfSlice[byte](isTrue, trueValue, falseValue)
}

// IfPtrByte is a typed convenience wrapper for IfPtr[byte].
func IfPtrByte(
	isTrue bool,
	trueValue, falseValue *byte,
) *byte {
	return IfPtr[byte](isTrue, trueValue, falseValue)
}

// NilDefByte is a typed convenience wrapper for NilDef[byte].
func NilDefByte(
	valuePointer *byte,
	defVal byte,
) byte {
	return NilDef[byte](valuePointer, defVal)
}

// NilDefPtrByte is a typed convenience wrapper for NilDefPtr[byte].
func NilDefPtrByte(
	valuePointer *byte,
	defVal byte,
) *byte {
	return NilDefPtr[byte](valuePointer, defVal)
}

// ValueOrZeroByte is a typed convenience wrapper for ValueOrZero[byte].
func ValueOrZeroByte(valuePointer *byte) byte {
	return ValueOrZero[byte](valuePointer)
}

// PtrOrZeroByte is a typed convenience wrapper for PtrOrZero[byte].
func PtrOrZeroByte(valuePointer *byte) *byte {
	return PtrOrZero[byte](valuePointer)
}

// NilValByte is a typed convenience wrapper for NilVal[byte].
func NilValByte(valuePointer *byte, onNil, onNonNil byte) byte {
	return NilVal[byte](valuePointer, onNil, onNonNil)
}

// NilValPtrByte is a typed convenience wrapper for NilValPtr[byte].
func NilValPtrByte(valuePointer *byte, onNil, onNonNil byte) *byte {
	return NilValPtr[byte](valuePointer, onNil, onNonNil)
}
