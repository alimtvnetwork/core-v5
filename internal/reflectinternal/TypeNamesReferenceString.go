package reflectinternal

import (
	"strings"

	"github.com/alimtvnetwork/core/constants"
)

func TypeNamesReferenceString(
	isFullName bool,
	anyItems ...any,
) string {
	return "Reference (Types): " + strings.Join(
		TypeNames(isFullName, anyItems...),
		constants.CommaSpace)
}
