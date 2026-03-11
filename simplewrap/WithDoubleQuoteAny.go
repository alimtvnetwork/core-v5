package simplewrap

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

// WithDoubleQuoteAny
//
//	Alias for ToJsonName
//
//	" + source + " , also take care of any double if available next.
func WithDoubleQuoteAny(source any) string {
	return fmt.Sprintf(
		constants.SprintValueDoubleQuotationFormat,
		toString(source))
}
