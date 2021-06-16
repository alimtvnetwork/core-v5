package stringslice

func FirstOrDefaultWith(slice []string, def string) (result string, isSuccess bool) {
	if len(slice) == 0 {
		return def, false
	}

	return slice[0], true
}
