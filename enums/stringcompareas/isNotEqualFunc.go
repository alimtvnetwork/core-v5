package stringcompareas

import "strings"

// isNotEqualFunc tided with NotEqual
var isNotEqualFunc = func(
	contentLine,
	notEqualText string,
	isCaseSensitive bool,
) bool {
	if isCaseSensitive {
		return contentLine != notEqualText
	}

	return !strings.EqualFold(
		notEqualText,
		contentLine)
}
