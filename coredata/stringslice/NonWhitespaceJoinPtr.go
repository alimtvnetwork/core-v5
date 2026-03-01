package stringslice

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

// Deprecated: Use NonWhitespaceJoin instead.
func NonWhitespaceJoinPtr(slice []string, joiner string) string {
	if len(slice) == 0 {
		return constants.EmptyString
	}

	return strings.Join(NonWhitespacePtr(slice), joiner)
}
