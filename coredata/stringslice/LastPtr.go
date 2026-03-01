package stringslice

import "gitlab.com/auk-go/core/constants"

// Deprecated: Use Last instead.
func LastPtr(slice []string) string {
	return slice[len(slice)-constants.One]
}
