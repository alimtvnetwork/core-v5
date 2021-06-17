package stringutil

import "strings"

func IsContainsPtrSimple(
	lines *[]string,
	findingString string,
	startsAtIndex int,
	isCaseSensitive bool,
) bool {
	if lines == nil || *lines == nil {
		return false
	}

	length := len(*lines)

	if length == 0 {
		return false
	}

	if !isCaseSensitive {
		// insensitive
		for i := startsAtIndex; i < length; i++ {
			if strings.EqualFold((*lines)[i], findingString) {
				return true
			}
		}

		return false
	}

	for i := startsAtIndex; i < length; i++ {
		if (*lines)[i] == findingString {
			return true
		}
	}

	return false
}
