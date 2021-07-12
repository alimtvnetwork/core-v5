package stringcompareas

import "strings"

// isAnyCharsFunc reports whether any Unicode
// code points in chars are within contentLine.
//
// Tided with AnyChars
var isAnyCharsFunc = func(
	contentLine,
	charsFind string,
	isCaseSensitive bool,
) bool {
	if isCaseSensitive {
		return strings.ContainsAny(
			contentLine,
			charsFind,
		)
	}

	return strings.ContainsAny(
		strings.ToLower(contentLine),
		strings.ToLower(charsFind),
	)
}
