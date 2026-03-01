package stringslice

import "gitlab.com/auk-go/core/constants"

// Deprecated: Use LastIndex instead (on non-pointer slice).
func LastIndexPtr(slice []string) int {
	return len(slice) - constants.One
}
