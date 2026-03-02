package conditional

// Deprecated: Use ValueOrZero[byte] instead.
func NilDefByte(
	valuePointer *byte,
) byte {
	return ValueOrZero[byte](valuePointer)
}

// Deprecated: Use PtrOrZero[byte] instead.
func NilDefBytePtr(
	valuePointer *byte,
) *byte {
	return PtrOrZero[byte](valuePointer)
}

// Deprecated: Use NilDef[byte] instead.
func NilByteVal(
	valuePointer *byte,
	defVal byte,
) byte {
	return NilDef[byte](valuePointer, defVal)
}

// Deprecated: Use NilDefPtr[byte] instead.
func NilByteValPtr(
	valuePointer *byte,
	defVal byte,
) *byte {
	return NilDefPtr[byte](valuePointer, defVal)
}
