package stringutil

import (
	"strings"
	"unicode"
)

// ReplaceWhiteSpacesToSingle
//
// Give "  some  nothing    \t" -> "some nothing"
func ReplaceWhiteSpacesToSingle(
	textContainsWhitespaces string,
) string {
	trimmedText := strings.TrimSpace(textContainsWhitespaces)

	if len(trimmedText) == 0 {
		return trimmedText
	}

	var sb strings.Builder
	sb.Grow(len(trimmedText))
	hasSpaceAlready := false
	isSpace := false

	for _, r := range trimmedText {
		isSpace = unicode.IsSpace(r)
		if r == '\n' || r == '\v' || r == '\f' || r == '\r' || isSpace && hasSpaceAlready {
			continue
		}

		sb.WriteRune(r)

		hasSpaceAlready = isSpace
	}

	return sb.String()
}
