package simplewrap

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
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
