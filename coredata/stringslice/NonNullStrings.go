package stringslice

func NonNullStrings(
	slice []string,
) []string {
	if slice == nil {
		return []string{}
	}

	result := make([]string, 0, len(slice))
	for _, s := range slice {
		if s != "" {
			result = append(result, s)
		}
	}

	return result
}
