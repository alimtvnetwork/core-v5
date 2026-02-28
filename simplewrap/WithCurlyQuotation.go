package simplewrap

import (
	"fmt"
	
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/internal/convertinternal"
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
