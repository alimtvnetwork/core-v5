package stringutil

// IsEndsRune searches for case sensitive terms
func IsEndsRune(
	content string,
	r rune,
) bool {
	length := len(content)

	if length == 0 {
		return false
	}

	runes := []rune(content)

	return runes[len(runes)-1] == r
}
