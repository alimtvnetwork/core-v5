package stringutil

// IsStarts searches for case sensitive terms
func IsStarts(
	content,
	startsWith string,
) bool {
	return IsStartsWith(
		content,
		startsWith,
		false)
}
