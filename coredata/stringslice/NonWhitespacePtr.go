package stringslice

// Deprecated: Use NonWhitespace instead.
func NonWhitespacePtr(
	slice []string,
) []string {
	if len(slice) == 0 {
		return []string{}
	}

	return NonWhitespace(slice)
}
