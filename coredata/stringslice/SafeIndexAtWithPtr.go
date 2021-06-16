package stringslice

func SafeIndexAtWithPtr(
	slice *[]string,
	index int,
	defaultVal string,
) string {
	if IsEmptyPtr(slice) {
		return defaultVal
	}

	return (*slice)[index]
}
