package stringutil

// IsStartsChar searches for case sensitive terms
func IsStartsChar(
	content string,
	char byte,
) bool {
	return len(content) > 0 && content[0] == char
}
