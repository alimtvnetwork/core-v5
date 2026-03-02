package conditional

// Deprecated: Use ValueOrZero[bool] instead.
func NilDefBool(
	valuePointer *bool,
) bool {
	return ValueOrZero[bool](valuePointer)
}

// Deprecated: Use PtrOrZero[bool] instead.
func NilDefBoolPtr(
	valuePointer *bool,
) *bool {
	return PtrOrZero[bool](valuePointer)
}

// Deprecated: Use NilDef[bool] instead.
func NilBoolVal(
	valuePointer *bool,
	defVal bool,
) bool {
	return NilDef[bool](valuePointer, defVal)
}

// Deprecated: Use NilDefPtr[bool] instead.
func NilBoolValPtr(
	valuePointer *bool,
	defVal bool,
) *bool {
	return NilDefPtr[bool](valuePointer, defVal)
}
