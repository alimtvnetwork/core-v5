package stringutil

func IsStartsAndEndsWith(
	content, startsWith, endsWith string,
	isIgnoreCase bool,
) bool {
	return IsStartsWith(
		content,
		startsWith,
		isIgnoreCase) &&
		IsEndsWith(
			content,
			endsWith,
			isIgnoreCase)

}

// IsStartsAndEnds case sensitive term
func IsStartsAndEnds(
	content, startsWith, endsWith string,
) bool {
	return IsStartsWith(
		content,
		startsWith,
		false) &&
		IsEndsWith(
			content,
			endsWith,
			false)

}
