package simplewrap

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/internal/convertinternal"
)

// WithCurlyQuotation
//
// Example : {\"%v\"}
func WithCurlyQuotation(
	source any,
) string {
	toStr := convertinternal.
		AnyTo.
		SmartString(source)

	return fmt.Sprintf(
		constants.CurlyQuotationWrapFormat,
		toStr)
}
