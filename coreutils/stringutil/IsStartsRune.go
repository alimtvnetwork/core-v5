package stringutil

// IsStartsRune searches for case sensitive terms
func IsStartsRune(
	content string,
	r rune,
) bool {
	return len(content) > 0 && []rune(content)[0] == r
}
