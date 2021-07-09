package stringutil

// IsStartsAndEndsChar case sensitive term
func IsStartsAndEndsChar(
	content string,
	startsWith, endsWith byte,
) bool {
	length := len(content)

	if length == 0 {
		return false
	}

	return content[0] == startsWith &&
		content[length-1] == endsWith

}
