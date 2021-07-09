package stringutil

// IsEndsChar searches for case sensitive terms
func IsEndsChar(
	content string,
	char byte,
) bool {
	length := len(content)

	if length == 0 {
		return false
	}

	return content[length-1] == char
}
