package conditional

// Deprecated: Use ValueOrZero[string] instead.
func NilDefStr(
	strPtr *string,
) string {
	return ValueOrZero[string](strPtr)
}

// Deprecated: Use PtrOrZero[string] instead.
func NilDefStrPtr(
	strPtr *string,
) *string {
	return PtrOrZero[string](strPtr)
}

// Deprecated: Use NilVal[string] instead.
func NilStr(
	strPtr *string,
	onNil string,
	onNonNil string,
) string {
	return NilVal[string](strPtr, onNil, onNonNil)
}

// NilOrEmptyStr checks for both nil and empty string.
// No generic replacement — string-specific behavior.
func NilOrEmptyStr(
	strPtr *string,
	onNilOrEmpty string,
	onNonNilOrNonEmpty string,
) string {
	if strPtr == nil || *strPtr == "" {
		return onNilOrEmpty
	}

	return onNonNilOrNonEmpty
}

// NilOrEmptyStrPtr checks for both nil and empty string, returns pointer.
// No generic replacement — string-specific behavior.
func NilOrEmptyStrPtr(
	strPtr *string,
	onNilOrEmpty string,
	onNonNilOrNonEmpty string,
) *string {
	if strPtr == nil || *strPtr == "" {
		return &onNilOrEmpty
	}

	return &onNonNilOrNonEmpty
}
