package utilstringinternal

func IsNullOrEmptyOrWhitespace(stringPtr *string) bool {
	return stringPtr == nil || IsEmptyOrWhitespace(*stringPtr)
}
