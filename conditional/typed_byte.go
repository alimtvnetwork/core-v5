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

// Deprecated: Use IfSliceByte instead.
func IfSlicePtrByte(
	isTrue bool,
	trueValue, falseValue []byte,
) []byte {
	return IfSlice[byte](isTrue, trueValue, falseValue)
}

// Deprecated: Use IfSlice[byte] with func wrappers instead.
func IfSlicePtrFuncByte(
	isTrue bool,
	trueValueFunc, falseValueFunc func() []byte,
) []byte {
	if isTrue {
		return trueValueFunc()
	}
	return falseValueFunc()
}

// IfPtrByte is a typed convenience wrapper for IfPtr[byte].
func IfPtrByte(
	isTrue bool,
	trueValue, falseValue *byte,
) *byte {
	return IfPtr[byte](isTrue, trueValue, falseValue)
}

// NilDefPtrByte is a typed convenience wrapper for NilDefPtr[byte].
// Note: NilDef[byte] typed wrapper is omitted due to naming conflict
// with the deprecated NilDefByte(ptr) in NilDefByte.go.
// Use NilDef[byte](ptr, defVal) directly.
func NilDefPtrByte(
	valuePointer *byte,
	defVal byte,
) *byte {
	return NilDefPtr[byte](valuePointer, defVal)
}
