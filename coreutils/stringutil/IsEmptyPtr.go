package stringutil

func IsEmptyPtr(str *string) bool {
	return str == nil || *str == ""
}
