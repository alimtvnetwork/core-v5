package stringutil

// IsEnds searches for case sensitive terms
func IsEnds(
	content,
	startsWith string,
) bool {
	return IsEndsWith(
		content,
		startsWith,
		false)
}
