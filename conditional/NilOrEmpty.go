package conditional

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
