package stringslice

func SafeIndexAtWith(
	slice []string,
	index int,
	defaultVal string,
) string {
	if len(slice) == 0 {
		return defaultVal
	}

	return slice[index]
}
