package stringutil

func IsNullOrEmptyPtr(stringPtr *string) bool {
	return stringPtr == nil || *stringPtr == ""
}
