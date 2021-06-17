package stringutil

import "strings"

func IsStartsWith(
	baseStr, startsWith string,
	isIgnoreCase bool,
) bool {
	if startsWith == "" {
		return true
	}

	if baseStr == "" {
		return startsWith == ""
	}

	if baseStr == startsWith {
		return true
	}

	basePathLength := len(baseStr)
	startsWithLength := len(startsWith)

	if startsWithLength > basePathLength {
		return false
	}

	if isIgnoreCase &&
		basePathLength == startsWithLength &&
		strings.EqualFold(baseStr, startsWith) {
		return true
	}

	if basePathLength <= startsWithLength {
		return false
	}

	remainingText := baseStr[:startsWithLength]

	if !isIgnoreCase {
		return startsWith == remainingText
	}

	return strings.EqualFold(startsWith, remainingText)
}
