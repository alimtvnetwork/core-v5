package stringutil

func IsAnyStartsWith(
	content string,
	isIgnoreCase bool,
	startsWithTerms ...string,
) bool {
	if len(content) > 0 && len(startsWithTerms) == 0 {
		return false
	}

	if len(content) == 0 && len(startsWithTerms) == 0 {
		return true
	}

	for _, term := range startsWithTerms {
		if IsStartsWith(
			content,
			term,
			isIgnoreCase) {
			return true
		}
	}

	return false
}
