package stringcompareas

import (
	"path/filepath"
	"strings"
)

// isGlobFunc matches content against a glob pattern using filepath.Match.
//
// Supports wildcards:
//   - '*'  matches any sequence of non-separator characters
//   - '?'  matches any single non-separator character
//   - '[' range ']' matches a character class
//
// When isIgnoreCase is true, both pattern and content are lowercased before matching.
//
// If the pattern is invalid, returns false (no panic).
var isGlobFunc = func(
	contentLine,
	globPattern string,
	isIgnoreCase bool,
) bool {
	if isIgnoreCase {
		contentLine = strings.ToLower(contentLine)
		globPattern = strings.ToLower(globPattern)
	}

	isMatch, err := filepath.Match(globPattern, contentLine)
	if err != nil {
		return false
	}

	return isMatch
}
