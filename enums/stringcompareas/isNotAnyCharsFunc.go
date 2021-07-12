package stringcompareas

import "strings"

// isNotAnyCharsFunc reports whether any
// Unicode code points in chars are NOT within contentLine.
//
// Tided with NotAnyChars
var isNotAnyCharsFunc = func(
	contentLine,
	charsFind string,
	isCaseSensitive bool,
) bool {
	if isCaseSensitive {
		return !strings.ContainsAny(
			contentLine,
			charsFind,
		)
	}

	return !strings.ContainsAny(
		strings.ToLower(contentLine),
		strings.ToLower(charsFind),
	)
}
