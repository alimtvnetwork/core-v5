package stringslice

import "gitlab.com/auk-go/core/constants"

// Deprecated: Use MakeDefault instead.
func MakeDefaultPtr(capacity int) []string {
	return make([]string, constants.Zero, capacity)
}
