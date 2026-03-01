package stringslice

import "gitlab.com/auk-go/core/constants"

// Deprecated: Use FirstOrDefault instead.
func FirstOrDefaultPtr(slice []string) string {
	if len(slice) == 0 {
		return constants.EmptyString
	}

	return slice[0]
}
