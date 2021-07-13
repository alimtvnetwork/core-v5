package stringslice

func NonEmptyIf(
	isNonEmpty bool,
	slice []string,
) []string {
	if isNonEmpty {
		return NonEmptySlice(slice)
	}

	return NonNullStrings(slice)
}
