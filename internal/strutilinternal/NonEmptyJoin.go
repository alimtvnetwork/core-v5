package strutilinternal

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

func NonEmptyJoin(slice []string, joiner string) string {
	if len(slice) == 0 {
		return constants.EmptyString
	}

	return strings.Join(NonEmptySlicePtr(slice), joiner)
}
