package reflectinternal

import (
	"strings"

	"github.com/alimtvnetwork/core/constants"
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
