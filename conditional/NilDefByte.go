package conditional

// Deprecated: Use NilDeref[byte] instead.
func NilDefByte(
	valuePointer *byte,
) byte {
	return NilDeref[byte](valuePointer)
}

// Deprecated: Use NilDerefPtr[byte] instead.
func NilDefBytePtr(
	valuePointer *byte,
) *byte {
	return NilDerefPtr[byte](valuePointer)
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
