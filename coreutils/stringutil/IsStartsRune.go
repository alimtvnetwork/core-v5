package stringutil

import "unicode/utf8"

// IsStartsRune searches for case sensitive terms
func IsStartsRune(
	content string,
	r rune,
) bool {
	firstRune, _ := utf8.DecodeRuneInString(content)

	return firstRune == r
}
