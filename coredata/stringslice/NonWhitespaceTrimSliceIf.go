package stringslice

func NonWhitespaceTrimSliceIf(
	isNonWhitespaceTrim bool,
	slice []string,
) []string {
	if isNonWhitespaceTrim {
		return NonWhitespaceTrimSlice(slice)
	}

	return NonNullStrings(slice)
}
