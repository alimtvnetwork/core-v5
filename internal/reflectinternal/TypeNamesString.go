package reflectinternal

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

// TypeNamesString
//
// Multiple type names as csv using TypeNames
func TypeNamesString(
	isFullName bool,
	anyItems ...any,
) string {
	return strings.Join(
		TypeNames(isFullName, anyItems...),
		constants.CommaSpace,
	)
}
