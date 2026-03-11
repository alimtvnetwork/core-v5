package simplewrap

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

// ToJsonName
//
//	Alias for WithDoubleQuoteAny
//
//	" + source + " , also take care of any double if available next.
func ToJsonName(source any) string {
	return fmt.Sprintf(
		constants.SprintValueDoubleQuotationFormat,
		toString(source))
}
