package stringutil

// IsDefinedPtr alias for NOT IsEmptyOrWhitespace
func IsDefinedPtr(str *string) bool {
	return !(str == nil || IsEmptyOrWhitespace(*str))
}
