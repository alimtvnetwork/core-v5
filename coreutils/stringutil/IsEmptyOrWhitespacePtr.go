package stringutil

func IsEmptyOrWhitespacePtr(stringPtr *string) bool {
	return stringPtr == nil || IsEmptyOrWhitespace(*stringPtr)
}
