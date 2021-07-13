package stringslice

func NonNullStrings(
	slice []string,
) []string {
	if slice == nil {
		return []string{}
	}

	return slice
}
