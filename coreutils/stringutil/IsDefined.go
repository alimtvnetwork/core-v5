package stringutil

import "strings"

// IsDefined alias for NOT IsEmptyOrWhitespace
func IsDefined(str string) bool {
	return !(str == "" || str == " " || str == "\n" || strings.TrimSpace(str) == "")
}
