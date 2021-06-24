package stringutil

import "strings"

// IsBlank alias for IsEmptyOrWhitespace
func IsBlank(str string) bool {
	return str == "" || str == " " || str == "\n" || strings.TrimSpace(str) == ""
}
