package stringutil

import "unicode/utf8"

// IsEndsRune searches for case sensitive terms
func IsEndsRune(
	content string,
	r rune,
) bool {
	lastOne, _ := utf8.DecodeLastRuneInString(content)

	return lastOne == r
}
