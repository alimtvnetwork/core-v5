package conditional

// Deprecated: Use NilDeref[int] instead.
func NilDefInt(
	valuePointer *int,
) int {
	return NilDeref[int](valuePointer)
}

// Deprecated: Use NilDerefPtr[int] instead.
func NilDefIntPtr(
	valuePointer *int,
) *int {
	return NilDerefPtr[int](valuePointer)
}

// Deprecated: Use NilDef[int] instead.
func NilDefValInt(
	valuePointer *int,
	defVal int,
) int {
	return NilDef[int](valuePointer, defVal)
}
