package stringutil

// IsBlankPtr alias for IsEmptyOrWhitespace
func IsBlankPtr(s *string) bool {
	return s == nil || IsBlank(*s)
}
