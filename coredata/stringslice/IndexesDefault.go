package stringslice

func IndexesDefault(slice []string, indexes ...int) (values []string) {
	length := len(slice)

	if length == 0 || len(indexes) == 0 {
		return []string{}
	}

	values = make([]string, len(indexes))

	inputIndex := 0
	for _, index := range indexes {
		values[inputIndex] = slice[index]
		inputIndex++
	}

	return values
}
