package simplewrap

import (
	"fmt"
	
	"gitlab.com/auk-go/core/constants"
)

// WithParenthesisQuotation
//
// (\"%v\")
func WithParenthesisQuotation(
	source any,
) string {
	return fmt.Sprintf(
		constants.ParenthesisQuotationWrap,
		toString(source))
}
