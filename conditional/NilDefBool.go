package conditional

// Deprecated: Use NilDeref[bool] instead.
func NilDefBool(
	valuePointer *bool,
) bool {
	return NilDeref[bool](valuePointer)
}

// Deprecated: Use NilDerefPtr[bool] instead.
func NilDefBoolPtr(
	valuePointer *bool,
) *bool {
	return NilDerefPtr[bool](valuePointer)
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
