package conditional

// Deprecated: Use ValueOrZero[int] instead.
func NilDefInt(
	valuePointer *int,
) int {
	return ValueOrZero[int](valuePointer)
}

// Deprecated: Use PtrOrZero[int] instead.
func NilDefIntPtr(
	valuePointer *int,
) *int {
	return PtrOrZero[int](valuePointer)
}

// Deprecated: Use NilDef[int] instead.
func NilDefValInt(
	valuePointer *int,
	defVal int,
) int {
	return NilDef[int](valuePointer, defVal)
}
