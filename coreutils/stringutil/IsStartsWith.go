package stringutil

import (
	"strings"
)

func IsStartsWith(
	content, startsWith string,
	isIgnoreCase bool,
) bool {
	if startsWith == "" {
		return true
	}

	if content == "" {
		return startsWith == ""
	}

	if content == startsWith {
		return true
	}

	basePathLength := len(content)
	startsWithLength := len(startsWith)

	if startsWithLength > basePathLength {
		return false
	}

	if isIgnoreCase &&
		basePathLength == startsWithLength &&
		strings.EqualFold(content, startsWith) {
		return true
	}

	if basePathLength <= startsWithLength {
		return false
	}

	remainingText := content[:startsWithLength]

	if !isIgnoreCase {
		return startsWith == remainingText
	}

	return strings.EqualFold(startsWith, remainingText)
}
