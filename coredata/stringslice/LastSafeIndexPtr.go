package stringslice

import "gitlab.com/auk-go/core/constants"

// Deprecated: Use LastSafeIndex instead (on non-pointer slice).
func LastSafeIndexPtr(slice []string) int {
	if IsEmptyPtr(slice) {
		return constants.InvalidNotFoundCase
	}

	return len(slice) - constants.One
}
