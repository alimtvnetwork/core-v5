package stringslice

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

// Deprecated: Use NonEmptyJoin instead.
func NonEmptyJoinPtr(slice []string, joiner string) string {
	if len(slice) == 0 {
		return constants.EmptyString
	}

	return strings.Join(NonEmptySlicePtr(slice), joiner)
}
