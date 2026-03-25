package enumimpl

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

// toJsonName
//
//	" + source + " , also take care of any double if available next.
func toJsonName(source any) string {
	return fmt.Sprintf(
		constants.SprintValueDoubleQuotationFormat,
		source)
}
