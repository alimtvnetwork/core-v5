package stringslice

import "strings"

func SplitContentsByWhitespace(
	input string,
) []string {
	return strings.Fields(input)
}
