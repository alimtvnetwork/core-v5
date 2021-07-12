package stringutil

func IsAnyEndsWith(
	content string,
	isIgnoreCase bool,
	endsWithTerms ...string,
) bool {
	if len(content) > 0 && len(endsWithTerms) == 0 {
		return false
	}

	if len(content) == 0 && len(endsWithTerms) == 0 {
		return true
	}

	for _, term := range endsWithTerms {
		if IsEndsWith(
			content,
			term,
			isIgnoreCase) {
			return true
		}
	}

	return false
}
