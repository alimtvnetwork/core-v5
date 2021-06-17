package stringutil

import "strings"

func IsEndsWith(
	baseStr,
	endsWith string,
	isIgnoreCase bool,
) bool {
	if endsWith == "" {
		return true
	}

	if baseStr == "" {
		return endsWith == ""
	}

	if baseStr == endsWith {
		return true
	}

	basePathLength := len(baseStr)
	endsWithLength := len(endsWith)

	if endsWithLength > basePathLength {
		return false
	}

	if isIgnoreCase &&
		basePathLength == endsWithLength &&
		strings.EqualFold(baseStr, endsWith) {
		return true
	}

	remainingLength := basePathLength - endsWithLength

	if remainingLength < 0 {
		return false
	}

	remainingText := baseStr[remainingLength:]

	if !isIgnoreCase {
		return endsWith == remainingText
	}

	return strings.EqualFold(endsWith, remainingText)
}
